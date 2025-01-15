package imageservice

import (
	"github.com/imagekit-developer/imagekit-go"
	"github.com/petmeds24/backend/config"
	"github.com/petmeds24/backend/pkg/rest/src/daos/handlers/userdao"
)

type ImageKit struct {
	imgKit  *imagekit.ImageKit
	userDao *userdao.UserDao
}

type ImgKit struct {
	Kit *imagekit.ImageKit
}

func NewImageKit(globalCfg *config.GlobalConfig) *ImageKit {
	cfg := globalCfg.GetConfig()
	ik := imagekit.NewFromParams(imagekit.NewParams{
		PrivateKey:  cfg.IMAGEKIT_PRIVATE_KEY,
		PublicKey:   cfg.IMAGEKIT_PUBLIC_KEY,
		UrlEndpoint: cfg.IMAGEKIT_URL_ENDPOINT,
	})
	return &ImageKit{
		imgKit:  ik,
		userDao: userdao.NewAuthDao(globalCfg),
	}
}

func NewImgKit(globalCfg *config.GlobalConfig) *ImgKit {
	cfg := globalCfg.GetConfig()
	ik := imagekit.NewFromParams(imagekit.NewParams{
		PrivateKey:  cfg.IMAGEKIT_PRIVATE_KEY,
		PublicKey:   cfg.IMAGEKIT_PUBLIC_KEY,
		UrlEndpoint: cfg.IMAGEKIT_URL_ENDPOINT,
	})
	return &ImgKit{
		Kit: ik,
	}
}
