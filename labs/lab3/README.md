# Laboratorio 2

## 📁 Estructura del repositorio

```
lab3/
├── config                  # Lógica de construcción de árbol sintáctico y graficación
│   └── PostfixToTree.go    # Convierte postfix a árbol, genera DOT y PNG
│   └── helpers.go          # Funciones para formato y expansión de regex
├── ejercicio1.go           # Ejecución principal: infix a postfix, arbol AST
├── expressions1.txt        # Expresiones infix de prueba (una por línea)
├── dotfiles                # Archivo Graphviz del árbol generado (ejemplo)
├── pngfiles                # Imagen del árbol sintáctico (ejemplo)
└── README.md               # Este archivo
```

---

## ⚙️ Requisitos previos

### Graphviz

- **macOS**
  brew install graphviz
- **Ubuntu/Debian**
  sudo apt install graphviz
- **Verifica instalación**
  dot -V

---

## 🛠️ Instalación

1. **Clona** tu repositorio (si aún no lo has hecho):
   ```bash
     git clone https://github.com/usuario/lab3-teocomp.git
     cd lab3
   ```
2. Ejecuta el laboratorio:
   ```bash
     cd labs/labs3
     go run ejercicio1.go
   ```

---

## ▶️ Ejecución

### Video en ejecución

[https://youtu.be/m2pA9_1GFxA](https://youtu.be/m2pA9_1GFxA)

### 🔹 Ejercicio 1 — Árbol Sintáctico de Expresiones Regulares

### Expresiones utilizadas:

- (a*|b*)+
- ((ε|a)|b*)*
- (a|b)abb(a|b)
- 0?(1?)?0\*

### Funcionamiento:

1. Expande + y ? usando:
   - a+ → aa\*
   - a? → (a|ε)
2. Inserta concatenación explícita con .
3. Convierte a postfix usando Shunting Yard
4. Construye árbol sintáctico (AST) con pila
5. Genera archivo .dot para Graphviz
6. Ejecuta dot para crear imagen .png

---
