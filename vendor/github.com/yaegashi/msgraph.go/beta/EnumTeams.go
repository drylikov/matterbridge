// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// TeamsAppDistributionMethod undocumented
type TeamsAppDistributionMethod string

const (
	// TeamsAppDistributionMethodVStore undocumented
	TeamsAppDistributionMethodVStore TeamsAppDistributionMethod = "store"
	// TeamsAppDistributionMethodVOrganization undocumented
	TeamsAppDistributionMethodVOrganization TeamsAppDistributionMethod = "organization"
	// TeamsAppDistributionMethodVSideloaded undocumented
	TeamsAppDistributionMethodVSideloaded TeamsAppDistributionMethod = "sideloaded"
	// TeamsAppDistributionMethodVUnknownFutureValue undocumented
	TeamsAppDistributionMethodVUnknownFutureValue TeamsAppDistributionMethod = "unknownFutureValue"
)

var (
	// TeamsAppDistributionMethodPStore is a pointer to TeamsAppDistributionMethodVStore
	TeamsAppDistributionMethodPStore = &_TeamsAppDistributionMethodPStore
	// TeamsAppDistributionMethodPOrganization is a pointer to TeamsAppDistributionMethodVOrganization
	TeamsAppDistributionMethodPOrganization = &_TeamsAppDistributionMethodPOrganization
	// TeamsAppDistributionMethodPSideloaded is a pointer to TeamsAppDistributionMethodVSideloaded
	TeamsAppDistributionMethodPSideloaded = &_TeamsAppDistributionMethodPSideloaded
	// TeamsAppDistributionMethodPUnknownFutureValue is a pointer to TeamsAppDistributionMethodVUnknownFutureValue
	TeamsAppDistributionMethodPUnknownFutureValue = &_TeamsAppDistributionMethodPUnknownFutureValue
)

var (
	_TeamsAppDistributionMethodPStore              = TeamsAppDistributionMethodVStore
	_TeamsAppDistributionMethodPOrganization       = TeamsAppDistributionMethodVOrganization
	_TeamsAppDistributionMethodPSideloaded         = TeamsAppDistributionMethodVSideloaded
	_TeamsAppDistributionMethodPUnknownFutureValue = TeamsAppDistributionMethodVUnknownFutureValue
)

// TeamsAsyncOperationStatus undocumented
type TeamsAsyncOperationStatus string

const (
	// TeamsAsyncOperationStatusVInvalid undocumented
	TeamsAsyncOperationStatusVInvalid TeamsAsyncOperationStatus = "invalid"
	// TeamsAsyncOperationStatusVNotStarted undocumented
	TeamsAsyncOperationStatusVNotStarted TeamsAsyncOperationStatus = "notStarted"
	// TeamsAsyncOperationStatusVInProgress undocumented
	TeamsAsyncOperationStatusVInProgress TeamsAsyncOperationStatus = "inProgress"
	// TeamsAsyncOperationStatusVSucceeded undocumented
	TeamsAsyncOperationStatusVSucceeded TeamsAsyncOperationStatus = "succeeded"
	// TeamsAsyncOperationStatusVFailed undocumented
	TeamsAsyncOperationStatusVFailed TeamsAsyncOperationStatus = "failed"
	// TeamsAsyncOperationStatusVUnknownFutureValue undocumented
	TeamsAsyncOperationStatusVUnknownFutureValue TeamsAsyncOperationStatus = "unknownFutureValue"
)

var (
	// TeamsAsyncOperationStatusPInvalid is a pointer to TeamsAsyncOperationStatusVInvalid
	TeamsAsyncOperationStatusPInvalid = &_TeamsAsyncOperationStatusPInvalid
	// TeamsAsyncOperationStatusPNotStarted is a pointer to TeamsAsyncOperationStatusVNotStarted
	TeamsAsyncOperationStatusPNotStarted = &_TeamsAsyncOperationStatusPNotStarted
	// TeamsAsyncOperationStatusPInProgress is a pointer to TeamsAsyncOperationStatusVInProgress
	TeamsAsyncOperationStatusPInProgress = &_TeamsAsyncOperationStatusPInProgress
	// TeamsAsyncOperationStatusPSucceeded is a pointer to TeamsAsyncOperationStatusVSucceeded
	TeamsAsyncOperationStatusPSucceeded = &_TeamsAsyncOperationStatusPSucceeded
	// TeamsAsyncOperationStatusPFailed is a pointer to TeamsAsyncOperationStatusVFailed
	TeamsAsyncOperationStatusPFailed = &_TeamsAsyncOperationStatusPFailed
	// TeamsAsyncOperationStatusPUnknownFutureValue is a pointer to TeamsAsyncOperationStatusVUnknownFutureValue
	TeamsAsyncOperationStatusPUnknownFutureValue = &_TeamsAsyncOperationStatusPUnknownFutureValue
)

