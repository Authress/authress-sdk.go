[Back to Repository](../README.md) [Authress API](https://authress.io/app/#/api)

## Documentation for API Endpoints

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccessRecords* | [**CreateClaim**](./AccessRecords.md#createclaim) | **Post** /v1/claims | Create resource Claim
*AccessRecords* | [**CreateRecord**](./AccessRecords.md#createrecord) | **Post** /v1/records | Create access record
*AccessRecords* | [**CreateRequest**](./AccessRecords.md#createrequest) | **Post** /v1/requests | Create access request
*AccessRecords* | [**DeleteRecord**](./AccessRecords.md#deleterecord) | **Delete** /v1/records/{recordId} | Deletes access record
*AccessRecords* | [**DeleteRequest**](./AccessRecords.md#deleterequest) | **Delete** /v1/requests/{requestId} | Deletes access request
*AccessRecords* | [**GetRecord**](./AccessRecords.md#getrecord) | **Get** /v1/records/{recordId} | Retrieve access record
*AccessRecords* | [**GetRecords**](./AccessRecords.md#getrecords) | **Get** /v1/records | List access records
*AccessRecords* | [**GetRequest**](./AccessRecords.md#getrequest) | **Get** /v1/requests/{requestId} | Retrieve access request
*AccessRecords* | [**GetRequests**](./AccessRecords.md#getrequests) | **Get** /v1/requests | List access requests
*AccessRecords* | [**RespondToAccessRequest**](./AccessRecords.md#respondtoaccessrequest) | **Patch** /v1/requests/{requestId} | Approve or deny access request
*AccessRecords* | [**UpdateRecord**](./AccessRecords.md#updaterecord) | **Put** /v1/records/{recordId} | Update access record
*Accounts* | [**DelegateAuthentication**](./Accounts.md#delegateauthentication) | **Post** /v1/identities | Link external provider
*Accounts* | [**GetAccount**](./Accounts.md#getaccount) | **Get** /v1/accounts/{accountId} | Retrieve account information
*Accounts* | [**GetAccountIdentities**](./Accounts.md#getaccountidentities) | **Get** /v1/identities | List linked external providers
*Accounts* | [**GetAccounts**](./Accounts.md#getaccounts) | **Get** /v1/accounts | List user Authress accounts
*Applications* | [**DelegateUserLogin**](./Applications.md#delegateuserlogin) | **Post** /v1/applications/{applicationId}/users/{userId}/delegation | Log user into third-party application
*Connections* | [**CreateConnection**](./Connections.md#createconnection) | **Post** /v1/connections | Create SSO connection
*Connections* | [**DeleteConnection**](./Connections.md#deleteconnection) | **Delete** /v1/connections/{connectionId} | Delete SSO connection
*Connections* | [**GetConnection**](./Connections.md#getconnection) | **Get** /v1/connections/{connectionId} | Retrieve SSO connection
*Connections* | [**GetConnectionCredentials**](./Connections.md#getconnectioncredentials) | **Get** /v1/connections/{connectionId}/users/{userId}/credentials | Retrieve user connection credentials
*Connections* | [**GetConnections**](./Connections.md#getconnections) | **Get** /v1/connections | List SSO connections
*Connections* | [**UpdateConnection**](./Connections.md#updateconnection) | **Put** /v1/connections/{connectionId} | Update SSO connection
*Extensions* | [**CreateExtension**](./Extensions.md#createextension) | **Post** /v1/extensions | Create extension
*Extensions* | [**DeleteExtension**](./Extensions.md#deleteextension) | **Delete** /v1/extensions/{extensionId} | Delete extension
*Extensions* | [**GetExtension**](./Extensions.md#getextension) | **Get** /v1/extensions/{extensionId} | Retrieve extension
*Extensions* | [**GetExtensions**](./Extensions.md#getextensions) | **Get** /v1/extensions | List extensions
*Extensions* | [**Login**](./Extensions.md#login) | **Get** / | OAuth Authorize
*Extensions* | [**RequestToken**](./Extensions.md#requesttoken) | **Post** /api/authentication/oauth/tokens | OAuth Token
*Extensions* | [**UpdateExtension**](./Extensions.md#updateextension) | **Put** /v1/extensions/{extensionId} | Update extension
*Groups* | [**CreateGroup**](./Groups.md#creategroup) | **Post** /v1/groups | Create group
*Groups* | [**DeleteGroup**](./Groups.md#deletegroup) | **Delete** /v1/groups/{groupId} | Deletes group
*Groups* | [**GetGroup**](./Groups.md#getgroup) | **Get** /v1/groups/{groupId} | Retrieve group
*Groups* | [**GetGroups**](./Groups.md#getgroups) | **Get** /v1/groups | List groups
*Groups* | [**UpdateGroup**](./Groups.md#updategroup) | **Put** /v1/groups/{groupId} | Update a group
*Invites* | [**CreateInvite**](./Invites.md#createinvite) | **Post** /v1/invites | Create user invite
*Invites* | [**DeleteInvite**](./Invites.md#deleteinvite) | **Delete** /v1/invites/{inviteId} | Delete invite
*Invites* | [**GetInvite**](./Invites.md#getinvite) | **Get** /v1/invites/{inviteId} | Retrieve invite
*Invites* | [**RespondToInvite**](./Invites.md#respondtoinvite) | **Patch** /v1/invites/{inviteId} | Accept invite
*ResourcePermissions* | [**GetPermissionedResource**](./ResourcePermissions.md#getpermissionedresource) | **Get** /v1/resources/{resourceUri} | Retrieve resource configuration
*ResourcePermissions* | [**GetPermissionedResources**](./ResourcePermissions.md#getpermissionedresources) | **Get** /v1/resources | List all resource configurations
*ResourcePermissions* | [**GetResourceUsers**](./ResourcePermissions.md#getresourceusers) | **Get** /v1/resources/{resourceUri}/users | List users with resource access
*ResourcePermissions* | [**UpdatePermissionedResource**](./ResourcePermissions.md#updatepermissionedresource) | **Put** /v1/resources/{resourceUri} | Update resource configuration
*Roles* | [**CreateRole**](./Roles.md#createrole) | **Post** /v1/roles | Create role
*Roles* | [**DeleteRole**](./Roles.md#deleterole) | **Delete** /v1/roles/{roleId} | Deletes role
*Roles* | [**GetRole**](./Roles.md#getrole) | **Get** /v1/roles/{roleId} | Retrieve role
*Roles* | [**GetRoles**](./Roles.md#getroles) | **Get** /v1/roles | List roles
*Roles* | [**UpdateRole**](./Roles.md#updaterole) | **Put** /v1/roles/{roleId} | Update role
*ServiceClients* | [**CreateClient**](./ServiceClients.md#createclient) | **Post** /v1/clients | Create service client
*ServiceClients* | [**DeleteAccessKey**](./ServiceClients.md#deleteaccesskey) | **Delete** /v1/clients/{clientId}/access-keys/{keyId} | Delete service client access key
*ServiceClients* | [**DeleteClient**](./ServiceClients.md#deleteclient) | **Delete** /v1/clients/{clientId} | Delete service client
*ServiceClients* | [**GetClient**](./ServiceClients.md#getclient) | **Get** /v1/clients/{clientId} | Retrieve service client
*ServiceClients* | [**GetClients**](./ServiceClients.md#getclients) | **Get** /v1/clients | List service clients
*ServiceClients* | [**RequestAccessKey**](./ServiceClients.md#requestaccesskey) | **Post** /v1/clients/{clientId}/access-keys | Generate service client access key
*ServiceClients* | [**UpdateClient**](./ServiceClients.md#updateclient) | **Put** /v1/clients/{clientId} | Update service client
*Tenants* | [**CreateTenant**](./Tenants.md#createtenant) | **Post** /v1/tenants | Create tenant
*Tenants* | [**DeleteTenant**](./Tenants.md#deletetenant) | **Delete** /v1/tenants/{tenantId} | Delete tenant
*Tenants* | [**GetTenant**](./Tenants.md#gettenant) | **Get** /v1/tenants/{tenantId} | Retrieve tenant
*Tenants* | [**GetTenants**](./Tenants.md#gettenants) | **Get** /v1/tenants | List tenants
*Tenants* | [**UpdateTenant**](./Tenants.md#updatetenant) | **Put** /v1/tenants/{tenantId} | Update tenant
*UserPermissions* | [**AuthorizeUser**](./UserPermissions.md#authorizeuser) | **Get** /v1/users/{userId}/resources/{resourceUri}/permissions/{permission} | Verify user authorization
*UserPermissions* | [**GetUserPermissionsForResource**](./UserPermissions.md#getuserpermissionsforresource) | **Get** /v1/users/{userId}/resources/{resourceUri}/permissions | Get user permissions for resource
*UserPermissions* | [**GetUserResources**](./UserPermissions.md#getuserresources) | **Get** /v1/users/{userId}/resources | List user resources
*UserPermissions* | [**GetUserRolesForResource**](./UserPermissions.md#getuserrolesforresource) | **Get** /v1/users/{userId}/resources/{resourceUri}/roles | Get user roles for resource
*Users* | [**DeleteUser**](./Users.md#deleteuser) | **Delete** /v1/users/{userId} | Delete a user
*Users* | [**GetUser**](./Users.md#getuser) | **Get** /v1/users/{userId} | Retrieve a user
*Users* | [**GetUsers**](./Users.md#getusers) | **Get** /v1/users | List users


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

