package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeC(t *testing.T) {
	expected_value := "C"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 70, Assignment)
	gradeCalculator.AddGrade("exam 1", 75, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 71, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeD(t *testing.T) {
	expected_value := "D"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 60, Assignment)
	gradeCalculator.AddGrade("exam 1", 65, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 61, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 50, Assignment)
	gradeCalculator.AddGrade("exam 1", 55, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 51, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestNoEssay(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestString(t *testing.T) {
	expected_value := "assignment"
	actual_value := Assignment.String()

	if expected_value != actual_value {
		t.Errorf("Expected Assignment.String() to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestPass(t *testing.T) {
	expected_value := "Pass"

	gradeCalculator := NewGradeCalculator()

	// 70/70/70 → weighted = 70 → Pass
	gradeCalculator.AddGrade("assignment avg", 70, Assignment)
	gradeCalculator.AddGrade("exam avg", 70, Exam)
	gradeCalculator.AddGrade("essay avg", 70, Essay)

	actual_value := gradeCalculator.CalculatePassFail()

	if expected_value != actual_value {
		t.Errorf("Expected GetPassFail to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestFail(t *testing.T) {
	expected_value := "Fail"

	gradeCalculator := NewGradeCalculator()

	// 69/69/69 → weighted = 69 → Fail
	gradeCalculator.AddGrade("assignment avg", 69, Assignment)
	gradeCalculator.AddGrade("exam avg", 69, Exam)
	gradeCalculator.AddGrade("essay avg", 69, Essay)

	actual_value := gradeCalculator.CalculatePassFail()

	if expected_value != actual_value {
		t.Errorf("Expected GetPassFail to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}
