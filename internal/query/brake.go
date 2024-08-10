package query

const CreateBrakeTable = `CREATE TABLE IF NOT EXISTS brake(
						InspectionID UInt32 not null,	
						BrakeFluidLevel String not null,
						BrakeConditionFront String not null,
						BrakeConditionRear  String not null,
						EmergencyBrake String not null,
						BrakeSummary String not null,
						Images Array(String))
						ENGINE = MergeTree()
						PRIMARY KEY (InspectionID)`

const InsertBrake = `INSERT INTO brake(
					InspectionID,
					BrakeFluidLevel,
					BrakeConditionFront,
					BrakeConditionRear,
					EmergencyBrake,
					BrakeSummary) VALUES (%d,'%s','%s','%s','%s','%s')`

const InsertBrakeImages = `ALTER TABLE brake
							UPDATE Images = %s
							WHERE InspectionID = %d`

const GetBrake = `SELECT * FROM brake WHERE InspectionID = %d`
