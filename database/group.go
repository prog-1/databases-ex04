package database

// Group represents a row storing student association with classes.
type Group struct {
	ClassID, StudentID int
}

// GroupTable is a group database.
type GroupTable struct {
	Data []Group
}

// Insert inserts a new mapping to the group database.
func (db *GroupTable) Insert(cID, sID int) Group {
	g := Group{
		ClassID:   cID,
		StudentID: sID,
	}
	db.Data = append(db.Data, g)
	return g
}

// AddStudentsToClass associates student(s) with a class.
func (db *GroupTable) AddStudentsToClass(c Class, s ...Student) {
	for _, st := range s {
		db.Insert(c.ID, st.ID)
	}
}
