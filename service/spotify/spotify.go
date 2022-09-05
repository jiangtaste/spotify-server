package spotify

import (
	"fmt"
	"spotify-server/core"

	"github.com/imroc/req/v3"
)

const (
	formContentType = "application/x-www-form-urlencoded"
	jsonContentType = "application/json; charset=utf-8"
)

var (
	adminClient *req.Client
)

// getAdmin 获取具有管理权限的client
func getAdmin() *req.Client {
	if adminClient == nil {
		// 从配置中获取 CLIENT_ID， SECRET
		clientID := core.Viper.GetString("spotify.client_id")
		secret := core.Viper.GetString("spotify.secret")

		client := req.C()
		client.SetCommonContentType(formContentType)
		client.SetCommonBasicAuth(clientID, secret)

		adminClient = client
	}

	return adminClient
}

// getClient 获取为web api服务的req client
func getClient(token string) *req.Client {
	c := req.C()
	c.SetCommonContentType(jsonContentType)
	c.SetCommonBearerAuthToken(token)
	return c
}

type SpotifySwipeTokenResp struct {
	SpotifyRefreshTokenResp
	// A token that can be sent to the Spotify Accounts apps in place of an authorization code.
	// (When the access code expires, send a POST request to the Accounts apps /api/token endpoint,
	// but use this code in place of an authorization code.
	// A new Access Token will be returned. A new refresh token might be returned too.)
	RefreshToken string `json:"refresh_token"`
}

type SpotifyRefreshTokenResp struct {
	// An Access Token that can be provided in subsequent calls, for example to Spotify Web API apps.
	AccessToken string `json:"access_token"`
	// How the Access Token may be used: always “Bearer”.
	TokenType string `json:"token_type"`
	// A space-separated list of scopes which have been granted for this access_token
	Scope string `json:"scope"`
	// The time period (in seconds) for which the Access Token is valid.
	ExpiresIn int64 `json:"expires_in"`
}

// SwipeToken 从 spotify 换取 access_token
func SwipeToken(code string, redirectURL string) *SpotifySwipeTokenResp {

	token := &SpotifySwipeTokenResp{}

	_, err := getAdmin().R().
		SetFormData(map[string]string{
			"code":         code,
			"redirect_uri": redirectURL,
			"grant_type":   GrantTypeAuthorizationCode,
		}).
		SetResult(token).
		Post("https://accounts.spotify.com/api/token")

	if err != nil {
		panic(err)
	}

	return token
}

// RefreshToken refresh access token from spotify
func RefreshToken(refreshToken string) *SpotifyRefreshTokenResp {
	token := &SpotifyRefreshTokenResp{}

	_, err := getAdmin().R().
		SetFormData(map[string]string{
			"refresh_token": refreshToken,
			"grant_type":    GrantTypeAuthorizationCode,
		}).
		SetResult(token).
		Post("https://accounts.spotify.com/api/token")

	if err != nil {
		panic(err)
	}

	return token
}

// GetCurrentUserProfile 获取制定用户的profile
func GetCurrentUserProfile(token string) (*CurrentUserProfile, error) {
	errMsg := &SpotifyError{}
	profile := &CurrentUserProfile{}

	resp, err := getClient(token).R().
		SetResult(profile).
		SetError(errMsg).
		Get("https://api.spotify.com/v1/me")

	if err != nil {
		panic(err)
	}

	if resp.IsSuccess() {
		return profile, nil
	} else {
		return nil, fmt.Errorf("spotify api error: %s", resp.Status)
	}
}

type SpotifyError struct {
	Error SpotifyErrorInfo `json:"error"`
}

type SpotifyErrorInfo struct {
	Status  int64  `json:"status"`
	Message string `json:"message"`
}

type CurrentUserProfile struct {
	// The country of the user, as set in the user's account profile.
	// An ISO 3166-1 alpha-2 country code.
	//  This field is only available when the current user has granted access to the user-read-private scope.
	Country string `json:"country"`

	// The name displayed on the user's profile. null if not available.
	DisplayName string `json:"display_name"`

	// The user's email address, as entered by the user when creating their account.
	// Important! This email address is unverified; there is no proof that it actually belongs to the user.
	// This field is only available when the current user has granted access to the user-read-email scope.
	Email string `json:"email"`

	// The user's explicit content settings.
	// This field is only available when the current user has granted access to the user-read-private scope.
	ExplicitContent ExplicitContent `json:"explicit_content"`

	// Known external URLs for this user.
	ExternalURLs ExternalURLs `json:"external_urls"`

	// Information about the followers of the user.
	Followers Followers `json:"followers"`

	// A link to the Web API endpoint for this user.
	Href string `json:"href"`

	// The Spotify user ID for the user.
	ID string `json:"id"`

	// The user's profile image.
	Images []Image `json:"images"`

	// The user's Spotify subscription level: "premium", "free", etc.
	// (The subscription level "open" can be considered the same as "free".)
	// This field is only available when the current user has granted access to the user-read-private scope.
	Product string `json:"product"`

	// The object type: "user"
	Type string `json:"type"`

	// The Spotify URI for the user.
	URI string `json:"uri"`
}

type ExplicitContent struct {
	// When true, indicates that explicit content should not be played.
	FilterEnabled bool `json:"filter_enabled"`
	// When true, indicates that the explicit content setting is locked and can't be changed by the user.
	FilterBocked bool `json:"filter_bocked"`
}

type ExternalURLs struct {
	// The Spotify URL for the object.
	Spotify string `json:"spotify"`
}

type Followers struct {
	// This will always be set to null, as the Web API does not support it at the moment.
	Href string `json:"href"`
	// The total number of followers.
	Total int64 `json:"total"`
}

type Image struct {
	// The source URL of the image.
	URL string `json:"url"`
	// The image height in pixels.
	Height int64 `json:"height"`
	// The image width in pixels.
	Width int64 `json:"width"`
}
