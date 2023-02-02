package schoolsystem

type Group struct {
	Name     string
	Pupils   []*Pupil
	Subjects []*Subject
}

func (group *Group) AddPupil(pupil *Pupil) {
	for _, pup := range group.Pupils {
		if pup == pupil {
			return
		}
	}

	group.Pupils = append(group.Pupils, pupil)

	for _, subject := range group.Subjects {
		pupil.Grades[subject.Name] = make([]int, 0)
	}
}

func (group *Group) AddSubjects(subjects ...*Subject) {
	group.Subjects = subjects
}

func NewGroup(name string) *Group {
	return &Group{
		Name:     name,
		Pupils:   make([]*Pupil, 0),
		Subjects: make([]*Subject, 0),
	}
}
