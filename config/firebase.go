package config

import (
	"os"

	firebase "firebase.google.com/go/v4"
)

func NewFirebaseConfig(cfg *Config) *firebase.Config {
	var firebaseProjectID string

	firebaseProjectID, ok := os.LookupEnv("FIREBASE_PROJECT_ID")
	if !ok {
		return &firebase.Config{ProjectID: cfg.GoogleProjectID}
	}
	return &firebase.Config{ProjectID: firebaseProjectID}
}
