// Signifies that this file is apart of a bundled package
package esepunittests

// Definition of a new struct type each portion of the grade is represented as a slice of an array called Grade
type GradeCalculator struct {
	assignments []Grade
	exams       []Grade
	essays      []Grade
}

// GradeType is defined to allow for each portion of the grade to be represented with ints
type GradeType int

// Each portion of the grade is assigned a type of GradeType and because it is based on int and iota is used to increment over the values 0, 1, and 2 are used for them
const (
	Assignment GradeType = iota // Assignment = 0
	Exam                        // Exam = 1
	Essay                       // Essay = 2
)

// To keep track of the actual sting name for each one a map is created with the name assigned to each portion
var gradeTypeName = map[GradeType]string{
	Assignment: "assignment", // Key: Assignment (Type GradeType and Value 0) / Pair: "assignment"
	Exam:       "exam",       // Key: Exam (Type GradeType and Value 1) / Pair: "exam"
	Essay:      "essay",      // Key: Essay (Type GradeType and Value 2) / Pair: "essay"
}

// Method to type GradeType that returns the string name of the variable that its called on (gt). Similar to a private function of a class and gt is similar to this or self
func (gt GradeType) String() string {
	return gradeTypeName[gt]
}

// A struct to hold one grade entry with a name, numeric score, and type.
type Grade struct {
	Name  string    // Name of the individual assignment
	Grade int       // Score on that assignment
	Type  GradeType // Category it belongs to (0 = Assignment, 1 = Exam, 2 = Essay)
}

// Function to create a new calulator and returns a pointer to it
func NewGradeCalculator() *GradeCalculator {
	return &GradeCalculator{

		// An empty slice is created for each assignment type
		assignments: make([]Grade, 0),
		exams:       make([]Grade, 0),
		essays:      make([]Grade, 0),
	}
}

// Method to type GradeCalculator that correlates the final grade number to a letter and returns it
func (gc *GradeCalculator) GetFinalGrade() string {

	// Helper function is used to get the overall weighted grade
	numericalGrade := gc.calculateNumericalGrade()

	if numericalGrade >= 90 {
		return "A"
	} else if numericalGrade >= 80 {
		return "B"
	} else if numericalGrade >= 70 {
		return "C"
	} else if numericalGrade >= 60 {
		return "D"
	}

	return "F"
}

// Method to type GradeCalculator that adds a grade in
func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {

	// Switch is used to select the correct branch to use based on gradeType
	switch gradeType {
	case Assignment:
		gc.assignments = append(gc.assignments, Grade{
			Name:  name,
			Grade: grade,
			Type:  Assignment,
		})
	case Exam:
		gc.exams = append(gc.exams, Grade{
			Name:  name,
			Grade: grade,
			Type:  Exam,
		})
	case Essay:
		gc.essays = append(gc.essays, Grade{
			Name:  name,
			Grade: grade,
			Type:  Essay,
		})
	}
}

// Method to type GradeCalculator that calculates the total grade (0–100 as an int).
func (gc *GradeCalculator) calculateNumericalGrade() int {

	// Compute the average for assignments, exams, and essays by passing each slice into computeAverage.
	assignment_average := computeAverage(gc.assignments)
	exam_average := computeAverage(gc.exams)
	essay_average := computeAverage(gc.essays)

	// Grades are converted to floats and weights are applied to each one
	weighted_grade := float64(assignment_average)*.5 + float64(exam_average)*.35 + float64(essay_average)*.15

	// Final grade value is changed back to an int and returned
	return int(weighted_grade)
}

func computeAverage(grades []Grade) int {

	// Get the amount of assignments in the slice
	assignCount := len(grades)

	// Check if there are no assignments
	if assignCount == 0 {

		// If so, return 0
		return 0
	}

	// Initialize the summing variable
	sum := 0

	// Loop through the grades and sum the value of the grade
	for _, grade := range grades {

		// Add the Grade component of the current assignment to the summing variable
		sum += grade.Grade
	}

	// Return the average by dividing the sum by the number of assignments
	return sum / assignCount
}

func (gc *GradeCalculator) CalculatePassFail() string {

	gradeValue := gc.calculateNumericalGrade()

	if gradeValue >= 70 {
		return "Pass"
	}
	return "Fail"
}
