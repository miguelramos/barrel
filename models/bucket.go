package models

import (
	"database/sql"
	"time"

	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/uuid"
	"github.com/gosimple/slug"
	"github.com/pkg/errors"
	"github.com/websublime/barrel/storage"
	"github.com/websublime/barrel/storage/namespace"
)

// Bucket model type
type Bucket struct {
	ID        uuid.UUID    `json:"id" db:"id"`
	Name      string       `json:"name" db:"name"`
	Bucket    nulls.String `json:"path,omitempty" db:"bucket"`
	OrgID     nulls.String `json:"orgId,omitempty" db:"org_id"`
	IsPrivate bool         `json:"isPrivate" db:"is_private"`
	Policy    string       `json:"policy" db:"-"`
	CreatedAt time.Time    `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time    `json:"updatedAt" db:"updated_at"`
	DeletedAt nulls.Time   `json:"deleteAt,omitempty" db:"deleted_at"`
}

// NewBucket creates new Bucket
func NewBucket(name string, private bool) (*Bucket, error) {
	uid, err := uuid.NewV4()
	if err != nil {
		return nil, errors.Wrap(err, "Error generating unique id")
	}

	bucket := &Bucket{
		ID:        uid,
		Name:      name,
		IsPrivate: private,
	}

	return bucket, nil
}

func (Bucket) TableName() string {
	tableName := "buckets"

	if namespace.GetNamespace() != "" {
		return namespace.GetNamespace() + "." + tableName
	}

	return tableName
}

func (b *Bucket) BeforeSave(tx *storage.Connection) error {
	if !b.Bucket.Valid {
		b.Bucket = nulls.NewString(slug.Make(b.Name))
	}

	return nil
}

func FindBucket(tx *storage.Connection, bucket string) (*Bucket, error) {
	buck := &Bucket{}
	if err := tx.Where("bucket = ?", bucket).First(buck); err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, errors.Wrap(err, "Bucket not found")
		}

		return nil, errors.Wrap(err, err.Error())
	}

	return buck, nil
}
