package ibmcloudauth

const (
	iamEndpointFieldDefault    = "https://iam.cloud.ibm.com"
	iamIdentityEndpointDefault = "https://iam.cloud.ibm.com/identity"
	userMgmtEndpointDefault    = "https://user-management.cloud.ibm.com/"
)

// IAM API paths
const (
	accessGroupMembershipCheck = "/v2/groups/%s/members/%s"
	serviceIDDetails           = "/v1/serviceids/%s"
	getUserProfile             = "/v2/accounts/%s/users/%s"
)

//Number of minutes to renew the admin token before expiration
const adminTokenRenewBeforeExpirationMinutes = 5

// request & response fields
const (
	accountIDField       = "account_id"
	identifierField      = "identifier"
	roleField            = "role"
	apiKeyField          = "api_key"
	tokenKeyField        = "token"
	redacted             = "<redacted>"
	iamIDField           = "IAM_ID"
	subjectField         = "subject"
	subjectTypeField     = "sub_type"
	serviceIDSubjectType = "ServiceId"
)
