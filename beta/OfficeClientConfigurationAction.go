// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// OfficeClientConfigurationCollectionUpdatePrioritiesRequestParameter undocumented
type OfficeClientConfigurationCollectionUpdatePrioritiesRequestParameter struct {
	// OfficeConfigurationPolicyIDs undocumented
	OfficeConfigurationPolicyIDs []string `json:"officeConfigurationPolicyIds,omitempty"`
	// OfficeConfigurationPriorities undocumented
	OfficeConfigurationPriorities []int `json:"officeConfigurationPriorities,omitempty"`
}

// OfficeClientConfigurationAssignRequestParameter undocumented
type OfficeClientConfigurationAssignRequestParameter struct {
	// OfficeConfigurationAssignments undocumented
	OfficeConfigurationAssignments []OfficeClientConfigurationAssignment `json:"officeConfigurationAssignments,omitempty"`
}

//
type OfficeClientConfigurationCollectionUpdatePrioritiesRequestBuilder struct{ BaseRequestBuilder }

// UpdatePriorities action undocumented
func (b *OfficeConfigurationClientConfigurationsCollectionRequestBuilder) UpdatePriorities(reqObj *OfficeClientConfigurationCollectionUpdatePrioritiesRequestParameter) *OfficeClientConfigurationCollectionUpdatePrioritiesRequestBuilder {
	bb := &OfficeClientConfigurationCollectionUpdatePrioritiesRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/updatePriorities"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type OfficeClientConfigurationCollectionUpdatePrioritiesRequest struct{ BaseRequest }

//
func (b *OfficeClientConfigurationCollectionUpdatePrioritiesRequestBuilder) Request() *OfficeClientConfigurationCollectionUpdatePrioritiesRequest {
	return &OfficeClientConfigurationCollectionUpdatePrioritiesRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *OfficeClientConfigurationCollectionUpdatePrioritiesRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type OfficeClientConfigurationAssignRequestBuilder struct{ BaseRequestBuilder }

// Assign action undocumented
func (b *OfficeClientConfigurationRequestBuilder) Assign(reqObj *OfficeClientConfigurationAssignRequestParameter) *OfficeClientConfigurationAssignRequestBuilder {
	bb := &OfficeClientConfigurationAssignRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/assign"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type OfficeClientConfigurationAssignRequest struct{ BaseRequest }

//
func (b *OfficeClientConfigurationAssignRequestBuilder) Request() *OfficeClientConfigurationAssignRequest {
	return &OfficeClientConfigurationAssignRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *OfficeClientConfigurationAssignRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]OfficeClientConfigurationAssignment, error) {
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
	var values []OfficeClientConfigurationAssignment
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
			value  []OfficeClientConfigurationAssignment
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
func (r *OfficeClientConfigurationAssignRequest) Post(ctx context.Context) ([]OfficeClientConfigurationAssignment, error) {
	return r.Paging(ctx, "POST", "", r.requestObject)
}
