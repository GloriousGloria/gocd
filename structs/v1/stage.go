package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	ak "github.com/jnnkrdb/gomw/middlewares/security/authorization/apikey"
)

type EnvironmentList []Environment

type Environment struct {
	Name              string `json:"name"`
	ClusterAPI        string `json:"clusterapi"`
	Subdomain         string `json:"subdomain"`
	AppendSuffix      bool   `json:"appendsuffix"`
	UseRegexSubdomain bool   `json:"useregexsubdomain"`
}

// receive lists from url
func (envList *EnvironmentList) GetFrom(url string, apikey string) (err error) {
	var request *http.Request
	if request, err = http.NewRequest(http.MethodGet, url, nil); err == nil {
		ak.SetAPIKeyForRequest(request)
		var response *http.Response
		if response, err = http.DefaultClient.Do(request); err == nil {
			defer response.Body.Close()
			if response.StatusCode == http.StatusOK {
				err = json.NewDecoder(response.Body).Decode(envList)
			} else {
				err = fmt.Errorf("%s", response.Header.Get("X-ERROR-MSG"))
			}
		}
	}
	return
}

// check the existance of a specific stage
func (el EnvironmentList) Contains(environment string) bool {
	for i := range el {
		if el[i].Name == environment {
			return true
		}
	}
	return false
}
