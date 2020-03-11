// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// AlertCollectionUpdateAlertsRequestParameter undocumented
type AlertCollectionUpdateAlertsRequestParameter struct {
	// Value undocumented
	Value []Alert `json:"value,omitempty"`
}

//
type AlertCollectionUpdateAlertsRequestBuilder struct{ BaseRequestBuilder }

// UpdateAlerts action undocumented
func (b *SecurityAlertsCollectionRequestBuilder) UpdateAlerts(reqObj *AlertCollectionUpdateAlertsRequestParameter) *AlertCollectionUpdateAlertsRequestBuilder {
	bb := &AlertCollectionUpdateAlertsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/updateAlerts"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type AlertCollectionUpdateAlertsRequest struct{ BaseRequest }

//
func (b *AlertCollectionUpdateAlertsRequestBuilder) Request() *AlertCollectionUpdateAlertsRequest {
	return &AlertCollectionUpdateAlertsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *AlertCollectionUpdateAlertsRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]Alert, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []Alert
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []Alert
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

//
func (r *AlertCollectionUpdateAlertsRequest) Post(ctx context.Context) ([]Alert, error) {
	return r.Paging(ctx, "POST", "", r.requestObject)
}
