package thecatapi

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type CatAPIResponse struct {
	URL string `json:"url"`
}

const (
	requestAPI        = "https://api.thecatapi.com/v1/images/search"
	mimeTypes         = "mime_types"
	gifType           = "gif"
	catapiTokenHeader = "x-api-key"
)

func GetCatImageURL(catAPIToken string, gif bool) (string, error) {
	req, err := http.NewRequest(http.MethodGet, requestAPI, nil)
	if err != nil {
		return "", errors.New("Can't create request")
	}
	req.Header.Set(catapiTokenHeader, catAPIToken)

	if gif {
		urlQuery := req.URL.Query()
		urlQuery.Add(mimeTypes, gifType)
		req.URL.RawQuery = urlQuery.Encode()
	}

	log.Printf("Request url: %s", req.URL.RequestURI())

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return "", errors.New("Can't get response from api")
	}
	defer resp.Body.Close()

	var catResps []CatAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&catResps); err != nil {
		return "", errors.New("Can't read body")
	}
	return catResps[0].URL, nil
}
