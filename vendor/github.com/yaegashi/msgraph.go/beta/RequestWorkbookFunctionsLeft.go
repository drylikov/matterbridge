// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsLeftRequestBuilder struct{ BaseRequestBuilder }

// Left action undocumented
func (b *WorkbookFunctionsRequestBuilder) Left(reqObj *WorkbookFunctionsLeftRequestParameter) *WorkbookFunctionsLeftRequestBuilder {
	bb := &WorkbookFunctionsLeftRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/left"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsLeftRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsLeftRequestBuilder) Request() *WorkbookFunctionsLeftRequest {
	return &WorkbookFunctionsLeftRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsLeftRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
