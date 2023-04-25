# assembly

```Makefile
# Binary names
APP_NAME=DEMO
APP_VERSION=V1.0.0
OUTDIR=./bin
BINARY_NAME_LINUX=$(OUTDIR)/${APP_NAME}
BINARY_NAME_WIN=$(OUTDIR)/${APP_NAME}.exe

# ldflags
PARAM_LDFLAGS="-X 'euler/pkg/assembly.CommitID=$(shell git log -1 --pretty=format:"%H")'\
-X 'euler/pkg/assembly.Name=${APP_NAME}'\
-X 'euler/pkg/assembly.Version=${APP_VERSION}'\
-X 'euler/pkg/assembly.CommitAuthor=$(shell git log -1 --pretty=format:"%an")'\
-X 'euler/pkg/assembly.CommitTime=$(shell git log -1 --pretty=format:"%cd" --date=unix)'\
-X 'euler/pkg/assembly.BuildTime=$(shell date +%s)'\
-X 'euler/pkg/assembly.BuildNumber=${BUILD_NUMBER}'\
-X 'euler/pkg/assembly.BuildJob=${JOB_NAME}'"

# linux
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -ldflags $(PARAM_LDFLAGS) -o $(BINARY_NAME_LINUX) -v
```