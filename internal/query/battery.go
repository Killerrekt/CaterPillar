package query

const CreateBatteryTable = `CREATE TABLE battery(
					InceptionID UInt32 not null,
					BatteryMake String not null,
					BatteryReplacementDate String not null,
					BatteryVoltage UInt32 not null,
					BatteryWaterlevel String not null,
					ConditionOfBattery Bool not null,
					RustInBattery Bool not null,
					BatterySummary String not null,
					BatteryImages []String not null)
					ENGINE = MergeTree()
					PRIMARY KEY (InspectionID)`

const GetBattery = `SELECT * FROM battery WHERE InspectionID = %d`

const InsertBattery = `INSERT INTO tire(
					InceptionID
					BatteryMake
					BatteryReplacementDate
					BatteryVoltage 
					BatteryWaterlevel 
					ConditionOfBattery 
					RustInBattery 
					BatterySummary) VALUES (%d,'%s','%s',%d,'%s',%s,%s,'%s')`

const InsertBatteryImages = `INSERT INTO tire (BatteryImages) values(%v)`
