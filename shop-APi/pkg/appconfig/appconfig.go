package appconfig

type JwtToken struct {
	AccessSecret string `json:"accessSecret"`
	AccessExpire int64  `json:"accessExpire"`
}

type AppConfig struct {
	JwtToken `json:"jwtToken"`
}
