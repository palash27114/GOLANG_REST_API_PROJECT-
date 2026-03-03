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

	// "github.com/palash27114/REST_API_PROJECT/internal/checker"
	"github.com/palash27114/REST_API_PROJECT/internal/config"
	"github.com/palash27114/REST_API_PROJECT/internal/http/handlers/student"
	"github.com/palash27114/REST_API_PROJECT/internal/storage/sqlite"
)

func main() {

	//load config
	cfg:=config.MustLoad()




	//db set up
	storage,err:=sqlite.New(cfg)
	if err!=nil{
		log.Fatal(err)
	}
	slog.Info("Storage initialised",slog.String("env",cfg.Env),slog.String("version","1.0.0"))

	
	//setup router
	router:=http.NewServeMux()


	router.HandleFunc("POST /api/students",student.New(storage))
	
	

	










	//setup server
	server:=http.Server{
		Addr:cfg.Address,
		Handler: router,
	}


	slog.Info("server started%s",slog.String("address",cfg.Address))
	fmt.Printf("started started %s",cfg.Address)

	done:=make(chan os.Signal,1)

	signal.Notify(done,os.Interrupt,syscall.SIGINT)



	go func(){
		err:=server.ListenAndServe()

	if err!=nil{
		panic("Problem 500 ")
	}
	}()

	<-done

	slog.Info("Shutting down server")
	ctx,cancel :=context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()

	if err:=server.Shutdown(ctx);err!=nil{
		slog.Error("failed to shutdown server",slog.String("error",err.Error()))
	}

	slog.Info("server shutdown successfully")



	
	




	
}