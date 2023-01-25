package config

const (
	MySQLDefaultDSN = "gorm:gorm@tcp(localhost:3306)/TTT?charset=utf8mb4&parseTime=True&loc=Local"
	JwtPrivateKey   = "config/jwtkey/sample_key"
	JwtPublicKey    = "config/jwtkey/sample_key.pub"
)
