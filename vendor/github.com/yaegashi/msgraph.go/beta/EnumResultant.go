// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// ResultantAppState undocumented
type ResultantAppState string

const (
	// ResultantAppStateVInstalled undocumented
	ResultantAppStateVInstalled ResultantAppState = "installed"
	// ResultantAppStateVFailed undocumented
	ResultantAppStateVFailed ResultantAppState = "failed"
	// ResultantAppStateVNotInstalled undocumented
	ResultantAppStateVNotInstalled ResultantAppState = "notInstalled"
	// ResultantAppStateVUninstallFailed undocumented
	ResultantAppStateVUninstallFailed ResultantAppState = "uninstallFailed"
	// ResultantAppStateVPendingInstall undocumented
	ResultantAppStateVPendingInstall ResultantAppState = "pendingInstall"
	// ResultantAppStateVUnknown undocumented
	ResultantAppStateVUnknown ResultantAppState = "unknown"
	// ResultantAppStateVNotApplicable undocumented
	ResultantAppStateVNotApplicable ResultantAppState = "notApplicable"
)

var (
	// ResultantAppStatePInstalled is a pointer to ResultantAppStateVInstalled
	ResultantAppStatePInstalled = &_ResultantAppStatePInstalled
	// ResultantAppStatePFailed is a pointer to ResultantAppStateVFailed
	ResultantAppStatePFailed = &_ResultantAppStatePFailed
	// ResultantAppStatePNotInstalled is a pointer to ResultantAppStateVNotInstalled
	ResultantAppStatePNotInstalled = &_ResultantAppStatePNotInstalled
	// ResultantAppStatePUninstallFailed is a pointer to ResultantAppStateVUninstallFailed
	ResultantAppStatePUninstallFailed = &_ResultantAppStatePUninstallFailed
	// ResultantAppStatePPendingInstall is a pointer to ResultantAppStateVPendingInstall
	ResultantAppStatePPendingInstall = &_ResultantAppStatePPendingInstall
	// ResultantAppStatePUnknown is a pointer to ResultantAppStateVUnknown
	ResultantAppStatePUnknown = &_ResultantAppStatePUnknown
	// ResultantAppStatePNotApplicable is a pointer to ResultantAppStateVNotApplicable
	ResultantAppStatePNotApplicable = &_ResultantAppStatePNotApplicable
)

var (
	_ResultantAppStatePInstalled       = ResultantAppStateVInstalled
	_ResultantAppStatePFailed          = ResultantAppStateVFailed
	_ResultantAppStatePNotInstalled    = ResultantAppStateVNotInstalled
	_ResultantAppStatePUninstallFailed = ResultantAppStateVUninstallFailed
	_ResultantAppStatePPendingInstall  = ResultantAppStateVPendingInstall
	_ResultantAppStatePUnknown         = ResultantAppStateVUnknown
	_ResultantAppStatePNotApplicable   = ResultantAppStateVNotApplicable
)

// ResultantAppStateDetail undocumented
type ResultantAppStateDetail string

