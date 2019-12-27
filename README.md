# inventoryflo-api
The backend orchestration layer to the InventoryFlo system.

## Development
This section describes the knowledge and tools needed to start making contributions to this project.
### Prerequisites
- Golang 1.12+
### Relevant Project Structure
```
.
|-- main.go (Entry point of the API. Lambda handler.)
|-- go.mod & go.sum (Go modules management files)
|-- pkg/ (Directory containing most of the code of the api)
    |-- auth/ (Module for creating and authorizing users)
    |-- data/ (Service layer that reads and writes data)
    |-- model/ (Contains all structs used in the code)
    |-- router/ (Routes client request to the appropriate modules)
    |-- util/ (Utilities used by other packages)
        |-- config (Environmental and configuration variables store)
        |-- db (Postgres abstraction module)
        |-- http (HTTP calls abstraction module)
|-- Makefile (GNU Make file with targets for build and deploying)
|-- ops/ (Directory containing deploy and secrets management scripts)
    |-- config.py (Configuration variables used during deployment)
    |-- deploy.py (Deploy script used by makefile)
    |-- rollback.py (Rollback script used by makefile)
    |-- encrypto.go (Used to encrypt secrets)
```
___
## Deployment
This section describes the knowledge and tools needed to deploy and manage the API.
### Prerequisites
- GNU Make
- Python3
- Boto3
- AWS Credentials configured locally https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-configure.html#cli-quick-configuration
### Make Targets
- **clean**: removes the binaries
- **build**: 1. *clean* 2. compiles and builds the binary
- **package**: 1. *build* 2. zips up the binary
- **deploy-[qa/prod]**: 1. *package* 2. deploys the code to qa or prod
- **rollback-[qa/prod]**: roll qa or prod back to the previous version
___
## Secrets Management
This section describes the knowledge and tools needed to encrypt and manage secrets needed by the code (e.g. db passwords and API tokens).
### Prerequisites
- Golang 1.12+
- AWS Credentials configured locally
- Have "AWS_DEFAULT_REGION" environment variable set to the region the API will be in. Currently 'us-east-2'
- Your secret must not have an '=' in it and must be less than 32 characters.
### Procedure
ENV will refer to the environment your are encrypting for i.e. qa or prod

*Encrypting*

1. ```cd ops/```
2. ```go build encrypt.go```
3. ```./encrypt <CIPHERTEXT_BLOB> <YOUR_SECRET>```
   - CIPHERTEXT_BLOB can be found in ops/config.py under env[ENV]['env-vars']['Variables']['data_key']
   - YOUR_SECRET is the secret you are trying to encrypt
4. The above step will print out your encrypted secret.

*Exposing your secret to the code*
1. Copy the encrypted secret
2. Open the file pkg/util/config/config.go
3. Find the *secrets* map.
4. Add the encrypted secret to the map in the appropriate ENV map with an appropriate name
5. To use your secret in the code, suppose you called the secret 'password' then you just call config.Get('password')

___
## Infrastructure
- Information about the architecture of the InventoryFlo System can be found here: https://github.com/herbal-goodness/inventoryflo-infra
- For getting access to the AWS account, contact Unoma Okorafor or Julius Phu
- Specific details about the resources used by this API can be found in ops/config.py