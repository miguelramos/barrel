package api

import (
	"fmt"
	"strings"

	iampolicy "github.com/minio/minio/pkg/iam/policy"
	"github.com/pkg/errors"
)

// CreateAnonymousPolicy read-only bucket objects
func CreateAnonymousPolicy(bucket string) (*iampolicy.Policy, error) {
	policyString := fmt.Sprintf(`{
		"Version": "2012-10-17",
		"Statement": [{
			"Action": [
				"s3:GetBucketLocation",
				"s3:GetObject"
			],
			"Resource": ["%s/*"],
			"Sid": ""
		}]
	}`, bucket)

	policy, err := iampolicy.ParseConfig(strings.NewReader(policyString))
	if err != nil {
		return nil, errors.Wrap(err, "Unable to parse policy")
	}

	return policy, nil
}

// CreateReaderPolicy read objects, tagging and list
func CreateReaderPolicy(bucket string) (*iampolicy.Policy, error) {
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
			"Resource": ["%s/*"],
			"Sid": ""
		}]
	}`, bucket)

	policy, err := iampolicy.ParseConfig(strings.NewReader(policyString))
	if err != nil {
		return nil, errors.Wrap(err, "Unable to parse policy")
	}

	return policy, nil
}

// CreateWriterPolicy reads and write objects, tagging
func CreateWriterPolicy(bucket string) (*iampolicy.Policy, error) {
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
			"Resource": ["%s/*"],
			"Sid": ""
		}]
	}`, bucket)

	policy, err := iampolicy.ParseConfig(strings.NewReader(policyString))
	if err != nil {
		return nil, errors.Wrap(err, "Unable to parse policy")
	}

	return policy, nil
}

// CreateManagerPolicy manage get,delete, put and notifications
func CreateManagerPolicy(bucket string) (*iampolicy.Policy, error) {
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
				"s3:GetBucketVersioning"
				"s3:ListenNotification"
			],
			"Resource": ["%s/*"],
			"Sid": ""
		}]
	}`, bucket)

	policy, err := iampolicy.ParseConfig(strings.NewReader(policyString))
	if err != nil {
		return nil, errors.Wrap(err, "Unable to parse policy")
	}

	return policy, nil
}

// CreateOwnerPolicy manage bucket
func CreateOwnerPolicy(bucket string) (*iampolicy.Policy, error) {
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
				"s3:GetBucketVersioning"
				"s3:ListenNotification"
			],
			"Resource": ["%s/*"],
			"Sid": ""
		}]
	}`, bucket)

	policy, err := iampolicy.ParseConfig(strings.NewReader(policyString))
	if err != nil {
		return nil, errors.Wrap(err, "Unable to parse policy")
	}

	return policy, nil
}

// CreateAdminPolicy all s3:+ roles
func CreateAdminPolicy(bucket string) (*iampolicy.Policy, error) {
	policyString := fmt.Sprintf(`{
		"Version": "2012-10-17",
		"Statement": [{
			"Action": [
				"s3:*"
			],
			"Resource": ["%s/*"],
			"Sid": ""
		}]
	}`, bucket)

	policy, err := iampolicy.ParseConfig(strings.NewReader(policyString))
	if err != nil {
		return nil, errors.Wrap(err, "Unable to parse policy")
	}

	return policy, nil
}
