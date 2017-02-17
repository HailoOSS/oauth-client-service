package types

type OauthProviderTokenInfo interface {
	IsValid() bool
}

type OauthProviderUserInfo interface {
	Email() string
	GivenName() string
	FamilyName() string
}
