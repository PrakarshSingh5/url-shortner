package main

import (
    "log"

    "github.com/gin-gonic/gin"

    "github.com/PrakarshSingh5/url-shortner/backend/internal/config"
    "github.com/PrakarshSingh5/url-shortner/backend/internal/database"
    "github.com/PrakarshSingh5/url-shortner/backend/internal/handlers"
    "github.com/PrakarshSingh5/url-shortner/backend/internal/middleware"
    "github.com/PrakarshSingh5/url-shortner/backend/internal/repository"
)

func main() {
    cfg, err := config.Load()
    if err != nil {
        log.Fatalf("failed to load configuration: %v", err)
    }

    db, err := database.New(cfg.DBPath)
    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }
    defer func() {
        if cerr := db.Close(); cerr != nil {
            log.Printf("error closing database: %v", cerr)
        }
    }()

    if err := database.Migrate(db); err != nil {
        log.Fatalf("failed to run migrations: %v", err)
    }

    repo := repository.NewURLRepository(db)
    handler := handlers.New(repo, cfg)

    router := gin.New()
    router.Use(gin.Recovery(), middleware.RequestLogger(),middleware.Ratelimitter(), middleware.CORSMiddleware(cfg.AllowedOrigins))

    handler.RegisterRoutes(router)

    log.Printf("starting server on %s", cfg.Addr())
    if err := router.Run(cfg.Addr()); err != nil {
        log.Fatalf("server exited with error: %v", err)
    }
}
