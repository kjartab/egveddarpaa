package main

import (
	"net/http"
	"html"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"fmt"
)

func main() {
	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)

	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(133700000)}
	sim := backends.NewSimulatedBackend(alloc)

	// deploy contract
	addr, _, contract, err := DeployWinnerTakesAll(auth, sim, big.NewInt(10), big.NewInt(time.Now().Add(2 * time.Minute).Unix()), big.NewInt(time.Now().Add(5 * time.Minute).Unix()))
	if err != nil {
		log.Fatalf("could not deploy contract: %v", err)
	}
	_ = addr
	_ = contract

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
