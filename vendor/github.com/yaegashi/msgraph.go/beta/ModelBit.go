// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// BitLockerFixedDrivePolicy undocumented
type BitLockerFixedDrivePolicy struct {
	// Object is the base model of BitLockerFixedDrivePolicy
	Object
	// EncryptionMethod Select the encryption method for fixed drives.
	EncryptionMethod *BitLockerEncryptionMethod `json:"encryptionMethod,omitempty"`
	// RequireEncryptionForWriteAccess This policy setting determines whether BitLocker protection is required for fixed data drives to be writable on a computer.
	RequireEncryptionForWriteAccess *bool `json:"requireEncryptionForWriteAccess,omitempty"`
	// RecoveryOptions This policy setting allows you to control how BitLocker-protected fixed data drives are recovered in the absence of the required credentials. This policy setting is applied when you turn on BitLocker.
	RecoveryOptions *BitLockerRecoveryOptions `json:"recoveryOptions,omitempty"`
}

// BitLockerRecoveryOptions undocumented
type BitLockerRecoveryOptions struct {
	// Object is the base model of BitLockerRecoveryOptions
	Object
	// BlockDataRecoveryAgent Indicates whether to block certificate-based data recovery agent.
	BlockDataRecoveryAgent *bool `json:"blockDataRecoveryAgent,omitempty"`
	// RecoveryPasswordUsage Indicates whether users are allowed or required to generate a 48-digit recovery password for fixed or system disk.
	RecoveryPasswordUsage *ConfigurationUsage `json:"recoveryPasswordUsage,omitempty"`
	// RecoveryKeyUsage Indicates whether users are allowed or required to generate a 256-bit recovery key for fixed or system disk.
	RecoveryKeyUsage *ConfigurationUsage `json:"recoveryKeyUsage,omitempty"`
	// HideRecoveryOptions Indicates whether or not to allow showing recovery options in BitLocker Setup Wizard for fixed or system disk.
	HideRecoveryOptions *bool `json:"hideRecoveryOptions,omitempty"`
	// EnableRecoveryInformationSaveToStore Indicates whether or not to allow BitLocker recovery information to store in AD DS.
	EnableRecoveryInformationSaveToStore *bool `json:"enableRecoveryInformationSaveToStore,omitempty"`
	// RecoveryInformationToStore Configure what pieces of BitLocker recovery information are stored to AD DS.
	RecoveryInformationToStore *BitLockerRecoveryInformationType `json:"recoveryInformationToStore,omitempty"`
	// EnableBitLockerAfterRecoveryInformationToStore Indicates whether or not to enable BitLocker until recovery information is stored in AD DS.
	EnableBitLockerAfterRecoveryInformationToStore *bool `json:"enableBitLockerAfterRecoveryInformationToStore,omitempty"`
}

// BitLockerRemovableDrivePolicy undocumented
type BitLockerRemovableDrivePolicy struct {
	// Object is the base model of BitLockerRemovableDrivePolicy
	Object
	// EncryptionMethod Select the encryption method for removable  drives.
	EncryptionMethod *BitLockerEncryptionMethod `json:"encryptionMethod,omitempty"`
	// RequireEncryptionForWriteAccess Indicates whether to block write access to devices configured in another organization.  If requireEncryptionForWriteAccess is false, this value does not affect.
	RequireEncryptionForWriteAccess *bool `json:"requireEncryptionForWriteAccess,omitempty"`
	// BlockCrossOrganizationWriteAccess This policy setting determines whether BitLocker protection is required for removable data drives to be writable on a computer.
	BlockCrossOrganizationWriteAccess *bool `json:"blockCrossOrganizationWriteAccess,omitempty"`
}

// BitLockerSystemDrivePolicy undocumented
type BitLockerSystemDrivePolicy struct {
	// Object is the base model of BitLockerSystemDrivePolicy
	Object
	// EncryptionMethod Select the encryption method for operating system drives.
	EncryptionMethod *BitLockerEncryptionMethod `json:"encryptionMethod,omitempty"`
	// StartupAuthenticationRequired Require additional authentication at startup.
	StartupAuthenticationRequired *bool `json:"startupAuthenticationRequired,omitempty"`
	// StartupAuthenticationBlockWithoutTpmChip Indicates whether to allow BitLocker without a compatible TPM (requires a password or a startup key on a USB flash drive).
	StartupAuthenticationBlockWithoutTpmChip *bool `json:"startupAuthenticationBlockWithoutTpmChip,omitempty"`
	// StartupAuthenticationTpmUsage Indicates if TPM startup is allowed/required/disallowed.
	StartupAuthenticationTpmUsage *ConfigurationUsage `json:"startupAuthenticationTpmUsage,omitempty"`
	// StartupAuthenticationTpmPinUsage Indicates if TPM startup pin is allowed/required/disallowed.
	StartupAuthenticationTpmPinUsage *ConfigurationUsage `json:"startupAuthenticationTpmPinUsage,omitempty"`
	// StartupAuthenticationTpmKeyUsage Indicates if TPM startup key is allowed/required/disallowed.
	StartupAuthenticationTpmKeyUsage *ConfigurationUsage `json:"startupAuthenticationTpmKeyUsage,omitempty"`
	// StartupAuthenticationTpmPinAndKeyUsage Indicates if TPM startup pin key and key are allowed/required/disallowed.
	StartupAuthenticationTpmPinAndKeyUsage *ConfigurationUsage `json:"startupAuthenticationTpmPinAndKeyUsage,omitempty"`
	// MinimumPinLength Indicates the minimum length of startup pin. Valid values 4 to 20
	MinimumPinLength *int `json:"minimumPinLength,omitempty"`
	// RecoveryOptions Allows to recover BitLocker encrypted operating system drives in the absence of the required startup key information. This policy setting is applied when you turn on BitLocker.
	RecoveryOptions *BitLockerRecoveryOptions `json:"recoveryOptions,omitempty"`
	// PrebootRecoveryEnableMessageAndURL Enable pre-boot recovery message and Url. If requireStartupAuthentication is false, this value does not affect.
	PrebootRecoveryEnableMessageAndURL *bool `json:"prebootRecoveryEnableMessageAndUrl,omitempty"`
	// PrebootRecoveryMessage Defines a custom recovery message.
	PrebootRecoveryMessage *string `json:"prebootRecoveryMessage,omitempty"`
	// PrebootRecoveryURL Defines a custom recovery URL.
	PrebootRecoveryURL *string `json:"prebootRecoveryUrl,omitempty"`
}