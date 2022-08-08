package routes

import (
	"api/src/controllers"
)

var usersRoutes = []Routes{
	{
		URI:                    "/users",
		Method:                 "POST",
		HandlerFunc:            controllers.Create,
		requiredAuthentication: false,
	},

	{
		URI:                    "/users",
		Method:                 "GET",
		HandlerFunc:            controllers.List,
		requiredAuthentication: false,
	},

	{
		URI:                    "/users/{id}",
		Method:                 "GET",
		HandlerFunc:            controllers.Get,
		requiredAuthentication: false,
	},

	{
		URI:                    "/users/{id}",
		Method:                 "PUT",
		HandlerFunc:            controllers.Update,
		requiredAuthentication: false,
	},

	{
		URI:                    "/users/{id}",
		Method:                 "DELETE",
		HandlerFunc:            controllers.Delete,
		requiredAuthentication: false,
	},
}
