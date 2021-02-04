package utils

const (
	ErrorServerUnknown          = "ESERVER_UNKNOWN"
	ErrorBucketModel            = "EBUCKET_MODEL"
	ErrorBucketBodyParse        = "EBUCKET_BODYPARSE"
	ErrorBucketExist            = "EBUCKET_EXIST"
	ErrorBucketCreation         = "EBUCKET_CREATE"
	ErrorBucketModelSave        = "EBUCKET_MODELSAVE"
	ErrorBucketModelTransaction = "EBUCKET_MODELTRANSACTION"
	ErrorOrgStatusForbidden     = "EORG_FORBIDDEN"
	ErrorOrgInvalidToken        = "EORG_INVALIDTOKEN"
	ErrorOrgInvalidIdentity     = "EORG_INVALIDIDENTITY"
	ErrorOrgUserFailure         = "EORG_USERFAILURE"
	ErrorOrgPolicyCreate        = "EORG_POLICYCREATE"
	ErrorOrgPolicyUser          = "EORG_POLICYUSER"
	ErrorInvalidBody            = "ERESPONSE_INVALIDBODY"
	ErrorParseJson              = "EJSON_PARSE"
	ErrorResourceForbidden      = "ERESOURCE_FORBIDDEN"
	ErrorResourceInvalidBody    = "ERESOURCE_BODYINVALID"
)

// contains checks if a string is present in a slice
func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
