package routes

import (
	"api/src/controllers"
	"net/http"
)

var routesUsers = []Route{
	{
		Uri: "/users",
		Method: http.MethodPost,
		Func: controllers.CreateUser,
		RequiresAutentication: false,
	},
	{
		Uri: "/users",
		Method: http.MethodGet,
		Func: controllers.GetUsers,
		RequiresAutentication: false,
	},
	{
		Uri: "/users/{userId}",
		Method: http.MethodGet,
		Func: controllers.GetUser,
		RequiresAutentication: false,
	},
	{
		Uri: "/users/{userId}",
		Method: http.MethodPut,
		Func: controllers.UpdateUser,
		RequiresAutentication: false,
	},
	{
		Uri: "/users/{userId}",
		Method: http.MethodDelete,
		Func: controllers.DeleteUser,
		RequiresAutentication: false,
	},
}