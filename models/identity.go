package models

import (
	"database/sql"

	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/pkg/errors"
	"github.com/websublime/barrel/storage"
	"github.com/websublime/barrel/storage/namespace"
)

type Identity struct {
	ID        uuid.UUID `json:"id" db:"id"`
	AccessKey string    `json:"accessKey" db:"key"`
	SecretKey string    `json:"secretKey" db:"secret"`
	IsAdmin   bool      `json:"isAdmin" db:"is_admin"`
}

func (Identity) TableName() string {
	tableName := "identities"

	if namespace.GetNamespace() != "" {
		return namespace.GetNamespace() + "." + tableName
	}

	return tableName
}

func NewIdentity(secret string, key string, isAdmin bool) (*Identity, error) {
	uid, err := uuid.NewV4()
	if err != nil {
		return nil, errors.Wrap(err, "Error generating unique id")
	}

	user := &Identity{
		ID:        uid,
		SecretKey: secret,
		AccessKey: key,
		IsAdmin:   isAdmin,
	}

	return user, nil
}

func (u *Identity) Validate() *validate.Errors {
	return validate.Validate(
		&validators.StringIsPresent{Field: u.AccessKey, Name: "AccessKey", Message: "Secret is missing"},
		&validators.StringIsPresent{Field: u.SecretKey, Name: "SecretKey", Message: "Key is missing"},
	)
}

func FindIdentityByKey(tx *storage.Connection, key string) (*Identity, error) {
	identity := &Identity{}
	if err := tx.Where("key = ?", key).First(identity); err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, errors.Wrap(err, "Identity not found")
		}

		return nil, errors.Wrap(err, err.Error())
	}

	return identity, nil
}
