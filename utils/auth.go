package utils

import "net/http"

func IsAuthenticated(req *http.Request) bool {
	exists := appConfig.Session.Exists(req.Context() , "user_id")

	return exists
}