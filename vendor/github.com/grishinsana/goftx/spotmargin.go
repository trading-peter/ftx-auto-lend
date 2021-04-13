package goftx

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"

	"github.com/grishinsana/goftx/models"
)

const (
	apiGetLendingInfo     = "/spot_margin/lending_info"
	apiGetLendingRates    = "/spot_margin/lending_rates"
	apiSubmitLendingOffer = "/spot_margin/offers"
)

type SpotMargin struct {
	client *Client
}

func (m *SpotMargin) GetLendingInfo() ([]*models.LendingInfo, error) {
	request, err := m.client.prepareRequest(Request{
		Auth:   true,
		Method: http.MethodGet,
		URL:    fmt.Sprintf("%s%s", apiUrl, apiGetLendingInfo),
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	response, err := m.client.do(request)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var result []*models.LendingInfo
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return result, nil
}

func (m *SpotMargin) GetLendingRates() ([]*models.LendingRate, error) {
	request, err := m.client.prepareRequest(Request{
		Auth:   true,
		Method: http.MethodGet,
		URL:    fmt.Sprintf("%s%s", apiUrl, apiGetLendingRates),
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	response, err := m.client.do(request)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var result []*models.LendingRate
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return result, nil
}

func (m *SpotMargin) SubmitLendingOffer(coin string, size decimal.Decimal, rate decimal.Decimal) error {
	body, err := json.Marshal(struct {
		Coin string          `json:"coin"`
		Size decimal.Decimal `json:"size"`
		Rate decimal.Decimal `json:"rate"`
	}{
		Coin: coin,
		Size: size,
		Rate: rate,
	})
	if err != nil {
		return errors.WithStack(err)
	}

	request, err := m.client.prepareRequest(Request{
		Auth:   true,
		Method: http.MethodPost,
		URL:    fmt.Sprintf("%s%s", apiUrl, apiSubmitLendingOffer),
		Body:   body,
	})
	if err != nil {
		return errors.WithStack(err)
	}

	_, err = m.client.do(request)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
