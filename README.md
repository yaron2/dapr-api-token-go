# Calling Dapr with an API token in self-hosted mode

## Prerequisites

* [Dapr installed locally](https://docs.dapr.io/getting-started/install-dapr-cli/)
* [Go](https://go.dev/doc/install)
* OpenSSL - Optional

## Getting Started

### 1. Create an API token

The API token can be any string value. The following example shows a random base64 token generation with OpenSSL:

```bash
openssl rand 16 | base64
```

### 2. Set the token and start Dapr

Take the token generated above and set it as an environment variable on the terminal session you will use to start Dapr:

```bash
export DAPR_API_TOKEN=<TOKEN>
```

Start a Dapr instance:

```bash
dapr run --app-id myapp --dapr-http-port 3500 --dapr-grpc-port 50001
```

### 3. Set the token and start the app

Set the token as an environment variable for your app to read when it sends requests to Dapr. This needs to be set in the terminal session you will use to start the `main.go` file:

*Note: you can choose any environment variable name when setting the token for your app*

```bash
export DAPR_TOKEN=<TOKEN>
```

Run the app:

```bash
go mod tidy
go run main.go
```

*Note: If you clear the `DAPR_TOKEN` env var for your app, you'll see the call to Dapr fails. Otherwise it succeeds.*

### Refreshing the token

To refresh the token, simply generate a new token (like in Step 1 above), kill the `daprd` and app process, export the new token and start the `daprd` and app processes again.