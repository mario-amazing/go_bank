package my_db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func DbConnect() *sql.DB {
	name := "go_develop"
	db, err := sql.Open("mysql", "root:dummy@/"+name)
	check(err)

	return db
}

func SetupDb() {
	db := DbConnect()

	defer db.Close()

	pwd, err := filepath.Abs("./")
	check(err)
	db_file_path := pwd + "/my_db/migration.sql"

	sql_execs, err := ioutil.ReadFile(db_file_path)
	check(err)
	sql_commands := strings.Split(string(sql_execs), ";")
	sql_commands = sql_commands[:len(sql_commands)-1]

	for _, c := range sql_commands {
		_, err = db.Exec(c + ";")
		check(err)
	}
	fmt.Println("DB setup complete")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
