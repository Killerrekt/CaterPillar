package query

const CreateExteriorTable = `CREATE TABLE IF NOT EXISTS exterior(
						InspectionID UInt32 not null,
						DamageToExterior Bool not null,
						ExteriorImages String,
						OilLeak Bool not null,
						ExteriorSummary String not null,
						Images Array(String))
						ENGINE = MergeTree()
						PRIMARY KEY (InspectionID)`

const InsertExterior = `INSERT INTO exterior(
					DamageToExterior,
					InspectionID,
					OilLeak,
					ExteriorSummary) VALUES(%s,%d,%s,'%s')`

const InsertExteriorImage = `ALTER TABLE exterior
					UPDATE ExteriorImages = '%s'
					WHERE InspectionID = %d`

const InsertExteriorImages = `ALTER TABLE exterior
					UPDATE Images = %s
					WHERE InceptionID = %d`

const GetExterior = `SELECT * FROM Exterior WHERE InspectionID = %d`
