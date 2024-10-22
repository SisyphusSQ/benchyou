build:
	@echo "--> Building..."
	@mkdir -p bin/
	GOARCH=amd64 GOOS=linux go build -o bin/mybenchx src/bench/mybenchx.go
	@chmod 755 bin/*

darwin:
	@echo "--> Building..."
	@mkdir -p bin/
	go build -o bin/mybenchx src/bench/mybenchx.go
	@chmod 755 bin/*

clean:
	@echo "--> Cleaning..."
	@go clean
	@rm -f bin/*

fmt:
	go fmt ./...

test:
	@echo "--> Testing..."
	@$(MAKE) testxcmd
	@$(MAKE) testxworker
	@$(MAKE) testxcommon
	@$(MAKE) testsysbench

testxcmd:
	go test -v xcmd
testxworker:
	go test -v xworker
testxcommon:
	go test -v xworker
testsysbench:
	go test -v sysbench

# code coverage
COVPKGS =	sysbench\
			xcmd\
			xworker\
			xcommon
coverage:
	go build -v -o bin/gotestcover \
	src/vendor/github.com/pierrre/gotestcover/*.go;
	gotestcover -coverprofile=coverage.out -v $(COVPKGS)
	go tool cover -html=coverage.out
.PHONY: build clean fmt test coverage
