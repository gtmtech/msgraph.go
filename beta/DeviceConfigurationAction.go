// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// DeviceConfigurationCollectionHasPayloadLinksRequestParameter undocumented
type DeviceConfigurationCollectionHasPayloadLinksRequestParameter struct {
	// PayloadIDs undocumented
	PayloadIDs []string `json:"payloadIds,omitempty"`
}

// DeviceConfigurationCollectionGetTargetedUsersAndDevicesRequestParameter undocumented
type DeviceConfigurationCollectionGetTargetedUsersAndDevicesRequestParameter struct {
	// DeviceConfigurationIDs undocumented
	DeviceConfigurationIDs []string `json:"deviceConfigurationIds,omitempty"`
}

// DeviceConfigurationAssignRequestParameter undocumented
type DeviceConfigurationAssignRequestParameter struct {
	// DeviceConfigurationGroupAssignments undocumented
	DeviceConfigurationGroupAssignments []DeviceConfigurationGroupAssignment `json:"deviceConfigurationGroupAssignments,omitempty"`
	// Assignments undocumented
	Assignments []DeviceConfigurationAssignment `json:"assignments,omitempty"`
}

// DeviceConfigurationWindowsPrivacyAccessControlsRequestParameter undocumented
type DeviceConfigurationWindowsPrivacyAccessControlsRequestParameter struct {
	// WindowsPrivacyAccessControls undocumented
	WindowsPrivacyAccessControls []WindowsPrivacyDataAccessControlItem `json:"windowsPrivacyAccessControls,omitempty"`
}

// DeviceConfigurationAssignedAccessMultiModeProfilesRequestParameter undocumented
type DeviceConfigurationAssignedAccessMultiModeProfilesRequestParameter struct {
	// AssignedAccessMultiModeProfiles undocumented
	AssignedAccessMultiModeProfiles []WindowsAssignedAccessProfile `json:"assignedAccessMultiModeProfiles,omitempty"`
}

//
type DeviceConfigurationCollectionHasPayloadLinksRequestBuilder struct{ BaseRequestBuilder }

