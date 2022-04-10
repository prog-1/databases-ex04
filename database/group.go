package database

// Group represents a row storing student association with classes.
type Group struct {
	ClassID, StudentID int
}

// GroupDB is a group database.
type GroupDB struct {
	Data []Group
}

// Insert inserts a new mapping to the group database.
func (db *GroupDB) Insert(cID, sID int) Group {
	g := Group{
		ClassID:   cID,
		StudentID: sID,
	}
	db.Data = append(db.Data, g)
	return g
}

// AddStudentsToClass associates student(s) with a class.
func (db *GroupDB) AddStudentsToClass(c Class, s ...Student) {
	for _, st := range s {
		db.Insert(c.ID, st.ID)
	}
}
