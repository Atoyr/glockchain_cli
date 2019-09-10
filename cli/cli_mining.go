package cli

import (
	"fmt"
	"log"
	"time"
	"github.com/atoyr/glockchain"
)

func (cli *CLI) mining(address string) {
	wallets := glockchain.NewWallets()
	wallet, err := wallets.GetWallet(address)
	if err != nil {
		log.Fatal(err)
	}
	t := time.Now()
	fmt.Printf("Execute mining at %d-%d-%d %d:%d:%d\n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	txpool, err := glockchain.NewTransactionPool()
	if err != nil {
		log.Fatal(err)
	}
	if len(txpool.Pool) == 0 {
		return
	}
	bc, tip, err := glockchain.GetBlockchain()
	if err != nil {
		log.Fatal(err)
	}
	block, err := glockchain.NewBlock(txpool.Pool, bc, tip)
	if err != nil {
		log.Fatal(err)
	}
	bc.AddBlock(block)
	txpool.ClearTransactionPool()
	t = time.Now()
	fmt.Printf("\n\nFinished mining at %d-%d-%d %d:%d:%d\n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	fmt.Println("Add Block ")
	fmt.Println(block)
	tx, _ := glockchain.NewCoinbaseTX(100, wallet)
	txp, _ := glockchain.NewTransactionPool()
	up, _ := glockchain.GetUTXOPool()
	txp.AddTransaction(tx)
	up.AddUTXO(tx)
}
