#  Sistema de Gestión de Alumnos en Go

Proyecto desarrollado en Go para gestionar alumnos y sus materias utilizando estructuras (`struct`), métodos y mapas (`map`).

El sistema permite:

- Agregar materias
- Buscar materias
- Actualizar calificaciones
- Eliminar materias
- Calcular promedio
- Verificar si una materia está aprobada
- Obtener la materia con mayor calificación

---

#  Tecnologías Utilizadas

- Go (Golang)
- Programación Orientada a Objetos con Structs y Métodos
- Maps en Go

---

# Estructura del Proyecto

```bash
.
├── main.go
└── README.md
```

---

#  Conceptos Implementados

## Structs
Se utilizan estructuras para representar:

- `Subject`
- `Student`

---

## Métodos
Se implementan métodos asociados a las estructuras para encapsular comportamiento.

Ejemplos:

```go
func (s Subject) IsApproved() bool
func (s *Student) AddSubject(name string, grade float64)
```

---

## Maps
Las materias del alumno se almacenan en un:

```go
map[string]Subject
```

Donde:
- La clave es el nombre de la materia
- El valor es la estructura `Subject`

---

#  Funcionalidades

##  Agregar Materia

Permite registrar una nueva materia con su calificación.

```go
student.AddSubject("Matemáticas", 90)
```

---

## 🔍 Buscar Materia

Busca una materia por nombre.

```go
student.FindSubject("Programación")
```

---

##  Actualizar Calificación

Actualiza la nota de una materia existente.

```go
student.UpdateGrade("Física", 75)
```

---

##  Eliminar Materia

Elimina una materia del mapa.

```go
student.DeleteSubject("Matemáticas")
```

---

##  Obtener Mayor Calificación

Devuelve la materia con la calificación más alta.

```go
student.HighestGradeSubject()
```

---

## 📊 Calcular Promedio

Calcula el promedio general del alumno.

```go
student.Average()
```

---

##  Verificar Aprobación

Una materia está aprobada si la calificación es mayor o igual a 70.

```go
sub.IsApproved()
```

---

#  Estructuras del Sistema

## Subject

Representa una materia.

| Campo | Tipo |
|---|---|
| Name | string |
| Grade | float64 |

---

## Student

Representa un alumno.

| Campo | Tipo |
|---|---|
| Name | string |
| Subjects | map[string]Subject |

---

#  Ejecución

## Clonar repositorio

```bash
git clone TU_REPOSITORIO
```

---

## Entrar al proyecto

```bash
cd nombre-del-proyecto
```

---

## Ejecutar

```bash
go run main.go
```

---

#  Ejemplo de Salida

```bash
Alumno: Carlos
- Matemáticas : 90
- Programación : 95
- Física : 65

Promedio: 83.33333333333333

--- Pruebas de Funcionalidad ---

Materia encontrada: Programación, Calificación: 95.00, Aprobado: true

Actualizando Física a 75...
Física actualizada: 75.00, Aprobado: true

Materia con mayor calificación: Programación (95.00)

Eliminando Matemáticas...
Alumno: Carlos
- Programación : 95
- Física : 75

Promedio: 85
```

---

#  Objetivos del Proyecto

Este proyecto fue desarrollado para practicar:

- Structs en Go
- Métodos
- Punteros
- Maps
- Encapsulación
- Organización de código
- Lógica de programación

---

#  Aprendizajes

Durante el desarrollo se aplicaron conceptos importantes de Go como:

- Receivers
- Uso de `map`
- Validación de existencia
- Manipulación de estructuras
- Programación modular

---

# Autor

Desarrollado por Marco Vizcarra.

---
