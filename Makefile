PACKAGES = $(shell go list ./... | grep -v /vendor/)

all: test_unit vet build

build_contract:
	abigen --sol=contract/contract.sol --pkg=main --out=contract/contract.go
