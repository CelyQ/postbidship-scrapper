package main

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	drive "google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"os"
)

func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}

	return config.Client(context.Background(), tok)
}

func saveToken(file string, tok *oauth2.Token) {
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(tok)
}

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	log.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code %v", err)
	}

	tok, err := config.Exchange(context.Background(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}
	return tok
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

func uploadFileToGoogleDrive(filename string) {

	clientID := os.Getenv("GOOGLE_DRIVE_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_DRIVE_CLIENT_SECRET")
	apiKey := os.Getenv("GOOGLE_DRIVE_API_KEY")
	parentFolderID := os.Getenv("GOOGLE_DRIVE_PARENT_FOLDER_ID")

	if clientID == "" || clientSecret == "" || apiKey == "" {
		log.Fatal("Missing required environment variables: GOOGLE_DRIVE_CLIENT_ID, GOOGLE_DRIVE_CLIENT_SECRET, or GOOGLE_DRIVE_API_KEY")
	}

	config := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  "urn:ietf:wg:oauth:2.0:oob",
		Scopes:       []string{drive.DriveScope},
		Endpoint:     google.Endpoint,
	}

	srv, err := drive.NewService(
		context.Background(),
		option.WithAPIKey(apiKey),
		option.WithHTTPClient(getClient(config)),
	)

	file := &drive.File{
		Name:     filename,
		MimeType: "application/vnd.google-apps.spreadsheet",
		Parents:  []string{parentFolderID},
	}

	existingFiles, _ := srv.Files.List().Do()
	for _, file := range existingFiles.Files {
		if file.Name == filename {
			_ = srv.Files.Delete(file.Id).Do()
		}
	}

	content, _ := os.Open(filename)
	if err != nil {
		log.Fatalf("Unable to read file %v", err)
	}

	_, err = srv.Files.Create(file).Media(content).Do()
	if err != nil {
		log.Fatalf("Unable to create file %v", err)

	}

}
