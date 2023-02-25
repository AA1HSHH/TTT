package config

const (
	MySQLDefaultDSN = "gorm:gorm@tcp(localhost:3306)/TTT?charset=utf8mb4&parseTime=True&loc=Local"
	JwtPrivateKey   = "config/jwtkey/sample_key"
	JwtPublicKey    = "config/jwtkey/sample_key.pub"
	DBDebug         = false
)
const (
	VideoCount = 10
)
const PlayUrlPrefix = "http://172.17.23.2:8080/static/"
const CoverUrlPrefix = "http://172.17.23.2:8080/static/"

// var Info Config

// type Path struct {
// 	FfmpegPath       string `toml:"ffmpeg_path"`
// 	StaticSourcePath string `toml:"static_source_path"`
// }

// type Config struct {
// 	Path `toml:"path"`
// }
