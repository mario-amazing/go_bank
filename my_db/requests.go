package my_db

// import (
// )

func Find(query, db_name string) sql.Result {
	db := DbConnect()
	defer db.Close()

	_, err := db.Exec("USE " + db_name)
	check(err)
	response, err := db.Exec(query)
	return response
}
