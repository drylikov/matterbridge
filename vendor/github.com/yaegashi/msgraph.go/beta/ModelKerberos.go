// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// KerberosSingleSignOnExtension undocumented
type KerberosSingleSignOnExtension struct {
	// SingleSignOnExtension is the base model of KerberosSingleSignOnExtension
	SingleSignOnExtension
	// Realm Gets or sets the case-sensitive realm name for this profile.
	Realm *string `json:"realm,omitempty"`
	// Domains Gets or sets a list of hosts or domain names for which the app extension performs SSO.
	Domains []string `json:"domains,omitempty"`
	// BlockAutomaticLogin Enables or disables Keychain usage.
	BlockAutomaticLogin *bool `json:"blockAutomaticLogin,omitempty"`
	// CacheName Gets or sets the Generic Security Services name of the Kerberos cache to use for this profile.
	CacheName *string `json:"cacheName,omitempty"`
	// CredentialBundleIDAccessControlList Gets or sets a list of app Bundle IDs allowed to access the Kerberos Ticket Granting Ticket.
	CredentialBundleIDAccessControlList []string `json:"credentialBundleIdAccessControlList,omitempty"`
	// DomainRealms Gets or sets a list of realms for custom domain-realm mapping. Realms are case sensitive.
	DomainRealms []string `json:"domainRealms,omitempty"`
	// IsDefaultRealm When true, this profile's realm will be selected as the default. Necessary if multiple Kerberos-type profiles are configured.
	IsDefaultRealm *bool `json:"isDefaultRealm,omitempty"`
	// PasswordBlockModification Enables or disables password changes.
	PasswordBlockModification *bool `json:"passwordBlockModification,omitempty"`
	// PasswordExpirationDays Overrides the default password expiration in days. For most domains, this value is calculated automatically.
	PasswordExpirationDays *int `json:"passwordExpirationDays,omitempty"`
	// PasswordExpirationNotificationDays Gets or sets the number of days until the user is notified that their password will expire (default is 15).
	PasswordExpirationNotificationDays *int `json:"passwordExpirationNotificationDays,omitempty"`
	// UserPrincipalName Gets or sets the principle user name to use for this profile. The realm name does not need to be included.
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
	// PasswordRequireActiveDirectoryComplexity Enables or disables whether passwords must meet Active Directory's complexity requirements.
	PasswordRequireActiveDirectoryComplexity *bool `json:"passwordRequireActiveDirectoryComplexity,omitempty"`
	// PasswordPreviousPasswordBlockCount Gets or sets the number of previous passwords to block.
	PasswordPreviousPasswordBlockCount *int `json:"passwordPreviousPasswordBlockCount,omitempty"`
	// PasswordMinimumLength Gets or sets the minimum length of a password.
	PasswordMinimumLength *int `json:"passwordMinimumLength,omitempty"`
	// PasswordMinimumAgeDays Gets or sets the minimum number of days until a user can change their password again.
	PasswordMinimumAgeDays *int `json:"passwordMinimumAgeDays,omitempty"`
	// PasswordRequirementsDescription Gets or sets a description of the password complexity requirements.
	PasswordRequirementsDescription *string `json:"passwordRequirementsDescription,omitempty"`
	// RequireUserPresence Gets or sets whether to require authentication via Touch ID, Face ID, or a passcode to access the keychain entry.
	RequireUserPresence *bool `json:"requireUserPresence,omitempty"`
	// ActiveDirectorySiteCode Gets or sets the Active Directory site.
	ActiveDirectorySiteCode *string `json:"activeDirectorySiteCode,omitempty"`
	// PasswordEnableLocalSync Enables or disables password syncing. This won't affect users logged in with a mobile account on macOS.
	PasswordEnableLocalSync *bool `json:"passwordEnableLocalSync,omitempty"`
	// BlockActiveDirectorySiteAutoDiscovery Enables or disables whether the Kerberos extension can automatically determine its site name.
	BlockActiveDirectorySiteAutoDiscovery *bool `json:"blockActiveDirectorySiteAutoDiscovery,omitempty"`
}
