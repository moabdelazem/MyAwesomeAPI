package main

import (
	"log"
	"time"

	"github.com/moabdelazem/sample-app/internal/db"
	"github.com/moabdelazem/sample-app/internal/env"
	"github.com/moabdelazem/sample-app/internal/store"
)

func main() {
	// Load the environment variables
	env.LoadEnv()

	// Define a new Config struct
	srvCfg := Config{
		Address: env.GetEnvVar("ADDRESS", ":8080", env.ParseString),
		DB: DBConfig{
			Address:             env.GetEnvVar("DB_ADDRESS", "", env.ParseString),
			MaxOpenConnections:  env.GetEnvVar("DB_MAX_OPEN_CONNECTIONS", 10, env.ParseInt),
			MaxIdleConnnections: env.GetEnvVar("DB_MAX_IDLE_CONNECTIONS", 10, env.ParseInt),
			MaxIdleTime:         env.GetEnvVar("DB_MAX_IDLE_TIME", time.Minute*15, time.ParseDuration),
		},
	}

	// Create a new connection to the database
	db, err := db.New(
		srvCfg.DB.Address,
		srvCfg.DB.MaxOpenConnections,
		srvCfg.DB.MaxIdleConnnections,
		srvCfg.DB.MaxIdleTime,
	)

	// Handle any errors
	if err != nil {
		log.Fatalf("error connecting to the database: %v", err)
	}

	// Log The Connection Of The Database
	defer db.Close()
	log.Println("Database connection pool established")

	// Create New PSQL Storage
	storage := store.NewStorage(db)

	// Create a new instance of the Application struct
	app := Application{
		Config:  srvCfg,
		Storage: storage,
	}

	// Call the Run method on the Application struct
	log.Fatal(app.Run(app.Mount()))
}
