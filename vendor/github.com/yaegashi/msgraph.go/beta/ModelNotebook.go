// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// Notebook undocumented
type Notebook struct {
	// OnenoteEntityHierarchyModel is the base model of Notebook
	OnenoteEntityHierarchyModel
	// IsDefault undocumented
	IsDefault *bool `json:"isDefault,omitempty"`
	// UserRole undocumented
	UserRole *OnenoteUserRole `json:"userRole,omitempty"`
	// IsShared undocumented
	IsShared *bool `json:"isShared,omitempty"`
	// SectionsURL undocumented
	SectionsURL *string `json:"sectionsUrl,omitempty"`
	// SectionGroupsURL undocumented
	SectionGroupsURL *string `json:"sectionGroupsUrl,omitempty"`
	// Links undocumented
	Links *NotebookLinks `json:"links,omitempty"`
	// Sections undocumented
	Sections []OnenoteSection `json:"sections,omitempty"`
	// SectionGroups undocumented
	SectionGroups []SectionGroup `json:"sectionGroups,omitempty"`
}

// NotebookLinks undocumented
type NotebookLinks struct {
	// Object is the base model of NotebookLinks
	Object
	// OneNoteClientURL undocumented
	OneNoteClientURL *ExternalLink `json:"oneNoteClientUrl,omitempty"`
	// OneNoteWebURL undocumented
	OneNoteWebURL *ExternalLink `json:"oneNoteWebUrl,omitempty"`
}