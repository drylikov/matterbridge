// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// ConnectionOperation undocumented
type ConnectionOperation struct {
	// Entity is the base model of ConnectionOperation
	Entity
	// Status undocumented
	Status *ConnectionOperationStatus `json:"status,omitempty"`
	// Error undocumented
	Error *ErrorDetail `json:"error,omitempty"`
}
