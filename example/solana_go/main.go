package main

import (
	"context"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

func main() {
	endpoint := rpc.TestNet_RPC
	client := rpc.New(endpoint)

	accountFrom, err := solana.PrivateKeyFromSolanaKeygenFile("/Users/cisco/.config/solana/id.json")
	if err != nil {
		panic(err)
	}

	latestBlockhash, err := client.GetLatestBlockhash(context.Background(), rpc.CommitmentFinalized)
	if err != nil {
		return
	}
	programId := solana.MustPublicKeyFromBase58("BJXgVUgXsSDiwW9yTtevgF29pKf1ssXEhZe7WZf2UuUb")
	var accounts []*solana.AccountMeta
	accounts = append(accounts, solana.Meta(programId))
	instruction := solana.NewInstruction(programId, accounts, []byte(""))

	txBuilder := solana.NewTransactionBuilder()
	txBuilder.SetRecentBlockHash(latestBlockhash.Value.Blockhash)
	txBuilder.SetFeePayer(accountFrom.PublicKey())
	txBuilder.AddInstruction(instruction)

}