var (
	_TeamsAsyncOperationStatusPInvalid            = TeamsAsyncOperationStatusVInvalid
	_TeamsAsyncOperationStatusPNotStarted         = TeamsAsyncOperationStatusVNotStarted
	_TeamsAsyncOperationStatusPInProgress         = TeamsAsyncOperationStatusVInProgress
	_TeamsAsyncOperationStatusPSucceeded          = TeamsAsyncOperationStatusVSucceeded
	_TeamsAsyncOperationStatusPFailed             = TeamsAsyncOperationStatusVFailed
	_TeamsAsyncOperationStatusPUnknownFutureValue = TeamsAsyncOperationStatusVUnknownFutureValue
)

// TeamsAsyncOperationType undocumented
type TeamsAsyncOperationType string

const (
	// TeamsAsyncOperationTypeVInvalid undocumented
	TeamsAsyncOperationTypeVInvalid TeamsAsyncOperationType = "invalid"
	// TeamsAsyncOperationTypeVCloneTeam undocumented
	TeamsAsyncOperationTypeVCloneTeam TeamsAsyncOperationType = "cloneTeam"
	// TeamsAsyncOperationTypeVArchiveTeam undocumented
	TeamsAsyncOperationTypeVArchiveTeam TeamsAsyncOperationType = "archiveTeam"
	// TeamsAsyncOperationTypeVUnarchiveTeam undocumented
	TeamsAsyncOperationTypeVUnarchiveTeam TeamsAsyncOperationType = "unarchiveTeam"
	// TeamsAsyncOperationTypeVCreateTeam undocumented
	TeamsAsyncOperationTypeVCreateTeam TeamsAsyncOperationType = "createTeam"
	// TeamsAsyncOperationTypeVUnknownFutureValue undocumented
	TeamsAsyncOperationTypeVUnknownFutureValue TeamsAsyncOperationType = "unknownFutureValue"
)

var (
	// TeamsAsyncOperationTypePInvalid is a pointer to TeamsAsyncOperationTypeVInvalid
	TeamsAsyncOperationTypePInvalid = &_TeamsAsyncOperationTypePInvalid
	// TeamsAsyncOperationTypePCloneTeam is a pointer to TeamsAsyncOperationTypeVCloneTeam
	TeamsAsyncOperationTypePCloneTeam = &_TeamsAsyncOperationTypePCloneTeam
	// TeamsAsyncOperationTypePArchiveTeam is a pointer to TeamsAsyncOperationTypeVArchiveTeam
	TeamsAsyncOperationTypePArchiveTeam = &_TeamsAsyncOperationTypePArchiveTeam
	// TeamsAsyncOperationTypePUnarchiveTeam is a pointer to TeamsAsyncOperationTypeVUnarchiveTeam
	TeamsAsyncOperationTypePUnarchiveTeam = &_TeamsAsyncOperationTypePUnarchiveTeam
	// TeamsAsyncOperationTypePCreateTeam is a pointer to TeamsAsyncOperationTypeVCreateTeam
	TeamsAsyncOperationTypePCreateTeam = &_TeamsAsyncOperationTypePCreateTeam
	// TeamsAsyncOperationTypePUnknownFutureValue is a pointer to TeamsAsyncOperationTypeVUnknownFutureValue
	TeamsAsyncOperationTypePUnknownFutureValue = &_TeamsAsyncOperationTypePUnknownFutureValue
)

var (
	_TeamsAsyncOperationTypePInvalid            = TeamsAsyncOperationTypeVInvalid
	_TeamsAsyncOperationTypePCloneTeam          = TeamsAsyncOperationTypeVCloneTeam
	_TeamsAsyncOperationTypePArchiveTeam        = TeamsAsyncOperationTypeVArchiveTeam
	_TeamsAsyncOperationTypePUnarchiveTeam      = TeamsAsyncOperationTypeVUnarchiveTeam
	_TeamsAsyncOperationTypePCreateTeam         = TeamsAsyncOperationTypeVCreateTeam
	_TeamsAsyncOperationTypePUnknownFutureValue = TeamsAsyncOperationTypeVUnknownFutureValue
)
