package requests

import (
	"context"
	"fmt"
	"github.com/IDL13/balance_ms/internal/config"
	"github.com/IDL13/balance_ms/pkg/postgresql"
	"os"
	"time"
)

const (
	timeLayout      = "2006-01-02T15:04:00"
	timeLayoutShort = "2006-01"
)

func TimeNow() string {
	now := time.Now().Format(timeLayoutShort)
	return now
}

func CreateUserTableRequest() error {
	conf := config.New().GetConf()

	conn, err := postgresql.NewClient(*conf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when trying to connect to the database:%e", err)
		os.Exit(1)
	}

	q := `CREATE TABLE Users (Id INT NOT NULL AUTO_INCREMENT, balance VARCHAR(255))`

	_, err = conn.Exec(context.Background(), q)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when trying to exec data in database:%e", err)
		os.Exit(1)
	}

	return nil
}

func CreateReserveTableRequest() error {
	conf := config.New().GetConf()

	conn, err := postgresql.NewClient(*conf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when trying to connect to the database:%e", err)
		os.Exit(1)
	}

	q := `CREATE TABLE Reserve (Id INT NOT NULL AUTO_INCREMENT, idService VARCHAR(255), idOrder VARCHAR(255), money VARCHAR(255))`

	_, err = conn.Exec(context.Background(), q)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when trying to exec data in database:%e", err)
		os.Exit(1)
	}

	return nil
}

func DellRequest(id int) error {
	conf := config.New().GetConf()

	conn, err := postgresql.NewClient(*conf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when trying to connect to the database:%e", err)
		os.Exit(1)
	}

	del := `DELETE FROM reserve WHERE id = $1`

	_, err = conn.Query(context.Background(), del, id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when trying to exec data in database:%e", err)
		os.Exit(1)
	}

	return nil
}
