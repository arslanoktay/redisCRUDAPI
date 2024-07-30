package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"redisCRUDAPI/application"
)

func main() {
	app := application.New(application.LoadConfig())

	// takes a context and signal, return context if signal created notify
	// contexts works like a tree, if parent context stops, childs will be stop aswell
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt) // background context is risky, only has to be use in main or so so. Can block many things
	//cancel func will be finish the context
	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		fmt.Println("failed to start app:", err)
	}

}
