package main

import (
	"fmt"
	assignment01bca "github.com/codesbyusman/assignment01bca"
)

func main() {
	
	// create a new block chain
	chain := assignment01bca.InitiateBlockChain()

	// hard coded transaction
	transaction := []string{"Alice sent to bob", "Trudy is here", "Another Transaction", "a small one"}

	// adding the transaction to the block chain
	for i := 0; i < len(transaction); i++ {
		chain.AddBlock(transaction[i])
	}

	// printing the block chain
	fmt.Printf("\n\n\t ::::::::::::::: Initial Chain ::::::::::::::: \n\n")
	chain.ListBlocks()

	fmt.Printf("\n\n\t :::::::::::::: Changing Chain ::::::::::::::: \n\n")

	// changing the block
	chain.ChangeBlock()

	// printing the block chain
	fmt.Printf("\n\n\t ::::::::::::::: Verifying Chain ::::::::::::::: \n\n")

	// verifying the chain
	if chain.VerifyChain() {
		chain.ListBlocks()
	}

	fmt.Println()
	fmt.Println()

}