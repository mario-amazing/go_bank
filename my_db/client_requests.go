package my_db

import (
	"cources/homework/models"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
)

func FindClients() []*models.Client {
	db := DbConnect()
	defer db.Close()

	response, err := db.Query("Select * FROM clients")
	check(err)

	var clients []*models.Client

	clients = scanClients(response)

	// fmt.Println(response *sql.Rows, client models.Client)
	return clients
}

func FindClient(id string) []*models.Client {
	db := DbConnect()
	defer db.Close()

	response, err := db.Query("Select * FROM clients WHERE id=?", id)
	check(err)

	var clients []*models.Client

	clients = scanClients(response)

	return clients
}

func scanClients(response *sql.Rows) []*models.Client {
	var clients []*models.Client

	for response.Next() {
		client := models.Client{}
		err := response.Scan(&client.Id, &client.Name, &client.Email, &client.Phone, &client.SmsNotification, &client.EmailNotification)
		check(err)

		fmt.Println(response)
		clients = append(clients, &client)
	}

	return clients
}

func CreateClient(client models.Client) []*models.Client {
	db := DbConnect()
	defer db.Close()
	c := strings.Join(models.ClientColumns(), ", ")
	v := strings.Join(GetClientValues(client), ", ")

	fmt.Println(c)
	fmt.Println(v)

	response, err := db.Query("INSERT INTO clients(%s) VALUES(%s);", c, v)
	check(err)

	var clients []*models.Client

	clients = scanClients(response)

	return clients
}

func GetClientValues(c models.Client) []string {
	v := reflect.ValueOf(c)
	values := make([]string, v.NumField())

	for i := 1; i <= v.NumField(); i++ {
		values[i] = fmt.Sprintf("%#v", v.Field(i+1).Interface())
	}

	return values
}