const (
	// ResultantAppStateDetailVNoAdditionalDetails undocumented
	ResultantAppStateDetailVNoAdditionalDetails ResultantAppStateDetail = "noAdditionalDetails"
	// ResultantAppStateDetailVDependencyFailedToInstall undocumented
	ResultantAppStateDetailVDependencyFailedToInstall ResultantAppStateDetail = "dependencyFailedToInstall"
	// ResultantAppStateDetailVDependencyWithRequirementsNotMet undocumented
	ResultantAppStateDetailVDependencyWithRequirementsNotMet ResultantAppStateDetail = "dependencyWithRequirementsNotMet"
	// ResultantAppStateDetailVDependencyPendingReboot undocumented
	ResultantAppStateDetailVDependencyPendingReboot ResultantAppStateDetail = "dependencyPendingReboot"
	// ResultantAppStateDetailVDependencyWithAutoInstallDisabled undocumented
	ResultantAppStateDetailVDependencyWithAutoInstallDisabled ResultantAppStateDetail = "dependencyWithAutoInstallDisabled"
	// ResultantAppStateDetailVSeeInstallErrorCode undocumented
	ResultantAppStateDetailVSeeInstallErrorCode ResultantAppStateDetail = "seeInstallErrorCode"
	// ResultantAppStateDetailVAutoInstallDisabled undocumented
	ResultantAppStateDetailVAutoInstallDisabled ResultantAppStateDetail = "autoInstallDisabled"
	// ResultantAppStateDetailVSeeUninstallErrorCode undocumented
	ResultantAppStateDetailVSeeUninstallErrorCode ResultantAppStateDetail = "seeUninstallErrorCode"
	// ResultantAppStateDetailVPendingReboot undocumented
	ResultantAppStateDetailVPendingReboot ResultantAppStateDetail = "pendingReboot"
	// ResultantAppStateDetailVInstallingDependencies undocumented
	ResultantAppStateDetailVInstallingDependencies ResultantAppStateDetail = "installingDependencies"
	// ResultantAppStateDetailVPowerShellScriptRequirementNotMet undocumented
	ResultantAppStateDetailVPowerShellScriptRequirementNotMet ResultantAppStateDetail = "powerShellScriptRequirementNotMet"
	// ResultantAppStateDetailVRegistryRequirementNotMet undocumented
	ResultantAppStateDetailVRegistryRequirementNotMet ResultantAppStateDetail = "registryRequirementNotMet"
	// ResultantAppStateDetailVFileSystemRequirementNotMet undocumented
	ResultantAppStateDetailVFileSystemRequirementNotMet ResultantAppStateDetail = "fileSystemRequirementNotMet"
	// ResultantAppStateDetailVPlatformNotApplicable undocumented
	ResultantAppStateDetailVPlatformNotApplicable ResultantAppStateDetail = "platformNotApplicable"
	// ResultantAppStateDetailVMinimumCPUSpeedNotMet undocumented
	ResultantAppStateDetailVMinimumCPUSpeedNotMet ResultantAppStateDetail = "minimumCpuSpeedNotMet"
	// ResultantAppStateDetailVMinimumLogicalProcessorCountNotMet undocumented
	ResultantAppStateDetailVMinimumLogicalProcessorCountNotMet ResultantAppStateDetail = "minimumLogicalProcessorCountNotMet"
	// ResultantAppStateDetailVMinimumPhysicalMemoryNotMet undocumented
	ResultantAppStateDetailVMinimumPhysicalMemoryNotMet ResultantAppStateDetail = "minimumPhysicalMemoryNotMet"
	// ResultantAppStateDetailVMinimumOsVersionNotMet undocumented
	ResultantAppStateDetailVMinimumOsVersionNotMet ResultantAppStateDetail = "minimumOsVersionNotMet"
	// ResultantAppStateDetailVMinimumDiskSpaceNotMet undocumented
	ResultantAppStateDetailVMinimumDiskSpaceNotMet ResultantAppStateDetail = "minimumDiskSpaceNotMet"
	// ResultantAppStateDetailVProcessorArchitectureNotApplicable undocumented
	ResultantAppStateDetailVProcessorArchitectureNotApplicable ResultantAppStateDetail = "processorArchitectureNotApplicable"
)

