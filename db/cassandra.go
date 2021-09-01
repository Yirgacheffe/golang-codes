package main

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
)

const (
	insert_stmt = `INSERT INTO tweet(timeline, id, text) VALUES(?, ?, ?)`
	select_stmt = `SELECT id, text FROM tweet WHERE timeline = ? LIMIT 1`
	search_iter = `SELECT id, text FROm tweet WHERE timeline = ?`
)

func main() {

	cluster := gocql.NewCluster("192.168.1.1", "192.168.1.2", "192.168.1.3")
	cluster.Keyspace = "exam"
	cluster.Consistency = gocql.Quorum()

	session, _ := cluster.CreateSession()
	defer session.Close()

	if err := session.Query(insert_stmt, "me", gocql.TimeUUID(), "hello world").Exec(); err != nil {
		log.Fatal(err)
	}

	var id gocql.UUID
	var text string

	if err := session.Query(select_stmt, "me").Consistency(gocql.One).Scan(&id, &text); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Tweet: ", id, text)

	iter := session.Query(search_iter, "me").Iter()
	for iter.Scan(&id, &text) {
		fmt.Print("Tweet: ", id, text)
	}

	if err := iter.Close(); err != nil {
		log.Fatal("-------------------------------------------------0", err)
	}

}

// create keyspace example with replication = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };
// create table example.tweet(timeline text, id UUID, text text, PRIMARY KEY(id));
// create index on example.tweet(timeline);
