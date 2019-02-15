# Health Check

This is just a very simple app that is meant to reply on root with a 200 ok status.

Intent is to be used as a very lightweight health check against a server.

## Run
To run it, there is only one flag `-port` to define which port the server will run on,
runs on port 8000 by default if no port is provided.

## Build  

To build you just need:  
```bash
go build -o health-check
```

Or it uses a docker [multi-stage build](https://docs.docker.com/develop/develop-images/multistage-build/)
so you can just do a `docker build -t health-check .` to get a local build.

## Development  
Not much to add here, this is a very basic go webserver, just one file.