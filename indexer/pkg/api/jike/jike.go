package jike

import (
	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/util"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/config"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/logger"
	jsoniter "github.com/json-iterator/go"
	"github.com/robfig/cron/v3"
)

//nolint:tagliatelle // format is required by Jike API
type RefreshTokenStruct struct {
	AccessToken  string `json:"x-jike-access-token"`
	RefreshToken string `json:"x-jike-refresh-token"`
	Success      bool   `json:"success"`
}

var (
	jsoni        = jsoniter.ConfigCompatibleWithStandardLibrary
	AccessToken  string
	RefreshToken string
)

func main() {
	c := cron.New()
	// everyday at 00:00, refresh Jike tokens
	c.AddFunc("0 0 * * *", func() { Login() })
	c.Start()
}

func Login() error {
	config.Setup()

	json, err := jsoni.MarshalToString(config.ThirdPartyConfig.Jike)

	if err != nil {
		logger.Fatalf("Jike Config read err: %v", err)

		return err
	}

	headers := map[string]string{
		"App-Version":  config.ThirdPartyConfig.Jike.AppVersion,
		"Content-Type": "application/json",
	}

	loginUrl := "https://api.ruguoapp.com/1.0/users/loginWithPhoneAndPassword"

	response, err := util.Post(loginUrl, headers, json)

	if err != nil {
		logger.Fatalf("Jike Login err: %v", err)

		return err
	}

	AccessToken = string(response.Header().Get("x-jike-access-token"))
	RefreshToken = string(response.Header().Get("x-jike-refresh-token"))

	return nil
}

func RefreshJikeToken() error {
	headers := map[string]string{
		"App-Version":          config.ThirdPartyConfig.Jike.AppVersion,
		"Content-Type":         "application/json",
		"x-jike-refresh-token": RefreshToken,
	}

	refreshUrl := "https://api.ruguoapp.com/app_auth_tokens.refresh"

	response, err := util.Get(refreshUrl, headers)

	if err != nil {
		logger.Fatalf("Jike RefreshToken err: %v", err)

		return err
	}

	token := new(RefreshTokenStruct)

	err = jsoni.Unmarshal(response, &token)

	if err != nil {
		if token.Success {
			AccessToken = token.AccessToken
			RefreshToken = token.RefreshToken

			return nil
		} else {
			logger.Fatalf("Jike RefreshToken err: %v", "Jike refresh token endpoint returned a failed response")

			return err
		}
	} else {
		logger.Fatalf("Jike RefreshToken err: %v", err)

		return err
	}
}
