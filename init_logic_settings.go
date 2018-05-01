package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type (
	Account struct {
		APIKey    string
		APISecret string
	}

	RuleConfiguration struct {
		ID                    int64
		Interval              int64
		MaximumVolume         int64
		TransactionVolume     int64
		VarianceOfTransaction float64
		BidPriceStepDown      float64
		MinimumBid            float64
		Enabled               bool
	}
)

func updateAccounts(c echo.Context) error {
	id, _ := strconv.ParseInt(c.FormValue("id"), 10, 64)
	if id == 1 {
		bot.accountOne.APIKey = c.FormValue("key")
		bot.accountOne.APISecret = c.FormValue("secret")
	} else if id == 2 {
		bot.accountTwo.APIKey = c.FormValue("key")
		bot.accountTwo.APISecret = c.FormValue("secret")
	} else {
		return jsonBadRequest(c, "error no such account.")
	}
	return jsonSuccess(c, echo.Map{
		"account": id,
		"key":     c.FormValue("key"),
	})
}

func updateSettings(c echo.Context) error {
	id, _ := strconv.ParseInt(c.FormValue("id"), 10, 64)
	interval, _ := strconv.ParseInt(c.FormValue("interval"), 10, 64)
	maximumVolume, _ := strconv.ParseInt(c.FormValue("maximumVolume"), 10, 64)
	transactionVolume, _ := strconv.ParseInt(c.FormValue("transactionVolume"), 10, 64)
	variance, _ := strconv.ParseFloat(c.FormValue("variance"), 64)
	stepDownPrice, _ := strconv.ParseFloat(c.FormValue("stepDownPrice"), 64)
	minimumBid, _ := strconv.ParseFloat(c.FormValue("minimumBid"), 64)

	if id == 1 {
		bot.ruleOne.Enabled = true
		bot.ruleOne.Interval = interval
		bot.ruleOne.MaximumVolume = maximumVolume
		bot.ruleOne.TransactionVolume = transactionVolume
		bot.ruleOne.VarianceOfTransaction = variance
		bot.ruleOne.BidPriceStepDown = stepDownPrice
		bot.ruleOne.MinimumBid = minimumBid
	} else if id == 2 {
		bot.ruleTwo.Enabled = true
		bot.ruleTwo.Interval = interval
		bot.ruleTwo.MaximumVolume = maximumVolume
		bot.ruleTwo.TransactionVolume = transactionVolume
		bot.ruleTwo.VarianceOfTransaction = variance
		bot.ruleTwo.BidPriceStepDown = stepDownPrice
		bot.ruleTwo.MinimumBid = minimumBid
	} else {
		return jsonBadRequest(c, "error no such account.")
	}

	return jsonSuccess(c, echo.Map{
		"account": id,
	})
}

func jsonSuccess(c echo.Context, o echo.Map) error {
	o["success"] = true
	return c.JSON(http.StatusOK, o)
}

func jsonBadRequest(c echo.Context, i interface{}) error {
	return c.JSON(http.StatusBadRequest, echo.Map{
		"success": false,
		"message": i,
	})
}

func jsonServerError(c echo.Context, i interface{}) error {
	return c.JSON(http.StatusInternalServerError, echo.Map{
		"success": false,
		"message": i,
	})
}
