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
	"github.com/kjartab/egveddarpaa/config"
	"log"
	"fmt"
)

func main() {

	cfg := config.LoadEnvConfig()

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

	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		// Todo: pass arguments in request
		contract.SubmitProject(&bind.TransactOpts{
			From:     auth.From,
			Signer:   auth.Signer,
			GasLimit: big.NewInt(2381623),
			Value:    big.NewInt(10),
		}, "test project", "http://www.example.com")
		fmt.Fprint(w, "Contract submitted")
	})

	http.HandleFunc("/mine", func(w http.ResponseWriter, r *http.Request) {
		sim.Commit()
		// Todo: return current balance
		fmt.Fprintf(w, "Mined sucessfully, %q", html.EscapeString(r.URL.Path))
		})

	/*http.HandleFunc("/count", func(w http.ResponseWriter, r *http.Request) {
		numOfProjects, _ = contract.NumberOfProjects(nil)
		fmt.Fprintf(w, "%v projects", html.EscapeString(numOfProjects))
		})*/


	log.Fatal(http.ListenAndServe(cfg.HttpAddress, nil))

}
