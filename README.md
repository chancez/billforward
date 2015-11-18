# go client for BillForward

[![GoDoc](https://godoc.org/github.com/authclub/billforward?status.png)](https://godoc.org/github.com/authclub/billforward)

*Status: Unstable*

Just sketching out the API, please reach out to paul@scaleft.com before using.

# Development

## Installation

### Checkout repository

Clone repository

```
git clone https://github.com/ecnahc515/billforward.git billforward-go-sdk
```

### Install Go

From [GoLang](https://golang.org/dl/)

#### Set `GOPATH`

Append the working directory of this project to your `GOPATH`.

You could put this in your `.bash_profile` or equivalent:

```bash
# path to where you cloned repository
export GOPATH="$GOPATH:$HOME/billforward-go-sdk"
```

You could also re-source the profile in the running command prompt:

```bash
source ~/.bash_profile
```

### Install `go-swagger`

Now that you have `go`, install [`go-swagger`](https://github.com/go-swagger/go-swagger):

```bash
go get -u github.com/go-swagger/go-swagger/cmd/swagger
```

## Generating client bindings

Run the script included in the repo to generate client bindings:

```bash
./generate_client
```