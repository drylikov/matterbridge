// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// Onenote undocumented
type Onenote struct {
	// Entity is the base model of Onenote
	Entity
	// Notebooks undocumented
	Notebooks []Notebook `json:"notebooks,omitempty"`
	// Sections undocumented
	Sections []OnenoteSection `json:"sections,omitempty"`
	// SectionGroups undocumented
	SectionGroups []SectionGroup `json:"sectionGroups,omitempty"`
	// Pages undocumented
	Pages []OnenotePage `json:"pages,omitempty"`
	// Resources undocumented
	Resources []OnenoteResource `json:"resources,omitempty"`
	// Operations undocumented
	Operations []OnenoteOperation `json:"operations,omitempty"`
}

// OnenoteEntityBaseModel undocumented
type OnenoteEntityBaseModel struct {
	// Entity is the base model of OnenoteEntityBaseModel
	Entity
	// Self undocumented
	Self *string `json:"self,omitempty"`
}

// OnenoteEntityHierarchyModel undocumented
type OnenoteEntityHierarchyModel struct {
	// OnenoteEntitySchemaObjectModel is the base model of OnenoteEntityHierarchyModel
	OnenoteEntitySchemaObjectModel
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// CreatedBy undocumented
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`
	// LastModifiedBy undocumented
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
}

// OnenoteEntitySchemaObjectModel undocumented
type OnenoteEntitySchemaObjectModel struct {
	// OnenoteEntityBaseModel is the base model of OnenoteEntitySchemaObjectModel
	OnenoteEntityBaseModel
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
}

// OnenoteOperation undocumented
type OnenoteOperation struct {
	// Operation is the base model of OnenoteOperation
	Operation
	// ResourceLocation undocumented
	ResourceLocation *string `json:"resourceLocation,omitempty"`
	// ResourceID undocumented
	ResourceID *string `json:"resourceId,omitempty"`
	// Error undocumented
	Error *OnenoteOperationError `json:"error,omitempty"`
	// PercentComplete undocumented
	PercentComplete *string `json:"percentComplete,omitempty"`
}

// OnenoteOperationError undocumented
type OnenoteOperationError struct {
	// Object is the base model of OnenoteOperationError
	Object
	// Code undocumented
	Code *string `json:"code,omitempty"`
	// Message undocumented
	Message *string `json:"message,omitempty"`
}

// OnenotePage undocumented
type OnenotePage struct {
	// OnenoteEntitySchemaObjectModel is the base model of OnenotePage
	OnenoteEntitySchemaObjectModel
	// Title undocumented
	Title *string `json:"title,omitempty"`
	// CreatedByAppID undocumented
	CreatedByAppID *string `json:"createdByAppId,omitempty"`
	// Links undocumented
	Links *PageLinks `json:"links,omitempty"`
	// ContentURL undocumented
	ContentURL *string `json:"contentUrl,omitempty"`
	// Content undocumented
	Content *Stream `json:"content,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Level undocumented
	Level *int `json:"level,omitempty"`
	// Order undocumented
	Order *int `json:"order,omitempty"`
	// UserTags undocumented
	UserTags []string `json:"userTags,omitempty"`
	// ParentSection undocumented
	ParentSection *OnenoteSection `json:"parentSection,omitempty"`
	// ParentNotebook undocumented
	ParentNotebook *Notebook `json:"parentNotebook,omitempty"`
}

// OnenotePagePreview undocumented
type OnenotePagePreview struct {
	// Object is the base model of OnenotePagePreview
	Object
	// PreviewText undocumented
	PreviewText *string `json:"previewText,omitempty"`
	// Links undocumented
	Links *OnenotePagePreviewLinks `json:"links,omitempty"`
}

// OnenotePagePreviewLinks undocumented
type OnenotePagePreviewLinks struct {
	// Object is the base model of OnenotePagePreviewLinks
	Object
	// PreviewImageURL undocumented
	PreviewImageURL *ExternalLink `json:"previewImageUrl,omitempty"`
}

// OnenotePatchContentCommand undocumented
type OnenotePatchContentCommand struct {
	// Object is the base model of OnenotePatchContentCommand
	Object
	// Action undocumented
	Action *OnenotePatchActionType `json:"action,omitempty"`
	// Target undocumented
	Target *string `json:"target,omitempty"`
	// Content undocumented
	Content *string `json:"content,omitempty"`
	// Position undocumented
	Position *OnenotePatchInsertPosition `json:"position,omitempty"`
}

// OnenoteResource undocumented
type OnenoteResource struct {
	// OnenoteEntityBaseModel is the base model of OnenoteResource
	OnenoteEntityBaseModel
	// Content undocumented
	Content *Stream `json:"content,omitempty"`
	// ContentURL undocumented
	ContentURL *string `json:"contentUrl,omitempty"`
}

// OnenoteSection undocumented
type OnenoteSection struct {
	// OnenoteEntityHierarchyModel is the base model of OnenoteSection
	OnenoteEntityHierarchyModel
	// IsDefault undocumented
	IsDefault *bool `json:"isDefault,omitempty"`
	// Links undocumented
	Links *SectionLinks `json:"links,omitempty"`
	// PagesURL undocumented
	PagesURL *string `json:"pagesUrl,omitempty"`
	// ParentNotebook undocumented
	ParentNotebook *Notebook `json:"parentNotebook,omitempty"`
	// ParentSectionGroup undocumented
	ParentSectionGroup *SectionGroup `json:"parentSectionGroup,omitempty"`
	// Pages undocumented
	Pages []OnenotePage `json:"pages,omitempty"`
}