package main

import (
	"fmt"
	"math"

	"github.com/jmoiron/sqlx"

	_ "modernc.org/sqlite"
)

type entries []struct {
	Class string `db:"class"`
	Count int    `db:"cnt"`
}

func studentCountPerClass(db *sqlx.DB) map[string]int {
	const selectStmt = `SELECT year || modifier as class, count(*) AS cnt
	FROM students as s
	INNER join groups as g 
	on s.id= g.student_id
	INNER join class as c 
	on c.id= g.class_id
	GROUP BY class`
	var entries entries
	db.Select(&entries, selectStmt)
	res := make(map[string]int)
	for _, v := range entries {
		res[v.Class] = v.Count
	}
	return res
}
func studentCountPerYear(db *sqlx.DB) map[int]int {
	const selectStmt = `SELECT year , count(*) AS cnt
	FROM students as s
	INNER join groups as g 
	on s.id= g.student_id
	INNER join class as c 
	on c.id= g.class_id
	GROUP BY year`
	var entries []struct {
		Year  int `db:"year"`
		Count int `db:"cnt"`
	}
	db.Select(&entries, selectStmt)
	r := make(map[int]int)
	for _, v := range entries {
		r[v.Year] = v.Count
	}
	return r
}

func lessonsPerYear(db *sqlx.DB, year int) (res []string) {
	const selectStmt = `SELECT lessons.name 
	FROM class
	JOIN timetable 
	ON class.id = timetable.class_id
	JOIN lessons 
	ON timetable.lesson_id = lessons.id
	WHERE year=? `
	var entries []struct {
		Name string `db:"name"`
	}
	db.Select(&entries, selectStmt, year)
	for _, v := range entries {
		res = append(res, v.Name)
	}
	return res
}

func examsPerClass(db *sqlx.DB, year int, mod string) (res []string) {
	const selectStmt = `SELECT DISTINCT name FROM class
	JOIN groups ON groups.class_id == class.id
	JOIN exams ON exams.student_id == groups.student_id
	JOIN lessons ON  lessons.id == exams.lesson_id
	WHERE year = ? AND modifier = ?`
	var entries []struct {
		Name string `db:"name"`
	}
	db.Select(&entries, selectStmt, year, mod)
	for _, v := range entries {
		res = append(res, v.Name)
	}
	return res
}
func averageGradeForStudents(db *sqlx.DB, firstName, lastName string) (grade float64) {
	const selectStmt = `SELECT AVG(grade) as average FROM students 
	JOIN exams ON exams.student_id = students.id
	WHERE name = ? AND surname = ?
	`
	var entry []struct {
		Grade float64 `db:"average"`
	}
	db.Select(&entry, selectStmt, firstName, lastName)
	grade = math.Round(entry[0].Grade*100) / 100
	return grade
}
func main() {
	db := sqlx.MustOpen("sqlite", "school.db")
	defer db.Close()
	fmt.Println(averageGradeForStudents(db, "Abbie", "Griffiths"))
}
