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
		NeedAuthentication: true,
	},
	{
		URI:                "/users/{user_id}",
		Method:             http.MethodGet,
		Func:               controllers.FindUser,
		NeedAuthentication: true,
	},
	{
		URI:                "/users/{user_id}",
		Method:             http.MethodPut,
		Func:               controllers.UpdateUser,
		NeedAuthentication: true,
	},
	{
		URI:                "/users/{user_id}",
		Method:             http.MethodDelete,
		Func:               controllers.DeleteUser,
		NeedAuthentication: true,
	},
	{
		URI:                "/users/{user_id}/follow",
		Method:             http.MethodPost,
		Func:               controllers.FollowUser,
		NeedAuthentication: true,
	},
	{
		URI:                "/users/{user_id}/unfollow",
		Method:             http.MethodPost,
		Func:               controllers.UnfollowUser,
		NeedAuthentication: true,
	},
	{
		URI:                "/users/{user_id}/followers",
		Method:             http.MethodGet,
		Func:               controllers.FindFollowers,
		NeedAuthentication: true,
	},
	{
		URI:                "/users/{user_id}/following",
		Method:             http.MethodGet,
		Func:               controllers.FindFollowing,
		NeedAuthentication: true,
	},
	{
		URI:                "/users/{user_id}/update-password",
		Method:             http.MethodPost,
		Func:               controllers.UpdatePassword,
		NeedAuthentication: true,
	},
}
