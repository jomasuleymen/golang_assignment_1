package schoolsystem

type Teacher struct {
	FirstName string
	LastName  string
	Subject   *Subject
}

func NewTeacher(firstName, lastName string, subject *Subject) *Teacher {
	return &Teacher{
		FirstName: firstName,
		LastName:  lastName,
		Subject:   subject,
	}
}
