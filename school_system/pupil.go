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

func (pupil *Pupil) PutMark(teacher *Teacher, mark int) {
	grades, ok := pupil.Grades[teacher.Subject.Name]
	if !ok {
		return
	}

	pupil.Grades[teacher.Subject.Name] = append(grades, mark)
}
