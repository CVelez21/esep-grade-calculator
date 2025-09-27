# ESEP Grade Calculator

This library implements a simple grade calculator for assignments, exams, and essays. It computes averages for each category, applies weights, and converts the final numeric grade into a letter grade.

- **Weights:** Assignments 50%, Exams 35%, Essays 15%  
- **Letter Scale:** A ≥ 90, B ≥ 80, C ≥ 70, D ≥ 60, F < 60  
- **Passing Rule:** C or greater

---

## Requirements

- **Go 1.24+**  
Install Go from [https://go.dev/dl/](https://go.dev/dl/) and verify with:

```bash
go version
```

Clone the repo and enter the project folder:

```bash
cd esep-grade-calculator/grade_calculator
```

---

## Usage Example

```go
gc := NewGradeCalculator()

gc.AddGrade("hw1", 90, Assignment)
gc.AddGrade("midterm", 85, Exam)
gc.AddGrade("ethics essay", 95, Essay)

fmt.Println(gc.GetFinalGrade()) // Output: "B"
```

---

## Running Tests

This project uses Go’s **built-in testing framework**.  

Run all tests:
```bash
go test ./...
```

Run with verbose output (see each test name and status):
```bash
go test -v ./...
```

Run with coverage and reports:
```bash
go test -coverprofile=coverage.txt ./...
go tool cover -func=coverage.txt
go tool cover -html=coverage.txt
```

---

## Go Language Notes for Beginners

If you are new to Go, here are the essentials needed to follow this codebase:

### Packages
- Every `.go` file starts with a `package` line.  
- Files in the same folder must use the same package name.  
- This project uses `package esepunittests`.

### Types
- **Structs** define composite types (like classes without methods).  
  ```go
  type Grade struct {
      Name  string
      Grade int
      Type  GradeType
  }
  ```

- **Slices** are Go’s dynamic arrays (like `std::vector` in C++).  
  They grow with `append` and have a length (`len(slice)`).

### Methods
- Functions attached to a type use a **receiver**:  
  ```go
  func (gc *GradeCalculator) GetFinalGrade() string { ... }
  ```
  - `(gc *GradeCalculator)` means this function is a method on `GradeCalculator`.  
  - Inside, `gc` acts like `this` in C++/Java of 'self' in Python.

### Enums
- Go doesn’t have built-in enums, but we can fake them with `const` + `iota`:  
  ```go
  type GradeType int
  const (
      Assignment GradeType = iota // 0
      Exam                        // 1
      Essay                       // 2
  )
  ```

### Loops
- The idiomatic way to loop over a slice is:
  ```go
  for _, g := range grades {
      sum += g.Grade
  }
  ```
  - `_` ignores the index.  
  - `g` is the element.

### Error Prevention
- Empty slices return length 0. Dividing by `len(slice)` will panic.  
  Always guard against empty slices if you compute averages:
  ```go
  if len(grades) == 0 { return 0 }
  ```

---
