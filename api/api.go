package api

import (
	"encoding/json"
	"fmt"
	"gofolio/services"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller interface {
	Routes(r *echo.Group)
}

func InitRoutes(e *echo.Group) {

	api := e.Group("")
	api.GET("/stock", getStockInfo)
	api.GET("/search", search)
	api.POST("/save", save)
}

func getStockInfo(c echo.Context) error {

	// temp
	baseRoute := "https://query1.finance.yahoo.com/v8/finance/chart"
	stockCode := "AAPL"
	qParam := "formatted=true&crumb=eAoBO3AGKQH&lang=en-US&region=US&includeAdjustedClose=true&interval=1d&period1=1696896000&period2=1697328000&events=capitalGain%7Cdiv%7Csplit&useYfid=true&corsDomain=finance.yahoo.com"
	endpoint := fmt.Sprintf("%s/%s?%s", baseRoute, stockCode, qParam)

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var stockRsp services.StockResponse
	err = json.Unmarshal(bodyBytes, &stockRsp)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, stockRsp)
}

func search(c echo.Context) error {
	baseRoute := "https://query2.finance.yahoo.com/v1/finance/search?quotesCount=8&newsCount=0&listsCount=0&q="
	searchParam := c.QueryParam("q")
	endpoint := fmt.Sprintf("%s%s", baseRoute, searchParam)

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var searchRsp services.SearchResponse
	err = json.Unmarshal(bodyBytes, &searchRsp)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, searchRsp)
}

func save(c echo.Context) error {
	sItem := new(services.StockItem)
	if err := c.Bind(sItem); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	id, _ := services.InsertStock(*sItem)

	return c.JSON(http.StatusOK, id)
	// insertID, err := services.InsertUser(*user)
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, err)
	// }
	// return c.JSON(http.StatusOK, insertID)
}
