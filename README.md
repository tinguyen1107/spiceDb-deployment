# Authzed SpiceDB

## Start

#### Steps

1. Clone this repo.
1. Run `make start`

#### Current Problem:

- Sleep time to waiting for postgre db start up. (set in `make file`)
- Did not set cert for spice db => just can connect to spice db via http protocol. Recommend to config tls for using grpc, it quite better than http request. You can un-comment this part and generate tls certificate + key in `encryptions` folder.

```
      # - --grpc-tls-cert-path
      # - /app/encryptions/cert
      # - --grpc-tls-key-path
      # - /app/encryptions/key
```

#### Good points:

- Using postgre db to store relationship data. You can also remove the part below to use `in-memory` database.

```(yaml)
services:
  ...
  spicedb:
    ...
    command:
      ...
      - --datastore-bootstrap-files
      - /app/schema.yaml
      - --http-enabled
```

## Connect to Spice DB

- Available ports:
  - 50051: `network = tcp, service = grpc` This port allow to connect with grpc, btw it required tls cert
  - 8443: `service = http, insecure = true` This port allow to connect with http, do not require tls cert
  - 9090: `service = http, insecure = true`
