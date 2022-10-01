
# ssh 1.0

Secure Shell from Direktiv

---
- #### Categories: network
- #### Image: gcr.io/direktiv/functions/ssh 
- #### License: [Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)
- #### Issue Tracking: https://github.com/direktiv-apps/ssh/issues
- #### URL: https://github.com/direktiv-apps/ssh
- #### Maintainer: [direktiv.io](https://www.direktiv.io) 
---

## About ssh

This function allows to start a secure shell (SSH) with a remote server. It supprots username/passname as well as SSH key authentication.  To sue SSH key authentication the certificate has to be provided by a secret useing the `files` input. The function can execute multiple commands.

### Example(s)
  #### Function Configuration
```yaml
functions:
- id: ssh
  image: gcr.io/direktiv/functions/ssh:1.0
  type: knative-workflow
```
   #### Password authentication
```yaml
- id: ssh 
  type: action
  action:
    function: ssh
    secrets: ["sshpwd"]
    input: 
      host:
        name: 10.100.6.16
        port: 2022
        verbose: true
      auth:
        username: direktiv
        password: jq(.secrets.sshpwd)
      commands:
      - command: date
  catch:
  - error: "*"
```
   #### SSH key authentication
```yaml
- id: ssh 
  type: action
  action:
    function: ssh
    secrets: ["sshcert"]
    input: 
      files:
      - name: cert
        data: |
          jq(.secrets.sshcert)
        mode: "0400"
      host:
        name: 10.100.6.16
      auth:
        username: direktiv
        certificate: jq(.secrets.sshcert)
      commands:
      - command: date
  catch:
  - error: "*"
```

   ### Secrets


- **sshcert**: Certificate to use for SSH key authentication
- **sshpwd**: Password for username / password authentication






### Request



#### Request Attributes
[PostParamsBody](#post-params-body)

### Response
  List of executed SSH commands.
#### Reponse Types
    
  

[PostOKBody](#post-o-k-body)
#### Example Reponses
    
```json
[
  {
    "result": "Sa 1. Okt 13:25:38 CEST 2021",
    "success": true
  }
]
```

### Errors
| Type | Description
|------|---------|
| io.direktiv.command.error | Command execution failed |
| io.direktiv.output.error | Template error for output generation of the service |
| io.direktiv.ri.error | Can not create information object from request |


### Types
#### <span id="post-o-k-body"></span> postOKBody

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| ssh | [][PostOKBodySSHItems](#post-o-k-body-ssh-items)| `[]*PostOKBodySSHItems` |  | |  |  |


#### <span id="post-o-k-body-ssh-items"></span> postOKBodySshItems

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| result | [interface{}](#interface)| `interface{}` | ✓ | |  |  |
| success | boolean| `bool` | ✓ | |  |  |


#### <span id="post-params-body"></span> postParamsBody

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| auth | [PostParamsBodyAuth](#post-params-body-auth)| `PostParamsBodyAuth` | ✓ | |  |  |
| commands | [][PostParamsBodyCommandsItems](#post-params-body-commands-items)| `[]*PostParamsBodyCommandsItems` |  | `[{"command":"echo Hello"}]`| Array of SSH commands. |  |
| files | [][DirektivFile](#direktiv-file)| `[]apps.DirektivFile` |  | | File to create before running commands. |  |
| host | [PostParamsBodyHost](#post-params-body-host)| `PostParamsBodyHost` | ✓ | |  |  |


#### <span id="post-params-body-auth"></span> postParamsBodyAuth

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| certificate | string| `string` |  | | Path to certificate for SSH key based authentication |  |
| password | string| `string` |  | | Password if username/password authentication |  |
| username | string| `string` | ✓ | | Username for authentication |  |


#### <span id="post-params-body-commands-items"></span> postParamsBodyCommandsItems

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| command | string| `string` |  | | Command to run |  |
| continue | boolean| `bool` |  | | Stops excecution if command fails, otherwise proceeds with next command |  |
| print | boolean| `bool` |  | `true`| If set to false the command will not print the full command with arguments to logs. |  |
| silent | boolean| `bool` |  | | If set to false the command will not print output to logs. |  |


#### <span id="post-params-body-host"></span> postParamsBodyHost

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| name | string| `string` | ✓ | | Hostname or IP to SSH to |  |
| port | integer| `int64` |  | `22`| SSH port |  |
| verbose | boolean| `bool` |  | | Enables debug output on the connection |  |

 
