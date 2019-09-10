package cli

import (
	"fmt"
	"github.com/atoyr/glockchain"
	"log"
)

func (cli *CLI) createTransaction(from, to string, amount int) {
	wallets := glockchain.NewWallets()
	wallet := wallets.Wallets[from]
	if wallet == nil {
		log.Fatal(glockchain.NewGlockchainError(94001))
	}
	tx, err := glockchain.NewTransaction(wallet, []byte(to), amount)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tx.String())
}

func (cli *CLI) printTransactionPool() {
	txp, err := glockchain.NewTransactionPool()
	if err != nil {
		log.Fatal(err)
	}
	for _, tx := range txp.Pool {
		fmt.Println(tx.String())
		fmt.Println()
	}
}
