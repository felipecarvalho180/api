package routes

import (
	"devbook-api/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:                "/users",
		Method:             http.MethodPost,
		Func:               controllers.CreateUser,
		NeedAuthentication: false,
	},
	{
		URI:                "/users",
		Method:             http.MethodGet,
		Func:               controllers.FindUsers,
		NeedAuthentication: false,
	},
	{
		URI:                "/users/{user_id}",
		Method:             http.MethodGet,
		Func:               controllers.FindUser,
		NeedAuthentication: false,
	},
	{
		URI:                "/users/{user_id}",
		Method:             http.MethodPut,
		Func:               controllers.UpdateUser,
		NeedAuthentication: false,
	},
	{
		URI:                "/users/{user_id}",
		Method:             http.MethodDelete,
		Func:               controllers.DeleteUser,
		NeedAuthentication: false,
	},
}
