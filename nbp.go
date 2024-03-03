package nbp

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/shopspring/decimal"
	"net/http"
	"strconv"
)

type Table byte

const (
	TableA Table = 'A'
	TableB Table = 'B'
	TableC Table = 'C'
)

const (
	urlCurrentRates      = "http://api.nbp.pl/api/exchangerates/tables/{table}/"
	urlNLatestRates      = "http://api.nbp.pl/api/exchangerates/tables/{table}/last/{topCount}/"
	urlTodayRates        = "http://api.nbp.pl/api/exchangerates/tables/{table}/today/"
	urlAtDateRates       = "http://api.nbp.pl/api/exchangerates/tables/{table}/{date}/"
	urlBetweenDatesRates = "http://api.nbp.pl/api/exchangerates/tables/{table}/{startDate}/{endDate}/"
)

type TableResult struct {
	Table         string         `json:"table"`
	No            string         `json:"no"`
	TradingDate   string         `json:"tradingDate"`
	EffectiveDate string         `json:"effectiveDate"`
	Rates         []ExchangeRate `json:"rates"`
}

type ExchangeRate struct {
	Currency string              `json:"currency"`
	Code     string              `json:"code"`
	Bid      decimal.NullDecimal `json:"bid"`
	Ask      decimal.NullDecimal `json:"ask"`
	Mid      decimal.NullDecimal `json:"mid"`
}

type Client struct {
	c *resty.Client
}

func New() *Client {
	return &Client{
		c: resty.New(),
	}
}

func (c *Client) GetNLastRates(ctx context.Context, table Table, n uint64) ([]*TableResult, error) {
	var results []*TableResult

	resp, err := c.c.R().
		SetContext(ctx).
		SetPathParam("table", string(table)).
		SetPathParam("topCount", strconv.FormatUint(n, 10)).
		SetResult(&results).
		Get(urlNLatestRates)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("request error: %s", resp.Status())
	}

	return results, nil
}

func (c *Client) GetCurrentRates(ctx context.Context, table Table) (*TableResult, error) {
	var results []*TableResult

	_, err := c.c.R().
		SetContext(ctx).
		SetContext(ctx).
		SetPathParam("table", string(table)).
		SetResult(&results).
		Get(urlCurrentRates)
	if err != nil {
		return nil, err
	}

	return results[0], nil
}
