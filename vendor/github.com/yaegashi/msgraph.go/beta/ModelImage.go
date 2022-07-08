// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// Image undocumented
type Image struct {
	// Object is the base model of Image
	Object
	// Height undocumented
	Height *int `json:"height,omitempty"`
	// Width undocumented
	Width *int `json:"width,omitempty"`
}

// ImageInfo undocumented
type ImageInfo struct {
	// Object is the base model of ImageInfo
	Object
	// IconURL undocumented
	IconURL *string `json:"iconUrl,omitempty"`
	// AlternativeText undocumented
	AlternativeText *string `json:"alternativeText,omitempty"`
	// AlternateText undocumented
	AlternateText *string `json:"alternateText,omitempty"`
	// AddImageQuery undocumented
	AddImageQuery *bool `json:"addImageQuery,omitempty"`
}