var (
	// ResultantAppStateDetailPNoAdditionalDetails is a pointer to ResultantAppStateDetailVNoAdditionalDetails
	ResultantAppStateDetailPNoAdditionalDetails = &_ResultantAppStateDetailPNoAdditionalDetails
	// ResultantAppStateDetailPDependencyFailedToInstall is a pointer to ResultantAppStateDetailVDependencyFailedToInstall
	ResultantAppStateDetailPDependencyFailedToInstall = &_ResultantAppStateDetailPDependencyFailedToInstall
	// ResultantAppStateDetailPDependencyWithRequirementsNotMet is a pointer to ResultantAppStateDetailVDependencyWithRequirementsNotMet
	ResultantAppStateDetailPDependencyWithRequirementsNotMet = &_ResultantAppStateDetailPDependencyWithRequirementsNotMet
	// ResultantAppStateDetailPDependencyPendingReboot is a pointer to ResultantAppStateDetailVDependencyPendingReboot
	ResultantAppStateDetailPDependencyPendingReboot = &_ResultantAppStateDetailPDependencyPendingReboot
	// ResultantAppStateDetailPDependencyWithAutoInstallDisabled is a pointer to ResultantAppStateDetailVDependencyWithAutoInstallDisabled
	ResultantAppStateDetailPDependencyWithAutoInstallDisabled = &_ResultantAppStateDetailPDependencyWithAutoInstallDisabled
	// ResultantAppStateDetailPSeeInstallErrorCode is a pointer to ResultantAppStateDetailVSeeInstallErrorCode
	ResultantAppStateDetailPSeeInstallErrorCode = &_ResultantAppStateDetailPSeeInstallErrorCode
	// ResultantAppStateDetailPAutoInstallDisabled is a pointer to ResultantAppStateDetailVAutoInstallDisabled
	ResultantAppStateDetailPAutoInstallDisabled = &_ResultantAppStateDetailPAutoInstallDisabled
	// ResultantAppStateDetailPSeeUninstallErrorCode is a pointer to ResultantAppStateDetailVSeeUninstallErrorCode
	ResultantAppStateDetailPSeeUninstallErrorCode = &_ResultantAppStateDetailPSeeUninstallErrorCode
	// ResultantAppStateDetailPPendingReboot is a pointer to ResultantAppStateDetailVPendingReboot
	ResultantAppStateDetailPPendingReboot = &_ResultantAppStateDetailPPendingReboot
	// ResultantAppStateDetailPInstallingDependencies is a pointer to ResultantAppStateDetailVInstallingDependencies
	ResultantAppStateDetailPInstallingDependencies = &_ResultantAppStateDetailPInstallingDependencies
	// ResultantAppStateDetailPPowerShellScriptRequirementNotMet is a pointer to ResultantAppStateDetailVPowerShellScriptRequirementNotMet
	ResultantAppStateDetailPPowerShellScriptRequirementNotMet = &_ResultantAppStateDetailPPowerShellScriptRequirementNotMet
	// ResultantAppStateDetailPRegistryRequirementNotMet is a pointer to ResultantAppStateDetailVRegistryRequirementNotMet
	ResultantAppStateDetailPRegistryRequirementNotMet = &_ResultantAppStateDetailPRegistryRequirementNotMet
	// ResultantAppStateDetailPFileSystemRequirementNotMet is a pointer to ResultantAppStateDetailVFileSystemRequirementNotMet
	ResultantAppStateDetailPFileSystemRequirementNotMet = &_ResultantAppStateDetailPFileSystemRequirementNotMet
	// ResultantAppStateDetailPPlatformNotApplicable is a pointer to ResultantAppStateDetailVPlatformNotApplicable
	ResultantAppStateDetailPPlatformNotApplicable = &_ResultantAppStateDetailPPlatformNotApplicable
	// ResultantAppStateDetailPMinimumCPUSpeedNotMet is a pointer to ResultantAppStateDetailVMinimumCPUSpeedNotMet
	ResultantAppStateDetailPMinimumCPUSpeedNotMet = &_ResultantAppStateDetailPMinimumCPUSpeedNotMet
	// ResultantAppStateDetailPMinimumLogicalProcessorCountNotMet is a pointer to ResultantAppStateDetailVMinimumLogicalProcessorCountNotMet
	ResultantAppStateDetailPMinimumLogicalProcessorCountNotMet = &_ResultantAppStateDetailPMinimumLogicalProcessorCountNotMet
	// ResultantAppStateDetailPMinimumPhysicalMemoryNotMet is a pointer to ResultantAppStateDetailVMinimumPhysicalMemoryNotMet
	ResultantAppStateDetailPMinimumPhysicalMemoryNotMet = &_ResultantAppStateDetailPMinimumPhysicalMemoryNotMet
	// ResultantAppStateDetailPMinimumOsVersionNotMet is a pointer to ResultantAppStateDetailVMinimumOsVersionNotMet
	ResultantAppStateDetailPMinimumOsVersionNotMet = &_ResultantAppStateDetailPMinimumOsVersionNotMet
	// ResultantAppStateDetailPMinimumDiskSpaceNotMet is a pointer to ResultantAppStateDetailVMinimumDiskSpaceNotMet
	ResultantAppStateDetailPMinimumDiskSpaceNotMet = &_ResultantAppStateDetailPMinimumDiskSpaceNotMet
	// ResultantAppStateDetailPProcessorArchitectureNotApplicable is a pointer to ResultantAppStateDetailVProcessorArchitectureNotApplicable
	ResultantAppStateDetailPProcessorArchitectureNotApplicable = &_ResultantAppStateDetailPProcessorArchitectureNotApplicable
)

