package app

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"net/http"
	"os/signal"
	"syscall"
)

func Run() {
	localserver := &http.Server{
		Addr: "localhost:8181",
	}
	ctx, cancelFunc := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancelFunc()

	conn, err := pgxpool.Connect(ctx, "postgresql://postgres:postgres@localhost/wb_tz_2")
	if err != nil {
		log.Fatal(err)
	}
	s := server.NewServer(ctx, localserver, conn)
	s.Init()
	s.Run(ctx)
	<-ctx.Done()
	conn.Close()
	localserver.Close()
}
