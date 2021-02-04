package config

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
)

type MetaClaimerRoles struct {
	AllowedRoles []string `json:"allowed_roles"`
	DefaultRule  string   `json:"default_rule"`
	UserID       string   `json:"user_id"`
}

type MetaClaimer struct {
	Claims MetaClaimerRoles `json:"claims"`
}

type UserMetaClaim struct {
	Meta    MetaClaimer            `json:"meta"`
	Profile map[string]interface{} `json:"profile"`
}

type GoTrueClaims struct {
	jwt.StandardClaims
	Email        string                 `json:"email"`
	AppMetaData  map[string]interface{} `json:"app_metadata"`
	UserMetaData UserMetaClaim          `json:"user_metadata"`
}

type Identity struct {
	ID        uuid.UUID `json:"id,omitempty"`
	UserID    uuid.UUID `json:"userID,omitempty"`
	AccessKey string    `json:"accessKey,omitempty"`
	SecretKey string    `json:"secretKey,omitempty"`
	Token     string    `json:"token,omitempty"`
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
