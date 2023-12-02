package util

import (
	"bytes"
	"github.com/k3a/html2text"
	"io"
	"net/http"
	"strings"
)

func SubmitAnswer(year string, day string, level string, answer string) (string, error) {
	url := "https://adventofcode.com/" + year + "/day/" + day + "/answer"
	payload := []byte("level=" + level + "&answer=" + answer)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return "", err
	}

	req.Header.Set("authority", "adventofcode.com")
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8")
	req.Header.Set("accept-language", "en-GB,en;q=0.7")
	req.Header.Set("cache-control", "max-age=0")
	req.Header.Set("content-type", "application/x-www-form-urlencoded")
	req.Header.Set("cookie", "session="+AocSession)
	req.Header.Set("origin", "https://adventofcode.com")
	req.Header.Set("referer", "https://adventofcode.com/")
	req.Header.Set("sec-fetch-dest", "document")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sec-fetch-user", "?1")
	req.Header.Set("sec-gpc", "1")
	req.Header.Set("upgrade-insecure-requests", "1")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var bodyBytes, errRead = io.ReadAll(resp.Body)
	if errRead != nil {
		return "", err
	}
	bodyString := string(bodyBytes)
	return bodyString, nil
}

func IsGlobalLeaderboardFull(year string, day string, level string) (bool, error) {
	urlInput := "https://adventofcode.com/" + year + "/leaderboard/day/" + day
	reqInput, err := http.NewRequest("GET", urlInput, nil)
	if err != nil {
		return false, err
	}
	reqInput.Header.Set("Cookie", "session="+AocSession)
	respInput, err := http.DefaultClient.Do(reqInput)
	if err != nil {
		return false, err
	}
	defer respInput.Body.Close()

	var bodyBytes, errRead = io.ReadAll(respInput.Body)
	if errRead != nil {
		return false, err
	}
	bodyString := string(bodyBytes)
	if level == "1" {
		return strings.Count(bodyString, "100)") >= 1, nil
	}
	if level == "2" {
		return strings.Count(bodyString, "100)") >= 2, nil
	}
	return false, nil
}
func DownloadInput(year string, day string) (string, error) {
	urlInput := "https://adventofcode.com/" + year + "/day/" + day + "/input"
	reqInput, err := http.NewRequest("GET", urlInput, nil)
	if err != nil {
		return "", err
	}
	reqInput.Header.Set("Cookie", "session="+AocSession)
	respInput, err := http.DefaultClient.Do(reqInput)
	if err != nil {
		return "", err
	}
	defer respInput.Body.Close()

	var bodyBytes, errRead = io.ReadAll(respInput.Body)
	if errRead != nil {
		return "", err
	}
	bodyString := string(bodyBytes)
	return bodyString, nil
}

func DownloadText(year string, day string) (string, error) {
	urlText := "https://adventofcode.com/" + year + "/day/" + day + ""
	reqText, err := http.NewRequest("GET", urlText, nil)
	if err != nil {
		return "", err
	}
	reqText.Header.Set("Cookie", "session="+AocSession)

	respText, err := http.DefaultClient.Do(reqText)
	if err != nil {
		return "", err
	}
	defer respText.Body.Close()
	var bodyBytes, errRead = io.ReadAll(respText.Body)
	if errRead != nil {
		return "", err
	}
	bodyTextString := string(bodyBytes)
	plainText := html2text.HTML2Text(bodyTextString)
	return plainText, nil
}
