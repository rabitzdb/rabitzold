package ingestion

//Dataspace structure
type Ds struct {
	Uid     uint64
	UserUid uint64
	Name    string
	Attr    string
}

//Dataspace attribute structure
type Dsa struct {
	Useruid uint64
	Name    string
	Field   string
}

/*
List dataspaces for client uid
*/
func Listds(userUid uint64) (dataspaces *[]Ds) {
	query := "SELECT uid, user_uid, name, attr name FROM dataspace WHERE user_uid = ?"
	stmt, err := db.Prepare(query)
	if err != nil {
		logger.Error("Error preparing statement: " + query)
		return nil
	}
	res, err := stmt.Query(userUid)
	if err != nil {
		logger.Error("Error executing query: " + query)
		return nil
	}
	var ds Ds
	for res.Next() {
		res.Scan(&ds.Uid, &ds.UserUid, &ds.Name, &ds.Attr)
		logger.Info("Dataspace: " + ds.Name)
	}
	return nil
}

/*
Get dataspace info
*/
func Getds(uid uint64) (dataspaces *Ds) {
	query := "SELECT Uid, UserUid, Name, Attr FROM dataspace WHERE Uid = ?"

	stmt, err := db.Prepare(query)
	if err != nil {
		logger.Error("Error preparing statement: " + query)
		return nil
	}
	res, err := stmt.Query(uid)
	if err != nil {
		logger.Error("Error executing query: " + query)
		return nil
	}
	var ds Ds
	res.Next()
	res.Scan(&ds.Uid, &ds.UserUid, &ds.Name, &ds.Attr)
	logger.Info("Get Dataspace: " + ds.Name)

	return nil
}

/*
Create a new dataspace
*/
func Insertds(dataspace Ds) {
	query := "INSERT INTO dataspace(UserUid,Name,Attr) VALUES(?,?,?)"
	insForm, err := db.Prepare(query)
	if err != nil {
		logger.Error("Error executing query: " + query)
	}
	insForm.Exec(dataspace.UserUid, dataspace.Name, dataspace.Attr)
}
