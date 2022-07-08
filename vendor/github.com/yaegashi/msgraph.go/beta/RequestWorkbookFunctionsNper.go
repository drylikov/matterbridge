// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsNperRequestBuilder struct{ BaseRequestBuilder }

// Nper action undocumented
func (b *WorkbookFunctionsRequestBuilder) Nper(reqObj *WorkbookFunctionsNperRequestParameter) *WorkbookFunctionsNperRequestBuilder {
	bb := &WorkbookFunctionsNperRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/nper"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsNperRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsNperRequestBuilder) Request() *WorkbookFunctionsNperRequest {
	return &WorkbookFunctionsNperRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsNperRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}