package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

// CRUDHandler .
type CRUDHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	Read(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func parseBody(r *http.Request) (map[string]interface{}, error) {
	if r.Method != http.MethodPost {
		return nil, nil
	}
	var body map[string]interface{}
	raw, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(raw, &body); err != nil {
		return nil, err
	}
	return body, nil
}

func parseIdentifier(r *http.Request, endpoint string) string {
	split := strings.Split(r.URL.Path, "/")
	lastIdx := len(split) - 1
	if strings.Contains(endpoint, split[lastIdx]) {
		return ""
	}
	return split[lastIdx]
}
