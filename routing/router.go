package routing

import (
	"net/http"

	"github.com/caeret/golang-web-skeleton/service"

	"github.com/caeret/golang-web-skeleton/app"
	routing "github.com/go-ozzo/ozzo-routing"
)

func Serve(logger app.Logger) {
	userService := &service.UserService{}
	userCtl := &UserCTL{
		userService: userService,
	}

	router := routing.New()
	router.Use(
		app.RoutePrepare(logger),
	)
	router.Post("/users", userCtl.CreateUser)
	logger.Info("serve http server.")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}
