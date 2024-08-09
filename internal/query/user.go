package query

const CreateUserTable = `CREATE TABLE IF NOT EXISTS users 
					(Name String not null,Password String not null,EmployeeID String not null)
					ENGINE = MergeTree()
					PRIMARY KEY (EmployeeID)`

const InsertUser = `INSERT INTO users (%s,%s,%d)`

const GetUser = `SELECT Name, EmployeeID FROM users WHERE EmployeeID = %d`
