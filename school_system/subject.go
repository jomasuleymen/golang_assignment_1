package schoolsystem

type Subject struct {
	Name string
}

func NewSubject(subjectName string) *Subject {
	return &Subject{Name: subjectName}
}
