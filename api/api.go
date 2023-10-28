package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type StockResponse struct {
	Chart Chart `json:"chart"`
}

type Chart struct {
	Result []Result `json:"result"`
	Error  string   `json:"error"`
}

type Result struct {
	Meta Meta `json:"meta"`
}

type Meta struct {
	Currency           string  `json:"currency"`
	Symbol             string  `json:"symbol"`
	RegularMarketPrice float32 `json:"regularMarketPrice"`
}

type SearchResponse struct {
	Quotes []Quote `json:"quotes"`
}

type Quote struct {
	Exchange  string `json:"exchange"`
	Symbol    string `json:"symbol"`
	QuoteType string `json:"quoteType"`
	ShortName string `json:"shortname"`
}

type Controller interface {
	Routes(r *echo.Group)
}

func InitRoutes(e *echo.Group) {

	api := e.Group("")
	api.GET("/stock", getStockInfo)
	api.GET("/search", search)

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

	var stockRsp StockResponse
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

	var searchRsp SearchResponse
	err = json.Unmarshal(bodyBytes, &searchRsp)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, searchRsp)
}
