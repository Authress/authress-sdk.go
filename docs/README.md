[Back to Repository](../README.md) [Authress API](https://authress.io/app/#/api)

## Documentation for API Endpoints

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccessRecordsAPI* | [**CreateClaim**](./AccessRecordsAPI.md#createclaim) | **Post** /v1/claims | Create resource Claim
*AccessRecordsAPI* | [**CreateRecord**](./AccessRecordsAPI.md#createrecord) | **Post** /v1/records | Create access record
*AccessRecordsAPI* | [**CreateRequest**](./AccessRecordsAPI.md#createrequest) | **Post** /v1/requests | Create access request
*AccessRecordsAPI* | [**DeleteRecord**](./AccessRecordsAPI.md#deleterecord) | **Delete** /v1/records/{recordId} | Deletes access record
*AccessRecordsAPI* | [**DeleteRequest**](./AccessRecordsAPI.md#deleterequest) | **Delete** /v1/requests/{requestId} | Deletes access request
*AccessRecordsAPI* | [**GetRecord**](./AccessRecordsAPI.md#getrecord) | **Get** /v1/records/{recordId} | Retrieve access record
*AccessRecordsAPI* | [**GetRecords**](./AccessRecordsAPI.md#getrecords) | **Get** /v1/records | List access records
*AccessRecordsAPI* | [**GetRequest**](./AccessRecordsAPI.md#getrequest) | **Get** /v1/requests/{requestId} | Retrieve access request
*AccessRecordsAPI* | [**GetRequests**](./AccessRecordsAPI.md#getrequests) | **Get** /v1/requests | List access requests
*AccessRecordsAPI* | [**RespondToAccessRequest**](./AccessRecordsAPI.md#respondtoaccessrequest) | **Patch** /v1/requests/{requestId} | Approve or deny access request
*AccessRecordsAPI* | [**UpdateRecord**](./AccessRecordsAPI.md#updaterecord) | **Put** /v1/records/{recordId} | Update access record
*AccountsAPI* | [**DelegateAuthentication**](./AccountsAPI.md#delegateauthentication) | **Post** /v1/identities | Link external provider
*AccountsAPI* | [**GetAccount**](./AccountsAPI.md#getaccount) | **Get** /v1/accounts/{accountId} | Retrieve account information
*AccountsAPI* | [**GetAccountIdentities**](./AccountsAPI.md#getaccountidentities) | **Get** /v1/identities | List linked external providers
*AccountsAPI* | [**GetAccounts**](./AccountsAPI.md#getaccounts) | **Get** /v1/accounts | List user Authress accounts
*ApplicationsAPI* | [**DelegateUserLogin**](./ApplicationsAPI.md#delegateuserlogin) | **Post** /v1/applications/{applicationId}/users/{userId}/delegation | Log user into third-party application
*ConnectionsAPI* | [**CreateConnection**](./ConnectionsAPI.md#createconnection) | **Post** /v1/connections | Create SSO connection
*ConnectionsAPI* | [**DeleteConnection**](./ConnectionsAPI.md#deleteconnection) | **Delete** /v1/connections/{connectionId} | Delete SSO connection
*ConnectionsAPI* | [**GetConnection**](./ConnectionsAPI.md#getconnection) | **Get** /v1/connections/{connectionId} | Retrieve SSO connection
*ConnectionsAPI* | [**GetConnectionCredentials**](./ConnectionsAPI.md#getconnectioncredentials) | **Get** /v1/connections/{connectionId}/users/{userId}/credentials | Retrieve user connection credentials
*ConnectionsAPI* | [**GetConnections**](./ConnectionsAPI.md#getconnections) | **Get** /v1/connections | List SSO connections
*ConnectionsAPI* | [**UpdateConnection**](./ConnectionsAPI.md#updateconnection) | **Put** /v1/connections/{connectionId} | Update SSO connection
*ExtensionsAPI* | [**CreateExtension**](./ExtensionsAPI.md#createextension) | **Post** /v1/extensions | Create extension
*ExtensionsAPI* | [**DeleteExtension**](./ExtensionsAPI.md#deleteextension) | **Delete** /v1/extensions/{extensionId} | Delete extension
*ExtensionsAPI* | [**GetExtension**](./ExtensionsAPI.md#getextension) | **Get** /v1/extensions/{extensionId} | Retrieve extension
*ExtensionsAPI* | [**GetExtensions**](./ExtensionsAPI.md#getextensions) | **Get** /v1/extensions | List extensions
*ExtensionsAPI* | [**Login**](./ExtensionsAPI.md#login) | **Get** / | OAuth Authorize
*ExtensionsAPI* | [**RequestToken**](./ExtensionsAPI.md#requesttoken) | **Post** /api/authentication/oauth/tokens | OAuth Token
*ExtensionsAPI* | [**UpdateExtension**](./ExtensionsAPI.md#updateextension) | **Put** /v1/extensions/{extensionId} | Update extension
*GroupsAPI* | [**CreateGroup**](./GroupsAPI.md#creategroup) | **Post** /v1/groups | Create group
*GroupsAPI* | [**DeleteGroup**](./GroupsAPI.md#deletegroup) | **Delete** /v1/groups/{groupId} | Deletes group
*GroupsAPI* | [**GetGroup**](./GroupsAPI.md#getgroup) | **Get** /v1/groups/{groupId} | Retrieve group
*GroupsAPI* | [**GetGroups**](./GroupsAPI.md#getgroups) | **Get** /v1/groups | List groups
*GroupsAPI* | [**UpdateGroup**](./GroupsAPI.md#updategroup) | **Put** /v1/groups/{groupId} | Update a group
*InvitesAPI* | [**CreateInvite**](./InvitesAPI.md#createinvite) | **Post** /v1/invites | Create user invite
*InvitesAPI* | [**DeleteInvite**](./InvitesAPI.md#deleteinvite) | **Delete** /v1/invites/{inviteId} | Delete invite
*InvitesAPI* | [**GetInvite**](./InvitesAPI.md#getinvite) | **Get** /v1/invites/{inviteId} | Retrieve invite
*InvitesAPI* | [**RespondToInvite**](./InvitesAPI.md#respondtoinvite) | **Patch** /v1/invites/{inviteId} | Accept invite
*ResourcePermissionsAPI* | [**GetPermissionedResource**](./ResourcePermissionsAPI.md#getpermissionedresource) | **Get** /v1/resources/{resourceUri} | Retrieve resource configuration
*ResourcePermissionsAPI* | [**GetPermissionedResources**](./ResourcePermissionsAPI.md#getpermissionedresources) | **Get** /v1/resources | List all resource configurations
*ResourcePermissionsAPI* | [**GetResourceUsers**](./ResourcePermissionsAPI.md#getresourceusers) | **Get** /v1/resources/{resourceUri}/users | List users with resource access
*ResourcePermissionsAPI* | [**UpdatePermissionedResource**](./ResourcePermissionsAPI.md#updatepermissionedresource) | **Put** /v1/resources/{resourceUri} | Update resource configuration
*RolesAPI* | [**CreateRole**](./RolesAPI.md#createrole) | **Post** /v1/roles | Create role
*RolesAPI* | [**DeleteRole**](./RolesAPI.md#deleterole) | **Delete** /v1/roles/{roleId} | Deletes role
*RolesAPI* | [**GetRole**](./RolesAPI.md#getrole) | **Get** /v1/roles/{roleId} | Retrieve role
*RolesAPI* | [**GetRoles**](./RolesAPI.md#getroles) | **Get** /v1/roles | List roles
*RolesAPI* | [**UpdateRole**](./RolesAPI.md#updaterole) | **Put** /v1/roles/{roleId} | Update role
*ServiceClientsAPI* | [**CreateClient**](./ServiceClientsAPI.md#createclient) | **Post** /v1/clients | Create service client
*ServiceClientsAPI* | [**DeleteAccessKey**](./ServiceClientsAPI.md#deleteaccesskey) | **Delete** /v1/clients/{clientId}/access-keys/{keyId} | Delete service client access key
*ServiceClientsAPI* | [**DeleteClient**](./ServiceClientsAPI.md#deleteclient) | **Delete** /v1/clients/{clientId} | Delete service client
*ServiceClientsAPI* | [**GetClient**](./ServiceClientsAPI.md#getclient) | **Get** /v1/clients/{clientId} | Retrieve service client
*ServiceClientsAPI* | [**GetClients**](./ServiceClientsAPI.md#getclients) | **Get** /v1/clients | List service clients
*ServiceClientsAPI* | [**RequestAccessKey**](./ServiceClientsAPI.md#requestaccesskey) | **Post** /v1/clients/{clientId}/access-keys | Generate service client access key
*ServiceClientsAPI* | [**UpdateClient**](./ServiceClientsAPI.md#updateclient) | **Put** /v1/clients/{clientId} | Update service client
*TenantsAPI* | [**CreateTenant**](./TenantsAPI.md#createtenant) | **Post** /v1/tenants | Create tenant
*TenantsAPI* | [**DeleteTenant**](./TenantsAPI.md#deletetenant) | **Delete** /v1/tenants/{tenantId} | Delete tenant
*TenantsAPI* | [**GetTenant**](./TenantsAPI.md#gettenant) | **Get** /v1/tenants/{tenantId} | Retrieve tenant
*TenantsAPI* | [**GetTenants**](./TenantsAPI.md#gettenants) | **Get** /v1/tenants | List tenants
*TenantsAPI* | [**UpdateTenant**](./TenantsAPI.md#updatetenant) | **Put** /v1/tenants/{tenantId} | Update tenant
*UserPermissionsAPI* | [**AuthorizeUser**](./UserPermissionsAPI.md#authorizeuser) | **Get** /v1/users/{userId}/resources/{resourceUri}/permissions/{permission} | Verify user authorization
*UserPermissionsAPI* | [**GetUserPermissionsForResource**](./UserPermissionsAPI.md#getuserpermissionsforresource) | **Get** /v1/users/{userId}/resources/{resourceUri}/permissions | Get user permissions for resource
*UserPermissionsAPI* | [**GetUserResources**](./UserPermissionsAPI.md#getuserresources) | **Get** /v1/users/{userId}/resources | List user resources
*UserPermissionsAPI* | [**GetUserRolesForResource**](./UserPermissionsAPI.md#getuserrolesforresource) | **Get** /v1/users/{userId}/resources/{resourceUri}/roles | Get user roles for resource
*UsersAPI* | [**DeleteUser**](./UsersAPI.md#deleteuser) | **Delete** /v1/users/{userId} | Delete a user
*UsersAPI* | [**GetUser**](./UsersAPI.md#getuser) | **Get** /v1/users/{userId} | Retrieve a user
*UsersAPI* | [**GetUsers**](./UsersAPI.md#getusers) | **Get** /v1/users | List users


## Documentation For Models

 - [AccessRecord](./AccessRecord.md)
 - [AccessRecordAccount](./AccessRecordAccount.md)
 - [AccessRecordCollection](./AccessRecordCollection.md)
 - [AccessRequest](./AccessRequest.md)
 - [AccessRequestCollection](./AccessRequestCollection.md)
 - [AccessRequestResponse](./AccessRequestResponse.md)
 - [AccessTemplate](./AccessTemplate.md)
 - [Account](./Account.md)
 - [AccountCollection](./AccountCollection.md)
 - [AccountLinks](./AccountLinks.md)
 - [ApplicationDelegation](./ApplicationDelegation.md)
 - [ClaimRequest](./ClaimRequest.md)
 - [Client](./Client.md)
 - [ClientAccessKey](./ClientAccessKey.md)
 - [ClientCollection](./ClientCollection.md)
 - [ClientOptions](./ClientOptions.md)
 - [CollectionLinks](./CollectionLinks.md)
 - [Connection](./Connection.md)
 - [ConnectionCollection](./ConnectionCollection.md)
 - [ConnectionData](./ConnectionData.md)
 - [ConnectionDefaultConnectionProperties](./ConnectionDefaultConnectionProperties.md)
 - [ConnectionUserDataConfiguration](./ConnectionUserDataConfiguration.md)
 - [Extension](./Extension.md)
 - [ExtensionApplication](./ExtensionApplication.md)
 - [ExtensionClient](./ExtensionClient.md)
 - [ExtensionCollection](./ExtensionCollection.md)
 - [Group](./Group.md)
 - [GroupCollection](./GroupCollection.md)
 - [Identity](./Identity.md)
 - [IdentityCollection](./IdentityCollection.md)
 - [IdentityRequest](./IdentityRequest.md)
 - [Invite](./Invite.md)
 - [Link](./Link.md)
 - [LinkedGroup](./LinkedGroup.md)
 - [Links](./Links.md)
 - [OAuthAuthorizeResponse](./OAuthAuthorizeResponse.md)
 - [OAuthTokenRequest](./OAuthTokenRequest.md)
 - [OAuthTokenResponse](./OAuthTokenResponse.md)
 - [Pagination](./Pagination.md)
 - [PaginationNext](./PaginationNext.md)
 - [PermissionCollection](./PermissionCollection.md)
 - [PermissionCollectionAccount](./PermissionCollectionAccount.md)
 - [PermissionObject](./PermissionObject.md)
 - [PermissionedResource](./PermissionedResource.md)
 - [PermissionedResourceCollection](./PermissionedResourceCollection.md)
 - [Resource](./Resource.md)
 - [ResourcePermission](./ResourcePermission.md)
 - [ResourceUsersCollection](./ResourceUsersCollection.md)
 - [Role](./Role.md)
 - [RoleCollection](./RoleCollection.md)
 - [Statement](./Statement.md)
 - [Tenant](./Tenant.md)
 - [TenantCollection](./TenantCollection.md)
 - [TenantConnection](./TenantConnection.md)
 - [TenantData](./TenantData.md)
 - [TokenRequest](./TokenRequest.md)
 - [User](./User.md)
 - [UserConnectionCredentials](./UserConnectionCredentials.md)
 - [UserIdentity](./UserIdentity.md)
 - [UserIdentityCollection](./UserIdentityCollection.md)
 - [UserResourcesCollection](./UserResourcesCollection.md)
 - [UserRole](./UserRole.md)
 - [UserRoleCollection](./UserRoleCollection.md)
 - [UserToken](./UserToken.md)

