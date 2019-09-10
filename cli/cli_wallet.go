package cli

import (
	"fmt"
	"log"
	"github.com/atoyr/glockchain"
)

func (cli *CLI) createWallet() {
	wallets := glockchain.NewWallets()
	address := wallets.CreateWallet()
	wallets.SaveToFile()

	fmt.Printf("address: %s\n", address)
}

func (cli *CLI) printWallets() {
	wallets := glockchain.NewWallets()
	for address := range wallets.Wallets {
		fmt.Printf("address: %s\n", address)
	}
}

func (cli *CLI) getBalance(address string) {
	if !glockchain.ValidateAddress([]byte(address)) {
		log.Panic("ERROR: Address is not valid")
	}
	utxopool, err := glockchain.GetUTXOPool()
	if err != nil {
		log.Fatal(err)
	}
	pubKeyHash := glockchain.AddressToPubKeyHash([]byte(address))
	balance, _ := utxopool.FindUTXOs(pubKeyHash)
	fmt.Printf("Balance of %s : %d \n", address, balance)
}

func (cli *CLI) getAllBalance() {
	wallets := glockchain.NewWallets()
	utxopool, err := glockchain.GetUTXOPool()
	if err != nil {
		log.Fatal(err)
	}
	for address := range wallets.Wallets {
		pubKeyHash := glockchain.AddressToPubKeyHash([]byte(address))
		balance, _ := utxopool.FindUTXOs(pubKeyHash)
		fmt.Printf("Balance of %s : %d \n", address, balance)
	}
}
