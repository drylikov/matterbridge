// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// Message undocumented
type Message struct {
	// OutlookItem is the base model of Message
	OutlookItem
	// ReceivedDateTime undocumented
	ReceivedDateTime *time.Time `json:"receivedDateTime,omitempty"`
	// SentDateTime undocumented
	SentDateTime *time.Time `json:"sentDateTime,omitempty"`
	// HasAttachments undocumented
	HasAttachments *bool `json:"hasAttachments,omitempty"`
	// InternetMessageID undocumented
	InternetMessageID *string `json:"internetMessageId,omitempty"`
	// InternetMessageHeaders undocumented
	InternetMessageHeaders []InternetMessageHeader `json:"internetMessageHeaders,omitempty"`
	// Subject undocumented
	Subject *string `json:"subject,omitempty"`
	// Body undocumented
	Body *ItemBody `json:"body,omitempty"`
	// BodyPreview undocumented
	BodyPreview *string `json:"bodyPreview,omitempty"`
	// Importance undocumented
	Importance *Importance `json:"importance,omitempty"`
	// ParentFolderID undocumented
	ParentFolderID *string `json:"parentFolderId,omitempty"`
	// Sender undocumented
	Sender *Recipient `json:"sender,omitempty"`
	// From undocumented
	From *Recipient `json:"from,omitempty"`
	// ToRecipients undocumented
	ToRecipients []Recipient `json:"toRecipients,omitempty"`
	// CcRecipients undocumented
	CcRecipients []Recipient `json:"ccRecipients,omitempty"`
	// BccRecipients undocumented
	BccRecipients []Recipient `json:"bccRecipients,omitempty"`
	// ReplyTo undocumented
	ReplyTo []Recipient `json:"replyTo,omitempty"`
	// ConversationID undocumented
	ConversationID *string `json:"conversationId,omitempty"`
	// ConversationIndex undocumented
	ConversationIndex *Binary `json:"conversationIndex,omitempty"`
	// UniqueBody undocumented
	UniqueBody *ItemBody `json:"uniqueBody,omitempty"`
	// IsDeliveryReceiptRequested undocumented
	IsDeliveryReceiptRequested *bool `json:"isDeliveryReceiptRequested,omitempty"`
	// IsReadReceiptRequested undocumented
	IsReadReceiptRequested *bool `json:"isReadReceiptRequested,omitempty"`
	// IsRead undocumented
	IsRead *bool `json:"isRead,omitempty"`
	// IsDraft undocumented
	IsDraft *bool `json:"isDraft,omitempty"`
	// WebLink undocumented
	WebLink *string `json:"webLink,omitempty"`
	// MentionsPreview undocumented
	MentionsPreview *MentionsPreview `json:"mentionsPreview,omitempty"`
	// InferenceClassification undocumented
	InferenceClassification *InferenceClassificationType `json:"inferenceClassification,omitempty"`
	// UnsubscribeData undocumented
	UnsubscribeData []string `json:"unsubscribeData,omitempty"`
	// UnsubscribeEnabled undocumented
	UnsubscribeEnabled *bool `json:"unsubscribeEnabled,omitempty"`
	// Flag undocumented
	Flag *FollowupFlag `json:"flag,omitempty"`
	// SingleValueExtendedProperties undocumented
	SingleValueExtendedProperties []SingleValueLegacyExtendedProperty `json:"singleValueExtendedProperties,omitempty"`
	// MultiValueExtendedProperties undocumented
	MultiValueExtendedProperties []MultiValueLegacyExtendedProperty `json:"multiValueExtendedProperties,omitempty"`
	// Attachments undocumented
	Attachments []Attachment `json:"attachments,omitempty"`
	// Extensions undocumented
	Extensions []Extension `json:"extensions,omitempty"`
	// Mentions undocumented
	Mentions []Mention `json:"mentions,omitempty"`
}

// MessageRule undocumented
type MessageRule struct {
	// Entity is the base model of MessageRule
	Entity
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Sequence undocumented
	Sequence *int `json:"sequence,omitempty"`
	// Conditions undocumented
	Conditions *MessageRulePredicates `json:"conditions,omitempty"`
	// Actions undocumented
	Actions *MessageRuleActions `json:"actions,omitempty"`
	// Exceptions undocumented
	Exceptions *MessageRulePredicates `json:"exceptions,omitempty"`
	// IsEnabled undocumented
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// HasError undocumented
	HasError *bool `json:"hasError,omitempty"`
	// IsReadOnly undocumented
	IsReadOnly *bool `json:"isReadOnly,omitempty"`
}

