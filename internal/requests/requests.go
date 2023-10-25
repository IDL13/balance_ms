package requests

import (
	"context"
	"fmt"
	"os"

	"github.com/IDL13/balance_ms/internal/CSV"
	"github.com/IDL13/balance_ms/internal/config"
	"github.com/IDL13/balance_ms/pkg/postgresql"
)

func New() Request {
	return &request{
		conf: config.New(),
	}
}

type DataStruct struct {
	Id        int
	IdService string
	IdOrder   string
	Money     string
}

//go:generate mockgen -source=requests.go -destination=../../mock/mock.go -package=mock_serv
type Request interface {
	AddBalanceRequest(id int64, balance string) (err error)
	GetBalanceRequest(id int64) (b string, err error)
	AddReserveRequest(id int64, idService, idOrder, money string) error
	GetReserveRequest() (re []DataStruct, err error)
}

type request struct {
	conf *config.Config
}

func (r request) AddBalanceRequest(id int64, balance string) (err error) {
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

func (r *request) GetBalanceRequest(id int64) (b string, err error) {
	var balance string

	conf := r.conf.GetConf()

	conn, err := postgresql.NewClient(*conf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when trying to connect to the database:%e", err)
		os.Exit(1)
	}

	q := `SELECT balance FROM users WHERE id = $1`

	row := conn.QueryRow(context.Background(), q, id)

	row.Scan(&balance)

	return balance, nil
}

func (r *request) AddReserveRequest(id int64, idService, idOrder, money string) error {
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

func (r *request) GetReserveRequest() (re []DataStruct, err error) {
	conf := r.conf.GetConf()

	conn, err := postgresql.NewClient(*conf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when trying to connect to the database:%e", err)
		os.Exit(1)
	}

	q := `SELECT id, idservice, money FROM Reserve`

	row, err := conn.Query(context.Background(), q)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when trying to exec data in database:%e", err)
		os.Exit(1)
	}

	var d DataStruct
	var request []DataStruct

	for row.Next() {
		err = row.Scan(&d.Id, &d.IdService, &d.Money)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when trying to write data in struct line 152:%e", err)
			os.Exit(1)
		}
		var record []string
		record = append(record, d.IdService, d.Money, TimeNow())
		err = CSV.WriteInCSV(record)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when trying to write data in CSV file")
			os.Exit(1)
		}

		err = DellRequest(d.Id)
		if err != nil {
			fmt.Println(err)
		}

		request = append(request, d)
	}
	return request, nil
}
