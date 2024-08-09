package query

const CreateUserTable = `CREATE TABLE IF NOT EXISTS users 
					(Name String not null,Password String not null,EmployeeID String not null)
					ENGINE = MergeTree()
					PRIMARY KEY (EmployeeID)`

const InsertUser = `INSERT INTO users VALUES('%s','%s','%s')`

const GetUser = `SELECT Name, EmployeeID, Password FROM users WHERE EmployeeID = '%s' LIMIT 1`
