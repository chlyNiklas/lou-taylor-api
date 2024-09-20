# lou-taylor-api

lou-taylor-api is the **future** backend for [lou-taylor.ch](http://lou-taylor.ch)


## open api

[oapi](./.openapi/README.md)

## configuration

The server can be configured using a toml [cofigfile](./config.example.toml) or 
using flags.


## config.toml

Per default lou-taylor-api searches for the config file at the path "./config.toml",
this can be overwritten with the config flag:

```sh
lou-taylor-api -config config.example.toml  
```

## flags

```
Usage of ./lou-taylor-api:
  -admin-name string
        administrator username (default "admin")
  -admin-password string
        administrator password (default "password")
  -base-ulr string
        base url of the api (default "localhost:8080")
  -config string
        path of the toml configuration file (default "./config.toml")
  -db-host string
        host of the PostrgreSQL server (default "localhost")
  -db-name string
        name of the database (default "event_db")
  -db-password string
        password for database (default "password")
  -db-port int
        port for the database connection (default 5432)
  -db-user string
        username for database (default "username")
  -image-max-width int
        maximum width for uploaded images (default 2096)
  -image-quality float
        image quality (e.g., 80 for 80%) (default 80)
  -image-save-path string
        path to save uploaded images (default ".tmp/")
  -jwt-secret string
        JWT secret key for signing tokens (default "my secret")
  -jwt-valid-period duration
        valid period for JWT tokens (default 24h0m0s)
```
