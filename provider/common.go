package provider

import (
	"fmt"
	common "github.com/HailoOSS/oauth-client-service/proto"
	facebook "github.com/HailoOSS/oauth-client-service/provider/facebook"
	google "github.com/HailoOSS/oauth-client-service/provider/google"
	types "github.com/HailoOSS/oauth-client-service/provider/types"
	"strings"
)

func CheckProviderToken(provider string, token string) (types.OauthProviderTokenInfo, error) {
	enum, ok := common.OauthProvider_value[strings.ToUpper(provider)]
	if !ok {
		return nil, fmt.Errorf("Invalid provider %s", provider)
	}

	switch common.OauthProvider(enum) {
	case common.OauthProvider_GOOGLE:
		return google.GoogleTokenInfo(token)
	}

	return nil, fmt.Errorf("No method available to check token with provider %s", provider)
}

func TokenUserInfo(provider string, token string) (types.OauthProviderUserInfo, error) {
	enum, ok := common.OauthProvider_value[strings.ToUpper(provider)]
	if !ok {
		return nil, fmt.Errorf("Invalid provider %s", provider)
	}

	switch common.OauthProvider(enum) {
	case common.OauthProvider_GOOGLE:
		return google.GoogleUserInfo(token)
	case common.OauthProvider_FACEBOOK:
		return facebook.FacebookUserInfo(token)
	}

	return nil, fmt.Errorf("No method available to check token with provider %s", provider)
}
