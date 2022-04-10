package database

import "time"

// Timetable represents a row storing information about a school timetable.
type Timetable struct {
	ID       int
	ClassID  int
	Day      time.Weekday
	LessonID int
}

// TimetableTable is a timetable database.
type TimetableTable struct {
	Data []Timetable
}

// Insert adds a new row to the timetable database.
func (db *TimetableTable) Insert(cID int, day time.Weekday, lID int) Timetable {
	tt := Timetable{
		ID:       len(db.Data) + 1,
		ClassID:  cID,
		Day:      day,
		LessonID: lID,
	}
	db.Data = append(db.Data, tt)
	return tt
}

// AddTimetable adds rows for several lessons to the timetable.
func (db *TimetableTable) AddTimetable(c Class, day time.Weekday, l ...Lesson) {
	for _, le := range l {
		db.Insert(c.ID, day, le.ID)
	}
}
