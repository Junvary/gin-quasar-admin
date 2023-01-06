package utils

import (
	"io"
	"net/http"
	"strings"
)

func HttpPost(url string, data []byte, headers map[string]string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		return nil, err
	}
	for key, header := range headers {
		req.Header.Set(key, header)
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
