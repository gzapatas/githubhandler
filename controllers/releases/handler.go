package releases

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
}

func (o Handler) NewRelease(c echo.Context) error {
	data, _ := ioutil.ReadAll(c.Request().Body)
	fmt.Println(string(data))
	return c.JSON(http.StatusCreated, nil)
}
