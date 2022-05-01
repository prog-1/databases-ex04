package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "modernc.org/sqlite"
)

func studentCountPerClass(db *sqlx.DB) map[string]int {
	const selectStmt = `SELECT class.year || class.modifier as class,
	count(*) AS students
	FROM groups 
	JOIN class ON class.id = groups.class_id
	GROUP BY class_id`

	var entries []struct {
		Class    string `db:"class"`
		Students int    `db:"students"`
	}

	db.Select(&entries, selectStmt)
	m := make(map[string]int)
	for _, v := range entries {
		m[v.Class] = v.Students
	}
	return m
}

func studentCountPerYear(db *sqlx.DB) map[string]int {
	const selectStmt = `SELECT class.year as year,
	count(*) AS cnt
	FROM groups 
	JOIN class
	ON class.id = groups.class_id
	GROUP BY year`

	var entries []struct {
		Year string `db:"year"`
		Cnt  int    `db:"cnt"`
	}

	db.Select(&entries, selectStmt)
	m := make(map[string]int)
	for _, v := range entries {
		m[v.Year] = v.Cnt
	}
	return m
}

func lessonsPerYear(db *sqlx.DB, year int) []string {
	const selectStmt = `SELECT lessons.name as lesson
	FROM class
	JOIN timetable 
	ON class.id = timetable.class_id
	JOIN lessons
	ON timetable.lesson_id = lessons.id
	WHERE class.year = ?
	GROUP BY lessons.name`

	var entries []struct {
		Lessons string `db:"lesson"`
	}

	db.Select(&entries, selectStmt, year)
	var s []string
	for _, v := range entries {
		s = append(s, v.Lessons)
	}
	return s
}

func examsPerClass(db *sqlx.DB, year int, mod string) []string {
	const selectStmt = `SELECT lessons.name as exams
	FROM class
	JOIN groups
	ON class.id = groups.class_id
	JOIN exams
	ON groups.student_id = exams.student_id
	JOIN lessons 
	ON exams.lesson_id = lessons.id
	WHERE class.year = ? AND class.modifier = ?
	GROUP BY lessons.name`

	var entries []struct {
		Exams string `db:"exams"`
	}

	db.Select(&entries, selectStmt, year, mod)
	var s []string
	for _, v := range entries {
		s = append(s, v.Exams)
	}
	return s
}

func averageGradeForStudents(db *sqlx.DB, firstName, lastName string) float64 {
	const selectStmt = `SELECT ROUND(AVG(exams.grade), 1) AS grade
	FROM exams
	JOIN students 
	ON exams.student_id = students.id
	WHERE students.name = ? AND students.surname = ?`

	var entries []struct {
		Grade float64 `db:"grade"`
	}

	db.Select(&entries, selectStmt, firstName, lastName)

	return entries[0].Grade
}

func main() {
	db := sqlx.MustOpen("sqlite", "school.db")
	defer db.Close()
	fmt.Println(studentCountPerClass(db))
	fmt.Println(studentCountPerYear(db))
	fmt.Println(lessonsPerYear(db, 9))
	fmt.Println(examsPerClass(db, 12, "a"))
	fmt.Println(averageGradeForStudents(db, "Abigail", "Tshuma"))

}
