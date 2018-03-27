package database

import "database/sql"

var db, _ = sql.Open("sqlite3", "test_db.db")