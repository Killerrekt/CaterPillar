package utils

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/ScoobieNoobie/Caterpillar/config"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

var CloudinaryClient *cloudinary.Cloudinary

func CreateCloudinaryClient() (err error) {
	CloudinaryClient, err = cloudinary.NewFromParams(config.ConfigData.CloudinaryCloud, config.ConfigData.CloudinaryPublic, config.ConfigData.CloudinaryPrivate)
	CloudinaryClient.Config.URL.Secure = true
	return
}

func UploadImage(file *multipart.FileHeader, key string) (string, error) {
	ctx := context.Background()
	res, err := CloudinaryClient.Upload.Upload(ctx, file, uploader.UploadParams{
		PublicID:       key,
		UniqueFilename: api.Bool(false),
		Overwrite:      api.Bool(true)})
	return res.URL, err
}

func GetImage(key string) (string, error) {
	res, err := CloudinaryClient.Image(key)
	if err != nil {
		return "", err
	}
	url, err := res.String()
	fmt.Println("url : ", url)
	return url, err
}
