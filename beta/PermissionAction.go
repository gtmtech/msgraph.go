// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// PermissionGrantRequestParameter undocumented
type PermissionGrantRequestParameter struct {
	// Roles undocumented
	Roles []string `json:"roles,omitempty"`
	// Recipients undocumented
	Recipients []DriveRecipient `json:"recipients,omitempty"`
}

//
type PermissionGrantRequestBuilder struct{ BaseRequestBuilder }

// Grant action undocumented
func (b *PermissionRequestBuilder) Grant(reqObj *PermissionGrantRequestParameter) *PermissionGrantRequestBuilder {
	bb := &PermissionGrantRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/grant"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type PermissionGrantRequest struct{ BaseRequest }

//
func (b *PermissionGrantRequestBuilder) Request() *PermissionGrantRequest {
	return &PermissionGrantRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *PermissionGrantRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([][]Permission, error) {
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
	var values [][]Permission
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
			value  [][]Permission
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
func (r *PermissionGrantRequest) Post(ctx context.Context) ([][]Permission, error) {
	return r.Paging(ctx, "POST", "", r.requestObject)
}
