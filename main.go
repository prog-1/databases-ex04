package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "modernc.org/sqlite"
)

func studentCountPerClass(db *sqlx.DB) map[string]int {
	const selectStmt = `SELECT year|| modifier as s, student_cnt FROM (SELECT class_id, count(*) AS student_cnt  FROM groups GROUP BY class_id)
	INNER JOIN class
	on class_id == class.id`
	var entries []struct {
		Class string `db:"s"`
		Count int    `db:"student_cnt"`
	}
	db.Select(&entries, selectStmt)
	r := make(map[string]int)
	for _, v := range entries {
		r[v.Class] = v.Count
	}
	return r
}
func studentCountPerYear(db *sqlx.DB) map[int]int {
	const selectStmt = `SELECT year, count(*) AS student_cnt FROM(
		SELECT * FROM groups
		INNER JOIN class
		on class_id == class.id
		) GROUP BY year`
	var entries []struct {
		Year  int `db:"year"`
		Count int `db:"student_cnt"`
	}
	db.Select(&entries, selectStmt)
	r := make(map[int]int)
	for _, v := range entries {
		r[v.Year] = v.Count
	}
	return r
}

func lessonsPerYear(db *sqlx.DB, year int) (r []string) {
	const selectStmt = `SELECT name FROM (
		SELECT * from class
		JOIN timetable on timetable.class_id == class.id
		JOIN lessons on timetable.lesson_id = lessons.id
		) WHERE year = ?
		GROUP by name `
	var entries []struct {
		Name string `db:"name"`
	}
	db.Select(&entries, selectStmt, year)
	for _, v := range entries {
		r = append(r, v.Name)
	}
	return r
}

func examsPerClass(db *sqlx.DB, year int, mod string) (r []string) {
	const selectStmt = `SELECT DISTINCT name FROM class
	JOIN groups ON groups.class_id == class.id
	JOIN exams ON exams.student_id == groups.student_id
	JOIN lessons ON exams.lesson_id == lessons.id
	WHERE year = ? AND modifier = ?`
	var entries []struct {
		Name string `db:"name"`
	}
	db.Select(&entries, selectStmt, year, mod)
	for _, v := range entries {
		r = append(r, v.Name)
	}
	return
}
func averageGradeForStudents(db *sqlx.DB, firstName, lastName string) float64 {
	const selectStmt = `SELECT AVG(grade) as AVG FROM students 
	JOIN exams ON exams.student_id = students.id
	WHERE name = ? AND surname = ?
	`
	var entries []struct {
		Grade float64 `db:"AVG"`
	}
	db.Select(&entries, selectStmt, firstName, lastName)
	return entries[0].Grade
}
func main() {
	db := sqlx.MustOpen("sqlite", "school.db")
	defer db.Close()
	fmt.Println(averageGradeForStudents(db, "Aaliyah", "Nhira"))
}
