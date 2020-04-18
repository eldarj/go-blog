package auth

import (
	"encoding/json"
	"go-blog/auth/model"
	a "go-blog/bootstrap/lib/action"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func (service *AuthenticationService) authGoogle() {
	println("Hello there")
}

type AuthenticationService struct{}

func (service *AuthenticationService) GetEndpoints() map[string]a.IAction {
	return map[string]a.IAction{
		"/_googleAuthSuccess": onGoogleAuthSuccess,
	}
}

var onGoogleAuthSuccess = a.ProtoAction{UnderlyingRun: func(w http.ResponseWriter, r *http.Request) {
	p, _ := url.Parse(r.RequestURI)
	params, _ := url.ParseQuery(p.RawQuery)

	parsedEmail, _ := strconv.ParseInt(strings.Join(params["id"], ""), 10, 64)

	accountModel := model.GoogleAccountModel{
		Id:         parsedEmail,
		IdToken:    strings.Join(params["idToken"], ""),
		Email:      strings.Join(params["email"], ""),
		FullName:   strings.Join(params["fullName"], ""),
		GivenName:  strings.Join(params["givenName"], ""),
		FamilyName: strings.Join(params["familyName"], ""),
		ImageUrl:   strings.Join(params["imageUrl"], ""),
	}

	jsonAccounModel, _ := json.Marshal(accountModel)
	w.Write(jsonAccounModel)
}}
