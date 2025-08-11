Laboratorio 4
📁 Estructura del repositorio

```
lab4/
├── ejercicio1/                   # Implementación del algoritmo de Thompson y simulación de AFN
│   ├── shunting_yard.go          # Conversión de expresión regular infix a postfix
│   ├── syntax_tree.go            # Construcción del Árbol Sintáctico (AST)
│   ├── thompson.go               # Construcción de AFN a partir del AST
│   ├── afn_simulation.go         # Simulación del AFN con una cadena de entrada
│   ├── helpers.go                # Funciones auxiliares (manejo de concatenación, +, ?, etc.)
│   └── main.go                   # Programa principal
├── ejercicio2/                   # Demostración con Lema de Bombeo
│   └── pumping_lemma.pdf         # Documento con la demostración
├── inputs/
│   └── expressions.txt           # Lista de expresiones regulares de prueba (una por línea)
├── outputs/
│   ├── afn_graphs/                # Imágenes PNG de los AFNs generados
│   └── simulation_results.txt     # Resultados de simulación de cadenas
└── README.md                     # Este archivo
```

---

⚙️ Requisitos previos
Go (Golang)
macOS

```
brew install go
```

Ubuntu/Debian

```
sudo apt install golang
```

Graphvi
macOS

```
brew install graphviz
```

Ubuntu/Debian

```
sudo apt install graphviz
```

Verifica instalación

```
dot -V
```

🛠️ Instalación
Clonar el repositorio:

```
git clone https://github.com/AscencioSIUU/TeoriaComputacion.git
cd labs/lab4
```

Ejecutar el laboratorio:

```
cd ejercicio1
go run main.go
```

---

▶️ Ejecución
Video de demostración
Enlace a YouTube — No listado

---

## 🔹 Ejercicio 1 — Algoritmo de Thompson y Simulación de AFN

Expresiones utilizadas:

```
(a*|b*)+
((ε|a)|b*)*
(a|b)abb(a|b)
0?(1?)?0\*
```

Funcionamiento:
Expande operadores:

```
a+ → aa\*
a? → (a|ε)
```

Inserta concatenación explícita con .

Convierte a postfix con Shunting Yard

Construye el Árbol Sintáctico (AST)

Aplica Algoritmo de Thompson para generar el AFN

Guarda un .dot y lo convierte a .png con Graphviz

## Simula la cadena w para determinar si pertenece a L(r)

## 🔹 Ejercicio 2 — Lema de Bombeo

Se demuestra que el lenguaje
A = { yy | y ∈ {0,1} }\*
no es regular.

Basado en el caso:
S = 0^p 1 0^p 1
con p = pumping length.
