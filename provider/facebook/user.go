package facebook

import (
	"encoding/json"
	"fmt"
	log "github.com/cihub/seelog"
	"github.com/HailoOSS/oauth-client-service/provider/types"
	"io/ioutil"
	"net/http"
)

const (
	FacebookUserInfoUrl = "https://graph.facebook.com/me"
)

type FacebookOauthUserInfo struct {
	UserEmail      string `json:"email"`
	UserGivenName  string `json:"first_name"`
	UserFamilyName string `json:"last_name"`
	UserGender     string `json:"gender"`
}

func (i *FacebookOauthUserInfo) Email() string {
	return i.UserEmail
}

func (i *FacebookOauthUserInfo) GivenName() string {
	return i.UserGivenName
}

func (i *FacebookOauthUserInfo) FamilyName() string {
	return i.UserFamilyName
}

func FacebookUserInfo(token string) (types.OauthProviderUserInfo, error) {
	var data FacebookOauthUserInfo

	req, err := http.NewRequest("GET", FacebookUserInfoUrl, nil)

	if err != nil {
		return &data, fmt.Errorf("Unable to create request: %s", err.Error())
	}

	req.Header.Add("Authorization", "Bearer "+token)
	rsp, err := http.DefaultClient.Do(req)

	if err != nil {
		return &data, fmt.Errorf("Unable to get user info: %s", err.Error())
	}

	body, err := ioutil.ReadAll(rsp.Body)

	if err != nil {
		return &data, fmt.Errorf("Unable to get user info: %s", err.Error())
	}

	log.Debugf("body: %+v", string(body))
	defer rsp.Body.Close()

	if rsp.StatusCode != 200 {
		return &data, fmt.Errorf("Invalid credentials")
	}

	var content FacebookOauthUserInfo
	err = json.Unmarshal(body, &content)

	if err != nil {
		return &data, fmt.Errorf("Unable to get user info: %s", err.Error())
	}

	log.Debugf("content: %+v", content)
	return &content, nil
}