// HasPayloadLinks action undocumented
func (b *DeviceManagementDeviceConfigurationsCollectionRequestBuilder) HasPayloadLinks(reqObj *DeviceConfigurationCollectionHasPayloadLinksRequestParameter) *DeviceConfigurationCollectionHasPayloadLinksRequestBuilder {
	bb := &DeviceConfigurationCollectionHasPayloadLinksRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/hasPayloadLinks"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// HasPayloadLinks action undocumented
func (b *WindowsDomainJoinConfigurationNetworkAccessConfigurationsCollectionRequestBuilder) HasPayloadLinks(reqObj *DeviceConfigurationCollectionHasPayloadLinksRequestParameter) *DeviceConfigurationCollectionHasPayloadLinksRequestBuilder {
	bb := &DeviceConfigurationCollectionHasPayloadLinksRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/hasPayloadLinks"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DeviceConfigurationCollectionHasPayloadLinksRequest struct{ BaseRequest }

//
func (b *DeviceConfigurationCollectionHasPayloadLinksRequestBuilder) Request() *DeviceConfigurationCollectionHasPayloadLinksRequest {
	return &DeviceConfigurationCollectionHasPayloadLinksRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DeviceConfigurationCollectionHasPayloadLinksRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([][]HasPayloadLinkResultItem, error) {
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
	var values [][]HasPayloadLinkResultItem
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
			value  [][]HasPayloadLinkResultItem
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
func (r *DeviceConfigurationCollectionHasPayloadLinksRequest) Post(ctx context.Context) ([][]HasPayloadLinkResultItem, error) {
	return r.Paging(ctx, "POST", "", r.requestObject)
}

//
type DeviceConfigurationCollectionGetTargetedUsersAndDevicesRequestBuilder struct{ BaseRequestBuilder }

// GetTargetedUsersAndDevices action undocumented
func (b *DeviceManagementDeviceConfigurationsCollectionRequestBuilder) GetTargetedUsersAndDevices(reqObj *DeviceConfigurationCollectionGetTargetedUsersAndDevicesRequestParameter) *DeviceConfigurationCollectionGetTargetedUsersAndDevicesRequestBuilder {
	bb := &DeviceConfigurationCollectionGetTargetedUsersAndDevicesRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/getTargetedUsersAndDevices"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// GetTargetedUsersAndDevices action undocumented
func (b *WindowsDomainJoinConfigurationNetworkAccessConfigurationsCollectionRequestBuilder) GetTargetedUsersAndDevices(reqObj *DeviceConfigurationCollectionGetTargetedUsersAndDevicesRequestParameter) *DeviceConfigurationCollectionGetTargetedUsersAndDevicesRequestBuilder {
	bb := &DeviceConfigurationCollectionGetTargetedUsersAndDevicesRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/getTargetedUsersAndDevices"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DeviceConfigurationCollectionGetTargetedUsersAndDevicesRequest struct{ BaseRequest }

//
func (b *DeviceConfigurationCollectionGetTargetedUsersAndDevicesRequestBuilder) Request() *DeviceConfigurationCollectionGetTargetedUsersAndDevicesRequest {
	return &DeviceConfigurationCollectionGetTargetedUsersAndDevicesRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DeviceConfigurationCollectionGetTargetedUsersAndDevicesRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([][]DeviceConfigurationTargetedUserAndDevice, error) {
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
	var values [][]DeviceConfigurationTargetedUserAndDevice
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
			value  [][]DeviceConfigurationTargetedUserAndDevice
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
func (r *DeviceConfigurationCollectionGetTargetedUsersAndDevicesRequest) Post(ctx context.Context) ([][]DeviceConfigurationTargetedUserAndDevice, error) {
	return r.Paging(ctx, "POST", "", r.requestObject)
}

//
type DeviceConfigurationAssignRequestBuilder struct{ BaseRequestBuilder }

// Assign action undocumented
func (b *DeviceConfigurationRequestBuilder) Assign(reqObj *DeviceConfigurationAssignRequestParameter) *DeviceConfigurationAssignRequestBuilder {
	bb := &DeviceConfigurationAssignRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/assign"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DeviceConfigurationAssignRequest struct{ BaseRequest }

//
func (b *DeviceConfigurationAssignRequestBuilder) Request() *DeviceConfigurationAssignRequest {
	return &DeviceConfigurationAssignRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DeviceConfigurationAssignRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([][]DeviceConfigurationAssignment, error) {
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
	var values [][]DeviceConfigurationAssignment
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
			value  [][]DeviceConfigurationAssignment
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
func (r *DeviceConfigurationAssignRequest) Post(ctx context.Context) ([][]DeviceConfigurationAssignment, error) {
	return r.Paging(ctx, "POST", "", r.requestObject)
}

//
type DeviceConfigurationWindowsPrivacyAccessControlsRequestBuilder struct{ BaseRequestBuilder }

// WindowsPrivacyAccessControls action undocumented
func (b *DeviceConfigurationRequestBuilder) WindowsPrivacyAccessControls(reqObj *DeviceConfigurationWindowsPrivacyAccessControlsRequestParameter) *DeviceConfigurationWindowsPrivacyAccessControlsRequestBuilder {
	bb := &DeviceConfigurationWindowsPrivacyAccessControlsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/windowsPrivacyAccessControls"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DeviceConfigurationWindowsPrivacyAccessControlsRequest struct{ BaseRequest }

//
func (b *DeviceConfigurationWindowsPrivacyAccessControlsRequestBuilder) Request() *DeviceConfigurationWindowsPrivacyAccessControlsRequest {
	return &DeviceConfigurationWindowsPrivacyAccessControlsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DeviceConfigurationWindowsPrivacyAccessControlsRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type DeviceConfigurationAssignedAccessMultiModeProfilesRequestBuilder struct{ BaseRequestBuilder }

// AssignedAccessMultiModeProfiles action undocumented
func (b *DeviceConfigurationRequestBuilder) AssignedAccessMultiModeProfiles(reqObj *DeviceConfigurationAssignedAccessMultiModeProfilesRequestParameter) *DeviceConfigurationAssignedAccessMultiModeProfilesRequestBuilder {
	bb := &DeviceConfigurationAssignedAccessMultiModeProfilesRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/assignedAccessMultiModeProfiles"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DeviceConfigurationAssignedAccessMultiModeProfilesRequest struct{ BaseRequest }

//
func (b *DeviceConfigurationAssignedAccessMultiModeProfilesRequestBuilder) Request() *DeviceConfigurationAssignedAccessMultiModeProfilesRequest {
	return &DeviceConfigurationAssignedAccessMultiModeProfilesRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DeviceConfigurationAssignedAccessMultiModeProfilesRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}
