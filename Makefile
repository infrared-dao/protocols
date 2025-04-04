.PHONY: all build

version := $(shell git rev-parse --short=12 HEAD)
timestamp := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

ROOT_DIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
BIN_DIR := $(ROOT_DIR)/bin
version := $(or $(version), $(shell cat /app/build-release | tr -d '\n'))

all: build

build: codegen lint
	rm -f $(BIN_DIR)/aquabera
	go build -o $(BIN_DIR)/aquabera -v -ldflags \
		"-X main.rev=$(version) -X main.bts=$(timestamp)" cmd/test/aquabera/main.go		
	rm -f $(BIN_DIR)/beraborrow
	go build -o $(BIN_DIR)/beraborrow -v -ldflags \
		"-X main.rev=$(version) -X main.bts=$(timestamp)" cmd/test/beraborrow/main.go
	rm -f $(BIN_DIR)/bex
	go build -o $(BIN_DIR)/bex -v -ldflags \
		"-X main.rev=$(version) -X main.bts=$(timestamp)" cmd/test/bex/main.go
	rm -f $(BIN_DIR)/bexv2
	go build -o $(BIN_DIR)/bexv2 -v -ldflags \
		"-X main.rev=$(version) -X main.bts=$(timestamp)" cmd/test/bexv2/main.go
	rm -f $(BIN_DIR)/bulla
	go build -o $(BIN_DIR)/bulla -v -ldflags \
		"-X main.rev=$(version) -X main.bts=$(timestamp)" cmd/test/bulla/main.go
	rm -f $(BIN_DIR)/burrbear
	go build -o $(BIN_DIR)/burrbear -v -ldflags \
		"-X main.rev=$(version) -X main.bts=$(timestamp)" cmd/test/burrbear/main.go	
	rm -f $(BIN_DIR)/d8x
	go build -o $(BIN_DIR)/d8x -v -ldflags \
		"-X main.rev=$(version) -X main.bts=$(timestamp)" cmd/test/d8x/main.go		
	rm -f $(BIN_DIR)/dolomite
	go build -o $(BIN_DIR)/dolomite -v -ldflags \
		"-X main.rev=$(version) -X main.bts=$(timestamp)" cmd/test/dolomite/main.go
	rm -f $(BIN_DIR)/kodiak
	go build -o $(BIN_DIR)/kodiak -v -ldflags \
		"-X main.rev=$(version) -X main.bts=$(timestamp)" cmd/test/kodiak/main.go
	rm -f $(BIN_DIR)/wasabi
	go build -o $(BIN_DIR)/wasabi -v -ldflags \
		"-X main.rev=$(version) -X main.bts=$(timestamp)" cmd/test/wasabi/main.go

lint:
	golangci-lint run

test: lint
	go test ./...

codegen:
	rm -rf internal/sc
	mkdir -p internal/sc
	abigen --abi assets/abis/erc20.abi --pkg sc --type ERC20 --out internal/sc/erc20.go
	abigen --abi assets/abis/4626.abi --pkg sc --type ERC4626 --out internal/sc/erc_4626.go
	abigen --abi assets/abis/croclperc20.abi --pkg sc --type CrocLPERC20 --out internal/sc/croc_lp_erc20.go
	abigen --abi assets/abis/crocquery.abi --pkg sc --type CrocQuery --out internal/sc/croc_query.go
	abigen --abi assets/abis/balancerbasepool.abi --pkg sc --type BalancerBasePool --out internal/sc/balancer_base_pool.go
	abigen --abi assets/abis/balancervault.abi --pkg sc --type BalancerVault --out internal/sc/balancer_vault.go
	abigen --abi assets/abis/beraborrowiw.abi --pkg sc --type BeraBorrowIW --out internal/sc/beraborrow_iw.go
	abigen --abi assets/abis/beraborrowcicv.abi --pkg sc --type BeraBorrowCICV --out internal/sc/beraborrow_cicv.go
	abigen --abi assets/abis/beraborrowsnect.abi --pkg sc --type BeraBorrowSNECT --out internal/sc/beraborrow_snect.go	
	abigen --abi assets/abis/bulla.abi --pkg sc --type Bulla --out internal/sc/bulla.go
	abigen --abi assets/abis/aquabera.abi --pkg sc --type AquaBera --out internal/sc/aquabera.go
	abigen --abi assets/abis/d8xpoolmanager.abi --pkg sc --type D8xPoolManager --out internal/sc/d8x_pool_manager.go
	abigen --abi assets/abis/d8xsharetoken.abi --pkg sc --type D8xShareToken --out internal/sc/d8x_share_token.go	
	abigen --abi assets/abis/aggregatorV3.abi --pkg sc --type AggregatorV3 --out internal/sc/aggregatorV3.go
	abigen --abi assets/abis/kodiakvaultv1.abi --pkg sc --type KodiakV1 --out internal/sc/kodiak_v1.go
	abigen --abi assets/abis/uniswapv2pair.abi --pkg sc --type UniswapV2 --out internal/sc/uniswap_v2.go	
	