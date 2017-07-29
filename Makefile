PACKAGES = $(shell go list ./... | grep -v /vendor/)

all: test_unit vet build

run:
	go build . && ./egveddarpaa

build_contract:
	abigen --sol=contract/contract.sol --pkg=contract --out=contract/contract.go
