package esepunittests

type GradingMode int

const (
    LetterGrade GradingMode = iota
    PassFail
)

type GradeCalculator struct {
    grades []Grade
    mode GradingMode
}

type GradeType int

const (
	Assignment GradeType = iota
	Exam
	Essay
)

var gradeTypeName = map[GradeType]string{
	Assignment: "assignment",
	Exam:       "exam",
	Essay:      "essay",
}

func (gt GradeType) String() string {
	return gradeTypeName[gt]
}

type Grade struct {
	Name  string
	Grade int
	Type  GradeType
}

func NewGradeCalculator(mode GradingMode) *GradeCalculator {
    return &GradeCalculator{
        grades: make([]Grade, 0),
        mode: mode,
    }
}

func (gc *GradeCalculator) GetFinalGrade() string {
	numericalGrade := gc.calculateNumericalGrade()

	if gc.mode == PassFail {
        if numericalGrade >= 70 {
            return "Pass"
        }
        return "Fail"
    }

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

func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
	gc.grades = append(gc.grades, Grade{
        Name:  name,
        Grade: grade,
        Type:  gradeType,
    })
}

func (gc *GradeCalculator) calculateNumericalGrade() int {
	assignment_average := computeAverage(gc.getGradesByType(Assignment))
	exam_average := computeAverage(gc.getGradesByType(Exam))
	essay_average := computeAverage(gc.getGradesByType(Essay))

	weighted_grade := float64(assignment_average)*.5 + float64(exam_average)*.35 + float64(essay_average)*.15

	return int(weighted_grade)
}

func computeAverage(grades []Grade) int {
	sum := 0

	for _, grade := range grades {
		sum += grade.Grade
	}

	return sum / len(grades)
}

func (gc *GradeCalculator) getGradesByType(gradeType GradeType) []Grade {
    grades := make([]Grade, 0)

    for _, grade := range gc.grades {
        if grade.Type == gradeType {
            grades = append(grades, grade)
        }
    }

    return grades
}