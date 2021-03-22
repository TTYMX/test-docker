package api

import (
    "net/http"
    "github.com/labstack/echo"
)

func Helloworld(c echo.Context) error {
    return c.JSON(http.StatusOK, "hello world fan")
}
