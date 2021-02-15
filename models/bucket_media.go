package models

import (
	"github.com/gobuffalo/uuid"
	"github.com/pkg/errors"
	"github.com/websublime/barrel/storage/namespace"
)

type BucketMedia struct {
	ID       uuid.UUID `json:"id" db:"id" primary_id:"id"`
	BucketID uuid.UUID `json:"bucketId" db:"bucket_id"`
	MediaID  uuid.UUID `json:"mediaId" db:"media_id"`
}

func (BucketMedia) TableName() string {
	tableName := "bucket_media"

	if namespace.GetNamespace() != "" {
		return namespace.GetNamespace() + "." + tableName
	}

	return tableName
}

func NewBucketMedia(bucket uuid.UUID, media uuid.UUID) (*BucketMedia, error) {
	uid, err := uuid.NewV4()
	if err != nil {
		return nil, errors.Wrap(err, "Error generating unique id")
	}

	bucketMedia := &BucketMedia{
		ID:       uid,
		BucketID: bucket,
		MediaID:  media,
	}

	return bucketMedia, nil
}
