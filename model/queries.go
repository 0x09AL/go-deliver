package model

var CreateTableSql = `CREATE TABLE payloads (
							id	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
							name	TEXT NOT NULL UNIQUE,
							content_type	TEXT,
							host_blacklist	TEXT,
							host_whitelist	TEXT,
							data_file	TEXT,
							data_b64	TEXT,
							type_id	INTEGER NOT NULL,
							guid	TEXT NOT NULL UNIQUE
						);
						`

var CreateHostSql = `CREATE TABLE hosts (
							id	INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
							type	TEXT NOT NULL,
							data	TEXT NOT NULL
						);`

var CreateTypesSql = `CREATE TABLE types (
							id	INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
							name	TEXT NOT NULL UNIQUE,
							type_template	TEXT
							content_type TEXT
						);`
var DeletePayloadQuery = "DELETE FROM payloads where name == ?;"

var InsertPayloadQuery = "INSERT INTO payloads VALUES (NULL,?,?,?,?,?,?,?,?);"

var GetPayloadsQuery = "SELECT id, name, guid, content_type,COALESCE(host_blacklist, '') as host_blacklist, COALESCE(host_whitelist, '') as host_whitelist FROM payloads;"

var GetPayloadQuery = `SELECT id,
						name,
						content_type,
						COALESCE(host_blacklist, '') as host_blacklist, 
						COALESCE(host_whitelist, '') as host_whitelist,
						COALESCE(data_file, '') as data_file, 
						COALESCE(data_b64, '') as data_b64 ,
						type_id 
						from payloads 
						WHERE guid=?`

var GetPayloadTypesQuery = "SELECT type_name , content_type FROM payload_types;"

var GetPayloadTypeId = "SELECT type_id , COALESCE(content_type, '') as content_type  FROM payload_types WHERE type_name == ?;";

var CreateHostQuery = "INSERT INTO hosts VALUES(NULL,?,?,?,?)"

var GetHostsQuery = "SELECT * from hosts;"