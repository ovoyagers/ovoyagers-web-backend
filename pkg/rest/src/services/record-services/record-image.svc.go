package recordservices

import (
	"context"
	"errors"

	"github.com/imagekit-developer/imagekit-go/api/media"
	"github.com/imagekit-developer/imagekit-go/api/uploader"
	recordmodel "github.com/petmeds24/backend/pkg/rest/src/models/record-model"
	log "github.com/sirupsen/logrus"
)

var ctx = context.Background()

func (rs *RecordService) InsertRecordImages(imgsMetadata recordmodel.MedicalRecordImageMetadata) ([]*uploader.UploadResponse, error) {
	var responses []*uploader.UploadResponse
	for _, img := range imgsMetadata.Image {
		resp, err := rs.imageKit.Kit.Uploader.Upload(ctx, img.ImageBytes, uploader.UploadParam{
			FileName: img.FileName,
			Folder:   "/" + imgsMetadata.UserId + "/" + imgsMetadata.PetId + "/medical-records",
		})
		if err != nil {
			return nil, err
		}
		responses = append(responses, resp)
	}
	return responses, nil
}

/**
 * DeleteRecordImages deletes multiple images associated with a record.
 *
 * Parameters:
 * - imageIds: a slice of strings containing the IDs of the images to be deleted
 *
 * Returns:
 * - An error if the deletion process encounters any issues, otherwise returns nil
 */
func (rs *RecordService) DeleteRecordImages(imageIds []string, recordId string) error {
	if len(imageIds) == 0 {
		return errors.New("no image ids provided")
	}
	resp, err := rs.imageKit.Kit.Media.DeleteBulkFiles(ctx, media.FileIdsParam{
		FileIds: imageIds,
	})

	if err != nil {
		return err
	}

	err = rs.recordDao.DeleteMedicalRecordById(recordId)
	if err != nil {
		return err
	}

	log.Info("Deleted Images: ", resp)
	return nil
}
