package schoolsystem

type Pupil struct {
	FirstName string
	LastName  string
	Grades    map[string][]int
}

func NewPupil(firstName, lastName string) *Pupil {
	return &Pupil{
		FirstName: firstName,
		LastName:  lastName,
		Grades:    make(map[string][]int),
	}
}

func (pupil *Pupil) PutMark(subject *Subject, mark int) {
	var subjectName string = subject.Name

	if grades, ok := pupil.Grades[subjectName]; ok {
		pupil.Grades[subjectName] = append(grades, mark)
	}
}
