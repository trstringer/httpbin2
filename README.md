# httpbin2

Relay web server for microservices environments and injecting quality of service issues.

## Usage

```
HTTP server for testing, demo'ing, and learning

Usage:
  httpbin2 [flags]
  httpbin2 [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  version     Application version

Flags:
      --delay-max-ms int       High end for response delay (default 0)
      --delay-min-ms int       Low end for response delay (default 0)
      --delay-rate int         Delay rate (default 100) (default 100)
  -h, --help                   help for httpbin2
      --message string         Message for server to return
      --message-hostname       Return hostname as message to requests
  -p, --port int               Listening port (default 8080)
      --relay strings          Relay routes to request on trigger (can be specified multiple times)
      --status-code int        HTTP status  code to return (default 200) (default 200)
      --status-code-rate int   Status code rate (default 100) (default 100)

Use "httpbin2 [command] --help" for more information about a command.
```

### Basic web server

Run a web server that returns a message:

```
$ httpbin2 --message "hello world"
```

Add hostname in the message, which is helpful in seeing which instance of the application served the request directly from the response:

```
$ httpbin2 --message "hello world" --message-hostname
```

### Failure injection

Inject a percentage of failed responses:

```
$ httpbin2 --message "some failures" --status-code 500 --status-code-rate 50
```

The above returns HTTP/500 half the time (and HTTP/200 the other half).

### Delay injection

Inject delays:

```
$ httpbin2 \
    --message "some delays" \
    --delay-min-ms 1000 \
    --delay-max-ms 5000 \
    --delay-rate 50
```

This delays half of the requests for a duration from 1 to 5 seconds.

### Multiple web server hops

You can chain web servers together with relays:

```
$ httpbin2 -p 8081 --message "first"
$ httpbin2 -p 8082 --message "second"
$ httpbin2 --relay "http://localhost:8081" --relay "http://localhost:8082"
```

This is helpful if you're trying to model multiple network hops.
