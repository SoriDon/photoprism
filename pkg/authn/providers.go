package authn

import (
	"strings"

	"github.com/photoprism/photoprism/pkg/clean"
	"github.com/photoprism/photoprism/pkg/list"
	"github.com/photoprism/photoprism/pkg/txt"
)

// ProviderType represents an authentication provider type.
type ProviderType string

// Standard authentication provider types.
const (
	ProviderUndefined         ProviderType = ""
	ProviderDefault           ProviderType = "default"
	ProviderClient            ProviderType = "client"
	ProviderClientCredentials ProviderType = "client_credentials"
	ProviderApplication       ProviderType = "application"
	ProviderAccessToken       ProviderType = "access_token"
	ProviderLocal             ProviderType = "local"
	ProviderLDAP              ProviderType = "ldap"
	ProviderLink              ProviderType = "link"
	ProviderNone              ProviderType = "none"
)

// RemoteProviders contains remote auth providers.
var RemoteProviders = list.List{
	string(ProviderLDAP),
}

// LocalProviders contains local auth providers.
var LocalProviders = list.List{
	string(ProviderLocal),
}

// Method2FAProviders contains auth providers that support Method2FA.
var Method2FAProviders = list.List{
	string(ProviderDefault),
	string(ProviderLocal),
	string(ProviderLDAP),
}

// ClientProviders contains all client auth providers.
var ClientProviders = list.List{
	string(ProviderClient),
	string(ProviderClientCredentials),
	string(ProviderApplication),
	string(ProviderAccessToken),
}

// Is compares the provider with another type.
func (t ProviderType) Is(provider ProviderType) bool {
	return t == provider
}

// IsNot checks if the provider is not the specified type.
func (t ProviderType) IsNot(provider ProviderType) bool {
	return t != provider
}

// IsUndefined checks if the provider is undefined.
func (t ProviderType) IsUndefined() bool {
	return t == ""
}

// IsRemote checks if the provider is external.
func (t ProviderType) IsRemote() bool {
	return list.Contains(RemoteProviders, string(t))
}

// IsLocal checks if local authentication is possible.
func (t ProviderType) IsLocal() bool {
	return list.Contains(LocalProviders, string(t))
}

// Supports2FA checks if the provider supports two-factor authentication with a passcode.
func (t ProviderType) Supports2FA() bool {
	return list.Contains(Method2FAProviders, string(t))
}

// IsClient checks if the authentication is provided for a client.
func (t ProviderType) IsClient() bool {
	return list.Contains(ClientProviders, string(t))
}

// IsApplication checks if the authentication is provided for an application.
func (t ProviderType) IsApplication() bool {
	return t == ProviderApplication
}

// IsDefault checks if this is the default provider.
func (t ProviderType) IsDefault() bool {
	return t.String() == ProviderDefault.String()
}

// String returns the provider identifier as a string.
func (t ProviderType) String() string {
	switch t {
	case "":
		return string(ProviderDefault)
	case "token":
		return string(ProviderLink)
	case "password":
		return string(ProviderLocal)
	case "oauth2", "client credentials":
		return string(ProviderClientCredentials)
	default:
		return string(t)
	}
}

// Equal checks if the type matches.
func (t ProviderType) Equal(s string) bool {
	return strings.EqualFold(s, t.String())
}

// NotEqual checks if the type is different.
func (t ProviderType) NotEqual(s string) bool {
	return !t.Equal(s)
}

// Pretty returns the provider identifier in an easy-to-read format.
func (t ProviderType) Pretty() string {
	switch t {
	case ProviderLDAP:
		return "LDAP/AD"
	case ProviderClient:
		return "Client"
	case ProviderAccessToken:
		return "Access Token"
	case ProviderClientCredentials:
		return "Client Credentials"
	default:
		return txt.UpperFirst(t.String())
	}
}

// Provider casts a string to a normalized provider type.
func Provider(s string) ProviderType {
	s = clean.TypeLower(s)
	switch s {
	case "", "-", "null", "nil", "0", "false":
		return ProviderDefault
	case "token", "url":
		return ProviderLink
	case "pass", "passwd", "password":
		return ProviderLocal
	case "ldap", "ad", "ldap/ad", "ldap\\ad":
		return ProviderLDAP
	case "oauth2", "client credentials":
		return ProviderClientCredentials
	default:
		return ProviderType(s)
	}
}
