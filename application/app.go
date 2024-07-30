package application

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

type App struct {
	router http.Handler
	rdb    *redis.Client
}

// ctor creating
func New() *App {
	app := &App{
		rdb: redis.NewClient(&redis.Options{}),
	}

	app.loadRoutes()

	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: a.router,
	}

	// check for is redis working
	err := a.rdb.Ping(ctx).Err()
	if err != nil {
		return fmt.Errorf("failed to connect to redis %w", err)
	}

	defer func() {
		if err := a.rdb.Close(); err != nil {
			fmt.Println("failed to close redis", err)
		}
	}()

	fmt.Println("server starts")

	ch := make(chan error, 1)

	go func() { // to not block main thread we're using goroutine
		err = server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("failed to start server: %w", err)
		}
		close(ch)
	}()

	select {
	case err = <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		return server.Shutdown(timeout)
	}
}

// Graceful Shutdown -> Using for redis, when a get signal shutdown db. Signtint with context library

// Context -> Using context to share values between scopes, deadlines or timeouts.
// For sharing value, funcs take same context as parameter, we can give value using context.WithValue(context,key,value). context.Value(key) will give the value
// Second is cancelation context.withcancel(context) will return context and cancel. if cancel triggers, context.done will be happen. We use select for that matters.
// timeouts are similiar, witchcancel(context, timeouts time) when time ups and context is still working, context.Done() will trigger
// context creating types  WithCancel, WithDeadline, WithTimeout, or WithValue
// When cancel happens stops all children contexts aswell
