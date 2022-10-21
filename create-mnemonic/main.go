package main

import (
	"fmt"

	"github.com/ElrondNetwork/elrond-sdk-erdgo/interactors"
)

func main() {
	w := interactors.NewWallet()
	mnemonic, err := w.GenerateMnemonic()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(mnemonic))
}
