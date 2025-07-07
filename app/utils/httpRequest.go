package utils

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type HTTPRequestParams struct {
	Method  string
	Headers map[string]string
	Url     string
}

type UrlParams struct {
	Protocol    string
	Host        string
	Port        string
	Path        string
	PathParams  []string
	QueryParams map[string]string
}

func ContructUrl(params UrlParams) string {
	baseURL := fmt.Sprintf("%s://%s", params.Protocol, params.Host)

	if params.Port != "" {
		baseURL = fmt.Sprintf("%s:%s", baseURL, params.Port)
	}

	fullPath := params.Path
	if len(params.PathParams) > 0 {
		fullPath = fmt.Sprintf(
			"%s/%s", strings.TrimSuffix(params.Path, "/"),
			strings.Join(params.PathParams, "/"),
		)
	}

	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return ""
	}
	parsedURL.Path = fullPath

	query := parsedURL.Query()
	for key, value := range params.QueryParams {
		query.Set(key, value)
	}
	parsedURL.RawQuery = query.Encode()

	return parsedURL.String()
}

func HTTPRequest(params HTTPRequestParams, requestBody io.Reader) (int, []byte, error) {
	request, err := http.NewRequest(params.Method, params.Url, requestBody)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	for key, value := range params.Headers {
		request.Header.Set(key, value)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return 0, nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return response.StatusCode, nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return response.StatusCode, body, nil
}
