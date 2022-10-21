package main

import (
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
	"os"
)

const txGastPrice = "1000000000"
const txGasLimit = "50000"
const txChainID = "yolllo-network"
const txVersion = "1"

func main() {
	if len(os.Args) != 6 {
		fmt.Println("error: not enough args")
		return
	}

	txNonce := os.Args[1]
	txValue := os.Args[2]
	txReceiver := os.Args[3]
	txSender := os.Args[4]
	txPkey := os.Args[5]

	privKey, err := hex.DecodeString(txPkey)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	signData := `{"nonce":` + txNonce +
		`,"value":"` + txValue +
		`","receiver":"` + txReceiver +
		`","sender":"` + txSender +
		`","gasPrice":` + txGastPrice +
		`,"gasLimit":` + txGasLimit +
		`,"chainID":"` + txChainID +
		`","version":` + txVersion + `}`
	signature := ed25519.Sign(ed25519.PrivateKey([]byte(privKey)), []byte(signData))
	fmt.Println(hex.EncodeToString(signature))
}
