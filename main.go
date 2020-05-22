package main

//go:generate go run github.com/prisma/prisma-client-go generate

import (
	"context"
	"log"
	"time"

	"github.com/steebchen/photon-example/complex_relations"
	"github.com/steebchen/photon-example/db"
	"github.com/steebchen/photon-example/simple"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	client := db.NewClient()
	err := client.Connect()
	check(err)
	defer func() {
		err := client.Disconnect()
		check(err)
	}()

	ctx := context.Background()

	log.Printf("")
	log.Printf("---")
	log.Printf("")

	log.Printf("")
	log.Printf("simple raw 1")
	log.Printf("")

	start := time.Now()
	simple.Raw(ctx, client)
	log.Printf("simple raw 1 %s", time.Now().Sub(start))

	log.Printf("")
	log.Printf("---")
	log.Printf("")

	log.Printf("")
	log.Printf("simple raw 2")
	log.Printf("")

	start = time.Now()
	simple.Raw(ctx, client)
	log.Printf("simple raw 2 %s", time.Now().Sub(start))

	log.Printf("")
	log.Printf("simple query")
	log.Printf("")

	start = time.Now()
	simple.Query(ctx, client)
	log.Printf("simple query %s", time.Now().Sub(start))

	log.Printf("")
	log.Printf("---")
	log.Printf("")

	log.Printf("")
	log.Printf("---")
	log.Printf("")

	log.Printf("")
	log.Printf("---")
	log.Printf("")

	log.Printf("")
	log.Printf("complex_relations raw")
	log.Printf("")

	start = time.Now()
	complex_relations.Raw(ctx, client)
	log.Printf("complex_relations raw %s", time.Now().Sub(start))

	log.Printf("")
	log.Printf("complex_relations query")
	log.Printf("")

	start = time.Now()
	complex_relations.Query(ctx, client)
	log.Printf("complex_relations query %s", time.Now().Sub(start))

	log.Printf("")
	log.Printf("---")
	log.Printf("")
}
