package utils

import (
	"context"
	"mime/multipart"

	"github.com/ScoobieNoobie/Caterpillar/config"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/admin"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

var CloudinaryClient *cloudinary.Cloudinary

func CreateCloudinaryClient() (err error) {
	CloudinaryClient, err = cloudinary.NewFromParams(config.ConfigData.CloudinaryCloud, config.ConfigData.CloudinaryPublic, config.ConfigData.CloudinaryPrivate)
	CloudinaryClient.Config.URL.Secure = true
	return
}

func UploadImage(file multipart.File) (string, error) {
	ctx := context.Background()
	res, err := CloudinaryClient.Upload.Upload(ctx, file, uploader.UploadParams{
		PublicID:       "some name",
		UniqueFilename: api.Bool(false),
		Overwrite:      api.Bool(true)})
	return res.URL, err
}

func GetImage(key string) (string, error) {
	ctx := context.Background()
	res, err := CloudinaryClient.Admin.Asset(ctx, admin.AssetParams{
		PublicID: key,
	})
	return res.URL, err
}
