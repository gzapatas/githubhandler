package releases

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
}

func (o Handler) NewRelease(c echo.Context) error {
	data, _ := ioutil.ReadAll(c.Request().Body)
	info := map[string]interface{}{}
	json.Unmarshal(data, &info)
	indent, _ := json.MarshalIndent(info, "", "\t")
	fmt.Println(string(indent))
	return c.JSON(http.StatusCreated, nil)
}
