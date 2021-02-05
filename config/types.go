package config

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofrs/uuid"
)

type GoTrueClaims struct {
	jwt.StandardClaims
	Email        string                 `json:"email"`
	AppMetaData  map[string]interface{} `json:"app_metadata"`
	UserMetaData map[string]interface{} `json:"user_metadata"`
}

type Identity struct {
	AccessKey string `json:"accessKey,omitempty"`
	SecretKey string `json:"secretKey,omitempty"`
}

func (identity *Identity) Validate() *validate.Errors {
	return validate.Validate(
		&validators.StringIsPresent{Field: identity.AccessKey, Name: "AccessKey", Message: "Identity accessKey is missing"},
		&validators.StringIsPresent{Field: identity.SecretKey, Name: "SecretKey", Message: "Identity secretKey is missing"},
	)
}

type GoTrueIdentity struct {
	Identity
	ID     uuid.UUID `json:"id,omitempty"`
	UserID uuid.UUID `json:"userID,omitempty"`
	Token  string    `json:"token,omitempty"`
}

type BucketPolicy struct {
	Bucket string `json:"bucket"`
}

type MetaFile struct {
	FileName          string `json:"filename"`
	FileSize          string `json:"filesize"`
	FileType          string `json:"filetype"`
	FileTypeExtension string `json:"extension"`
	MIMEType          string `json:"mimetype"`
}
