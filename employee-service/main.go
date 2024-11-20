package main

import (
	"context"
	"employee-service/config"
	"employee-service/handlers"
	"employee-service/proto"
	"employee-service/repositories"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func NewDBPool(connString string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.Connect(context.Background(), connString)
	if err != nil {
		return nil, err
	}
	return pool, nil
}

func RunMigrations(connStr string) error {
	m, err := migrate.New(
		fmt.Sprintf("file://%s", "./migrations"),
		connStr,
	)
	if err != nil {
		return fmt.Errorf("Error run migrations: %w", err)
	}

	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("Error up migrations: %w", err)
	}
	log.Println("Migrations applied successfully")
	return nil
}

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %w", err)
	}

	if err = RunMigrations(cfg.PostgresURL()); err != nil {
		log.Fatalf("Fail run migrations: %w", err)
	}

	pool, err := NewDBPool(cfg.PostgresURL())
	if err != nil {
		log.Fatalf("Error connecting to database: %w", err)
	}
	defer pool.Close()

	employeeRepo := repositories.NewEmployeeRepository(pool)
	employeeHandler := handlers.NewEmployeeHandler(*employeeRepo)

	grpcServer := grpc.NewServer()
	proto.RegisterEmployeeServiceServer(grpcServer, employeeHandler)

	reflection.Register(grpcServer)

	listener, err := net.Listen(cfg.ServicesNetworkType, cfg.EmployeePort)
	if err != nil {
		log.Fatalf("Failed to listen: %w", err)
	}
	log.Printf("Employee service started on %s", cfg.EmployeePort)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %w", err)
	}
}
