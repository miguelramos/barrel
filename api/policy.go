package api

import (
	"fmt"
	"strings"

	"github.com/minio/minio/pkg/bucket/policy"
	"github.com/minio/minio/pkg/bucket/policy/condition"
	iampolicy "github.com/minio/minio/pkg/iam/policy"
	"github.com/pkg/errors"
)

//PolicyType types for policies
type PolicyType string

//Types of policies to apply when creating
const (
	AnonymousPolicy PolicyType = "ANONYMOUS_POLICY"
	ReaderPolicy    PolicyType = "READER_POLICY"
	WriterPolicy    PolicyType = "WRITER_POLICY"
	ManagerPolicy   PolicyType = "MANAGER_POLICY"
	OwnerPolicy     PolicyType = "OWNER_POLICY"
	AdminPolicy     PolicyType = "ADMIN_POLICY"
)

//GetPolicyType get defined type
func GetPolicyType(policy string) PolicyType {
	types := map[string]PolicyType{
		"ANONYMOUS_POLICY": AnonymousPolicy,
		"READER_POLICY":    ReaderPolicy,
		"WRITER_POLICY":    WriterPolicy,
		"MANAGER_POLICY":   ManagerPolicy,
		"OWNER_POLICY":     OwnerPolicy,
		"ADMIN_POLICY":     AdminPolicy,
	}

	return types[policy]
}

// CreatePolicy create one of the Policy types
func CreatePolicy(policy PolicyType, bucket string) (*iampolicy.Policy, string, error) {
	switch policy {
	case AnonymousPolicy:
		return CreateAnonymousPolicy(bucket)
	case ReaderPolicy:
		return CreateReaderPolicy(bucket)
	case WriterPolicy:
		return CreateWriterPolicy(bucket)
	case ManagerPolicy:
		return CreateManagerPolicy(bucket)
	case OwnerPolicy:
		return CreateOwnerPolicy(bucket)
	case AdminPolicy:
		return CreateAdminPolicy(bucket)
	}

	return nil, "", errors.New("Invalid policy type")
}

// CreateAnonymousPolicy read-only bucket objects
func CreateAnonymousPolicy(bucket string) (*iampolicy.Policy, string, error) {
	policyString := fmt.Sprintf(`{
		"Version": "2012-10-17",
		"Statement": [{
			"Action": [
				"s3:GetObject"
			],
			"Effect": "Allow",
			"Principal": {"AWS": ["*"]},
			"Resource": ["arn:aws:s3:::%s/*"],
			"Sid": ""
		}]
	}`, bucket)

	p := iampolicy.Policy{
		Version: iampolicy.DefaultVersion,
		Statements: []iampolicy.Statement{
			iampolicy.NewStatement(
				policy.Allow,
				iampolicy.NewActionSet(iampolicy.GetObjectAction),
				iampolicy.NewResourceSet(iampolicy.NewResource(fmt.Sprintf("%s/*", bucket), "")),
				condition.NewFunctions(),
			),
		},
	}

	fmt.Print(p.Validate())

	policy, err := iampolicy.ParseConfig(strings.NewReader(policyString))
	if err != nil {
		return nil, policyString, errors.Wrap(err, err.Error())
	}

	return policy, policyString, nil
}

// CreateReaderPolicy read objects, tagging and list
func CreateReaderPolicy(bucket string) (*iampolicy.Policy, string, error) {
	policyString := fmt.Sprintf(`{
		"Version": "2012-10-17",
		"Statement": [{
			"Action": [
				"s3:GetBucketLocation",
				"s3:GetObject",
				"s3:ListBucket",
				"s3:GetBucketTagging",
				"s3:GetObjectTagging"
			],
			"Effect": "Allow",
			"Principal": {"AWS": ["*"]},
			"Resource": ["arn:aws:s3:::%s/*"],
			"Sid": ""
		}]
	}`, bucket)

	policy, err := iampolicy.ParseConfig(strings.NewReader(policyString))
	if err != nil {
		return nil, policyString, errors.Wrap(err, "Unable to parse policy")
	}

	return policy, policyString, nil
}

