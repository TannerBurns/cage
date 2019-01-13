package routes

import (
	"net/http"

	"../controllers"
	"../models"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() (*mux.Router, *models.Logger) {
	controller := &controllers.Controller{Name: "API.Controller"}
	controller.Manager = models.NewManager()
	controller.Logger = models.NewLogger()
	api_route := "/house/api/v1"

	AuthRoutes := Routes{
		Route{
			"Index",
			"GET",
			api_route + "/status",
			controller.Index,
		},
		Route{
			"Authentication",
			"POST",
			api_route + "/authentication/token",
			controller.GetToken,
		},
	}
	PlayerRoutes := Routes{
		Route{
			"Search",
			"GET",
			api_route + "/search",
			controllers.AuthenticationMiddleware(controller.Search),
		},
		Route{
			"CreatePlayer",
			"POST",
			api_route + "/player",
			controllers.AuthenticationMiddleware(controller.CreatePlayer),
		},
		Route{
			"ReadPlayer",
			"GET",
			api_route + "/player/{id}",
			controllers.AuthenticationMiddleware(controller.GetPlayer),
		},
		Route{
			"UpdatePlayer",
			"PUT",
			api_route + "/player/{id}",
			controllers.AuthenticationMiddleware(controller.UpdatePlayer),
		},
		Route{
			"CreateMembership",
			"POST",
			api_route + "/player/{id}/membership",
			controllers.AuthenticationMiddleware(controller.CreateMembership),
		},
		Route{
			"ReadMembership",
			"GET",
			api_route + "/player/{id}/membership",
			controllers.AuthenticationMiddleware(controller.GetMembership),
		},
		Route{
			"AddPlayTime",
			"PUT",
			api_route + "/player/{id}/membership",
			controllers.AuthenticationMiddleware(controller.AddPlayTime),
		},
		Route{
			"CheckIn",
			"GET",
			api_route + "/player/{id}/checkin",
			controllers.AuthenticationMiddleware(controller.CheckIn),
		},
		Route{
			"CheckOut",
			"Get",
			api_route + "/player/{id}/checkout",
			controllers.AuthenticationMiddleware(controller.CheckOut),
		},
	}
	EmployeeRoutes := Routes{
		Route{
			"CreateEmployee",
			"POST",
			api_route + "/employee",
			controllers.AuthenticationMiddleware(controller.CreateEmployee),
		},
		Route{
			"ReadEmployee",
			"GET",
			api_route + "/employee/{id}",
			controllers.AuthenticationMiddleware(controller.GetEmployee),
		},
		Route{
			"UpdateEmployee",
			"PUT",
			api_route + "/employee/{id}",
			controllers.AuthenticationMiddleware(controller.UpdateEmployee),
		},
		Route{
			"CreateLogin",
			"POST",
			api_route + "/employee/{id}/login",
			controllers.AuthenticationMiddleware(controller.CreateLogin),
		},
		/*Route{
			"UpdatePassword",
			"PUT",
			api_route + "/employee/{id}/login",
			controllers.AuthenticationMiddleware(controller.UpdatePassword),
		},*/
		Route{
			"AddRole",
			"POST",
			api_route + "/employee/{id}/roles",
			controllers.AuthenticationMiddleware(controller.AddRole),
		},
		Route{
			"GetRoles",
			"GET",
			api_route + "/employee/{id}/roles",
			controllers.AuthenticationMiddleware(controller.GetRoles),
		},
		Route{
			"GetRoster",
			"GET",
			api_route + "/roster",
			controllers.AuthenticationMiddleware(controller.GetRoster),
		},
	}

	Routes := []Routes{AuthRoutes, PlayerRoutes, EmployeeRoutes}

	router := mux.NewRouter().StrictSlash(true)
	for _, routes := range Routes {
		for _, route := range routes {
			var handler http.Handler
			handler = route.HandlerFunc

			router.
				Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				Handler(handler)
		}
	}
	return router, controller.Logger
}
