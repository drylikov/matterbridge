// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// User is navigation property
func (b *AadUserConversationMemberRequestBuilder) User() *UserRequestBuilder {
	bb := &UserRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/user"
	return bb
}