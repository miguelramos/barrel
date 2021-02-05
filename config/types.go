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

type CannedPolicy struct {
	Name   string `json:"name"`
	Policy string `json:"policy"`
	Bucket string `json:"bucket"`
}

func (canned *CannedPolicy) Validate() *validate.Errors {
	return validate.Validate(
		&validators.StringIsPresent{Field: canned.Name, Name: "Name", Message: "Policy name is missing"},
		&validators.StringIsPresent{Field: canned.Policy, Name: "Policy", Message: "Policy type is missing"},
		&validators.StringIsPresent{Field: canned.Bucket, Name: "Bucket", Message: "Bucket is missing"},
	)
}

func (canned *CannedPolicy) ValidatePolicy() *validate.Errors {
	return validate.Validate(
		&validators.StringIsPresent{Field: canned.Policy, Name: "Policy", Message: "Policy type is missing"},
		&validators.StringIsPresent{Field: canned.Bucket, Name: "Bucket", Message: "Bucket is missing"},
	)
}

type IdentityPolicy struct {
	AccessKey  string `json:"accessKey"`
	PolicyName string `json:"policyName"`
}

func (ip *IdentityPolicy) Validate() *validate.Errors {
	return validate.Validate(
		&validators.StringIsPresent{Field: ip.AccessKey, Name: "AccessKey", Message: "Access key is missing"},
		&validators.StringIsPresent{Field: ip.PolicyName, Name: "PolicyName", Message: "Policy name is missing"},
	)
}

type MetaFile struct {
	FileName          string `json:"filename"`
	FileSize          string `json:"filesize"`
	FileType          string `json:"filetype"`
	FileTypeExtension string `json:"extension"`
	MIMEType          string `json:"mimetype"`
}
