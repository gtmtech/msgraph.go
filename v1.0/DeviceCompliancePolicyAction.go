// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// DeviceCompliancePolicyAssignRequestParameter undocumented
type DeviceCompliancePolicyAssignRequestParameter struct {
	// Assignments undocumented
	Assignments []DeviceCompliancePolicyAssignment `json:"assignments,omitempty"`
}

// DeviceCompliancePolicyScheduleActionsForRulesRequestParameter undocumented
type DeviceCompliancePolicyScheduleActionsForRulesRequestParameter struct {
	// DeviceComplianceScheduledActionForRules undocumented
	DeviceComplianceScheduledActionForRules []DeviceComplianceScheduledActionForRule `json:"deviceComplianceScheduledActionForRules,omitempty"`
}

//
type DeviceCompliancePolicyAssignRequestBuilder struct{ BaseRequestBuilder }

// Assign action undocumented
func (b *DeviceCompliancePolicyRequestBuilder) Assign(reqObj *DeviceCompliancePolicyAssignRequestParameter) *DeviceCompliancePolicyAssignRequestBuilder {
	bb := &DeviceCompliancePolicyAssignRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/assign"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DeviceCompliancePolicyAssignRequest struct{ BaseRequest }

//
func (b *DeviceCompliancePolicyAssignRequestBuilder) Request() *DeviceCompliancePolicyAssignRequest {
	return &DeviceCompliancePolicyAssignRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DeviceCompliancePolicyAssignRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]DeviceCompliancePolicyAssignment, error) {
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
	var values []DeviceCompliancePolicyAssignment
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
			value  []DeviceCompliancePolicyAssignment
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
func (r *DeviceCompliancePolicyAssignRequest) Post(ctx context.Context) ([]DeviceCompliancePolicyAssignment, error) {
	return r.Paging(ctx, "POST", "", r.requestObject)
}

//
type DeviceCompliancePolicyScheduleActionsForRulesRequestBuilder struct{ BaseRequestBuilder }

// ScheduleActionsForRules action undocumented
func (b *DeviceCompliancePolicyRequestBuilder) ScheduleActionsForRules(reqObj *DeviceCompliancePolicyScheduleActionsForRulesRequestParameter) *DeviceCompliancePolicyScheduleActionsForRulesRequestBuilder {
	bb := &DeviceCompliancePolicyScheduleActionsForRulesRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/scheduleActionsForRules"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DeviceCompliancePolicyScheduleActionsForRulesRequest struct{ BaseRequest }

//
func (b *DeviceCompliancePolicyScheduleActionsForRulesRequestBuilder) Request() *DeviceCompliancePolicyScheduleActionsForRulesRequest {
	return &DeviceCompliancePolicyScheduleActionsForRulesRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DeviceCompliancePolicyScheduleActionsForRulesRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}
