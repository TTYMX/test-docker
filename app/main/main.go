
package main

import (
	"net/http"

    "app/api"


    //导入echo包
	"github.com/labstack/echo"

)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "mengxianfan hello welcome to my docker go")
	})


    e.GET("/hello", api.Helloworld)
	e.Start(":8080")

    e.Logger.Fatal(e.Start(":1323"))
}
