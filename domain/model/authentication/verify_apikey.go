package authentication

type VeriryAPIKeyResponse struct {
	IsAPIKeyValid    bool `json:"isApiKeyValid"`
	IsIPAddressValid bool `json:"isIpAddressValid"`
}

type VeriryTokenResponse struct {
	OperatorID *string `json:"operatorId"`
}
