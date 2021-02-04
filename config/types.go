package config

import (
	"github.com/dgrijalva/jwt-go"
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
