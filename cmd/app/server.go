package app

import (
	"github.com/ashtishad/banking/db"
	"github.com/ashtishad/banking/internal/domain"
	"github.com/ashtishad/banking/internal/rest"
	"github.com/ashtishad/banking/internal/service"
	"github.com/ashtishad/banking/pkg/lib"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

const port = ":5000"

func Start() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	// Database connection pool config
	db := db.GetDbClient()
	defer db.Close()

	customerDbConn := domain.NewCustomerRepoDb(db)

	// Wire up the handlers
	ch := rest.CustomerHandlers{Service: service.NewCustomerService(customerDbConn)}

	// Server configurations
	srv := &http.Server{
		Addr:           port,
		Handler:        r,
		IdleTimeout:    100 * time.Second,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// Route URL mapping
	getRouteMappings(r, ch)

	// Middlewares: custom logger and recovery middlewares
	r.Use(gin.LoggerWithFormatter(lib.Logger))
	r.Use(gin.CustomRecovery(lib.Recover))

	// start server
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("Could not listen on %s: %v", port, err)
			return
		}
	}()

	// graceful shutdown
	lib.GracefulShutdown(srv)
}
