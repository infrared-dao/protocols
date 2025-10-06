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
	rm -f $(BIN_DIR)/concrete
	go build -o $(BIN_DIR)/concrete -v -ldflags \
		"-X main.rev=$(version) -X main.bts=$(timestamp)" cmd/test/concrete/main.go			
	rm -f $(BIN_DIR)/d2
	go build -o $(BIN_DIR)/d2 -v -ldflags \
		"-X main.rev=$(version) -X main.bts=$(timestamp)" cmd/test/d2/main.go
	rm -f $(BIN_DIR)/d8x
	go build -o $(BIN_DIR)/d8x -v -ldflags \
		"-X main.rev=$(version) -X main.bts=$(timestamp)" cmd/test/d8x/main.go		
	rm -f $(BIN_DIR)/dolomite
	go build -o $(BIN_DIR)/dolomite -v -ldflags \
		"-X main.rev=$(version) -X main.bts=$(timestamp)" cmd/test/dolomite/main.go
	rm -f $(BIN_DIR)/etherfi
	go build -o $(BIN_DIR)/etherfi -v -ldflags \
		"-X main.rev=$(version) -X main.bts=$(timestamp)" cmd/test/etherfi/main.go
	rm -f $(BIN_DIR)/euler
	go build -o $(BIN_DIR)/euler -v -ldflags \
		"-X main.rev=$(version) -X main.bts=$(timestamp)" cmd/test/euler/main.go
	rm -f $(BIN_DIR)/ivx
	go build -o $(BIN_DIR)/ivx -v -ldflags \
		"-X main.rev=$(version) -X main.bts=$(timestamp)" cmd/test/ivx/main.go
	rm -f $(BIN_DIR)/kodiak
	go build -o $(BIN_DIR)/kodiak -v -ldflags \
		"-X main.rev=$(version) -X main.bts=$(timestamp)" cmd/test/kodiak/main.go
	rm -f $(BIN_DIR)/paddlefi
	go build -o $(BIN_DIR)/paddlefi -v -ldflags \
		"-X main.rev=$(version) -X main.bts=$(timestamp)" cmd/test/paddlefi/main.go		
	rm -f $(BIN_DIR)/pendle
	go build -o $(BIN_DIR)/pendle -v -ldflags \
		"-X main.rev=$(version) -X main.bts=$(timestamp)" cmd/test/pendle/main.go		
	rm -f $(BIN_DIR)/solvbtc
	go build -o $(BIN_DIR)/solvbtc -v -ldflags \
		"-X main.rev=$(version) -X main.bts=$(timestamp)" cmd/test/solvbtc/main.go
	rm -f $(BIN_DIR)/steer
	go build -o $(BIN_DIR)/steer -v -ldflags \
		"-X main.rev=$(version) -X main.bts=$(timestamp)" cmd/test/steer/main.go
	rm -f $(BIN_DIR)/wasabee
	go build -o $(BIN_DIR)/wasabee -v -ldflags \
		"-X main.rev=$(version) -X main.bts=$(timestamp)" cmd/test/wasabee/main.go
	rm -f $(BIN_DIR)/wasabi
	go build -o $(BIN_DIR)/wasabi -v -ldflags \
		"-X main.rev=$(version) -X main.bts=$(timestamp)" cmd/test/wasabi/main.go
    rm -f $(BIN_DIR)/winnieswap
	go build -o $(BIN_DIR)/winnieswap -v -ldflags \
		"-X main.rev=$(version) -X main.bts=$(timestamp)" cmd/test/winnieswap/main.go
	go build -o $(BIN_DIR)/webera -v -ldflags \
		"-X main.rev=$(version) -X main.bts=$(timestamp)" cmd/test/webera/main.go


lint:
	golangci-lint run

test: lint
	go test ./...

