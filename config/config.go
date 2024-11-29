package config

import (
	"errors"
	"os"
	"strconv"

	"data-spaces-backend/extension/logger"
)

// Config
// Summary: This is structure which defines Config
type Config struct {
	Env    string
	Server struct {
		Port                  string
		RedirectURLAfterLogin string
		Host                  string
	}
	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		Database string
		Sslmode  string
	}
	GoogleAuth struct {
		RedirectURL string
	}
	LogLevel               string
	ZapLogLevel            string
	GoogleProjectID        string
	IsTraceabilityAccess   bool
	TraceabilityBaseURL    string
	TraceabilityAPIVersion string
	TraceabilityAPIKey     string
	AuthenticaterURL       string
	DataSpaceApikey        string
	LocalServerIPAddress   string
}

var (
	ErrEnvNotDefined    = errors.New("GO_ENV not defined")
	ErrReadConfigFile   = errors.New("config file read error")
	ErrConfigFileFormat = errors.New("config file formant error")
)

// NewConfig
// Summary: This is function which is used to get the configuration from environment variables
// output: (*Config) pointer of Config struct
// output: (error) error object
func NewConfig() (*Config, error) {
	var cfg Config

	var err error

	current := &cfg

	current.Env = os.Getenv("GO_ENV")
	current.Server.Port = os.Getenv("SERVER_PORT")
	current.Server.RedirectURLAfterLogin = os.Getenv("SERVER_REDIRECT_URL_AFTER_LOGIN")
	current.Server.Host = os.Getenv("SERVER_HOST")

	current.Database.Host = os.Getenv("DB_HOST")
	current.Database.Port = os.Getenv("DB_PORT")
	current.Database.User = os.Getenv("DB_USER")
	current.Database.Password = os.Getenv("DB_PASSWORD")
	current.Database.Database = os.Getenv("DB_DATABASE")
	current.Database.Sslmode = os.Getenv("DB_SSLMODE")

	current.GoogleAuth.RedirectURL = os.Getenv("GOOGLE_REDIRECT_URL")

	current.LogLevel = os.Getenv("ECHO_LOG_LEVEL")
	current.ZapLogLevel = os.Getenv("ZAP_LOG_LEVEL")

	cfg.GoogleProjectID = os.Getenv("GOOGLE_PROJECT_ID")

	if current.IsTraceabilityAccess, err = strconv.ParseBool(os.Getenv("IS_TRACEABILITY_ACCESS")); err != nil {
		logger.Set(nil).Errorf(err.Error())

		return nil, ErrReadConfigFile
	}
	current.TraceabilityBaseURL = os.Getenv("TRACEABILITY_BASE_URL")
	current.TraceabilityAPIVersion = os.Getenv("TRACEABILITY_API_VERSION")
	current.TraceabilityAPIKey = os.Getenv("TRACEABILITY_API_KEY")

	current.AuthenticaterURL = os.Getenv("AUTHENTICATER_URL")

	current.DataSpaceApikey = os.Getenv("DATA_SPACE_APIKEY")

	current.LocalServerIPAddress = os.Getenv("LOCAL_SERVER_IP_ADDRESS")
	return current, nil
}
