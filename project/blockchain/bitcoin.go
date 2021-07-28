package main

import (
        "fmt"
        "time"
        "crypto/sha256"
)

type Block struct {
     timestamp     time.Time
     transactions  []string
     prevHash      []byte
     Hash          []byte
}

func NewBlock(transactions []string, prevHash []byte) *Block {
     currentTime := time.Now()
     
     return &Block{
            timestamp: currentTime,
            transactions: transactions,
            prevHash: prevHash,
            Hash: NewHash(currentTime, transactions, prevHash),
     }    
}

func NewHash(cT time.Time, tX []string, pH []byte) []byte {

     input := append(pH, cT.String()...)

     for transaction := range tX {
         input = append(input, string(rune(transaction))...)
     }

     hash := sha256.Sum256(input)
     return hash[:]

}



func main() {
     genesisTransactions := []string{"Izzy sent Will 50 bitcoin", "Will sent Izzy 30 bitcoin"}
     genesisBlock := NewBlock(genesisTransactions, []byte{})
     fmt.Println("--- First Block ---")
     printBlockInformation(genesisBlock)
     
     block2Transactions := []string{"John sent Izzy 30 bitcoin"}
     block2 := NewBlock(block2Transactions, genesisBlock.Hash)
     fmt.Println("--- Second Block ---")
     printBlockInformation(block2)

     block3Transactions := []string{"Onkar sent Uttara 30 bitcoin"}
     block3 := NewBlock(block3Transactions, block2.Hash)
     fmt.Println("--- Third Block ---")
     printBlockInformation(block3)


}
 

func printBlockInformation(block *Block) {
     fmt.Printf("\t time: %s \n", block.timestamp.String())
     fmt.Printf("\t prevHash: %x \n", block.prevHash)
     fmt.Printf("\t Hash: %x \n", block.Hash)
     printTransaction(block)
}

func printTransaction(block *Block) {
     for i, tx := range block.transactions {
         fmt.Printf("\t %d Transaction: %s \n", i, tx)
     }

}


     
