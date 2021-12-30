package cdn

import (
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/admin"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/cloudinary/cloudinary-go/config"
	"github.com/cloudinary/cloudinary-go/logger"
)

func NewFromConfiguration(configuration config.Configuration) (*Cloudinary, error) {
	logger := logger.New()

	return &Cloudinary{
		Config: configuration,
		Admin: admin.API{
			Config: configuration,
			Logger: logger,
		},
		Upload: uploader.API{
			Config: configuration,
			Logger: logger,
		},
		Logger: logger,
	}, nil
}

func CdnSetting() (*Cloudinary, error) {
	cld, err := cloudinary.NewFromParams("p3l-neopeople", "641612681972414", "Iq91j0N_CvezM-n9IMBo8YKfiIM")
	if err != nil {
		return nil, err
	}
	// fmt.Println("cld: ", cld)
	return NewFromConfiguration(cld.Config)
}
