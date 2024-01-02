package http

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"code.sajari.com/docconv"
)

func GetDataFromUrl() (url string, err error) {
	initResp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to make GET request: %s", err)
	}
	defer initResp.Body.Close()

	finalURL := strings.Split(initResp.Request.URL.String(), ";")
	fmt.Println(finalURL[0])

	resp, err := http.Get(finalURL[0])
	if err != nil {
		log.Fatalf("Failed to make GET request: %s", err)
	}
	defer resp.Body.Close()

	fileData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %s", err)
	}

	tmpFile, err := os.CreateTemp("", "docxfile*.docx")
	if err != nil {
		log.Fatalf("Failed to create temporary file: %s", err)
	}

	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.Write(fileData)
	if err != nil {
		log.Fatalf("Failed to write to temporary file: %s", err)
	}

	res, err := docconv.ConvertPath(tmpFile.Name())
	if err != nil {
		log.Fatalf("Failed to convert temporary file: %s", err)
	}

	return res.Body, nil
}
