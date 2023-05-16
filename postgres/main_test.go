package postgres_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jarri-abidi/vehicle-tracking/postgres"
	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
)

var conn *pgx.Conn

func TestMain(m *testing.M) {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %v", err)
	}

	// uses pool to try to connect to Docker
	if err = pool.Client.Ping(); err != nil {
		log.Fatalf("Could not connect to docker: %v", err)
	}

	container, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "alpine",
		Env: []string{
			"POSTGRES_USER=test",
			"POSTGRES_PASSWORD=test",
			"POSTGRES_DB=vehicle-tracking",
			"listen_addresses = '*'",
		},
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	})
	if err != nil {
		log.Fatalf("Could not start docker container: %v", err)
	}

	options := []string{"sslmode=disable"}
	connString := fmt.Sprintf("postgres://test:test@localhost:%s/vehicle-tracking?%s", container.GetPort("5432/tcp"), strings.Join(options, "&"))

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err = pool.Retry(func() error {
		var err error
		conn, err = pgx.Connect(context.Background(), connString)
		if err != nil {
			return err
		}

		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		return conn.Ping(ctx)
	}); err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	if err = postgres.Migrate("file://migrations", connString); err != nil {
		log.Fatalf("could not perform migrations: %v", err)
	}

	code := m.Run()

	// can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(container); err != nil {
		log.Fatalf("Could not purge resource: %v", err)
	}

	os.Exit(code)
}
