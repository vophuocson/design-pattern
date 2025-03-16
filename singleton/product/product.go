package product

import "sync"

var (
	DbConnOnce sync.Once
	conn       *DbConnection
)

type DbConnection struct {
}

func (s *DbConnection) Show() {}

func GetConnection() *DbConnection {
	DbConnOnce.Do(func() {
		if conn == nil {
			conn = &DbConnection{}
		}
	})
	return conn
}
