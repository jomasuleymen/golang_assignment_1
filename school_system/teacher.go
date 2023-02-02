package schoolsystem

type Teacher struct {
	FirstName string
	LastName  string
	Subject   *Subject
}

func (teacher *Teacher) PutMark(pupil *Pupil, mark int) {
	pupil.PutMark(teacher.Subject, mark)
}

func NewTeacher(firstName, lastName string, subject *Subject) *Teacher {
	return &Teacher{
		FirstName: firstName,
		LastName:  lastName,
		Subject:   subject,
	}
}
