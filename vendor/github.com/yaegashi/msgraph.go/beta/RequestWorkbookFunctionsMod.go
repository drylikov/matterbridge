// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsModRequestBuilder struct{ BaseRequestBuilder }

// Mod action undocumented
func (b *WorkbookFunctionsRequestBuilder) Mod(reqObj *WorkbookFunctionsModRequestParameter) *WorkbookFunctionsModRequestBuilder {
	bb := &WorkbookFunctionsModRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/mod"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsModRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsModRequestBuilder) Request() *WorkbookFunctionsModRequest {
	return &WorkbookFunctionsModRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsModRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
