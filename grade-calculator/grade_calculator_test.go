package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator(LetterGrade)

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

	gradeCalculator := NewGradeCalculator(LetterGrade)

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

	gradeCalculator := NewGradeCalculator(LetterGrade)

	gradeCalculator.AddGrade("open source assignment", 70, Assignment)
	gradeCalculator.AddGrade("exam 1", 71, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 75, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeD(t *testing.T) {
	expected_value := "D"

	gradeCalculator := NewGradeCalculator(LetterGrade)

	gradeCalculator.AddGrade("open source assignment", 60, Assignment)
	gradeCalculator.AddGrade("exam 1", 61, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 65, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator(LetterGrade)

	gradeCalculator.AddGrade("open source assignment", 15, Assignment)
	gradeCalculator.AddGrade("exam 1", 30, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 41, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGradeTypeString(t *testing.T) {
    expectedValue := "assignment"

    actualValue := Assignment.String()

    if expectedValue != actualValue {
        t.Errorf("Expected Assignment.String() to return '%s'; got '%s'", expectedValue, actualValue)
    }
}

func TestPassFailPassingGrade(t *testing.T) {
    expectedValue := "Pass"

    gradeCalculator := NewGradeCalculator(PassFail)

    gradeCalculator.AddGrade("assignment", 75, Assignment)
    gradeCalculator.AddGrade("exam", 75, Exam)
    gradeCalculator.AddGrade("essay", 75, Essay)

    actualValue := gradeCalculator.GetFinalGrade()

    if expectedValue != actualValue {
        t.Errorf("Expected GetFinalGrade to return '%s'; got '%s' instead", expectedValue, actualValue)
    }
}

func TestPassFailFailingGrade(t *testing.T) {
    expectedValue := "Fail"

    gradeCalculator := NewGradeCalculator(PassFail)

    gradeCalculator.AddGrade("assignment", 65, Assignment)
    gradeCalculator.AddGrade("exam", 65, Exam)
    gradeCalculator.AddGrade("essay", 65, Essay)

    actualValue := gradeCalculator.GetFinalGrade()

    if expectedValue != actualValue {
        t.Errorf("Expected GetFinalGrade to return '%s'; got '%s' instead", expectedValue, actualValue)
    }
}
