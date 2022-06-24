package config

import (
	"os"
	"time"

	"admin.server/logger"
	"github.com/joho/godotenv"
)

const AccessTokenExpire = time.Minute * 1

const RefreshTokenExpire = time.Hour * (24 * 30)

const AccessCookieExpire = time.Minute * 1

const RefreshCookieExpire = time.Hour * (24 * 30)

func EnvCloudName() string {
    err := godotenv.Load()
    if err != nil {
        logger.Logger.ErrorLogger.Println(err)
    }
    return os.Getenv("CLOUDINARY_CLOUD_NAME")
}

func EnvCloudAPIKey() string {
    err := godotenv.Load()
    if err != nil {
        logger.Logger.ErrorLogger.Println(err)
    }
    return os.Getenv("CLOUDINARY_API_KEY")
}

func EnvCloudAPISecret() string {
    err := godotenv.Load()
    if err != nil {
        logger.Logger.ErrorLogger.Println(err)
    }
    return os.Getenv("CLOUDINARY_API_SECRET")
}


func Secret() string {
    err := godotenv.Load()
    if err != nil {
        logger.Logger.ErrorLogger.Println(err)
    }
    return os.Getenv("JWT_SECRET")
}

