package helpers

import (
	"context"
	"time"

	"admin.server/config"
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

func UploadFileToCloudinary(file interface{}, folder string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cld, err := cloudinary.NewFromParams(config.EnvCloudName(), config.EnvCloudAPIKey(), config.EnvCloudAPISecret())
	if err != nil {
		return "", err
	}

	uploadParam, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{Folder: folder})

	if err != nil {
		return "", err
	}

	image := uploadParam.SecureURL

	return image, nil
}
