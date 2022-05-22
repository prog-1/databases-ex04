package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "modernc.org/sqlite"
)

func studentCountPerClass(db *sqlx.DB) map[string]int {
	var entries []struct {
		Class    string `db:"class"`
		Students int    `db:"students"`
	}

	const selectStmt = `SELECT class.year || class.modifier, 
	count(*) AS student_cnt
	FROM groups 
	JOIN class ON (class.id = groups.class_id)
	GROUP BY class_id`

	db.Select(&entries, selectStmt)
	
	res := make(map[string]int)

	for _, v := range entries {
		res[v.Class] = v.Students
	}
	return res
}

func studentCountPerYear(db *sqlx.DB) map[string]int {
	var entries []struct {
		Year string `db:"year"`
		Cnt  int    `db:"cnt"`
	}

	const selectStmt = `SELECT year, 
	count(*) AS student_cnt
	FROM groups
	JOIN class ON (class.id = groups.class_id)
	GROUP BY year`

	db.Select(&entries, selectStmt)

	res := make(map[string]int)
	for _, v := range entries {
		res[v.Year] = v.Cnt
	}
	return res
}

func averageGradeForStudents(db *sqlx.DB, firstName, lastName string) float64 {
	var entries []struct {
		Grade float64 `db:"grade"`
	}

	const selectStmt = `SELECT class.year, class.modifier, lessons.name, 
	avg(exam.grade) AS average_grade
	FROM exam
	JOIN groups
	ON exam.student_id=groups.student_id
	JOIN class
	ON groups.class_id=class.id
	JOIN lessons
	ON exam.lesson_id=lessons.id
	GROUP BY year, modifier`

	return entries[0].Grade
}

func main() {
	db := sqlx.MustOpen("sqlite", "school.db")
	defer db.Close()
	fmt.Println(studentCountPerClass(db))
	fmt.Println(studentCountPerYear(db))
}