// MessageRuleActions undocumented
type MessageRuleActions struct {
	// Object is the base model of MessageRuleActions
	Object
	// MoveToFolder undocumented
	MoveToFolder *string `json:"moveToFolder,omitempty"`
	// CopyToFolder undocumented
	CopyToFolder *string `json:"copyToFolder,omitempty"`
	// Delete undocumented
	Delete *bool `json:"delete,omitempty"`
	// PermanentDelete undocumented
	PermanentDelete *bool `json:"permanentDelete,omitempty"`
	// MarkAsRead undocumented
	MarkAsRead *bool `json:"markAsRead,omitempty"`
	// MarkImportance undocumented
	MarkImportance *Importance `json:"markImportance,omitempty"`
	// ForwardTo undocumented
	ForwardTo []Recipient `json:"forwardTo,omitempty"`
	// ForwardAsAttachmentTo undocumented
	ForwardAsAttachmentTo []Recipient `json:"forwardAsAttachmentTo,omitempty"`
	// RedirectTo undocumented
	RedirectTo []Recipient `json:"redirectTo,omitempty"`
	// AssignCategories undocumented
	AssignCategories []string `json:"assignCategories,omitempty"`
	// StopProcessingRules undocumented
	StopProcessingRules *bool `json:"stopProcessingRules,omitempty"`
}

// MessageRulePredicates undocumented
type MessageRulePredicates struct {
	// Object is the base model of MessageRulePredicates
	Object
	// Categories undocumented
	Categories []string `json:"categories,omitempty"`
	// SubjectContains undocumented
	SubjectContains []string `json:"subjectContains,omitempty"`
	// BodyContains undocumented
	BodyContains []string `json:"bodyContains,omitempty"`
	// BodyOrSubjectContains undocumented
	BodyOrSubjectContains []string `json:"bodyOrSubjectContains,omitempty"`
	// SenderContains undocumented
	SenderContains []string `json:"senderContains,omitempty"`
	// RecipientContains undocumented
	RecipientContains []string `json:"recipientContains,omitempty"`
	// HeaderContains undocumented
	HeaderContains []string `json:"headerContains,omitempty"`
	// MessageActionFlag undocumented
	MessageActionFlag *MessageActionFlag `json:"messageActionFlag,omitempty"`
	// Importance undocumented
	Importance *Importance `json:"importance,omitempty"`
	// Sensitivity undocumented
	Sensitivity *Sensitivity `json:"sensitivity,omitempty"`
	// FromAddresses undocumented
	FromAddresses []Recipient `json:"fromAddresses,omitempty"`
	// SentToAddresses undocumented
	SentToAddresses []Recipient `json:"sentToAddresses,omitempty"`
	// SentToMe undocumented
	SentToMe *bool `json:"sentToMe,omitempty"`
	// SentOnlyToMe undocumented
	SentOnlyToMe *bool `json:"sentOnlyToMe,omitempty"`
	// SentCcMe undocumented
	SentCcMe *bool `json:"sentCcMe,omitempty"`
	// SentToOrCcMe undocumented
	SentToOrCcMe *bool `json:"sentToOrCcMe,omitempty"`
	// NotSentToMe undocumented
	NotSentToMe *bool `json:"notSentToMe,omitempty"`
	// HasAttachments undocumented
	HasAttachments *bool `json:"hasAttachments,omitempty"`
	// IsApprovalRequest undocumented
	IsApprovalRequest *bool `json:"isApprovalRequest,omitempty"`
	// IsAutomaticForward undocumented
	IsAutomaticForward *bool `json:"isAutomaticForward,omitempty"`
	// IsAutomaticReply undocumented
	IsAutomaticReply *bool `json:"isAutomaticReply,omitempty"`
	// IsEncrypted undocumented
	IsEncrypted *bool `json:"isEncrypted,omitempty"`
	// IsMeetingRequest undocumented
	IsMeetingRequest *bool `json:"isMeetingRequest,omitempty"`
	// IsMeetingResponse undocumented
	IsMeetingResponse *bool `json:"isMeetingResponse,omitempty"`
	// IsNonDeliveryReport undocumented
	IsNonDeliveryReport *bool `json:"isNonDeliveryReport,omitempty"`
	// IsPermissionControlled undocumented
	IsPermissionControlled *bool `json:"isPermissionControlled,omitempty"`
	// IsReadReceipt undocumented
	IsReadReceipt *bool `json:"isReadReceipt,omitempty"`
	// IsSigned undocumented
	IsSigned *bool `json:"isSigned,omitempty"`
	// IsVoicemail undocumented
	IsVoicemail *bool `json:"isVoicemail,omitempty"`
	// WithinSizeRange undocumented
	WithinSizeRange *SizeRange `json:"withinSizeRange,omitempty"`
}
