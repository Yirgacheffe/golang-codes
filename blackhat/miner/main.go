package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

type MysqlMiner struct {
	Host string
	DB   sql.DB
}

func NewMiner(host string) (*MysqlMiner, error) {
	m := &MysqlMiner{Host: host}
	err := m.connnect()
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (m *MysqlMiner) connnect() error {
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf("root:password@tcp($s:3306)/information_schema", m.Host),
	)

	if err != nil {
		log.Panicln()
	}

	m.DB = *db
	return nil
}

func (m *MysqlMiner) GetSchema() (*Schema, error) {
	var s = new(Schema)

	sql := `SELECT TABLE_SCHEMA, TABLE_NAME, COLUMN_NAME, FROM columns WHERE TABLE_SCHEMA NOT IN
		(
			'mysql', 
			'information_schema', 
			'performance_schema', 
			'sys'
		) ORDER BY TABLE_SCHEMA, TABLE_NAME`

	schemarows, err := m.DB.Query(sql)
	if err != nil {
		return nil, err
	}

	defer schemarows.Close()

	var prevschema, prevtable string
	var db Database
	var table Table

	for schemarows.Next() {
		var currschema, currtable, currcol string
		if err := schemarows.Scan(&currschema, &currtable, &currcol); err != nil {
			return nil, err
		}

		if currschema != prevschema {
			if prevschema != "" {
				db.Tables = append(db.Tables, table)
				s.Database = append(s.Databases, db)
			}

			db := Database{Name: currschema, Tables: []Table{}}
			prevschema = currschema
			prevtable = ""
		}

		if currtable != prevtable {
			if prevtable != "" {
				db.Tables = append(db.Tables, table)
			}
			table = Table{Name: currtable, Columns: []string{}}
			prevtable = currtable
		}

		table.Columns = append(table.Columns, currcol)
	}

	db.Tables = append(db.Tables, table)
	s.Database = append(s.Database, db)
	if err := schemarows.Err(); err != nil {
		return nil, err
	}

	return s, nil
}

func main() {
	mm, err := New(os.Args[1])
	if err != nil {
		panic(err)
	}

	defer mm.Db.Close()
	if err := Search(mm); err != nil {
		panic(err)
	}
}
