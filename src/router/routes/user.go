package routes

import (
	"api/src/controllers"
	"net/http"
)

var routesUsers = []Route{
	{
		Uri: "/usuarios",
		Method: http.MethodPost,
		Func: controllers.CreateUser,
		RequiresAutentication: false,
	},
	{
		Uri: "/usuarios",
		Method: http.MethodGet,
		Func: controllers.GetUsers,
		RequiresAutentication: false,
	},
	{
		Uri: "/usuarios/{usuarioId}",
		Method: http.MethodGet,
		Func: controllers.GetUser,
		RequiresAutentication: false,
	},
	{
		Uri: "/usuarios/{usuarioId}",
		Method: http.MethodPut,
		Func: controllers.UpdateUser,
		RequiresAutentication: false,
	},
	{
		Uri: "/usuarios/{usuarioId}",
		Method: http.MethodDelete,
		Func: controllers.DeleteUser,
		RequiresAutentication: false,
	},
}