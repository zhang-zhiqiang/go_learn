package main

import (
    "context"
    "fmt"
    "golang.org/x/sync/errgroup"
    "net/http"
    "os"
    "os/signal"
)

func main() {
    ctx := context.Background()

    _, cancel := context.WithCancel(ctx)

    g, errCtx := errgroup.WithContext(ctx)

    app := &http.Server{Addr: ":8080"}
    g.Go(func() error {
        return ServeApp(app)
    })

    g.Go(func() error {
        <-errCtx.Done()
        return app.Shutdown(errCtx)
    })

    sigs := make(chan os.Signal, 1)
    signal.Notify(sigs)

    g.Go(func() error {
        for {
            select {
            case <-errCtx.Done():
                return errCtx.Err()
            case <-sigs:
                cancel()
            }
        }
    })

}

func ServeApp(app *http.Server) error {
    http.HandleFunc("/", HelloWord)

    return app.ListenAndServe()
}

func HelloWord(resp http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(resp, "Hello World")
}
