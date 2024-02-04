package authress

import (
	"net/http"

	. "github.com/authress/authress-sdk.go/apis"
)

// Authress manages communication with the Authress API vv1
// In most cases there should be only one, shared, Authress.
type AuthressClient struct {
	authressSettings *AuthressSettings

	// API Services
	AccessRecords       *AccessRecordsApi
	Accounts            *AccountsApi
	Applications        *ApplicationsApi
	Connections         *ConnectionsApi
	Extensions          *ExtensionsApi
	Groups              *GroupsApi
	Invites             *InvitesApi
	ResourcePermissions *ResourcePermissionsApi
	Roles               *RolesApi
	ServiceClients      *ServiceClientsApi
	Tenants             *TenantsApi
	UserPermissions     *UserPermissionsApi
	Users               *UsersApi
}

// New creates a new Authress API client.
func NewAuthressClient(authressSettings AuthressSettings) *AuthressClient {
	httpClient := HttpClient{
		InternalClient: http.DefaultClient,
		ClientConfiguration: &Configuration{
			Version:   GetBuildInfo().Version,
			UserAgent: authressSettings.UserAgent,
			Scheme:    authressSettings.AuthressApiUrl.Scheme,
			Host:      authressSettings.AuthressApiUrl.Host,
		},
	}

	authressClient := AuthressClient{
		// API Services
		AccessRecords: &AccessRecordsApi{
			Client: &httpClient,
		},
		Accounts:            &AccountsApi{Client: &httpClient},
		Applications:        &ApplicationsApi{Client: &httpClient},
		Connections:         &ConnectionsApi{Client: &httpClient},
		Extensions:          &ExtensionsApi{Client: &httpClient},
		Groups:              &GroupsApi{Client: &httpClient},
		Invites:             &InvitesApi{Client: &httpClient},
		ResourcePermissions: &ResourcePermissionsApi{Client: &httpClient},
		Roles:               &RolesApi{Client: &httpClient},
		ServiceClients:      &ServiceClientsApi{Client: &httpClient},
		Tenants:             &TenantsApi{Client: &httpClient},
		UserPermissions:     &UserPermissionsApi{Client: &httpClient},
		Users:               &UsersApi{Client: &httpClient},
	}

	return &authressClient
}
