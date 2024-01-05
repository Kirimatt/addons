package http

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"

	"code.sajari.com/docconv"
)

func GetDataFromUrl(url string) (text *string, finalUrl *string, err error) {
	initResp, err := http.Get(url)
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to make GET request: %w", err)
	}
	defer initResp.Body.Close()

	finalURL := strings.Split(initResp.Request.URL.String(), ";")
	fmt.Printf("Redirected to url: %s \n", finalURL[0])

	resp, err := http.Get(finalURL[0])
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to make GET request: %w", err)
	}
	defer resp.Body.Close()

	fileData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to read response body: %w", err)
	}

	tmpFile, err := os.CreateTemp("", "docxfile*.docx")
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to create temporary file: %w", err)
	}
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.Write(fileData)
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to write to temporary file: %w", err)
	}

	res, err := docconv.ConvertPath(tmpFile.Name())
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to convert temporary file: %w", err)
	}

	return &res.Body, &finalURL[0], nil
}

func GetUrlsFromSearch(searchUrl string, placeholderUrl string) (urls *map[string]string, err error) {
	resp, err := http.Get(searchUrl)
	if err != nil {
		return nil, fmt.Errorf("Failed to make GET request: %w", err)
	}
	defer resp.Body.Close()

	html, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read response body: %w", err)
	}

	regex := regexp.MustCompile(`<a href="/rus/docs/([^"]+)">[\s\S]*?<span class="status ([^"]+)">([^<]+)</span>`)
	matches := regex.FindAllStringSubmatch(string(html), -1)

	var result map[string]string = make(map[string]string)
	for _, match := range matches {
		result[fmt.Sprintf(placeholderUrl, match[1])] = match[2]
		unstructedData.WriteString(match[1])
		unstructedData.WriteString(match[2])
	}

	return &result, unstructedData.String(), nil
}
