package schoolsystem

import "fmt"

func StartApp() {
	mathSubject := NewSubject("Math")
	mathTeacher := NewTeacher("Aigul", "Imanbayeva", mathSubject)

	group1 := NewGroup("11B")
	pupil1 := NewPupil("Ainash", "Suleymenova")
	pupil2 := NewPupil("Zhomart", "Suleymen")

	group1.AddSubjects(mathSubject)
	group1.AddPupil(pupil1)
	group1.AddPupil(pupil2)

	pupil1.PutMark(mathTeacher, 5)
	pupil1.PutMark(mathTeacher, 2)
	pupil1.PutMark(mathTeacher, 4)

	pupil2.PutMark(mathTeacher, 5)
	pupil2.PutMark(mathTeacher, 5)
	pupil2.PutMark(mathTeacher, 5)

	fmt.Println(pupil2)
}
