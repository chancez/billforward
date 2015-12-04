# go client for BillForward

[![GoDoc](https://godoc.org/github.com/authclub/billforward?status.png)](https://godoc.org/github.com/authclub/billforward)

*Status: Unstable*

Just sketching out the API, please reach out to paul@scaleft.com before using.

# Development

## Installation

### Install Go

From [GoLang](https://golang.org/dl/)

#### Set `GOPATH`

Ensure there is a folder allocated on your computer as your Go workspace.

For example `~/go`.

Append the path to that workspace to your `GOPATH`.

You could put this in your `.bash_profile` or equivalent:

```bash
export GOPATH="$HOME/go"
```

### Clone repository

Understand that GOPATH uses directory structure as a means of namespacing, so you will need to create some directories under your `GOPATH` (for example under `~/go`):

`src/github.com/authclub/`

Clone repository into `authclub`:

```
git clone https://github.com/ecnahc515/billforward.git
```

Such that a path exists like:

```bash
$GOPATH/src/github.com/authclub/billforward
```

#### Set `GOROOT`

Find where `go` is installed. This should get you pretty close:

```bash
which go
```

Export `GOROOT` to the root of the Go package. Something like this:

```bash
export GOROOT="/usr/local/go"
```

#### Update your PATH to include Go executables:

This will give you access to Go on command prompt (if you've not already), and any Go modules you install:

```bash
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```

#### Re-source your profile

To get the updated environment variables exported to your running command prompt:

```bash
source ~/.bash_profile
```

### Install `go-swagger`

Now that you have `go`, install [`go-swagger`](https://github.com/go-swagger/go-swagger):

```bash
go get -u github.com/go-swagger/go-swagger/cmd/swagger
cd $GOPATH/src/github.com/go-swagger/go-swagger/cmd/swagger
git checkout f88d96a
rm $GOPATH/bin/swagger
go install github.com/go-swagger/go-swagger/cmd/swagger
```

We re-install with a specific known working version to ensure consistency between
builds. In the future this shouldn't be necessary.

## Generating client bindings

Run the script included in the repo to generate client bindings:

```bash
./generate_client
```