package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

// PostgresRepository -
type PostgresRepository struct {
	db *sqlx.DB
}

// NewPostgresRepository - Initializes a new Repo for connecting to Postgres
func NewPostgresRepository(db *sqlx.DB) *PostgresRepository {
	return &PostgresRepository{db}
}

// Create -
func (r *PostgresRepository) Create(ctx context.Context, user *User) error {
	user.ID = uuid.NewV4().String()
	query := "insert into users (id, name, email, company, password) values ($1, $2, $3, $4, $5)"
	_, err := r.db.ExecContext(ctx, query, user.ID, user.Name, user.Email, user.Company, user.Password)
	return err
}

// GetByID -
func (r *PostgresRepository) GetByID(ctx context.Context, id string) (*User, error) {
	var user *User
	if err := r.db.GetContext(ctx, &user, "select * from users where id = $1", id); err != nil {
		return user, err
	}
	return user, nil
}

// GetByEmail -
func (r *PostgresRepository) GetByEmail(ctx context.Context, email string) (*User, error) {
	var user *User
	if err := r.db.GetContext(ctx, &user, "select * from users where email = $1", email); err != nil {
		return user, err
	}
	return user, nil
}

// GetAll -
func (r *PostgresRepository) GetAll(ctx context.Context) ([]*User, error) {
	users := make([]*User, 0)
	if err := r.db.GetContext(ctx, users, "select * from users"); err != nil {
		return users, err
	}
	return users, nil
}

// NewConnection returns a new database connection instance
func NewConnection() (*sqlx.DB, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")
	conn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable", host, user, dbName, password)
	log.Println(conn)
	db, err := sqlx.Connect("postgres", conn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
