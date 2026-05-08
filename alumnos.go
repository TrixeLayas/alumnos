package main

import "fmt"

/* Leer alumnos.md */

// Materia
type Subject struct {
	Name  string
	Grade float64
}

// Función para saber si está aprobado (>= 70)
func (s Subject) IsApproved() bool {
	return s.Grade >= 70
}

// Alumno
type Student struct {
	Name     string
	Subjects map[string]Subject
}

// Agregar materia
func (s *Student) AddSubject(name string, grade float64) {
	if s.Subjects == nil {
		s.Subjects = make(map[string]Subject)
	}

	s.Subjects[name] = Subject{
		Name:  name,
		Grade: grade,
	}
}

// Funcion para buscar una materia por nombre
func (s Student) FindSubject(name string) (Subject, bool) {
	sub, exists := s.Subjects[name]
	return sub, exists
}

// Funcion eliminar materia
func (s *Student) DeleteSubject(name string) {
	delete(s.Subjects, name)
}

// Funcion actualizar calificación
func (s *Student) UpdateGrade(name string, newGrade float64) {
	if sub, exists := s.Subjects[name]; exists {
		sub.Grade = newGrade
		s.Subjects[name] = sub
	}
}

// Funcion mostrar materia con mayor calificación
func (s Student) HighestGradeSubject() (Subject, bool) {
	if len(s.Subjects) == 0 {
		return Subject{}, false
	}

	var highest Subject
	first := true

	for _, sub := range s.Subjects {
		if first || sub.Grade > highest.Grade {
			highest = sub
			first = false
		}
	}

	return highest, true
}

// Promedio
func (s Student) Average() float64 {
	total := 0.0

	for _, sub := range s.Subjects {
		total += sub.Grade
	}

	if len(s.Subjects) == 0 {
		return 0
	}

	return total / float64(len(s.Subjects))
}

// Mostrar info
func (s Student) Print() {
	fmt.Println("Alumno:", s.Name)

	for _, sub := range s.Subjects {
		fmt.Println("-", sub.Name, ":", sub.Grade)
	}

	fmt.Println("Promedio:", s.Average())
}

func main() {
	student := Student{Name: "Carlos"}

	student.AddSubject("Matemáticas", 90)
	student.AddSubject("Programación", 95)
	student.AddSubject("Física", 65)

	student.Print()

	// Probar búsqueda
	fmt.Println("\n--- Pruebas de Funcionalidad ---")
	if sub, exists := student.FindSubject("Programación"); exists {
		fmt.Printf("Materia encontrada: %s, Calificación: %.2f, Aprobado: %v\n", sub.Name, sub.Grade, sub.IsApproved())
	}

	// Probar actualización
	fmt.Println("\nActualizando Física a 75...")
	student.UpdateGrade("Física", 75)
	if sub, exists := student.FindSubject("Física"); exists {
		fmt.Printf("Física actualizada: %.2f, Aprobado: %v\n", sub.Grade, sub.IsApproved())
	}

	// Probar materia con mayor calificación
	if highest, exists := student.HighestGradeSubject(); exists {
		fmt.Printf("\nMateria con mayor calificación: %s (%.2f)\n", highest.Name, highest.Grade)
	}

	// Probar eliminación
	fmt.Println("\nEliminando Matemáticas...")
	student.DeleteSubject("Matemáticas")
	student.Print()
}
