package query

const CreateEngineTable = `CREATE TABLE IF NOT EXISTS engine(
	InspectionID UInt32 not null,
	DamageEngine Bool not null,
	EngineOilCondition String not null,
	EngineOilColor String not null,
	BrakeFluidCondition String not null,
	BrakeFluidColor String not null,
	OilLeak Bool not null,
	Summary String not null,
	DmgEngineImg String not null,
	Images Array(String))
	ENGINE = MergeTree()
	PRIMARY KEY (InspectionID)`

const InsertEngine = `INSERT INTO engine(
					InspectionID,
					DamageEngine,
					EngineOilCondition,
					EngineOilColor,
					BrakeFluidCondition,
					BrakeFluidColor,
					OilLeak,
					Summary) VALUES (%d , %s,'%s','%s','%s','%s',%s,'%s')`

const InsertEngineImages = `ALTER TABLE engine
		UPDATE Images = %s
		WHERE InspectionID = %d`

const InsertEngineImage = `ALTER TABLE engine
		UPDATE DmgEngineImd = '%s'
		WHERE InspectionID = %d`

const GetEngine = `SELECT * FROM engine WHERE InspectionID = %d`
