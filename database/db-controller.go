package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

type Database struct {
	Conn *pgx.Conn
}

func NewDb() Database {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	database := Database{
		Conn: conn,
	}

	return database
}

func (db *Database) CheckPassword(username, passwordHash string) bool {
	var result string
	db.Conn.QueryRow(context.Background(), "SELECT password_hash FROM prizm.admin WHERE username = $1", username).Scan(&result)

	return result == passwordHash
}

func (db *Database) GetClients(serverIp string) ([]Client, error) {
	var rows pgx.Rows

	if serverIp == "" {
		rows, _ = db.Conn.Query(context.Background(), "SELECT * FROM prizm.clients")
	} else {
		rows, _ = db.Conn.Query(context.Background(), "SELECT * FROM prizm.clients WHERE server_ip = $1", serverIp)
	}

	defer rows.Close()

	var clientsList []Client
	var client Client

	for rows.Next() {
		err := rows.Scan(&client.Id, &client.Name, &client.ServerIp, &client.PublicKey)
		if err != nil {
			fmt.Println(err)
			return []Client{}, err
		}

		clientsList = append(clientsList, client)
	}

	return clientsList, nil
}

func (db *Database) GetServers(countryCode string) ([]Server, error) {
	var rows pgx.Rows

	if countryCode == "" {
		rows, _ = db.Conn.Query(context.Background(), "SELECT * FROM prizm.servers WHERE")
	} else {
		rows, _ = db.Conn.Query(context.Background(), "SELECT * FROM prizm.servers WHERE country = $1", countryCode)
	}

	defer rows.Close()

	var serversList []Server
	var server Server

	for rows.Next() {
		err := rows.Scan(&server.ServerIp, &server.ClientsNum, &server.Country)
		if err != nil {
			fmt.Println(err)
			return []Server{}, err
		}

		serversList = append(serversList, server)
	}

	return serversList, nil
}
