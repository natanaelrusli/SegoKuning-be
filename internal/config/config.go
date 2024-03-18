package config

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	App        AppConfig
	HttpServer HttpServerConfig
	Database   DatabaseConfig
	Jwt        JwtConfig
}

type AppConfig struct {
	Environment  string
	BCryptCost   int
	WalletPrefix string
}

type HttpServerConfig struct {
	Host              string
	Port              int
	GracePeriod       int
	MaxUploadFileSize int64
}

type DatabaseConfig struct {
	Port                  int
	Host                  string
	DbName                string
	Username              string
	Password              string
	Sslmode               string
	MaxIdleConn           int
	MaxOpenConn           int
	MaxConnLifetimeMinute int
}

type JwtConfig struct {
	Issuer        string
	AllowedAlgs   []string
	TokenDuration int
	SecretKey     string
}

func GetEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func InitConfig() *Config {
	GetEnv()

	return &Config{
		App:        initAppConfig(),
		Database:   initDbConfig(),
		HttpServer: initHttpServerConfig(),
		Jwt:        initJwtConfig(),
	}
}

func initJwtConfig() JwtConfig {
	issuer := os.Getenv("JWT_ISSUER")
	secretKey := os.Getenv("JWT_SECRET_KEY")
	allowedAlgs := strings.Split(os.Getenv("JWT_ALLOWED_ALGS"), ",")

	duration, err := strconv.ParseInt(os.Getenv("JWT_DURATION"), 10, 32)
	if err != nil {
		log.Fatal("cannot parse JWT_DURATION")
	}

	return JwtConfig{
		Issuer:        issuer,
		AllowedAlgs:   allowedAlgs,
		TokenDuration: int(duration),
		SecretKey:     secretKey,
	}
}

func initDbConfig() DatabaseConfig {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	sslMode := os.Getenv("DB_SSL_MODE")

	port, err := strconv.ParseInt(os.Getenv("DB_PORT"), 10, 32)
	if err != nil {
		log.Fatal("cannot parse DB_PORT")
	}

	maxIdleConn, err := strconv.ParseInt(os.Getenv("DB_MAX_IDLE_CONN"), 10, 32)
	if err != nil {
		log.Fatal("cannot parse DB_MAX_IDLE_CONN")
	}

	maxOpenConn, err := strconv.ParseInt(os.Getenv("DB_MAX_OPEN_CONN"), 10, 32)
	if err != nil {
		log.Fatal("cannot parse DB_MAX_OPEN_CONN")
	}

	connMaxLifetime, err := strconv.ParseInt(os.Getenv("DB_CONN_MAX_LIFETIME"), 10, 32)
	if err != nil {
		log.Fatal("cannot parse DB_CONN_MAX_LIFETIME")
	}

	return DatabaseConfig{
		Port:                  int(port),
		Host:                  host,
		DbName:                name,
		Username:              user,
		Password:              password,
		Sslmode:               sslMode,
		MaxIdleConn:           int(maxIdleConn),
		MaxOpenConn:           int(maxOpenConn),
		MaxConnLifetimeMinute: int(connMaxLifetime),
	}
}

func initHttpServerConfig() HttpServerConfig {
	host := os.Getenv("HTTP_SERVER_HOST")

	port, err := strconv.ParseInt(os.Getenv("HTTP_SERVER_PORT"), 10, 32)
	if err != nil {
		log.Fatal("cannot parse HTTP_SERVER_PORT")
	}

	gracePeriod, err := strconv.ParseInt(os.Getenv("HTTP_SERVER_GRACE_PERIOD"), 10, 32)
	if err != nil {
		log.Fatal("cannot parse HTTP_SERVER_GRACE_PERIOD")
	}

	maxUploadFileSizeKb, err := strconv.ParseInt(os.Getenv("HTTP_MAX_UPLOAD_FILE_SIZE_KB"), 10, 32)
	if err != nil {
		log.Fatal("cannot parse HTTP_SERVER_GRACE_PERIOD")
	}

	return HttpServerConfig{
		Host:              host,
		Port:              int(port),
		GracePeriod:       int(gracePeriod),
		MaxUploadFileSize: maxUploadFileSizeKb * 1024,
	}
}

func initAppConfig() AppConfig {
	environment := os.Getenv("APP_ENVIRONMENT")

	bcryptCost, err := strconv.ParseInt(os.Getenv("APP_BCRYPT_COST"), 10, 32)
	if err != nil {
		log.Fatal("cannot parse APP_BCRYPT_COST")
	}

	walletPrefix := os.Getenv("APP_WALLET_PREFIX")

	return AppConfig{
		Environment:  environment,
		BCryptCost:   int(bcryptCost),
		WalletPrefix: walletPrefix,
	}
}
