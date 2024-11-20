package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/joho/godotenv"
)

var Map map[string]string

type AccountsConfig struct {
	AppCode          string `env:"APP_CODE"`
	AppVersion       string `env:"APP_VERSION"`
	AppPort          string `env:"APP_PORT"`
	LocalRedisPort   string `env:"LOCAL_REDIS_PORT"`
	LocalRedisHost   string `env:"LOCAL_REDIS_HOST"`
	LocalRedisPass   string `env:"LOCAL_REDIS_PASS"`
	CentralRedisPort string `env:"CENTRAL_REDIS_PORT"`
	CentralRedisHost string `env:"CENTRAL_REDIS_HOST"`
	CentralRedisPass string `env:"CENTRAL_REDIS_PASS"`
	TimeZone         string `env:"TIMEZONE"`
	AWSBucket        string `env:"AWS_BUCKET"`
	UserApiHealth    string `env:"USER_API_HEALTH"`
	AppEnv           string `env:"APP_ENV"`
	LogPath          string `env:"LOG_PATH"`
	LogFormat        string `env:"LOG_FORMAT"`

	MongoDBURI                 string `env:"MONGODB_URI"`
	MongoDBName                string `env:"MONGODB_DB_NAME"`
	MongoUserCollection        string `env:"MONGODB_USER_COLLECTION"`
	MongoInviteCollection      string `env:"MONGODB_INVITE_COLLECTION"`
	MongoApplicationCollection string `env:"MONGODB_APPLICATION_COLLECTION"`
	MongoUsername              string `env:"MONGODB_USERNAME"`
	MongoPassword              string `env:"MONGODB_PASSWORD"`
	MongoTimeout               string `env:"MONGODB_TIMEOUT"`
	MongoDBInstanceLocation    string `env:"MONGODB_INSTANCE_LOCATION"`
	MongoDBURISrv              string `env:"MONGODB_URI_SRV"`
	MongoClientCollection      string `env:"MONGODB_CLIENT_COLLECTION"`
	MongoAuthCodeCollection    string `env:"MONGODB_AUTH_CODE_COLLECTION"`
	MongoEcdsaKeys             string `env:"MONGODB_ECDSA_KEYS"`

	// ClientId     string `env:"CLIENT_ID"`
	// ClientSecret string `env:"CLIENT_SECRET"`

	// PublicKey  string `env:"PUBLIC_KEY"`
	// PrivateKey string `env:"PRIVATE_KEY"`

	KeyId     string `env:"KEY_ID"`
	KeySecret string `env:"KEY_SECRET"`

	Dns string `env:"DNS"`
}

func InitConfig() (AccountsConfig, error) {
	config := AccountsConfig{}
	// for local development load from `local`
	// for docker development load from `file`
	Load("file")
	ParseToStruct(&config)
	fmt.Printf("Config: %+v\n", config)
	return config, nil
}

func Load(source string) {
	wd := GetRoot()
	var env map[string]string

	if source == "file" {
		//load .env file
		env, err := godotenv.Read(wd + "/.env")
		if err != nil {
			log.Fatalf("Error loading " + wd + "/.env file")
		}
		Map = env
		log.Println(".env loaded")
		return
	}

	if source == "local" {
		//load .env file
		env, err := godotenv.Read("../build/.env")
		if err != nil {
			log.Fatalf("Error loading /.env file : %v", err)
		}
		Map = env
		log.Println(".env loaded")
		return
	}

	// load OS variables
	variables := os.Environ()
	env = make(map[string]string)
	for _, variable := range variables {
		// split by equal sign
		keyVal := strings.Split(variable, "=")
		env[keyVal[0]] = keyVal[1]
	}
	Map = env
	log.Println("OS variables loaded")
}

func GetRoot() string {
	ex, err := os.Executable()
	if err != nil {
		log.Println(err)
	}

	wd := filepath.Dir(ex)
	log.Println("Working Dir: ", wd)

	return wd
}

// ParseToStruct parses the environment variables from the Map into
// the fields of the provided struct.
func ParseToStruct[T any](obj *T) {
	// set environment variables before parsing
	for key, value := range Map {
		os.Setenv(key, value)
	}

	objValue := reflect.ValueOf(obj).Elem()
	objType := objValue.Type()

	for i := 0; i < objValue.NumField(); i++ {
		field := objType.Field(i)
		envVar, exists := Map[field.Tag.Get("env")]
		if exists {
			fieldValue := objValue.Field(i)
			if fieldValue.CanSet() {
				fieldValue.SetString(envVar)
			}
		}
	}
}
