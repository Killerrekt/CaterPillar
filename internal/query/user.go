package query

const CreateUserTable = `CREATE TABLE IF NOT EXISTS users 
					(Name string not null,Password string not null,EmployeeID string not null unique)
					ENGINE = MergeTree()
					PRIMARY KEY (employee_id)`

const InsertUser = `INSERT INTO users (%s,%s,%d)`

const GetUser = `SELECT name, employee_id FROM users WHERE employee_id = %d`
