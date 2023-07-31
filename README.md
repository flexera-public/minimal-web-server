# Minimal Web Server

A minimal web server that you can use for testing purposes!

Usage:

```
$ cd minimal-web-server
$ go run main.go
```

This will start an HTTP server that will listen on 0.0.0.0:8080.

Then you can make HTTP request to the root path:

```
$ curl http://localhost:8080/
```

You can also run it using Docker:

```
$ docker build -t minimal-web-server:latest .
$ docker run -p 8080:8080 minimal-web-server:latest
```

Docker repository: https://gallery.ecr.aws/m9j9u5c4/minimal-web-server
