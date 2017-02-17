package facebook

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
	FacebookTokenInfoUrl = "https://graph.facebook.com/oauth/access_token_info"
)

type FacebookOauthTokenInfo struct {
	Expires int64  `json:"expires_in"`
	Access  string `json:"token_type"`
}

func (i *FacebookOauthTokenInfo) IsValid() bool {
	if i.Expires > 0 {
		return true
	}

	return false
}

func (i *FacebookOauthTokenInfo) Email() string {
	return ""
}

func FacebookTokenInfo(token string) (types.OauthProviderTokenInfo, error) {
	var data FacebookOauthTokenInfo

	rsp, err := http.PostForm(FacebookTokenInfoUrl, url.Values{"access_token": {token}})

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
