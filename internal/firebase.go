package internal

import (
	"context"
	"fmt"
	"os"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"firebase.google.com/go/v4/db"
	"google.golang.org/api/option"
)

var FirebaseApp *firebase.App

func InitFirebaseApp() error {
	file := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	if file == "" {
		return fmt.Errorf("GOOGLE_APPLICATION_CREDENTIALS not set")
	}

	opt := option.WithCredentialsFile(file)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return fmt.Errorf("error initializing firebase app: %v", err)
	}

	FirebaseApp = app
	return nil
}

func GetAuthClient() (*auth.Client, error) {
	if FirebaseApp == nil {
		if err := InitFirebaseApp(); err != nil {
			return nil, err
		}
	}

	client, err := FirebaseApp.Auth(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error getting Auth client: %v", err)
	}

	return client, nil
}

func GetRealTimeDBClient() (*db.Client, error) {
	if FirebaseApp == nil {
		if err := InitFirebaseApp(); err != nil {
			return nil, err
		}
	}

	url := os.Getenv("REALTIME_DB_URL")

	client, err := FirebaseApp.DatabaseWithURL(context.Background(), url)
	if err != nil {
		return nil, fmt.Errorf("error getting RealtimeDB Client: %v", err)
	}

	return client, nil
}
