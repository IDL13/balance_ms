package requests

import (
	"context"
	"fmt"
	"github.com/IDL13/balance_ms/internal/CSV"
	"github.com/IDL13/balance_ms/internal/config"
	"github.com/IDL13/balance_ms/pkg/postgresql"
	"os"
)

func New() *Request {
	return &Request{
		conf: config.New(),
	}
}

type DataStruct struct {
	Id        int
	idService string
	idOrder   string
	money     string
}

type Request struct {
	conf *config.Config
}

func (r *Request) CreateUserTableRequest() error {
	conf := r.conf.GetConf()

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

func (r *Request) CreateReserveTableRequest() error {
	conf := r.conf.GetConf()

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

func (r *Request) AddBalanceRequest(id int64, balance string) (err error) {
	conf := r.conf.GetConf()

	conn, err := postgresql.NewClient(*conf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when trying to connect to the database:%e", err)
		os.Exit(1)
	}

	q := `INSERT INTO Users (Id, balance) VALUES ($, $)`

	_, err = conn.Exec(context.Background(), q, balance)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when trying to exec data in database:%e", err)
		os.Exit(1)
	}

	return nil
}

func (r *Request) GetBalanceRequest(balance string) (b string, err error) {
	conf := r.conf.GetConf()

	conn, err := postgresql.NewClient(*conf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when trying to connect to the database:%e", err)
		os.Exit(1)
	}

	q := `SELECT balance FROM Users WHERE Id = $`

	row, err := conn.Query(context.Background(), q, balance)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when trying to exec data in database:%e", err)
		os.Exit(1)
	}
	row.Scan(&balance)

	return balance, nil
}

func (r *Request) AddReserveRequest(id int, idService, idOrder, money string) error {
	conf := r.conf.GetConf()

	conn, err := postgresql.NewClient(*conf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when trying to connect to the database:%e", err)
		os.Exit(1)
	}

	q := `INSERT INTO Reserve (Id, idService, idOrder, money) VALUES ($, $))`

	_, err = conn.Exec(context.Background(), q)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when trying to exec data in database:%e", err)
		os.Exit(1)
	}

	return nil
}

func (r *Request) GetReserveRequest() (re []DataStruct, err error) {
	conf := r.conf.GetConf()

	conn, err := postgresql.NewClient(*conf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when trying to connect to the database:%e", err)
		os.Exit(1)
	}

	q := `SELECT * FROM Reserve`

	row, err := conn.Query(context.Background(), q)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when trying to exec data in database:%e", err)
		os.Exit(1)
	}

	var d DataStruct
	var request []DataStruct

	for row.Next() {
		err = row.Scan(&d.Id, &d.idService, &d.idOrder)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when trying to write data in struct line 115:%e", err)
			os.Exit(1)
		}
		var record []string
		record = append(record, d.idService, d.money, TimeNow())
		err = CSV.WriteInCSV(record)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when trying to write data in CSV file")
			os.Exit(1)
		}

		request = append(request, d)
	}

	return request, nil
}
