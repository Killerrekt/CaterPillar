package query

const CreateHeaderTable = `CREATE TABLE IF NOT EXISTS header(
							TruckSerialNumber String not null,
							TruckModel String not null,
							InspectionID UInt32 not null,
							InspectorName String not null,
							InspectorEmployeeID String not null,
							DateAndTime DateTime not null,
							Location String not null,
							ClientName String not null,
							ClientID String not null)
							ENGINE = MergeTree()
							PRIMARY KEY (InspectionID, DateAndTime)`

const GetHeader = `SELECT * FROM header WHERE InspectorEmployeeID = '%s'`

const InsertHeader = `INSERT INTO header VALUES('%s','%s','%d','%s','%s',now(),'%s','%s','%s')`
