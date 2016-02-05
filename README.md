Go HTTP2 Demo
==========

* Intro
* Structure
* Certificates
* Getting started
* Build
* Caveats
* .gitignore notes


# Intro


This project is a demo of HTTP2 (with TLS) in Golang.  
In order for me to experiment with reproducible builds it also uses [gb](https://getgb.io) as the
build/vendoring/retrival mechanism instead of `go get` or `godeps`. (For one import, http2)

| Demo             | Executable Name
|------------------|:---------------|
| HTTP2 with TLS   |    tls         |
| HTTP2 no TLS     |    notls       |
| HTTP1.1 with TLS |    http1.1     |


# Structure
The layout of this repository uses the `gb` convention of creating a `src` directory with "my code"
and a `vendor` directory for "other peoples code".  
The `bin` directory is where `gb build` will output the compiled binaries for each demo.

A [gb plugin named "fetch"](https://getgb.io/news/gb-vendor-2015-05-26/) is used to retrieve the code we need to build.

In this demo, the only thing not in the standard library is http2 itself.

# Creating certificate and key for the demos
You can create one in each demo directory under `src` or simply copy-paste the same pair to each directory.

Create a certificate and private key to be used for the http2 demo. No password on the private key with the -nodes option.  
`openssl req -x509 -newkey rsa:2048 -nodes -keyout localhost.key -out localhost.cert -days 365`

# Getting started
NOTE! gb projects do not live in $GOPATH  
In fact you do not even need $GOPATH or $GOBIN environment variables set to run gb build!

First time usage:  
`gb vendor fetch golang.org/x/net/http2`  
Note that this will create a manifest file in vendor dir if you want to review it.

Additional usage from then on:  
`gb vendor update --all`  
You will notice the time stamps of various files in the vendor/src/directory get updated with the current time each time this is ran.

### Build all the demos
`gb build all`

This will place each demo in the `bin` directory.  
Run each binary to see what they are doing behind the scenes.

```
cd bin
# run these one at a time or all at once in different shells if desired
./tls
./notls
./http1.1
```

# Caveats
Note that http2 without TLS is supported in the specification, but not supported in any widely
dispersed user agents, so I have aimed to try to make it work here.

# A note on .gitignore
The .gitignore file excludes a couple project dirs that we would probably want checked into source code in a production environment.  
Other than, that this project layout would work the same everywhere, even without a $GOPATH.

```
.DS_Store
*.cert
*.key
# in production environment the directories below would not be in .gitignore
vendor/
pkg
```
