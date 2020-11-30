package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"errors"

	"github.com/Eydzhpee08/gosql/pkg/types"
	_ "github.com/jackc/pgx/v4/stdlib"

)

func main() {
	dsn := "postgres://app:pass@localhost:5432/db"

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Println(err)
		os.Exit(1)
		return
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Println(err)
		}
	}()

	ctx := context.Background()
	_, err = db.ExecContext(ctx, `
	CREATE TABLE IF NOT EXISTS customers (
		id			BIGSERIAL 	PRIMARY KEY,
		name 		TEXT 		NOT NULL,
		phone 		TEXT 		NOT NULL UNIQUE,
		active      BOOLEAN 	NOT NULL DEFAULT TRUE,
		created 	TIMESTAMP	NOT NULL DEFAULT CURRENT_TIMESTAMP
	);
`)
	if err != nil {
		log.Print(err)
		return
	}

	customer := &types.Customer{}

	id := 1
	newPhone := "+992000000099"

	err = db.QueryRowContext(ctx, `
	UPDATE customers SET phone = $2 WHERE id= $1 RETURNING id,name, phone, active, created
	`, id,newPhone).Scan(&customer.ID, &customer.Name, &customer.Phone,&customer.Active, &customer.Created)

	if errors.Is(err, sql.ErrNoRows) {
		log.Print("No rows")
		return
	}

	if err != nil {
		log.Print(err)
		return
	}
}
