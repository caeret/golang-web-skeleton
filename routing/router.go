package routing

import (
	"net/http"

	"github.com/caeret/golang-web-skeleton/app"
	routing "github.com/go-ozzo/ozzo-routing"
)

func Serve(logger app.Logger, container app.Container) {
	router := routing.New()
	router.Use(
		app.RoutePrepare(logger, container),
	)
	router.Get("/", func(c *routing.Context) error {
		panic("")
		return c.Write(app.GetRequestScope(c).RequestID())
	})
	logger.Info("serve http server.")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}
