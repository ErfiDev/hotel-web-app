package utils

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func ClientError(w http.ResponseWriter , status int) {
	appConfig.InfoLog.Println("client error with status:" , status)
	http.Error(w , http.StatusText(status) , status)
}

func ServerError(w http.ResponseWriter , err error) {
	trace := fmt.Sprintf("%s\n%s" , err.Error() , debug.Stack())

	appConfig.ErrorLog.Println(trace)

	http.Error(w , http.StatusText(http.StatusInternalServerError) , http.StatusInternalServerError)

}