// CreateWriterPolicy reads and write objects, tagging
func CreateWriterPolicy(bucket string) (*iampolicy.Policy, string, error) {
	policyString := fmt.Sprintf(`{
		"Version": "2012-10-17",
		"Statement": [{
			"Action": [
				"s3:GetBucketLocation",
				"s3:GetObject",
				"s3:ListBucket",
				"s3:GetBucketTagging",
				"s3:GetObjectTagging",
				"s3:PutObjectTagging",
				"s3:PutBucketTagging",
				"s3:PutObject"
			],
			"Effect": "Allow",
			"Principal": {"AWS": ["*"]},
			"Resource": ["arn:aws:s3:::%s/*"],
			"Sid": ""
		}]
	}`, bucket)

	policy, err := iampolicy.ParseConfig(strings.NewReader(policyString))
	if err != nil {
		return nil, policyString, errors.Wrap(err, "Unable to parse policy")
	}

	return policy, policyString, nil
}

// CreateManagerPolicy manage get,delete, put and notifications
func CreateManagerPolicy(bucket string) (*iampolicy.Policy, string, error) {
	policyString := fmt.Sprintf(`{
		"Version": "2012-10-17",
		"Statement": [{
			"Action": [
				"s3:GetBucketLocation",
				"s3:GetObject",
				"s3:ListBucket",
				"s3:GetBucketTagging",
				"s3:GetObjectTagging",
				"s3:PutObjectTagging",
				"s3:PutBucketTagging",
				"s3:PutObject",
				"s3:DeleteObject",
				"s3:DeleteObjectTagging",
				"s3:GetBucketVersioning",
				"s3:ListenNotification"
			],
			"Effect": "Allow",
			"Principal": {"AWS": ["*"]},
			"Resource": ["arn:aws:s3:::%s/*"],
			"Sid": ""
		}]
	}`, bucket)

	policy, err := iampolicy.ParseConfig(strings.NewReader(policyString))
	if err != nil {
		return nil, policyString, errors.Wrap(err, "Unable to parse policy")
	}

	return policy, policyString, nil
}

// CreateOwnerPolicy manage bucket
func CreateOwnerPolicy(bucket string) (*iampolicy.Policy, string, error) {
	policyString := fmt.Sprintf(`{
		"Version": "2012-10-17",
		"Statement": [{
			"Action": [
				"s3:CreateBucket",
				"s3:DeleteBucket",
				"s3:ForceDeleteBucket",
				"s3:GetBucketLocation",
				"s3:GetObject",
				"s3:ListBucket",
				"s3:ListAllMyBuckets",
				"s3:GetBucketTagging",
				"s3:GetObjectTagging",
				"s3:PutObjectTagging",
				"s3:PutBucketTagging",
				"s3:PutObject",
				"s3:DeleteObject",
				"s3:DeleteObjectTagging",
				"s3:GetBucketVersioning",
				"s3:ListenNotification"
			],
			"Effect": "Allow",
			"Principal": {"AWS": ["*"]},
			"Resource": ["arn:aws:s3:::%s/*"],
			"Sid": ""
		}]
	}`, bucket)

	policy, err := iampolicy.ParseConfig(strings.NewReader(policyString))
	if err != nil {
		return nil, policyString, errors.Wrap(err, "Unable to parse policy")
	}

	return policy, policyString, nil
}

// CreateAdminPolicy all s3:+ roles
func CreateAdminPolicy(bucket string) (*iampolicy.Policy, string, error) {
	policyString := fmt.Sprintf(`{
		"Version": "2012-10-17",
		"Statement": [{
			"Action": [
				"s3:*"
			],
			"Effect": "Allow",
			"Principal": {"AWS": ["*"]},
			"Resource": ["arn:aws:s3:::%s/*"],
			"Sid": ""
		}]
	}`, bucket)

	policy, err := iampolicy.ParseConfig(strings.NewReader(policyString))
	if err != nil {
		return nil, policyString, errors.Wrap(err, "Unable to parse policy")
	}

	return policy, policyString, nil
}
