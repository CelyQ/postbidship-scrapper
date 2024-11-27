package main

import (
	"context"
	"flag"
	"github.com/chromedp/chromedp"
	"log"
	"os"
)

type App struct {
	url      string
	email    string
	password string
}

func main() {
	email := os.Getenv("EMAIL")
	password := os.Getenv("PASSWORD")

	if email == "" || password == "" {
		log.Fatal("EMAIL and PASSWORD environment variables are required")
	}


	var verbose bool
	flag.BoolVar(&verbose, "verbose", false, "verbose output")
	flag.Parse()

	opts := append(
		chromedp.DefaultExecAllocatorOptions[:],
		chromedp.NoFirstRun,
		chromedp.NoDefaultBrowserCheck,
		chromedp.Flag("headless", !verbose),
		chromedp.Flag("disable-extensions", true),
	)
	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	app := App{
		url:      "https://web.postbidship.com/sign-in",
		email:    email,
		password: password,
	}

	parser := NewParser(ctx, app)

	parser.InputAccountCredentials()

	headers := parser.GetTableHeaders()

	var combinedRows [][]string

	currentPage := 1
	totalNumberOfPages := parser.GetPages()

	for currentPage <= totalNumberOfPages {
		rows := parser.GetTableRows()

		combinedRows = append(combinedRows, rows...)
		currentPage++

		if currentPage > totalNumberOfPages {
			break
		}

		parser.NextPage()

	}

	tableMap := make(map[string][]string)
	for i, header := range headers {
		tableMap[header] = []string{}

		for _, row := range combinedRows {
			tableMap[header] = append(tableMap[header], row[i])
		}
	}

	filename := createExcelFile(tableMap)

	uploadFileToGoogleDrive(filename)

	err := os.Remove(filename)
	if err != nil {
		return
	}
}
