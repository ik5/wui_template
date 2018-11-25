package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/ik5/wui_template/config"
	"github.com/ik5/wui_template/db"
	"github.com/ik5/wui_template/gqlserver"
	"github.com/ik5/wui_template/logging"
	"github.com/ik5/wui_template/restserver"
	"github.com/spf13/viper"
)

func initServer(server *http.Server) {
	r := chi.NewRouter()
	cors := cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"X-API-TAG"},
		AllowCredentials: true,
		MaxAge:           3000, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Mount("/r", restserver.RestRouter())
	r.Mount("/gql", gqlserver.GQLRouter())

	addr := viper.GetString("address")
	port := viper.GetInt("port")
	addr = fmt.Sprintf("%s:%d", addr, port)
	server = &http.Server{Addr: addr, Handler: r}
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}

var logFile *os.File

func initLogger() {
	logging.InitLog(viper.GetString("syslog_socket_type"),
		viper.GetString("syslog_address"),
		viper.GetString("syslog_tag"),
		config.SyslogLevel(),
	)

	debug := viper.GetString("env") == "debug"
	var err error
	if debug {
		logFile, err = os.OpenFile(viper.GetString("log_file"), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			logging.Logger.SetOutput(os.Stdout)
		} else {
			logging.Logger.SetOutput(logFile)
		}
	}

	// Set the default logger to use our logger :)
	middleware.DefaultLogger = middleware.RequestLogger(&middleware.DefaultLogFormatter{
		Logger: logging.Logger, NoColor: true,
	})

}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	configPath := []string{"."}
	err := config.Init("yaml", "config", configPath)
	if err != nil {
		panic(err)
	}

	initLogger()
	if logFile != nil {
		defer logFile.Close()
	}

	err = db.Init()
	if err != nil {
		logging.Logger.Fatalf("Unable to init database: %s", err)
		os.Exit(-1)
	}
	defer db.DB.Close()

	quit := make(chan bool, 1)
	server := &http.Server{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		select {
		case <-c:
			fmt.Fprintf(os.Stderr, "Ctrl+C")
			server.Shutdown(ctx)
			quit <- true
		case <-ctx.Done():
			fmt.Fprintf(os.Stderr, "Context Done")
			server.Shutdown(ctx)
			quit <- true
		}
	}()

	go initServer(server)
	<-quit
}
