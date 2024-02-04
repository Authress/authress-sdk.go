package authress

import (
	"net/http"
)

// APIClient manages communication with the Authress API vv1
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services

	AccessRecords *AccessRecordsApi

	Applications *ApplicationsApi

	Connections *ConnectionsApi

	Extensions *ExtensionsApi

	Groups *GroupsApi

	Invites *InvitesApi

	ResourcePermissions *ResourcePermissionsApi

	Roles *RolesApi

	ServiceClients *ServiceClientsApi

	Tenants *TenantsApi

	UserPermissions *UserPermissionsApi

	Users *UsersApi
}

type service struct {
	client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = http.DefaultClient
	}

	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	// API Services
	c.AccessRecords = (*AccessRecordsApi)(&c.common)
	c.Accounts = (*AccountsApi)(&c.common)
	c.Applications = (*ApplicationsApi)(&c.common)
	c.Connections = (*ConnectionsApi)(&c.common)
	c.Extensions = (*ExtensionsApi)(&c.common)
	c.Groups = (*GroupsApi)(&c.common)
	c.Invites = (*InvitesApi)(&c.common)
	c.ResourcePermissions = (*ResourcePermissionsApi)(&c.common)
	c.Roles = (*RolesApi)(&c.common)
	c.ServiceClients = (*ServiceClientsApi)(&c.common)
	c.Tenants = (*TenantsApi)(&c.common)
	c.UserPermissions = (*UserPermissionsApi)(&c.common)
	c.Users = (*UsersApi)(&c.common)

	return c
}

// Allow modification of underlying config for alternate implementations and testing
// Caution: modifying the configuration while live can cause data races and potentially unwanted behavior
func (c *APIClient) GetConfig() *Configuration {
	return c.cfg
}
