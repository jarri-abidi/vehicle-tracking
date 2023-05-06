package postgres_test

import (
	"context"
	"fmt"
	"log"
	"os"
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
		log.Fatalf("Could not construct pool: %s", err)
	}

	// uses pool to try to connect to Docker
	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	container, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "11",
		Env: []string{
			"POSTGRES_USER=test",
			"POSTGRES_PASSWORD=test",
			"POSTGRES_DB=karma",
			"listen_addresses = '*'",
		},
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	})

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		var err error
		conn, err = pgx.Connect(context.Background(), fmt.Sprintf("postgres://test:test@127.0.0.1:%s/karma?sslmode=disable", container.GetPort("5432/tcp")))
		if err != nil {
			return err
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()
		return conn.Ping(ctx)
	}); err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}

	err = postgres.Migrate("file://migrations", "127.0.0.1", container.GetPort("5432/tcp"), "test", "test", "sslmode=disable")
	if err != nil {
		log.Fatalf("could not perform migrations: %s", err)
	}

	code := m.Run()

	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(container); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}
