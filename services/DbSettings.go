package services

import (
	"log"
	"os"

	model "kpt_api/model"

	"github.com/joho/godotenv"
)

func GetDbSettings() *model.DbSettings {
	DB_URI_LOCAL, DB_NAME, GPU_COLLECTION_NAME, CPU_COLLECTION_NAME, GPU_FILTER_COLLECTION_NAME := getEnvs()
	DbSettings := new(model.DbSettings)
	DbSettings.ConnectUri = DB_URI_LOCAL
	DbSettings.DatabaseName = DB_NAME
	DbSettings.CollectionName = append(DbSettings.CollectionName, GPU_COLLECTION_NAME)
	DbSettings.CollectionName = append(DbSettings.CollectionName, GPU_FILTER_COLLECTION_NAME)
	DbSettings.CollectionName = append(DbSettings.CollectionName, CPU_COLLECTION_NAME)
	return DbSettings
}

func getEnvs() (string, string, string, string, string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("DB_URI_LOCAL"), os.Getenv("DB_NAME"), os.Getenv("GPU_COLLECTION_NAME"), os.Getenv("CPU_COLLECTION_NAME"), os.Getenv("GPU_FILTER_COLLECTION_NAME")
}
