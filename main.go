package main

import (
    "fmt"

    "github.com/gocql/gocql"
)

func main() {
    // Connect to the Cassandra cluster
    cluster := gocql.NewCluster("localhost")
    cluster.Keyspace = "test"
    session, err := cluster.CreateSession()
    if err != nil {
        panic(err)
    }
    defer session.Close()

    // Create the keyspace and table
    if err := session.Query("CREATE KEYSPACE IF NOT EXISTS test WITH REPLICATION = {'class' : 'SimpleStrategy', 'replication_factor' : 1}").Exec(); err != nil {
        panic(err)
    }
    if err := session.Query("CREATE TABLE IF NOT EXISTS test.users (id UUID, name text, age int, PRIMARY KEY (id))").Exec(); err != nil {
        panic(err)
    }

    // Insert some data
    id := gocql.TimeUUID()
    name := "Alice"
    age := 30
    if err := session.Query("INSERT INTO test.users (id, name, age) VALUES (?, ?, ?)", id, name, age).Exec(); err != nil {
        panic(err)
    }

    // Read the data back
    var retrievedName string
    var retrievedAge int
    if err := session.Query("SELECT name, age FROM test.users WHERE id = ?", id).Scan(&retrievedName, &retrievedAge); err != nil {
        panic(err)
    }
    fmt.Printf("Retrieved name: %s, age: %d\n", retrievedName, retrievedAge)
}