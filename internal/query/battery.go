package query

const CreateBatteryTable = `CREATE TABLE IF NOT EXISTS battery(
					InspectionID UInt32 not null,
					BatteryMake String not null,
					BatteryReplacementDate String not null,
					BatteryVoltage UInt32 not null,
					BatteryWaterlevel String not null,
					ConditionOfBattery Bool not null,
					RustInBattery Bool not null,
					BatterySummary String not null,
					BatteryImages Array(String))
					ENGINE = MergeTree()
					PRIMARY KEY (InspectionID)`

const GetBattery = `SELECT * FROM battery WHERE InspectionID = %d`

const InsertBattery = `INSERT INTO battery(
					InspectionID,
					BatteryMake,
					BatteryReplacementDate,
					BatteryVoltage ,
					BatteryWaterlevel, 
					ConditionOfBattery ,
					RustInBattery ,
					BatterySummary) VALUES (%d,'%s','%s',%d,'%s',%s,%s,'%s')`

const InsertBatteryImages = `ALTER TABLE battery UPDATE Images = %s WHERE InspectionID = %d`

const GetBatteryImage = `SELECT Images FROM battery WHERE InspectionID = %d`
