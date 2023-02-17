package main

import (
	"fmt"
	"os"
)

func main() {
	wd, _ := os.Getwd()
	wallet := walletapi.Wallet{
		Filename:   wd + "/test.wallet",
		Password:   "password",
		DaemonHost: "public.turtlenode.io",
		DaemonPort: 11898,
	}
	W := walletapi.InitWalletAPI("password", "127.0.0.1", "8070")

	fmt.Println("====CreateWallet====")
	fmt.Println(W.CreateWallet(&wallet))

	fmt.Println("====Status====")
	fmt.Println(W.Status())

	fmt.Println("====Node====")
	fmt.Println(W.Node())

	fmt.Println("====ViewKey====")
	fmt.Println(W.ViewKey())

	fmt.Println("====PrimaryAddress===")
	primary, _ := W.PrimaryAddress()
	fmt.Println(primary)

	fmt.Println("====Addresses====")
	addresses, _ := W.Addresses()
	fmt.Println(addresses)

	fmt.Println("====ValidateAddress====")
	fmt.Println(W.ValidateAddress(primary))

	fmt.Println("====GetBalance====")
	fmt.Println(W.GetBalance())

	fmt.Println("====GetBalances====")
	fmt.Println(W.GetBalances())

	fmt.Println("====GetKeys====")
	fmt.Println(W.GetKeys(primary))

	fmt.Println("====GetMnemonic====")
	fmt.Println(W.GetMnemonic(primary))

	/* example sending an advanced transaction
	fmt.Println("====SendTransactionAdvanced====")
	fmt.Println(W.SendTransactionAdvanced(
		[]map[string]interface{}{
			map[string]interface{}{
				"address": "TRTLuySpDqd2fcvq5vx7Jiayw6yao7JHXFPuia5V83cVREtQSKyvWpxX9vamnUcG35BkQy6VfwUy5CsV9YNomioPGGyVhK3YXLq",
				"amount":  100,
			},
		},
		nil, nil, []string{primary}, nil, nil, nil,
	))
	*/

	fmt.Println("====GetAllTransactions====")
	fmt.Println(W.GetAllTransactions())

	fmt.Println("====GetTransactionByHash====")
	fmt.Println(W.GetTransactionByHash("invalid hash"))

	fmt.Println("===GetTransactionPrivateKey====")
	fmt.Println(W.GetTransactionPrivateKey("invalid hash"))

	addr, _ := W.CreateAddress()
	fmt.Println("Created address:", addr)

	fmt.Println("====GetAddressTransactionsByStartHeight - primary wallet address====")
	fmt.Println(W.GetAddressTransactionsByStartHeight(primary, 1000))

	fmt.Println("====GetAddressTransactionsByStartHeight - new wallet address====")
	fmt.Println(W.GetAddressTransactionsByStartHeight(addr["address"], 1000))

	W.DeleteAddress(addr["address"])
	fmt.Println("Deleted address:", addr)

	W.CloseWallet()
	fmt.Println("====Done====")
}
