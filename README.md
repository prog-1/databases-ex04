# Relational Databases

## Definition

A [relational database](https://en.wikipedia.org/wiki/Relational_database) is a digital database, where data is organizes into one or more [tables](https://en.wikipedia.org/wiki/Table_(database)) (or "relations") of [columns](https://en.wikipedia.org/wiki/Column_(database)) and [rows](https://en.wikipedia.org/wiki/Row_(database)), with a unique [key](https://en.wikipedia.org/wiki/Relational_database#Keys) identifying each row.

## Database Example

1. Table `Students`. Columns: `ID`, `Name`.

   Row examples:

   ```
   (1, "Jaroslavs")
   (2, "Pavels")
   ```

2. Table `Class`. Columns: `ID`, `Year`, `Modifier`.

   Row examples:

   ```
   (2, 10, "a")
   (3, 10, "b")
   ```

3. Table: `Groups`. Columns: `Class ID`, `Student ID`.

   Row examples:

   ```
   (3, 1)
   (3, 2)
   ```

4. Table: `Lessons`. Columns: `ID`, `Name`.

   Row examples:

   ```
   (5, "Math")
   (10, "Sport")
   (7, "Programming")
   ```

5. Table `Timetable`. Columns: `ID`, `Class ID`, `Day`, `Lesson ID`.

   Row example:

   ```
   (1, 3, "Monday", 5)
   (2, 3, "Monday", 7)
   ```