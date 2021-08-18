package main

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"google.golang.org/grpc/credentials/oauth"
	"log"
	"os"
)

// Retrieves a token from a local file.
func tokenFromFile(file string) *oauth2.Token {
	f, err := os.Open(file)
	if err != nil {
		log.Fatalf("Error decoding token: %v", err)
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok
}

func retrieveToken() oauth.TokenSource {
	token := tokenFromFile("token.json")
	return oauth.TokenSource{TokenSource: oauth2.StaticTokenSource(token)}
}

func configService() *sheets.Service {
	srv, err := sheets.NewService(context.Background(), option.WithTokenSource(retrieveToken()))

	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	return srv
}

func main() {

	spreadsheetId := "176FYD8uB5-S7SAXEAf4XWETk0oXJvGGJgeJnbz9WL3g"
	readRange := "A2:A4"
	resp, err := configService().Spreadsheets.Values.Get(spreadsheetId, readRange).Do()

	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
		return
	}

	for _, row := range resp.Values {
		if len(row) > 0 {
			fmt.Printf("%s\n", row[0])
		} else {
			fmt.Println("{Empty row}")
		}

	}

}
