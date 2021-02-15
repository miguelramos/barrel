package models

import (
	"time"

	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/pkg/errors"
	"github.com/websublime/barrel/storage/namespace"
)

type Media struct {
	ID          uuid.UUID      `json:"id" db:"id"`
	URL         string         `json:"url" db:"url"`
	Owner       nulls.UUID     `json:"ownerId" db:"owner"`
	BucketFile  nulls.String   `json:"bucketFile" db:"bucket_file"`
	Metafile    nulls.String   `json:"metadata" db:"meta_file"`
	CreatedAt   time.Time      `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time      `json:"updatedAt" db:"updated_at"`
	DeletedAt   nulls.Time     `json:"deletedAt" db:"deleted_at"`
	BucketMedia []*BucketMedia `json:"medias,omitempty" many_to_many:"bucket_media" db:"-" fk_id:"media_id" primary_id:"id"`
}

func (Media) TableName() string {
	tableName := "medias"

	if namespace.GetNamespace() != "" {
		return namespace.GetNamespace() + "." + tableName
	}

	return tableName
}

func NewMedia(url string, owner nulls.UUID, bucketFile nulls.String, metadata nulls.String) (*Media, error) {
	uid, err := uuid.NewV4()
	if err != nil {
		return nil, errors.Wrap(err, "Error generating unique id")
	}

	media := &Media{
		ID:         uid,
		URL:        url,
		Owner:      owner,
		BucketFile: bucketFile,
		Metafile:   metadata,
	}

	return media, nil
}

func (m *Media) Validate() *validate.Errors {
	return validate.Validate(
		&validators.StringIsPresent{Field: m.URL, Name: "URL", Message: "Url is missign"},
	)
}
