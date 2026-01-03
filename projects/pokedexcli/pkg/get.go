package pkg

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Get(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []byte{}, fmt.Errorf("error creating request: %v", err)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return []byte{}, fmt.Errorf("network error: %v", err)
	}
	defer res.Body.Close()

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("error reading response body: %v", err)
	}

	return bytes, nil
}

func Unmarshal[T any](bytes []byte, v *T) error {
	if err := json.Unmarshal(bytes, v); err != nil {
		return fmt.Errorf("error parsing JSON: %v", err)
	}

	return nil
}
