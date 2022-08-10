package routes

import (
	"devbookapi/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:        "/users",
		Method:     http.MethodPost,
		Function:   controllers.InsertUser,
		ShouldAuth: false,
	},
	{
		URI:        "/users",
		Method:     http.MethodGet,
		Function:   controllers.GetUsers,
		ShouldAuth: false,
	},
	{
		URI:        "/users/{userId}",
		Method:     http.MethodGet,
		Function:   controllers.GetUser,
		ShouldAuth: false,
	},
	{
		URI:        "/users/{userId}",
		Method:     http.MethodPut,
		Function:   controllers.UpdateUser,
		ShouldAuth: false,
	},
	{
		URI:        "/users/{userId}",
		Method:     http.MethodDelete,
		Function:   controllers.DeleteUser,
		ShouldAuth: false,
	},
}