codegen:
	rm -rf internal/sc
	mkdir -p internal/sc
	abigen --abi assets/abis/erc20.abi --pkg sc --type ERC20 --out internal/sc/erc20.go
	abigen --abi assets/abis/4626.abi --pkg sc --type ERC4626 --out internal/sc/erc_4626.go
	abigen --abi assets/abis/concretevault.abi --pkg sc --type ConcreteVault --out internal/sc/concrete_vault.go
	abigen --abi assets/abis/croclperc20.abi --pkg sc --type CrocLPERC20 --out internal/sc/croc_lp_erc20.go
	abigen --abi assets/abis/crocquery.abi --pkg sc --type CrocQuery --out internal/sc/croc_query.go
	abigen --abi assets/abis/balancerbasepool.abi --pkg sc --type BalancerBasePool --out internal/sc/balancer_base_pool.go
	abigen --abi assets/abis/balancervault.abi --pkg sc --type BalancerVault --out internal/sc/balancer_vault.go
	abigen --abi assets/abis/beraborrowiw.abi --pkg sc --type BeraBorrowIW --out internal/sc/beraborrow_iw.go
	abigen --abi assets/abis/beraborrowcicv.abi --pkg sc --type BeraBorrowCICV --out internal/sc/beraborrow_cicv.go
	abigen --abi assets/abis/beraborrowsnect.abi --pkg sc --type BeraBorrowSNECT --out internal/sc/beraborrow_snect.go
	abigen --abi assets/abis/bulla.abi --pkg sc --type Bulla --out internal/sc/bulla.go
	abigen --abi assets/abis/aquabera.abi --pkg sc --type AquaBera --out internal/sc/aquabera.go
	abigen --abi assets/abis/d2vault.abi --pkg sc --type D2Vault --out internal/sc/d2_vault.go
	abigen --abi assets/abis/d8xpoolmanager.abi --pkg sc --type D8xPoolManager --out internal/sc/d8x_pool_manager.go
	abigen --abi assets/abis/d8xsharetoken.abi --pkg sc --type D8xShareToken --out internal/sc/d8x_share_token.go	
	abigen --abi assets/abis/etherfiaccountant.abi --pkg sc --type EtherfiAccountant --out internal/sc/etherfi_accountant.go
	abigen --abi assets/abis/etherfivault.abi --pkg sc --type EtherfiVault --out internal/sc/etherfi_vault.go
	abigen --abi assets/abis/eulervault.abi --pkg sc --type EulerVault --out internal/sc/euler_vault.go	
	abigen --abi assets/abis/aggregatorV3.abi --pkg sc --type AggregatorV3 --out internal/sc/aggregatorV3.go
	abigen --abi assets/abis/ivxlpmonitor.abi --pkg sc --type IVXLPMonitor --out internal/sc/ivx_lp_monitor.go
	abigen --abi assets/abis/uniswapv2pair.abi --pkg sc --type UniswapV2 --out internal/sc/uniswap_v2.go
	abigen --abi assets/abis/kodiakisland.abi --pkg sc --type KodiakIsland --out internal/sc/kodiak_island.go
	abigen --abi assets/abis/alphaprovault.abi --pkg sc --type AlphaProVault --out internal/sc/alpha_pro_vault.go	
	abigen --abi assets/abis/paddlefi.abi --pkg sc --type Paddlefi --out internal/sc/paddlefi.go
	abigen --abi assets/abis/pendlelpwrapper.abi --pkg sc --type PendleWrapper --out internal/sc/pendle_wrapper.go	
	abigen --abi assets/abis/solvbtc.abi --pkg sc --type SolvBTC --out internal/sc/solvbtc.go	
	abigen --abi assets/abis/steerpool.abi --pkg sc --type SteerPool --out internal/sc/steer_pool.go		
	abigen --abi assets/abis/wasabeevault.abi --pkg sc --type WasabeeVault --out internal/sc/wasabee_vault.go
    abigen --abi assets/abis/stickyvault.abi --pkg sc --type StickyVault --out internal/sc/sticky_vault.go
	abigen --abi assets/abis/weberavault.abi --pkg sc --type WeberaVault --out internal/sc/webera_vault.go
