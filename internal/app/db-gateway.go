package app

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)

type DbGateway struct {
	sync.Mutex
	cnf Database
	db  *sql.DB
}

func NewDbGateway(cnf Database) *DbGateway {
	g := &DbGateway{
		cnf: cnf,
	}

	return g
}

func (g *DbGateway) Connection() (*sql.DB, error) {
	g.Mutex.Lock()
	defer g.Mutex.Unlock()

	if nil != g.db {
		return g.db, nil
	}

	db, err := sql.Open(g.cnf.Driver, g.cnf.Dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	g.db = db

	return db, nil
}

func (g *DbGateway) Close() error {
	g.Mutex.Lock()
	defer g.Mutex.Unlock()

	if nil != g.db {
		return g.db.Close()
	}

	return nil
}
