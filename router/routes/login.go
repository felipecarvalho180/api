package routes

import (
	"devbook-api/controllers"
	"net/http"
)

var loginRouter = Route{
	URI:                "/login",
	Method:             http.MethodPost,
	Func:               controllers.Login,
	NeedAuthentication: false,
}
