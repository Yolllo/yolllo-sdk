package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"yolllo-sdk-cmd/utils"

	"github.com/ElrondNetwork/elrond-sdk-erdgo/data"
	"github.com/ElrondNetwork/elrond-sdk-erdgo/interactors"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("error: not enough args")
		return
	}
	mnemonicStr := os.Args[1]
	indexStr := os.Args[2]
	index, err := utils.StringToUint32(indexStr)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	w := interactors.NewWallet()
	privateKey := w.GetPrivateKeyFromMnemonic(data.Mnemonic(mnemonicStr), 0, index)
	address, err := w.GetAddressFromPrivateKey(privateKey)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	if len(privateKey) == 32 {
		privateKey = append(privateKey, address.AddressBytes()...)
	}
	fmt.Println(hex.EncodeToString(privateKey))
}
