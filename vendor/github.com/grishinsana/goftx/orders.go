package goftx

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/grishinsana/goftx/models"
	"github.com/pkg/errors"
)

const (
	apiGetOpenOrders     = "/orders"
	apiGetOrderStatus    = "/orders/%d"
	apiGetOrdersHistory  = "/orders/history"
	apiGetTriggerOrders  = "/conditional_orders"
	apiGetOrderTriggers  = "/conditional_orders/%d/triggers"
	apiPlaceTriggerOrder = "/conditional_orders"
	apiPlaceOrder        = "/orders"
	apiCancelOrders      = "/orders"
)

type Orders struct {
	client *Client
}

func (o *Orders) GetOpenOrders(market string) ([]*models.Order, error) {
	request, err := o.client.prepareRequest(Request{
		Auth:   true,
		Method: http.MethodGet,
		URL:    fmt.Sprintf("%s%s", apiUrl, apiGetOpenOrders),
		Params: map[string]string{
			"market": market,
		},
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	response, err := o.client.do(request)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var result []*models.Order
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return result, nil
}

func (o *Orders) GetOrderStatus(orderID int64) (*models.Order, error) {
	request, err := o.client.prepareRequest(Request{
		Auth:   true,
		Method: http.MethodGet,
		URL:    fmt.Sprintf("%s%s", apiUrl, fmt.Sprintf(apiGetOrderStatus, orderID)),
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	response, err := o.client.do(request)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var result *models.Order
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return result, nil
}

func (o *Orders) GetOrdersHistory(params *models.GetOrdersHistoryParams) ([]*models.Order, error) {
	queryParams, err := PrepareQueryParams(params)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	request, err := o.client.prepareRequest(Request{
		Auth:   true,
		Method: http.MethodGet,
		URL:    fmt.Sprintf("%s%s", apiUrl, apiGetOrdersHistory),
		Params: queryParams,
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	response, err := o.client.do(request)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var result []*models.Order
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return result, nil
}

func (o *Orders) GetOpenTriggerOrders(params *models.GetOpenTriggerOrdersParams) ([]*models.TriggerOrder, error) {
	queryParams, err := PrepareQueryParams(params)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	request, err := o.client.prepareRequest(Request{
		Auth:   true,
		Method: http.MethodGet,
		URL:    fmt.Sprintf("%s%s", apiUrl, apiGetTriggerOrders),
		Params: queryParams,
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	response, err := o.client.do(request)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var result []*models.TriggerOrder
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return result, nil
}

func (o *Orders) GetOrderTriggers(orderID int64) ([]*models.Trigger, error) {
	request, err := o.client.prepareRequest(Request{
		Auth:   true,
		Method: http.MethodGet,
		URL:    fmt.Sprintf("%s%s", apiUrl, fmt.Sprintf(apiGetOrderTriggers, orderID)),
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	response, err := o.client.do(request)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var result []*models.Trigger
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return result, nil
}

func (o *Orders) PlaceOrder(orderParams models.PlaceOrderParams) (*models.Order, error) {
	body, err := json.Marshal(orderParams)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	request, err := o.client.prepareRequest(Request{
		Auth:   true,
		Method: http.MethodPost,
		URL:    fmt.Sprintf("%s%s", apiUrl, apiPlaceOrder),
		Body:   body,
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	response, err := o.client.do(request)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var result *models.Order
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return result, nil
}

func (o *Orders) PlaceTriggerOrder(orderParams interface{}) (*models.TriggerOrder, error) {
	body, err := json.Marshal(orderParams)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	request, err := o.client.prepareRequest(Request{
		Auth:   true,
		Method: http.MethodPost,
		URL:    fmt.Sprintf("%s%s", apiUrl, apiPlaceTriggerOrder),
		Body:   body,
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	response, err := o.client.do(request)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var result *models.TriggerOrder
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return result, nil
}

func (o *Orders) CancelAllOrders(market string) error {
	return o.cancelOrders(struct {
		Market string `json:"market"`
	}{
		Market: market,
	})
}

func (o *Orders) CancelAllLimitOrders(market string) error {
	return o.cancelOrders(struct {
		Market          string `json:"market"`
		LimitOrdersOnly bool   `json:"limitOrdersOnly"`
	}{
		Market:          market,
		LimitOrdersOnly: true,
	})
}

func (o *Orders) CancelAllConditionalOrders(market string) error {
	return o.cancelOrders(struct {
		Market                string `json:"market"`
		ConditionalOrdersOnly bool   `json:"conditionalOrdersOnly"`
	}{
		Market:                market,
		ConditionalOrdersOnly: true,
	})
}

func (o *Orders) cancelOrders(req interface{}) error {
	body, err := json.Marshal(req)

	if err != nil {
		return errors.WithStack(err)
	}

	request, err := o.client.prepareRequest(Request{
		Auth:   true,
		Method: http.MethodDelete,
		URL:    fmt.Sprintf("%s%s", apiUrl, apiCancelOrders),
		Body:   body,
	})
	if err != nil {
		return errors.WithStack(err)
	}

	_, err = o.client.do(request)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
