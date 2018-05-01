package models

import (
	_ "errors"
)

// var (
// 	queryClients      db.Query = "SELECT * FROM clients"
// 	queryClientByID   db.Query = "SELECT * FROM clients WHERE id=:id"
// 	queryClientInsert db.Query = "INSERT INTO clients (name, email, phone, sms_notification, email_notification) VALUES (:name, :email, :phone, :sms_notification, :email_notification)"
// 	queryClientUpdate db.Query = "UPDATE persons SET first_name=:first_name, last_name=:last_name, gender=:gender, modified_date=:modified_date, active=:active WHERE person_id=:person_id"
// )

type Client struct {
	Id                int64  `db:"id"    json:"id"`
	Name              string `db:"name"  json:"name"`
	Email             string `db:"email" json:"email"`
	Phone             string `db:"phone" json:"phone"`
	SmsNotification   bool   `db:"sms_notification" json:"sms_notification"`
	EmailNotification bool   `db:"email_notification" json:"email_notification"`
}

func ClientColumns() []string {
	f := []string{"name", "email", "phone", "sms_notification", "email_notification"}
	return f
}
