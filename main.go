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
	result := make(map[string]int)
	db.Select(&entries, `SELECT year || modifier AS class, 
	COUNT(*) AS students
	FROM class
	JOIN groups ON class.id = groups.class_id
	GROUP BY class_id`)
	for _, v := range entries {
		result[v.Class] = v.Students
	}
	return result
}

func studentCountPerYear(db *sqlx.DB) map[int]int {
	var entries []struct {
		Year     int `db:"year"`
		Students int `db:"students"`
	}
	result := make(map[int]int)
	db.Select(&entries, `SELECT year, 
	COUNT(*) AS students
	FROM class
	JOIN groups ON class.id = groups.class_id
	GROUP BY year`)
	for _, v := range entries {
		result[v.Year] = v.Students
	}
	return result
}

func lessonsPerYear(db *sqlx.DB, year int) []string {
	var entries []struct {
		Name string `db:"name"`
	}
	var result []string
	db.Select(&entries, `SELECT lessons.name
	FROM class
	JOIN timetable ON class.id = timetable.class_id
	JOIN lessons ON timetable.lesson_id = lessons.id
	WHERE class.year = ?
	GROUP BY lessons.name`, year)
	for _, v := range entries {
		result = append(result, v.Name)
	}
	return result
}

func examsPerClass(db *sqlx.DB, year int, mod string) []string {
	var entries []struct {
		Name string `db:"name"`
	}
	var result []string
	db.Select(&entries, `SELECT lessons.name 
	FROM class
	JOIN groups ON class.id = groups.class_id
	JOIN exams ON groups.student_id = exams.student_id
	JOIN lessons ON exams.lesson_id = lessons.id
	WHERE class.year = ? AND class.modifier = ?
	GROUP BY lessons.name`, year, mod)
	for _, v := range entries {
		result = append(result, v.Name)
	}
	return result
}

func averageGradeForStudents(db *sqlx.DB, firstName, lastName string) float64 {
	var entries []struct {
		AverageGrade float64 `db:"average_grade"`
	}
	db.Select(&entries, `SELECT ROUND(AVG(exams.grade), 1) AS average_grade
	FROM students
	JOIN exams ON students.id = exams.student_id
	WHERE students.name = ? AND students.surname = ?`, firstName, lastName)
	return entries[0].AverageGrade
}

func main() {
	db := sqlx.MustOpen("sqlite", "school.db")
	defer db.Close()

	studentCountPerClass := studentCountPerClass(db)
	fmt.Println("Student count per class:")
	for i, v := range studentCountPerClass {
		fmt.Printf(" - %s: %v\n", i, v)
	}

	studentCountPerYear := studentCountPerYear(db)
	fmt.Println("Student count per year:")
	for i, v := range studentCountPerYear {
		fmt.Printf(" - %v: %v\n", i, v)
	}

	fmt.Printf("Lesson per year %v:\n", year)
	lessonsPerYear := lessonsPerYear(db, year)
	for _, v := range lessonsPerYear {
		fmt.Printf(" - %s\n", v)
	}

	fmt.Printf("Exams per class %s:\n", fmt.Sprintf("%v%s", year, mod))
	examsPerClass := examsPerClass(db, year, mod)
	for _, v := range examsPerClass {
		fmt.Printf(" - %s\n", v)
	}

	averageGradeForStudent := averageGradeForStudents(db, firstName, lastName)
	fmt.Printf("Average grade for student %s %s: %v", firstName, lastName, averageGradeForStudent)
}

const year, mod, firstName, lastName = 10, "a", "Aaliyah", "Nhira"
