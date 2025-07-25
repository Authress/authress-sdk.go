## Start Development

* Install GO
* `go install github.com/go-delve/delve/cmd/dlv@latest`
* run commands from makefile
  * make integration

## Using the OpenAPI Generator to generate new models

#### Start container
```sh
podman pull docker://openapitools/openapi-generator-online 
```

#### Start container at port 8888 and save the container id
```sh
CID=$(podman run -d -p 8888:8080 openapitools/openapi-generator-online)
sleep 10

# Execute an HTTP request to generate a Ruby client
curl -X POST --header 'Content-Type: application/json' --header 'Accept: application/json' -d '{"openAPIUrl": "https://api.authress.io/", "options": { "useSingleRequestParameter": true, "packageName": "authress", "packageVersion": "99.99.99" } }' 'http://localhost:8888/api/gen/clients/go'


# RESPONSE: { "link":"http://localhost:8888/api/gen/download/b11ce82c-37dd-448b-b51b-d42bf3f04c2e" }
```

### Download the generated zip file
```sh

wget RESPONSE_LINK

# Unzip the file
unzip SHA

# Shutdown the openapi generator image
podman stop $CID && podman rm $CID
```

### Common review items
* [x] Inputs to Constructor are (string: authress_api_url, string: service_client_access_key)
* [ ] authress_api_url should sanitize https:// prefix and remove trailing `/`s
* [x] Add authentication to the configuration class.
* [x] Change configuration class name to be `AuthressSettings`
  * Specify all the inputs to be consistent across languages
* [x] constructors for classes should only have relevant input properties, for instances `links` are not required.
* [x] Update documentation
  * Make sure markdown is valid
  * Remove unnecessary links
  * Add first class examples to readme.md + api documentation
  * create alias types for UserId, RoleId, TenantId, GroupId, Action properties (OR convert every instance of the type assumption to be as string)
* [x] Remove any unnecessary validations from object and parameter injection, often there are some even when properties are allowed to be null
* [ ] The service client code to generate a JWT from private key needs to be added
* [ ] Add UnauthorizedError type to the authorizeUser function
* [x] GET query parameters should be an object
* [x] Top level tags from the API should accessible from the base class: `authressClient.accessRecords.getRecords(...)`
* [ ] Automatic Retry
  * [ ] Automatic fallback to cache
* [ ] `OptimisticPerformanceHandler` - Automatic fallback to cache on timeout reached
* [ ] In memory caching for authorization checks - memoize
* [x] Unsigned int for all limits
* [ ] readonly properties are never specified as required for request bodies
* [x] Update Documentation for the API
* [ ] Validate all enums are enums and can be null when they should be.
* [x] Remove LocalHost from the docs
* [x] Tests
* [x] If-unmodified-since should called `expectedLastModifiedTime`, accept string or dateTime and convert this to an ISO String
* [ ] Update OAuth2 openapi authentication references in the documentation
* [x] Find all incorrect references to version 99.99.99
* [x] Inject in the User-Agent 