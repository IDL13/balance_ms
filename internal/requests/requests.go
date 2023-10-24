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
	IdService string
	IdOrder   string
	Money     string
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

	q := `INSERT INTO users (id, balance) VALUES ($1, $2)`

	_, err = conn.Exec(context.Background(), q, id, balance)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when trying to exec data in database:%e", err)
		os.Exit(1)
	}

	return nil
}

func (r *Request) GetBalanceRequest(id int64) (b string, err error) {
	var balance string

	conf := r.conf.GetConf()

	conn, err := postgresql.NewClient(*conf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when trying to connect to the database:%e", err)
		os.Exit(1)
	}

	q := `SELECT balance FROM Users WHERE Id = $1`

	row, err := conn.Query(context.Background(), q, id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when trying to exec data in database:%e", err)
		os.Exit(1)
	}
	row.Scan(&balance)

	return balance, nil
}

func (r *Request) AddReserveRequest(id int64, idService, idOrder, money string) error {
	conf := r.conf.GetConf()

	conn, err := postgresql.NewClient(*conf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when trying to connect to the database:%e", err)
		os.Exit(1)
	}

	q := `INSERT INTO reserve (Id, idService, idOrder, money) VALUES ($1, $2, $3, $4)`

	_, err = conn.Exec(context.Background(), q, id, idService, idOrder, money)
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
		err = row.Scan(&d.Id, &d.IdService, &d.IdOrder)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when trying to write data in struct line 115:%e", err)
			os.Exit(1)
		}
		var record []string
		record = append(record, d.IdService, d.Money, TimeNow())
		err = CSV.WriteInCSV(record)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when trying to write data in CSV file")
			os.Exit(1)
		}

		request = append(request, d)
	}

	return request, nil
}
