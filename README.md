# w3bstream

## Arch

![w3bstream](__doc__/modules_and_dataflow.png)

## Features

1. wasm applet management
2. wasm runtime instance deployment
3. interact with wasm (a word count demo)

## How to run

### Dependencies:

- os : macOS(11.0+)
- docker: to start a postgres

### Setup:
Initiate protocols and database

```sh
make run_depends # start postgres and mqtt
make migrate     # create or update schema
```

boot up the w3bstream server
```sh
make run_server
```


### Account Creation:

Please keep the terminal alive, and open a new terminal for following interactions with the server

```sh
make create_admin // if admin already created, skip this step
```

you are expected to see the output like below
```sh
> username: admin
> password: {$password}
> please remember it
```

### Login (fetch auth token):

```sh
curl -X PUT localhost:8888/srv-applet-mgr/v0/login -d '{"username":"admin","password":"{password}"}'
```

you are expected to see the output like below

```json
{
  "accountID": "{account_id}",
  "expireAt": "2022-09-23T07:20:08.099601+08:00",
  "issuer": "srv-applet-mgr",
  "token": "{token}"
}
```

### Project creation:
```sh
curl -X POST localhost:8888/srv-applet-mgr/v0/project -H "Authorization: Bearer {token}" -d '{"name":"{project_name}","version":"0.0.1"}'
```
you are expected to see the output like below
```json
{
  "accountID": "{account_id}",
  "createdAt": "2022-09-23T07:26:52.013626+08:00",
  "name": "{project_name}",
  "projectID": "{project_id}",
  "updatedAt": "2022-09-23T07:26:52.013626+08:00",
  "version": "0.0.1"
}
```

### Applet Deploy
Please prepare the compiled wasm file before uploaded to the server. Wasm examples are given under the folder `./pkg/modules/vm/testdata` for reference

1. upload wasm
```sh
curl -X POST localhost:8888/srv-applet-mgr/v0/applet -H "Authorization: Bearer {token}" -F info='{"projectID":"{project_id}","appletName":"{your_applet_name}"}' -F file=@{path_to_wasm_file}
```

you are expected to see the output like below

```json
{
  "appletID": "{applet_id}",
  "config": null,
  "createdAt": "2022-09-23T07:37:08.101494+08:00",
  "name": "{applet_name}",
  "projectID": "{project_id}",
  "updatedAt": "2022-09-23T07:37:08.101494+08:00"
}
```

2. deploy applet
```sh
curl -X POST localhost:8888/srv-applet-mgr/v0/deploy/applet/{applet_id} -H "Authorization: Bearer {token}"
```

### Run applet 
```sh
curl -X PUT localhost:8888/srv-applet-mgr/v0/deploy/{instance_id}/START -H "Authorization: Bearer {token}" 
```

### Push an event to server
```sh
curl --location --request POST 'localhost:8888/srv-applet-mgr/v0/event/{project_id}/{applet_id}/start' \
--header 'publisher: {publisher_id}' \
--header 'Content-Type: text/plain' \
--data-raw 'input a test sentence'
```
`publisher_id` is the identifier of the events' publisher