var (
	_ResultantAppStateDetailPNoAdditionalDetails                = ResultantAppStateDetailVNoAdditionalDetails
	_ResultantAppStateDetailPDependencyFailedToInstall          = ResultantAppStateDetailVDependencyFailedToInstall
	_ResultantAppStateDetailPDependencyWithRequirementsNotMet   = ResultantAppStateDetailVDependencyWithRequirementsNotMet
	_ResultantAppStateDetailPDependencyPendingReboot            = ResultantAppStateDetailVDependencyPendingReboot
	_ResultantAppStateDetailPDependencyWithAutoInstallDisabled  = ResultantAppStateDetailVDependencyWithAutoInstallDisabled
	_ResultantAppStateDetailPSeeInstallErrorCode                = ResultantAppStateDetailVSeeInstallErrorCode
	_ResultantAppStateDetailPAutoInstallDisabled                = ResultantAppStateDetailVAutoInstallDisabled
	_ResultantAppStateDetailPSeeUninstallErrorCode              = ResultantAppStateDetailVSeeUninstallErrorCode
	_ResultantAppStateDetailPPendingReboot                      = ResultantAppStateDetailVPendingReboot
	_ResultantAppStateDetailPInstallingDependencies             = ResultantAppStateDetailVInstallingDependencies
	_ResultantAppStateDetailPPowerShellScriptRequirementNotMet  = ResultantAppStateDetailVPowerShellScriptRequirementNotMet
	_ResultantAppStateDetailPRegistryRequirementNotMet          = ResultantAppStateDetailVRegistryRequirementNotMet
	_ResultantAppStateDetailPFileSystemRequirementNotMet        = ResultantAppStateDetailVFileSystemRequirementNotMet
	_ResultantAppStateDetailPPlatformNotApplicable              = ResultantAppStateDetailVPlatformNotApplicable
	_ResultantAppStateDetailPMinimumCPUSpeedNotMet              = ResultantAppStateDetailVMinimumCPUSpeedNotMet
	_ResultantAppStateDetailPMinimumLogicalProcessorCountNotMet = ResultantAppStateDetailVMinimumLogicalProcessorCountNotMet
	_ResultantAppStateDetailPMinimumPhysicalMemoryNotMet        = ResultantAppStateDetailVMinimumPhysicalMemoryNotMet
	_ResultantAppStateDetailPMinimumOsVersionNotMet             = ResultantAppStateDetailVMinimumOsVersionNotMet
	_ResultantAppStateDetailPMinimumDiskSpaceNotMet             = ResultantAppStateDetailVMinimumDiskSpaceNotMet
	_ResultantAppStateDetailPProcessorArchitectureNotApplicable = ResultantAppStateDetailVProcessorArchitectureNotApplicable
)