package google

import (
	"encoding/json"
	"fmt"
	log "github.com/cihub/seelog"
	"github.com/HailoOSS/oauth-client-service/provider/types"
	"github.com/HailoOSS/oauth-client-service/utils"
	"io/ioutil"
	"net/http"
)

const (
	GoogleUserInfoUrl = "https://www.googleapis.com/plus/v1/people/me"
)

type GoogleOauthUserInfo struct {
	UserEmail      string `json:"email"`
	UserGivenName  string `json:"givenName"`
	UserFamilyName string `json:"familyName"`
	UserGender     string `json:"gender"`
}

func (i *GoogleOauthUserInfo) Email() string {
	return i.UserEmail
}

func (i *GoogleOauthUserInfo) GivenName() string {
	return i.UserGivenName
}

func (i *GoogleOauthUserInfo) FamilyName() string {
	return i.UserFamilyName
}

func convertToGoogleOauthUserInfo(content map[string]interface{}) GoogleOauthUserInfo {
	var user GoogleOauthUserInfo

	if len(content) == 0 {
		return user
	}

	emails := utils.GetArrayInterfaceField(content, "emails")
	// get the first email we find
	for _, data := range emails {
		email := utils.GetMap(data)

		user.UserEmail = utils.GetStringValue(email, "value")
		if user.UserEmail != "" {
			// got one, stop here
			break
		}
	}

	var empty interface{}
	name := utils.GetInterfaceField(content, "name")
	if name != empty {
		data := utils.GetMap(name)

		user.UserGivenName = utils.GetStringValue(data, "givenName")
		user.UserFamilyName = utils.GetStringValue(data, "familyName")
	}

	user.UserGender = utils.GetStringValue(content, "gender")

	return user
}

func GoogleUserInfo(token string) (types.OauthProviderUserInfo, error) {
	var data GoogleOauthUserInfo

	req, err := http.NewRequest("GET", GoogleUserInfoUrl, nil)

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

	var content map[string]interface{}
	err = json.Unmarshal(body, &content)

	if err != nil {
		return &data, fmt.Errorf("Unable to get user info: %s", err.Error())
	}

	log.Debugf("content: %+v", content)
	data = convertToGoogleOauthUserInfo(content)
	log.Debugf("data: %+v", data)
	return &data, nil
}
