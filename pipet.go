package pipet

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/jinzhu/gorm"
)

type Pipet struct {
	Driver        string
	ReadDatabases []*gorm.DB
	WriteDatabase *gorm.DB
}

func Init(driver string) *Pipet {
	p := Pipet{
		Driver: driver,
	}

	return &p
}

func (s *Pipet) SetWrite(host string, port string, user string, password string, database string) {
	log.Print("initialize write db...")
	db, err := gorm.Open(s.Driver, fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true&loc=Local", user, password, host, port, database))
	if err != nil {
		panic("Failed to connect to write database")
	}
	log.Print("successfuly connected to write db")

	s.WriteDatabase = db

	if len(s.ReadDatabases) < 1 {
		s.ReadDatabases = append(s.ReadDatabases, db)
	}
}

func (s *Pipet) SetRead(host string, port string, user string, password string, database string) {
	log.Print("initialize read db...")
	db, err := gorm.Open(s.Driver, fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true&loc=Local", user, password, host, port, database))
	if err != nil {
		panic("failed to connect to read database")
	}
	log.Print("successfuly connected to read db")

	s.ReadDatabases = append(s.ReadDatabases, db)

	if s.WriteDatabase == nil {
		s.WriteDatabase = db
	}
}

func (s *Pipet) Where(query interface{}, args ...interface{}) *gorm.DB {
	db := s.selectRead()
	return db.Where(query, args)
}

func (s *Pipet) Save(value interface{}) *gorm.DB {
	return s.WriteDatabase.Save(value)
}

func (s *Pipet) Create(value interface{}) *gorm.DB {
	return s.WriteDatabase.Create(value)
}

func (s *Pipet) Exec(sql string, values ...interface{}) *gorm.DB {
	return s.WriteDatabase.Exec(sql, values)
}

func (s *Pipet) Conn() *gorm.DB {
	return s.Instance("write")
}

func (s *Pipet) Instance(conn string) *gorm.DB {
	if conn == "read" {
		return s.selectRead()
	}

	return s.WriteDatabase
}

// Private functions
func (s *Pipet) selectRead() *gorm.DB {
	i := rand.Intn(len(s.ReadDatabases))
	db := s.ReadDatabases[i]
	return db
}
