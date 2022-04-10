package database

// Student represents a row storing information about a student.
type Student struct {
	ID   int
	Name string
}

// StudentDB is a class database.
type StudentDB struct {
	Data []Student
}

// Insert adds a new row to the student database.
func (db *StudentDB) Insert(name string) Student {
	s := Student{
		ID:   len(db.Data) + 1,
		Name: name,
	}
	db.Data = append(db.Data, s)
	return s
}
