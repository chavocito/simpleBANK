package main

import (
	"context"
	"log"
	"simple-bank/ent"
	"simple-bank/ent/migrate"

	_ "github.com/lib/pq"
)

func main() {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=admin dbname=simpleBANK password=secret sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	ctx := context.Background()

	createAccount(client, ctx)

	err = client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithGlobalUniqueID(true),
	)

	client.Account.Get()

	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

func createAccount(client *ent.Client, ctx context.Context) {
	account, err := client.Account.
		Create().
		SetOwner("Aboabi-Ba").
		SetAge(50).
		SetBalance(1500.00).
		SetCurrency("cedi").
		SetCountryCode(233).
		Save(ctx)

	if err != nil {
		log.Fatalf("error creating account: %v", err)
	}

	log.Printf("account created: %v", account)
}
