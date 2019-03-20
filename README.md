# Phil

This project is a companion to [punxatawny](https://github.com/WinterTechForum/punxatawny).
It is a simple "hello world" style web service app written in Go.  

## Build

### Dependencies
Dependencies are described via `dep` in `Gopkg.toml`

```bash
dep ensure
```

### Local

```bash
go build phil.go
```

### For Docker

We ran into a few problems getting this app to run in a standard Alpine container.
For the most part, this came down to Alpine uses musl rather than glibc.  
Googling found [this article](https://blog.codeship.com/building-minimal-docker-containers-for-go-applications/)
that described statically linking the entire Go app, 
then building a Docker image based on the "scratch" container.  The file `build.sh`
in this directory encapsulates building the go app and creating the Docker container.

```bash
./build.sh
```
