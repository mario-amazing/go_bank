package models

import (
	_ "errors"
)

type Client struct {
	Id                int64  `db:"id"    json:"id"`
	Name              string `db:"name"  json:"name"`
	Email             string `db:"email" json:"email"`
	Phone             string `db:"phone" json:"phone"`
	SmsNotification   bool   `db:"sms_notification" json:"sms_notification"`
	EmailNotification bool   `db:"email_notification" json:"email_notification"`
}
