package output

type VerifyAPIKey struct {
	IsAPIKeyValid    bool
	IsIPAddressValid bool
}

type VerifyToken struct {
	OperatorID *string
}
