# SQL & Go

## Using sqlite from Go

There is an existing standard library [database/sql] that provides an interface
around SQL (or SQL-like) databases.

However, this library is overly verbose to use, and it does not allow to select
into structures and slices, which is quite convenient to use. Because of that
we will use [jmoiron/sqlx] instead.

> NOTE: [jmoiron/sqlx] is an optional library, and you could use standard
  [database/sql] to complete your home assignments.

Using generic SQL libraries require [drivers] that are implement for specific
SQL service or library. A driver is usually imported for its side-effects, which
is registering the driver name. To ensure Go does not remove unused import, you
have to set its name to '_' e.g.

```golang
import _ "modernc.org/sqlite"
```

There are multiple sqlite [drivers] available for Go, but we will be using
[modernc.org/sqlite](https://modernc.org/sqlite) since it features a pure Go
implementation that does not require C compiler. And this makes it easier to
use for our purpose.

You can find an example of a small program that creates the `test` table,
inserts some data and fetches it at [yarcat/playground/modernc-sqlite].

[drivers]: https://github.com/golang/go/wiki/SQLDrivers
[database/sql]: https://pkg.go.dev/database/sql
[yarcat/playground/modernc-sqlite]: https://github.com/yarcat/playground/tree/master/modernc-sqlite
[jmoiron/sqlx]: https://jmoiron.github.io/sqlx/

## Exercises

To implement the exercises you will need to use the same database and tables as
created in:

- https://github.com/prog-1/databases-ex02 (for students, classes, etc)
- https://github.com/prog-1/databases-ex03 (for exams)

The exercises are exactly the same as in [databases-ex01], but must be
implemented using SQL.

[databases-ex01]: https://github.com/prog-1/databases-ex01

> NOTE: The functions will accept `*sql.DB` if you use `sql` instead of `sqlx`.

1. Implement a function `func studentCountPerClass(db *sqlx.DB) map[string]int`
   that returns the number of students in every class. The returned value is a map
   "class -> count", e.g. `map[string]int{"10b": 2}`.
2. Implement a function func `studentCountPerYear(db *sqlx.DB) map[int]int` that
   returns the number of students for every year. The returned value is a map
   "year -> count", e.g. `map[int]int{10: 2}`.
3. Implement a function `func lessonsPerYear(db *sqlx.DB, year int) []string` that
   returns a slice of unique subjects that students learn for a given year, e.g.
   `[]string{"Math", "Sport", "Programming"}`.
4. Write a function `func examsPerClass(db *sqlx.DB, year int, mod string) []string`
   that returns a list of exams for a given class and modifier.
5. Write a function `func averageGradeForStudents(db *sqlx.DB, name string) float64`
   that returns an average grade for a given student.
