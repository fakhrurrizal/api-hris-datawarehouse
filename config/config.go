package config

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName                     string
	AppKey                      string
	BaseUrl                     string
	Environtment                string
	DatabaseUsername            string
	DatabasePassword            string
	DatabaseHost                string
	DatabasePort                string
	DatabaseName                string
	DatabasePlannerName         string
	PathDB                      string
	CacheURL                    string
	CachePassword               string
	LoggerLevel                 string
	ContextTimeout              int
	Port                        string
	GoogleClientID              string
	GoogleClientSecret          string
	EnableCSRF                  bool
	EnableEncodeID              bool
	EnableDatabaseAutomigration bool
	EnableAPIKey                bool
	APIGeolocationAPIKey        string
	RunLocalDatabaseVia         string
	OnBehalfOf                  string
	AccountNumber               string
	WhatsappContact             string
	SpecialApiKey               string
	ServerAPI                   string
	EmailVerificationUrl        string
	EmailVerificationApiKey     string
	EnableIDDuplicationHandling bool
	EnableEmailVerification     bool
	FrontEndUrl                 string
	ContactEmail                string
	MailMailer                  string
	MailHost                    string
	MailPort                    int
	MailUsername                string
	MailPassword                string
		APIKey                      string
	MailEncryption              string
}

func LoadConfig() (config *Config) {

	if err := godotenv.Load(RootPath() + `/.env`); err != nil {
		fmt.Println(err)
	}

	appName := os.Getenv("APP_NAME")
	appKey := os.Getenv("APP_KEY")
	baseurl := os.Getenv("BASE_URL")
	environment := strings.ToUpper(os.Getenv("ENVIRONMENT"))
	contactEmail := strings.ToUpper(os.Getenv("CONTACT_EMAIL"))
	databaseUsername := os.Getenv("DATABASE_USERNAME")
	databasePassword := os.Getenv("DATABASE_PASSWORD")
	databaseHost := os.Getenv("DATABASE_HOST")
	databasePort := os.Getenv("DATABASE_PORT")
	databaseName := os.Getenv("DATABASE_NAME")
	databasePlannerName := os.Getenv("DATABASE_PLANNER_NAME")
	PathDB := os.Getenv("PATH_DB")
	cacheURL := os.Getenv("CACHE_URL")
	cachePassword := os.Getenv("CACHE_PASSWORD")
	loggerLevel := os.Getenv("LOGGER_LEVEL")
	contextTimeout, _ := strconv.Atoi(os.Getenv("CONTEXT_TIMEOUT"))
	port := os.Getenv("PORT")
	googleClientID := os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	enableCSRF, _ := strconv.ParseBool(os.Getenv("ENABLE_CSRF"))
	enableDatabaseAutomigration, _ := strconv.ParseBool(os.Getenv("ENABLE_DATABASE_AUTOMIGRATION"))
	enableApiKey, _ := strconv.ParseBool(os.Getenv("ENABLE_API_KEY"))
	enableEncodeID, _ := strconv.ParseBool(os.Getenv("ENABLE_ENCODE_ID"))
	apiGeolocationAPIKey := os.Getenv("APIGEOLOCATION_API_KEY")
	runLocalDatabaseVia := strings.ToUpper(os.Getenv("RUN_LOCAL_DATABASE_VIA"))
	specialApiKey := os.Getenv("SPECIAL_API_KEY")
	emailVerificationUrl := os.Getenv("EMAIL_VERIFICATION_URL")
	emailVerificationApiKey := os.Getenv("EMAIL_VERIFICATION_API_KEY")
	enableEmailVerification, _ := strconv.ParseBool(os.Getenv("ENABLE_EMAIL_VERIFICATION"))
	accountNumber := os.Getenv("ACCOUNT_NUMBER")
	whatsappContact := os.Getenv("WHATSAPP_CONTACT")
	serverAPI := os.Getenv("SERVER_API")
	frontendurl := os.Getenv("FRONT_END_URL")
	mailMailer := os.Getenv("MAIL_MAILER")
	mailHost := os.Getenv("MAIL_HOST")
	mailPort, _ := strconv.Atoi(os.Getenv("MAIL_PORT"))
	mailUsername := os.Getenv("MAIL_USERNAME")
	mailPassword := os.Getenv("MAIL_PASSWORD")
	mailEncryption := os.Getenv("MAIL_ENCRYPTION")
	enableIDDuplicationHandling, _ := strconv.ParseBool(os.Getenv("ENABLE_ID_UPLICATION_HANDLING"))

	return &Config{
		AppName:                     appName,
		AppKey:                      appKey,
		BaseUrl:                     baseurl,
		Environtment:                environment,
		PathDB:                      PathDB,
		FrontEndUrl:                 frontendurl,
		Port:                        port,
		EnableDatabaseAutomigration: enableDatabaseAutomigration,
		RunLocalDatabaseVia:         runLocalDatabaseVia,
		EnableAPIKey:                enableApiKey,
		ContactEmail:                contactEmail,
		AccountNumber:               accountNumber,
		WhatsappContact:             whatsappContact,
		EnableIDDuplicationHandling: enableIDDuplicationHandling,
		SpecialApiKey:               specialApiKey,
		ServerAPI:                   serverAPI,
		DatabaseUsername:            databaseUsername,
		DatabasePassword:            databasePassword,
		EnableEncodeID:              enableEncodeID,
		DatabaseHost:                databaseHost,
		DatabasePort:                databasePort,
		DatabaseName:                databaseName,
		DatabasePlannerName:         databasePlannerName,
		CacheURL:                    cacheURL,
		CachePassword:               cachePassword,
		LoggerLevel:                 loggerLevel,
		ContextTimeout:              contextTimeout,
		GoogleClientID:              googleClientID,
		GoogleClientSecret:          googleClientSecret,
		EnableCSRF:                  enableCSRF,
		APIGeolocationAPIKey:        apiGeolocationAPIKey,
		EmailVerificationUrl:        emailVerificationUrl,
		EmailVerificationApiKey:     emailVerificationApiKey,
		EnableEmailVerification:     enableEmailVerification,
		MailMailer:                  mailMailer,
		MailHost:                    mailHost,
		MailUsername:                mailUsername,
		MailPassword:                mailPassword,
		MailPort:                    mailPort,
		MailEncryption:              mailEncryption,
	}
}

func RootPath() string {
	projectDirName := os.Getenv("DIR_NAME")
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))
	return string(rootPath)
}
