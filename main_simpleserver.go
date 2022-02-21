package main

import (
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.Static("/", "static")
	e.HTTPErrorHandler = notfoundHandler(e)
	go e.Logger.Fatal(e.Start(":1325"))
}

func notfoundHandler(e *echo.Echo) func(error, echo.Context) {
	return func(err error, c echo.Context) {
		code, ok := errorCode(err)
		if !ok || code != 404 {
			e.DefaultHTTPErrorHandler(err, c)
			return
		}
		// read .html
		if er := tryRead(c, http.StatusOK, "static", c.Request().URL.Path+".html"); er == nil {
			return
		}
		// 404.html
		if er := tryRead(c, http.StatusNotFound, "static", "404.html"); er == nil {
			return
		}

		e.DefaultHTTPErrorHandler(err, c)
	}
}

func errorCode(err error) (int, bool) {
	ee, ok := err.(*echo.HTTPError)
	if !ok {
		return -1, false
	}

	return ee.Code, true
}

func tryRead(c echo.Context, code int, elem ...string) error {
	path := filepath.Join(elem...)
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	c.HTML(code, string(bytes))
	return nil
}
