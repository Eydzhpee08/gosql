package customers

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

)

var ErrNotFound = errors.New("item not found")

var ErrInternal = errors.New("internal error")

type Service struct {
	db *sql.DB
	items []*Customer
}

func NewService(db *sql.DB) *Service {
	return &Service{db: db}
}

type Customer struct {
	ID      int64
	Name    string
	Phone   string
	Active  string
	Created time.Time
}

func (s *Service) ByID(ctx context.Context, id int64) (*Customer, error) {
	item := &Customer{}
	err := s.db.QueryRowContext(ctx, `
		SELECT id, name, phone, active, created FROM customers WHERE id = $1
	`,id).Scan(&item.ID,&item.Name, &item.Phone,&item.Active, &item.Created)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}

	if err != nil {
		log.Print(err)
		return nil, ErrInternal
	}

	return item,nil

}
func (s *Service) All(ctx context.Context) ([]*Customer, error) {
	
	items := s.items
		rows,err := s.db.QueryContext(ctx, `
			SELECT id, name, phone, active, created FROM customers ORDER BY id DESC
		`)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}

	for rows.Next() {
		item := &Customer{}
		err := rows.Scan(&item.ID,&item.Name, &item.Phone, &item.Active, &item.Created)
		if err != nil {
			log.Print(err)
			return nil, ErrNotFound
		}
		items = append(items, item)
	}
	if err != nil {
		log.Print(err)
		return nil, ErrInternal
	}
	
	return items,nil
}
func (s *Service) AllActive(ctx context.Context) ([]*Customer, error) {
	
	items := s.items
		rows,err := s.db.QueryContext(ctx, `
			SELECT id, name, phone, active, created FROM customers WHERE active
		`)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}

	for rows.Next() {
		item := &Customer{}
		err := rows.Scan(&item.ID,&item.Name, &item.Phone, &item.Active, &item.Created)
		if err != nil {
			log.Print(err)
			return nil, ErrNotFound
		}
		items = append(items, item)
	}
	if err != nil {
		log.Print(err)
		return nil, ErrInternal
	}
	
	return items,nil
}
func (s *Service) Save(ctx context.Context, item *Customer) (*Customer, error) {
	items := &Customer{}
	err := s.db.QueryRowContext(ctx, `
	INSERT INTO customers(name,phone) VALUES($1,$2) ON CONFLICT (phone) DO UPDATE SET name = excluded.name RETURNING id, name, phone, active, created
	`,item.Name, item.Phone).Scan(&items.ID,&items.Name, &items.Phone,&items.Active, &items.Created)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}

	if err != nil {
		log.Print(err)
		return nil, ErrInternal
	}

	return items,nil

}
