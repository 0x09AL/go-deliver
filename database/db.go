package database

import "database/sql"

var db_filename = "test_db.db"

var db, _ = sql.Open("sqlite3", db_filename)

