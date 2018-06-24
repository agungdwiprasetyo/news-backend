package config

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	sessions "github.com/gin-gonic/contrib/sessions"
	env "github.com/joho/godotenv"
)

var (
	JWTpublicKey  *rsa.PublicKey
	JWTprivateKey *rsa.PrivateKey
	Sessions      sessions.CookieStore
)

func AppConfig(appPath string) error {
	// init .env variable
	if strings.TrimSpace(appPath) == "" {
		return fmt.Errorf("Error: Cannot find config path %s", appPath)
	}
	appPath = strings.TrimRight(appPath, "/")
	if err := env.Load(appPath + "/.env"); err != nil {
		return fmt.Errorf("Error: %v", err)
	}
	os.Setenv("APP_DIR", appPath)

	// init JWT key
	keyPath := fmt.Sprintf("%s/%s", appPath, os.Getenv("JWT_PRIVATE_KEY"))
	signBytes, err := ioutil.ReadFile(keyPath)
	JWTprivateKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		return err
	}
	keyPath = fmt.Sprintf("%s/%s", appPath, os.Getenv("JWT_PUBLIC_KEY"))
	verifyBytes, err := ioutil.ReadFile(keyPath)
	JWTpublicKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		return err
	}

	// // init cloud computing storage
	// cloud.Init(appPath)

	return nil
}
