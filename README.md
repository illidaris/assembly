# assembly

```Makefile
# Binary names
APP_NAME=DEMO
APP_VERSION=V1.0.0
ARGS=argsparam
HIDE_ARGS=hideargsparam
OUTDIR=./bin
BINARY_NAME_LINUX=$(OUTDIR)/${APP_NAME}
BINARY_NAME_WIN=$(OUTDIR)/${APP_NAME}.exe

# ldflags
PARAM_LDFLAGS="-X 'github.com/illidaris/assembly.CommitID=$(shell git log -1 --pretty=format:"%H")'\
-X 'github.com/illidaris/assembly.Name=${APP_NAME}'\
-X 'github.com/illidaris/assembly.Version=${APP_VERSION}'\
-X 'github.com/illidaris/assembly.CommitAuthor=$(shell git log -1 --pretty=format:"%an")'\
-X 'github.com/illidaris/assembly.CommitTime=$(shell git log -1 --pretty=format:"%cd" --date=unix)'\
-X 'github.com/illidaris/assembly.BuildTime=$(shell date +%s)'\
-X 'github.com/illidaris/assembly.Args=${ARGS}'" \
-X 'github.com/illidaris/assembly.HideArgs=${HIDE_ARGS}'" \
-X 'github.com/illidaris/assembly.BuildNumber=${BUILD_NUMBER}'\
-X 'github.com/illidaris/assembly.BuildJob=${JOB_NAME}'"

# linux
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -ldflags $(PARAM_LDFLAGS) -o $(BINARY_NAME_LINUX) -v
```