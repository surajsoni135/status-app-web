package route

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	api "mysms4you.com/app/src/repo/route/api"
	ui "mysms4you.com/app/src/repo/route/ui"
)

func Handler() http.Handler {
	router := httprouter.New()

	/**
	 * UI Routes
	 */
	router.GET("/", ui.GetHome())
	router.GET("/img/:param", ui.ServerStaticResource())
	router.GET("/css/:param", ui.ServerStaticResource())
	router.GET("/scripts/:param", ui.ServerStaticResource())

	/**
	 * APIs Routes
	 */
	router.GET("/api/v1/main.php", api.GetV1Main())

	return router
}
