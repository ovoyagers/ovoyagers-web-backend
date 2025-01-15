package petservice

import (
	"context"
	"errors"

	"github.com/imagekit-developer/imagekit-go/api"
	"github.com/imagekit-developer/imagekit-go/api/media"
	"github.com/imagekit-developer/imagekit-go/api/uploader"
	"github.com/petmeds24/backend/pkg/rest/src/models/petmodel"
)

var ctx = context.Background()

// InsertNewPetImage uploads a new pet image to imagekit, given the image metadata.
//
// The image is uploaded to a folder in the format: /<user_id>/<pet_id>.
//
// The function returns the upload response from imagekit, which contains the URL
// of the uploaded image. If an error occurs, the function returns the error.
func (p *PetService) InsertNewPetImage(imgMetadata petmodel.PetProfilePicture) (*uploader.UploadResponse, error) {
	// Upload the image to imageKit using the image metadata. The image is uploaded
	// to a folder in the format: /<user_id>/<pet_id>. This folder structure is
	// used to keep all the images associated with a particular user and pet
	// organized.
	resp, err := p.imageKit.Kit.Uploader.Upload(ctx, imgMetadata.ImageBytes, uploader.UploadParam{
		FileName: imgMetadata.FileName,
		Folder:   "/" + imgMetadata.UserId + "/" + imgMetadata.PetId,
	})
	// Handle error
	//
	// If an error occurs during the upload process, return the error.
	// This error will be propagated upwards to the calling code.
	if err != nil {
		return nil, err
	}

	// Print upload response
	//
	// If the upload is successful, return the upload response from imageKit.
	// This response contains the URL of the uploaded image.
	return resp, nil
}

// UpdatePetImage updates the pet's image by first deleting the existing image
// if a fileId is provided, and then uploading the new image using the provided
// image metadata. The new image is uploaded to a folder structured as:
// /<user_id>/<pet_id>. If the upload is successful, the pet's profile picture
// information is updated in the database. Returns an error if any operation fails.
func (p *PetService) UpdatePetImage(imgMetadata petmodel.PetProfilePicture) error {
	// First, get the fileId of the existing image so we can delete it later.
	profilePicture, err := p.petDao.GetProfilePictureByPetId(imgMetadata.PetId)
	if err != nil {
		// If we can't get the existing image, return the error.
		return err
	}

	// If there is an existing image, delete it now.
	if profilePicture.FileId != "" {
		if _, err := p.DeletePetImage(profilePicture.FileId); err != nil {
			// If deleting the image fails, return the error.
			return err
		}
	}

	// Upload the new image to imageKit.
	resp, err := p.InsertNewPetImage(imgMetadata)
	if err != nil {
		// If the upload fails, return the error.
		return err
	}

	// Now that the image is uploaded, create a ProfilePicture struct to hold
	// the information about the new image.
	petImage := petmodel.ProfilePicture{
		FileId:       resp.Data.FileId,
		Url:          resp.Data.Url,
		Name:         resp.Data.Name,
		ThumbnailUrl: resp.Data.ThumbnailUrl,
	}

	// Update the pet's profile picture information in the database.
	if err := p.petDao.UpdateProfilePicture(imgMetadata.PetId, petImage); err != nil {
		// If updating the profile picture fails, return the error.
		return err
	}

	// If all operations were successful, return nil.
	return nil
}

// DeletePetImage deletes an image associated with a pet using its unique fileId.
// It communicates with the imageKit service to remove the image file.
// Returns the API response from imageKit if successful, or an error if the operation fails.
func (p *PetService) DeletePetImage(fileId string) (*api.Response, error) {
	// Call the DeleteFile method on the Media interface of imageKit.
	// This function takes a context and a fileId, and attempts to delete the file
	// associated with that fileId from the imageKit server.
	resp, err := p.imageKit.Kit.Media.DeleteFile(ctx, fileId)

	// Check if an error occurred during the deletion process.
	// If an error is returned, it indicates that the deletion failed,
	// so we return nil for the response and propagate the error upwards.
	if err != nil {
		return nil, err
	}

	// If no error occurred, the deletion was successful.
	// Return the response from imageKit, which contains details of the operation.
	return resp, nil
}

// DeletePetFolder is responsible for deleting a folder associated with a pet
// using the provided userId and petId. This function interacts with the imageKit
// service to perform the deletion operation. It returns the API response from
// imageKit if the operation is successful, or an error if the operation fails.
func (p *PetService) DeletePetFolder(userId string, petId string) (*api.Response, error) {
	// Construct the folder path using the userId and petId.
	folderPath := "/" + userId + "/" + petId

	// Call the DeleteFolder method on the Media interface of imageKit.
	// The DeleteFolder method requires a context and a DeleteFolderParam
	// containing the path of the folder to be deleted.
	resp, err := p.imageKit.Kit.Media.DeleteFolder(ctx, media.DeleteFolderParam{
		FolderPath: folderPath,
	})

	// Check if an error occurred during the folder deletion process.
	// If an error is returned, it indicates that the deletion failed,
	// and we return nil for the response and propagate the error upwards.
	if err != nil {
		return nil, err
	}

	// If no error occurred, the deletion was successful.
	// Return the response from imageKit, which contains details of the operation.
	return resp, nil
}

// DeletePetProfilePicture is responsible for deleting a pet's profile picture.
// It takes two parameters: a fileId for the image to be deleted, and a petId
// to identify the pet whose profile picture should be deleted.
// The function first deletes the image associated with the fileId using the
// DeletePet method. If this operation is successful, it then calls the
// DeletePetProfilePicture method on the petDao to delete the profile picture
// associated with the petId from the database. If either of these operations
// fail, the error is propagated upwards.
func (p *PetService) DeletePetProfilePicture(fileId string, petId string) error {
	// First, delete the image associated with the fileId.
	if fileId == "" {
		return errors.New("fileId is empty")
	}
	if _, err := p.DeletePetImage(fileId); err != nil {
		return err
	}

	// Then, delete the profile picture associated with the petId from the database.
	if err := p.petDao.DeletePetProfilePicture(petId); err != nil {
		return err
	}

	// If both operations were successful, return nil.
	return nil
}
