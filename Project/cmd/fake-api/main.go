package main
import (
	// "fmt"
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/warisamir/golang/Project/internal/config"
)

func main() {
	cfg:=config.MustLoad()
	//router setup
	router:=http.NewServeMux()
	router.HandleFunc("GET /", func(rw http.ResponseWriter,r *http.Request){
		rw.Write([]byte("Welcome to server api"))
	})
	server:=http.Server{
		Addr:cfg.HTTPServer.Addr,
		Handler: router,
	}
	slog.Info("server started %s",slog.String("address",cfg.Addr))
	fmt.Printf("server started %s ,%s",cfg.HTTPServer,cfg.Addr)
	done :=make(chan os.Signal,1)
	signal.Notify(done,os.Interrupt,syscall.SIGINT,syscall.SIGTERM)
	go func() {
		err:= server.ListenAndServe()
		if err!=nil{
			log.Fatal("failed to start server")
		}
	}()
	<- done 
	slog.Info("shutting down the server")
	ctx,cancel:= context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()
	err:=server.Shutdown(ctx); if err!=nil{
		slog.Error("Failed to shutdown server",slog.String("error", err.Error()))
	}
	slog.Info("server shutdown successfully")
}