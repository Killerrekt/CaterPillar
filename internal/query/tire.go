package query

const CreateTireTable = `CREATE TABLE IF NOT EXISTS tire(
						InspectionID UInt32 not null,
						TirePressureLeftFront UInt32 not null,
						TirePressureRightFront UInt32 not null,
						TirePressureLeftRear UInt32 not null,
						TirePressureRightRear UInt32 not null,
						LeftFrontCondition String not null,
						LeftRearCondition String not null,
						RightFrontCondition String not null,
						RightRearCondition String not null,
						TireSummary String not null,
						Images Array(String))
						ENGINE = MergeTree()
						PRIMARY KEY (InspectionID)`

const InsertTire = `INSERT INTO tire(InspectionID,
						TirePressureLeftFront,
						TirePressureRightFront,
						TirePressureLeftRear,
						TirePressureRightRear,
						LeftFrontCondition,
						LeftRearCondition,
						RightFrontCondition,
						RightRearCondition,
						TireSummary) 
						VALUES(%d,%d,%d,%d,%d,'%s','%s','%s','%s','%s')`

const InsertTireImage = `ALTER TABLE tire UPDATE Images = %s WHERE InspectionID = %d`

const GetTireImage = `SELECT Images FROM tire WHERE InspectionID = %d`

const GetTire = `SELECT * FROM tire WHERE InspectionID = %d`
