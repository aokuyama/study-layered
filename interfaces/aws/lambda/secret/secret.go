package secret

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const (
	endpoint      = "http://localhost:2773/systemsmanager/parameters/get"
	headerSecrets = "X-Aws-Parameters-Secrets-Token"
)

type resultFromExtension struct {
	Parameter struct {
		ARN              string
		DateType         string
		LastModifiedDate time.Time
		Name             string
		Selector         string
		SourceResult     *string
		Type             string
		Value            string
		Version          int
	}
	ResultMetadata any
}

func getSecretByUsingExtension(key string) (string, error) {
	query := url.Values{}
	query.Add("name", key)
	query.Add("withDecryption", "true")
	queryStr := query.Encode()

	url := fmt.Sprintf("%s?%s", endpoint, queryStr)
	req, err := http.NewRequestWithContext(context.Background(), "GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Add(headerSecrets, os.Getenv("AWS_SESSION_TOKEN"))

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(res.Body); err != nil {
		return "", err
	}
	bodyString := buf.String()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get parameter by using extension. statusCode:%d body:%s", res.StatusCode, bodyString)
	}

	exRes := resultFromExtension{}
	if err := json.Unmarshal([]byte(bodyString), &exRes); err != nil {
		return "", err
	}

	return exRes.Parameter.Value, nil
}

func SetEnvBySecretParam(key, prefix string) error {
	paramKey := "/" + prefix + "/app/secret/" + strings.ToLower(key)

	v, err := getSecretByUsingExtension(paramKey)
	if err != nil {
		return err
	}
	os.Setenv(key, v)
	return nil
}
