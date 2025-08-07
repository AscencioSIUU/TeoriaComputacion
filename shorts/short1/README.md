# Examen Corto No. 1 — Teoría de la Computación

**Universidad del Valle de Guatemala**
**Facultad de Ingeniería**
**Curso: Teoría de la Computación**
**Modalidad: Trabajo en casa**
**Entrega individual**

## 📁 Estructura del repositorio

```
short1/
├── config/                 # Funciones auxiliares
│   └── config.go           # Funciones comunes como construcción de expresiones
├── go.mod                  # Archivo de configuración del módulo Go
├── main.go                 # Selector de ejecución por consola
├── problem1.go             # Desarrollo del Problema 1 (AFN y AFD con Thompson y subconjuntos)
├── problem2.go             # Resolución con el Lema de Arden
├── problem3.go             # Demostración con Lema de Bombeo
└── README.md               # Este archivo
```

---

## 🧪 Descripción general

Este examen corto se centra en la implementación de algoritmos clásicos de teoría de autómatas utilizando el lenguaje **Go**. Se divide en tres problemas principales:

### 🧩 Problema 1 — Construcción de Autómatas

Implementa:

- Algoritmo de Thompson para:
  - `letter = [Aa–Bb]`
  - `digit = [0–1]`
  - `digits = digit+`
  - `id = letter(letter|digit)*`
  - `number = digits(.digits)?(E[+-]?digits)?`
- Tabla de transiciones del AFN para `id`
- AFD utilizando la construcción por subconjuntos

📄 Archivo: `problem1.go`

---

### 🔀 Problema 2 — Lema de Arden

[Problema 2](./shorts/short1/problem2.png)

A partir de un autómata definido, se construye la **expresión regular equivalente** utilizando el **Lema de Arden**.

📄 Archivo: `problem2.go`

---

### 📉 Problema 3 — Lema de Bombeo

Demostración formal de que el lenguaje
`L = {a^n b^n # $}`
**no es regular**, usando el **Lema de Bombeo**.

📄 Archivo: `problem3.go`

---

## ▶️ Ejecución del programa

1. Asegúrate de tener Go instalado (`go version`)
2. Ejecuta el programa desde l`a raíz del proyecto:

```bash
go run main.go
``
```
