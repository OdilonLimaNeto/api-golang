package routes

import (
	"api/src/controllers"
	"net/http"
)

var usersRoutes = []Route{
	{
		URI:                    "/users",
		Method:                 http.MethodPost,
		HandlerFunc:            controllers.Create,
		requiredAuthentication: false,
	},

	{
		URI:                    "/users",
		Method:                 http.MethodGet,
		HandlerFunc:            controllers.List,
		requiredAuthentication: false,
	},

	{
		URI:                    "/users/{id}",
		Method:                 http.MethodGet,
		HandlerFunc:            controllers.Get,
		requiredAuthentication: false,
	},

	{
		URI:                    "/users/{id}",
		Method:                 http.MethodPut,
		HandlerFunc:            controllers.Update,
		requiredAuthentication: false,
	},

	{
		URI:                    "/users/{id}",
		Method:                 http.MethodDelete,
		HandlerFunc:            controllers.Delete,
		requiredAuthentication: false,
	},
}
