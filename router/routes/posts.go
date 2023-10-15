package routes

import (
	"devbook-api/controllers"
	"net/http"
)

var postsRoutes = []Route{
	{
		URI:                "/posts",
		Method:             http.MethodPost,
		Func:               controllers.CreatePost,
		NeedAuthentication: true,
	},
	{
		URI:                "/posts",
		Method:             http.MethodGet,
		Func:               controllers.GetPosts,
		NeedAuthentication: true,
	},
	{
		URI:                "/posts/{post_id}",
		Method:             http.MethodGet,
		Func:               controllers.GetPost,
		NeedAuthentication: true,
	},
	{
		URI:                "/posts/{post_id}",
		Method:             http.MethodPut,
		Func:               controllers.UpdatePost,
		NeedAuthentication: true,
	},
	{
		URI:                "/posts/{post_id}",
		Method:             http.MethodDelete,
		Func:               controllers.DeletePost,
		NeedAuthentication: true,
	},
	{
		URI:                "/users/{user_id}/posts",
		Method:             http.MethodGet,
		Func:               controllers.GetPostByUserID,
		NeedAuthentication: true,
	},
	{
		URI:                "/posts/{post_id}/like",
		Method:             http.MethodPost,
		Func:               controllers.Like,
		NeedAuthentication: true,
	},
	{
		URI:                "/posts/{post_id}/unlike",
		Method:             http.MethodPost,
		Func:               controllers.Unlike,
		NeedAuthentication: true,
	},
}
