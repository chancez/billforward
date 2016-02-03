# go client for BillForward

[![GoDoc](https://godoc.org/github.com/authclub/billforward?status.png)](https://godoc.org/github.com/authclub/billforward)

*Status: Unstable*

Just sketching out the API, please reach out to paul@scaleft.com before using.

# Development

## Installation

We expect you to have followed https://golang.org/doc/install to setup your Go environment.
At a minimum you must have a valid `$GOPATH` setup.

### Clone repository

```
git clone https://github.com/authclub/billforward.git $GOPATH/src/github.com/authclub/billforward
```

### Install `go-swagger`

Now that you have `go`, install [`go-swagger`](https://github.com/go-swagger/go-swagger):

```bash
go get -u github.com/go-swagger/go-swagger/cmd/swagger
```

## Generating client bindings

Run the script included in the repo to generate client bindings:

```bash
cd $GOPATH/src/github.com/authclub/billforward
./generate_client
```

After generating the client bindings run `go test ./...` to verify the generated
code is valid. There are no tests, but this step validates that `go-swagger` has
created syntactically valid Go code.

Once you've done some minimal validation, commit your changes to
`./generate_client` in it's own commit, and then create a new commit for the
changes to the bindings in `models` and `client`.

