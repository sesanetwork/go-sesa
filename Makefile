.PHONY: all
all: sesa

GOPROXY ?= "https://proxy.golang.org,direct"
.PHONY: sesa
sesa:
	GIT_COMMIT=`git rev-list -1 HEAD 2>/dev/null || echo ""` && \
	GIT_DATE=`git log -1 --date=short --pretty=format:%ct 2>/dev/null || echo ""` && \
	GOPROXY=$(GOPROXY) \
	go build \
	    -ldflags "-s -w -X github.com/sesanetwork/go-sesa/cmd/sesa/launcher.gitCommit=$${GIT_COMMIT} -X github.com/sesanetwork/go-sesa/cmd/sesa/launcher.gitDate=$${GIT_DATE}" \
	    -o build/sesa \
	    ./cmd/sesa


TAG ?= "latest"
NET ?= "mainnet"
.PHONY: sesa-image
sesa-image:
	curl -O https://raw.githubusercontent.com/sesanetwork/sesa-genesis/main/$(NET).g
	docker build \
    	    --network=host \
    	    -f ./docker/Dockerfile.sesa -t "sesa:$(TAG)" .

.PHONY: test
test:
	go test ./...

.PHONY: coverage
coverage:
	go test -coverprofile=cover.prof $$(go list ./... | grep -v '/gossip/contract/' | grep -v '/gossip/emitter/mock' | xargs)
	go tool cover -func cover.prof | grep -e "^total:"

.PHONY: fuzz
fuzz:
	CGO_ENABLED=1 \
	mkdir -p ./fuzzing && \
	go run github.com/dvyukov/go-fuzz/go-fuzz-build -o=./fuzzing/gossip-fuzz.zip ./gossip && \
	go run github.com/dvyukov/go-fuzz/go-fuzz -workdir=./fuzzing -bin=./fuzzing/gossip-fuzz.zip


.PHONY: clean
clean:
	rm -fr ./build/*
