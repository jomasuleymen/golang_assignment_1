package schoolsystem

import "fmt"

func StartApp() {
	mathSubject := NewSubject("Math")
	mathTeacher := NewTeacher("Aigul", "Imanbayeva", mathSubject)

	group1 := NewGroup("11B")
	pupil := NewPupil("Zhomart", "Suleymen")

	group1.AddSubjects(mathSubject)
	group1.AddPupil(pupil)

	mathTeacher.PutMark(pupil, 4)
	pupil.PutMark(mathSubject, 5)

	fmt.Println(pupil)
}
