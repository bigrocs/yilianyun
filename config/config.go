package config

type Config struct {
	ClientId     string `json:"client_id"`     // 开发者ID
	ClientSecret string `json:"client_secret"` // 开发者密钥
	AccessToken  string `json:"access_token"`  // 产品KEY
	Sandbox      bool   `json:"sandbox"`       // 沙盒
}
