Laboratorio 4
📁 Estructura del repositorio

```
lab4/
├── cmd/
│   └── lab4/
│       └── main.go            # Orquestación: lectura, pipeline, DOT/PNG, simulación
├── config/
│   └── config.go              # Expand (+,?), Format (.), Infix→Postfix (Shunting Yard), helpers
├── graphviz/
│   └── dot.go                 # Exportar NFA a DOT y generar PNG (Graphviz)
├── nfa/
│   └── simulate.go            # Simulador AFN (cierre-ε y transición por símbolo)
├── regex/
│   └── ast.go                 # Construcción del AST desde postfix
├── thompson/
│   └── nfa.go                 # Construcción de AFN a partir del AST (algoritmo de Thompson)
├── docs/
│   └── Ejercicio2.pdf         # Demostración (Lema de Bombeo) — Ejercicio 2
├── dotout/                    # Salida: archivos .dot generados (se crea en runtime)
├── pngout/                    # Salida: archivos .png generados (se crea en runtime)
├── input.txt                  # Entradas: "regex;w" (una por línea)
└── README.md                  # Este archivo
```

---

⚙️ Requisitos previos
- Go ≥ 1.20
- Graphviz (necesario para generar PNG desde DOT)

macOS
```
brew install go graphviz
```

Ubuntu/Debian
```
sudo apt update
sudo apt install golang graphviz -y
```

Verifica Graphviz:
```
dot -V
```

🛠️ Instalación
Clonar el repositorio:

```
git clone https://github.com/AscencioSIUU/TeoriaComputacion.git
cd TeoriaComputacion/labs/lab4
```

Ejecutar el laboratorio:

```
cd ejercicio1
go run main.go
```

---

▶️ Ejecución
Video de demostración

https://youtu.be/hj7xuHoQCto

---

## 🔹 Ejercicio 1 — Algoritmo de Thompson y Simulación de AFN

Expresiones utilizadas:
```
(a*|b*)+;aaaa
((ε|a)|b*)*;abbb
(a|b)*abb(a|b)*;xxabbx
0?(1?)?0*;0100
```

¿Qué hace cada parte?

- config/config.go
     - ExpandRegexExtensions: `X+ → X.X*`, `X? → (X|ε)` (sin dejar +/? en la expresión).
     - FormatRegex: inserta . para concatenaciones implícitas.
     - InfixToPostfix: convierte infix → postfix (Shunting Yard).
     - Utilidades (IsAlphanumeric, precedencias, normalización de ε).
- regex/ast.go
     - Construye el AST desde la notación postfix usando una pila.
     - Nodos: Literal, Concat, Union, Star.
- thompson/nfa.go
     - Implementa Thompson: a partir del AST genera un NFA con un único Start y Accept.
     - Usa transiciones por símbolo y transiciones ε.

- nfa/simulate.go
     - Simulador AFN: cierre-ε + transición por símbolo a lo largo de w.
     - Acepta si el `Accept` está en el conjunto de estados actuales al final.
- graphviz/dot.go
     - WriteDOT: exporta el NFA a formato DOT.
     - GeneratePNGFromDot: ejecuta dot para producir el PNG.
- cmd/lab4/main.go
     - Lee input.txt (formato `regex;w`).
     - Pipeline: expand → format → postfix → AST → Thompson.
     - Exporta .dot y .png.
     - Simula w y muestra `sí`/`no`.


🔗 Referencias

- Go
     - bufio.Scanner (lectura línea a línea): https://pkg.go.dev/bufio
     - flag (argumentos CLI): https://pkg.go.dev/flag
     - fmt (salida formateada): https://pkg.go.dev/fmt
     - os, os/exec (archivos, procesos externos como Graphviz): https://pkg.go.dev/os · https://pkg.go.dev/os/exec
     - path/filepath (rutas portables): https://pkg.go.dev/path/filepath
     - strings (manipulación de strings): https://pkg.go.dev/strings
     - unicode (clases de caracteres y runes): https://pkg.go.dev/unicode
     - builtin.rune (tipo de datos para caracteres Unicode): https://pkg.go.dev/builtin#rune
     - Structs en Go (definición y uso): https://go.dev/tour/moretypes/2
     - Composite literals (inicialización de structs y slices): https://go.dev/doc/effective_go#composite_literals
     - Métodos en structs: https://go.dev/doc/effective_go#methods
- Teoría de Computación
     - Árbol sintáctico abstracto (AST): https://en.wikipedia.org/wiki/Abstract_syntax_tree
     - Autómata finito no determinista (NFA): https://en.wikipedia.org/wiki/Nondeterministic_finite_automaton
     - Algoritmo de Thompson (construcción de AFN a partir de regex): https://en.wikipedia.org/wiki/Thompson%27s_construction
- Graphviz
     - Lenguaje DOT (formato de grafos): https://graphviz.org/doc/info/lang.html
