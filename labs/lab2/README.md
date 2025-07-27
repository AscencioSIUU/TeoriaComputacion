# Laboratorio 2

## 📁 Estructura del repositorio

```
lab2/
├── exercise2.go         # Verificador de expresiones balanceadas
├── expressions2.txt     # Expresiones de prueba para el ejercicio 2
├── exercise3.go         # Shunting Yard: infix → postfix
├── expressions3.txt     # Expresiones de prueba para el ejercicio 3
└── README.md            # Este archivo
```

---

## ⚙️ Requisitos previos

- **Go** (versión 1.18 o superior) instalado en tu máquina.
- **Git** instalado (para clonar o versionar).
- Línea de comandos (Terminal, PowerShell, etc.).

---

## 🛠️ Instalación

1. **Clona** tu repositorio (si aún no lo has hecho):
   ```bash
   https://github.com/AscencioSIUU/TeoriaComputacion.git
   cd labs/lab2
   ```
2. Asegúrate de tener Go en tu PATH:
   ```bash
   go version
   ```
   Debe mostrar algo como `go version go1.24.1 darwin/arm64`.

---

## ▶️ Ejecución

### 🔹 Ejercicio 2 — Verificador de expresiones balanceadas

1. Coloca tus expresiones en `expressions2.txt`, una por línea.
2. Para **ejecutar directamente** sin compilar:
   ```bash
   go run exercise2.go
   ```
3. Para **compilar** y luego ejecutar:
   ```bash
   go build -o verifier exercise2.go
   ./verifier
   ```

El programa leerá `expressions2.txt`, mostrará paso a paso las operaciones de pila y el resultado de cada línea.

---

### 🔹 Ejercicio 3 — Shunting Yard (infix → postfix)

1. Coloca tus expresiones en `expressions3.txt`, una por línea.
2. Para **ejecutar directamente**:
   ```bash
   go run exercise3.go
   ```
3. Para **compilar** y luego ejecutar:
   ```bash
   go build -o shunting exercise3.go
   ./shunting
   ```

El programa leerá `expressions3.txt`, mostrará los pasos del algoritmo y la conversión a notación postfix.

---

## 📄 Archivos de entrada

- **expressions2.txt**
  Contiene ejemplos como:
  ```
  a(a|b)*b+a?
  A(a|b)bB*[az]b]
  (a*b*c*d*(a|e|i|o|u))e*f*g*h){1,2}
  ^[aZ].com{5,30}
  ([[az][AZ]](((((.|;)|;)|.)|.)){10,20})*)
  ```
- **expressions3.txt**
  Infix de expresiones regulares (mismas del ejercicio 1), p. ej.:
  ```
  (a|b)c
  (a|b)*abb(a|b)*
  (a*|b*)*
  0?(1?)?0*
  …etc.
  ```

---

## 🔍 Investigación

En esta sección incluye tu análisis teórico y referencias según el enunciado:

1. **Explicación del algoritmo de Shunting Yard**

   - ¿Cómo funciona?
   - Complejidad temporal y espacial.

2. **Método de Thompson**

   - Comparación con otros métodos de construcción de autómatas.
   - Ventajas y limitaciones.

3. **Casos de prueba adicionales**

   - Listado de expresiones (con y sin balanceo).
   - Resultados esperados vs. obtenidos.

4. **Referencias bibliográficas**
   - Artículos, libros o recursos web que consultaste.

> _Sigue el formato APA 7ª edición para tus referencias._

---
