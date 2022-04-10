package database

// Lesson represents a row storing information about a school lesson.
type Lesson struct {
	ID   int
	Name string
}

// LessonDB is a class database.
type LessonDB struct {
	Data []Lesson
}

// Insert adds a new row to the lesson database.
func (db *LessonDB) Insert(name string) Lesson {
	l := Lesson{
		ID:   len(db.Data) + 1,
		Name: name,
	}
	db.Data = append(db.Data, l)
	return l
}
