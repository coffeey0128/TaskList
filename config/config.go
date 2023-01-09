package config

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

var ()

func GetBasePath() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}

func InitEnv() {
	remoteBranch := os.Getenv("REMOTE_BRANCH")

	if remoteBranch == "" {
		// load env from .env file
		envPath := GetBasePath() + "/.env"
		err := godotenv.Load(envPath)

		if err != nil {
			log.Panicln(err)
		}
	}
}
