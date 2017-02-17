package google

import (
	"encoding/json"
	"fmt"
	log "github.com/cihub/seelog"
	"github.com/HailoOSS/oauth-client-service/provider/types"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	GoogleTokenInfoUrl = "https://www.googleapis.com/oauth2/v1/tokeninfo"
)

type GoogleOauthTokenInfo struct {
	Audience  string `json:"audience"`
	IssuedTo  string `json:"issued_to"`
	UserId    string `json:"user_id"`
	Scope     string `json:"scope"`
	Expires   int64  `json:"expires_in"`
	UserEmail string `json:"email"`
	Access    string `json:"access_type"`
	Verified  bool   `json:"verified_email"`
}

func (i *GoogleOauthTokenInfo) IsValid() bool {
	if i.UserEmail == "" || !i.Verified {
		return false
	}

	return true
}

func GoogleTokenInfo(token string) (types.OauthProviderTokenInfo, error) {
	var data GoogleOauthTokenInfo

	rsp, err := http.PostForm(GoogleTokenInfoUrl, url.Values{"access_token": {token}})

	if err != nil {
		return &data, fmt.Errorf("Unable to verify token: %s", err.Error())
	}

	body, err := ioutil.ReadAll(rsp.Body)

	if err != nil {
		return &data, fmt.Errorf("Unable to verify token: %s", err.Error())
	}

	log.Debugf("body: %+v", string(body))
	defer rsp.Body.Close()

	if rsp.StatusCode == 400 {
		return &data, fmt.Errorf("Invalid token")
	}

	err = json.Unmarshal(body, &data)

	if err != nil {
		return &data, fmt.Errorf("Unable to verify token: %s", err.Error())
	}

	log.Debugf("data: %+v", data)
	return &data, nil
}
