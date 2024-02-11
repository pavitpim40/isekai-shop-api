package model

type (
	UserInfo struct {
		Id            string `json:"id"`
		Email         string `json:"email"`
		VerifiedEmail bool   `json:"verified_email"`
		Name          string `json:"name"`
		GivenName     string `json:"given_name"`
		FamilyName    string `json:"family_name"`
		Picture       string `json:"picture"`
		Locale        string `json:"locale"`
	}

	UserPassport struct {
		RefreshToken string `json:"refresh_token"`
	}

	CreateUserInfo struct {
		Id      string
		Email   string
		Name    string
		Picture string
		*UserPassport
	}
)
