package env


type SecretKey struct {
	ForToken string `json:"for_token"`
	ForCoder string `json:"for_coder"`
	ForHasher string `json:"for_hasher"`
}
