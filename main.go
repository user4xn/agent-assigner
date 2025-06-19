package main

import (
	"agent-assigner/internal/api"
	"agent-assigner/internal/consumer"
	"agent-assigner/internal/factory"
	"agent-assigner/pkg/util"
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {
	var (
		m string
	)

	flag.StringVar(
		&m,
		"m",
		"",
		`This flag is used for mode`,
	)

	flag.Parse()

	ctx := context.Background()

	if m != "" {
		// consumer mode will consume messages from the queue
		if m == "consumer" {
			fmt.Println("Starting consumer")
			con := consumer.NewConsumer(factory.NewFactory(ctx).BuildConsumerFactory())
			con.Init()
			return
		}

		fmt.Println("Error: Unknown mode")
		return
	}

	// initiate router framework
	r := chi.NewRouter()

	// add CORS middleware
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Origin", "X-Requested-With"},
		ExposedHeaders:   []string{"Link", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	r.Use(cors.Handler)

	// setup api router group bind to chi router & factory
	api.NewAPI(r, factory.NewFactory(ctx).BuildRestFactory())

	// start server
	port := util.GetEnv("SERVER_PORT", "8080")
	log.Printf("Starting server on port %s", port)

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
