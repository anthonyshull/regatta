![logo](logo.svg)

Server software for conducting virtual ergometer regattas.

## INSTALLATION

## TERMINAL USER INTERFACE

*future release*

`rui` is a [terminal user interface](https://en.wikipedia.org/wiki/Text-based_user_interface) into Regatta that can be installed separately.
```
%> ...
```
You can attach to any running Regatta instance you control:
```
%> rui <user>:<password>@<uri>
```

## DEVELOPMENT
We use the [gorilla](https://www.gorillatoolkit.org/) toolkit with [JSON-RPC](https://www.jsonrpc.org/specification) and [Cassandra](https://cassandra.apache.org/).
Run the server in development mode:
```
%> go run ./cmd/server
```
### PHILOSOPHY

### STRUCTURE

### BUILDS
We use [mage](https://magefile.org/).

### CASSANDRA
Start a local Cassandra instance:
```
%> docker run --name cassandra -p 9042:9042 -d cassandra:latest
```
Access the instance:
```
%> docker exec -it cassandra cqlsh
```
Create the keyspace:
```cassandraql
CREATE KEYSPACE regatta WITH REPLICATION = {'class': 'SimpleStrategy', 'replication_factor': 1};
```
### MIGRATIONS

We use [journey](https://github.com/db-journey/journey/v2).
Raw migration commands:
```
%> journey --url cassandra://127.0.0.1:9042/regatta --path ./internal/migrations migrate create my_migration
%> journey --url cassandra://127.0.0.1:9042/regatta --path ./internal/migrations migrate up
```
The Mage equivalents executed from */internal/build*:
```
%> MIGRATION=my_migration mage migrate:create
%> mage migrate:up
```
### WEBSOCKETS
You can use any websocket client. We use [websocat](https://github.com/vi/websocat).
```
%> wget https://github.com/vi/websocat/releases/download/v1.5.0/websocat_1.5.0_ssl1.1_amd64.deb
%> sudo dpkg -i websocat_1.5.0_ssl1.1_amd64.deb
```
Get updates from a regatta by id:
```
%> websocat ws://127.0.0.1:9999/results/regatta/123
```