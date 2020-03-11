// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// RoleScopeTagCollectionGetRoleScopeTagsByIDRequestParameter undocumented
type RoleScopeTagCollectionGetRoleScopeTagsByIDRequestParameter struct {
	// RoleScopeTagIDs undocumented
	RoleScopeTagIDs []string `json:"roleScopeTagIds,omitempty"`
}

// RoleScopeTagAssignRequestParameter undocumented
type RoleScopeTagAssignRequestParameter struct {
	// Assignments undocumented
	Assignments []RoleScopeTagAutoAssignment `json:"assignments,omitempty"`
}

//
type RoleScopeTagCollectionGetRoleScopeTagsByIDRequestBuilder struct{ BaseRequestBuilder }

// GetRoleScopeTagsByID action undocumented
func (b *DeviceAndAppManagementRoleAssignmentRoleScopeTagsCollectionRequestBuilder) GetRoleScopeTagsByID(reqObj *RoleScopeTagCollectionGetRoleScopeTagsByIDRequestParameter) *RoleScopeTagCollectionGetRoleScopeTagsByIDRequestBuilder {
	bb := &RoleScopeTagCollectionGetRoleScopeTagsByIDRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/getRoleScopeTagsById"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// GetRoleScopeTagsByID action undocumented
func (b *DeviceManagementRoleScopeTagsCollectionRequestBuilder) GetRoleScopeTagsByID(reqObj *RoleScopeTagCollectionGetRoleScopeTagsByIDRequestParameter) *RoleScopeTagCollectionGetRoleScopeTagsByIDRequestBuilder {
	bb := &RoleScopeTagCollectionGetRoleScopeTagsByIDRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/getRoleScopeTagsById"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type RoleScopeTagCollectionGetRoleScopeTagsByIDRequest struct{ BaseRequest }

//
func (b *RoleScopeTagCollectionGetRoleScopeTagsByIDRequestBuilder) Request() *RoleScopeTagCollectionGetRoleScopeTagsByIDRequest {
	return &RoleScopeTagCollectionGetRoleScopeTagsByIDRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *RoleScopeTagCollectionGetRoleScopeTagsByIDRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([][]RoleScopeTag, error) {
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
	var values [][]RoleScopeTag
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
			value  [][]RoleScopeTag
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
func (r *RoleScopeTagCollectionGetRoleScopeTagsByIDRequest) Post(ctx context.Context) ([][]RoleScopeTag, error) {
	return r.Paging(ctx, "POST", "", r.requestObject)
}

//
type RoleScopeTagAssignRequestBuilder struct{ BaseRequestBuilder }

// Assign action undocumented
func (b *RoleScopeTagRequestBuilder) Assign(reqObj *RoleScopeTagAssignRequestParameter) *RoleScopeTagAssignRequestBuilder {
	bb := &RoleScopeTagAssignRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/assign"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type RoleScopeTagAssignRequest struct{ BaseRequest }

//
func (b *RoleScopeTagAssignRequestBuilder) Request() *RoleScopeTagAssignRequest {
	return &RoleScopeTagAssignRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *RoleScopeTagAssignRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([][]RoleScopeTagAutoAssignment, error) {
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
	var values [][]RoleScopeTagAutoAssignment
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
			value  [][]RoleScopeTagAutoAssignment
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
func (r *RoleScopeTagAssignRequest) Post(ctx context.Context) ([][]RoleScopeTagAutoAssignment, error) {
	return r.Paging(ctx, "POST", "", r.requestObject)
}
