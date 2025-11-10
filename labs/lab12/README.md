# Laboratorio 12 - Programación Funcional en Python

## Descripción General
Este laboratorio explora conceptos fundamentales de **programación funcional** en Python utilizando funciones de orden superior como `sorted()`, `map()`, `filter()`, y expresiones `lambda`. Contiene 4 ejercicios que demuestran diferentes técnicas funcionales aplicadas a estructuras de datos comunes.

## Video Explicativo
[Video Explicativo](https://youtu.be/zcawpWmNeR0)

## Requisitos
- **Python 3.6+** (requiere f-strings y características modernas)
- No requiere bibliotecas externas (solo módulos estándar)

## Cómo Ejecutar

### Desde la raíz del proyecto:
```bash
python3 src/main.py
```

### Desde el directorio `src`:
```bash
cd src
python3 main.py
```

### Ejecutar ejercicios individuales:
En el archivo `src/main.py`, descomenta el ejercicio que quieras ejecutar en la sección `if __name__ == "__main__":`:

```python
if __name__ == "__main__":  
    ejercicio1()  # Activo por defecto
    # ejercicio2()  # Descomentar para ejecutar
    # ejercicio3()  # Descomentar para ejecutar
    # ejercicio4()  # Descomentar para ejecutar
```

## Ejercicios y Funciones

### **Ejercicio 1: Ordenamiento de Diccionarios**

#### Función: `order_by_key(data, key)`
**¿Qué hace?**  
Ordena una lista de diccionarios basándose en el valor de una clave específica.

**Parámetros:**
- `data`: Lista de diccionarios a ordenar
- `key`: Nombre de la clave por la cual ordenar (string)

**Retorna:**  
Lista ordenada de diccionarios

**Implementación:**
```python
def order_by_key(data, key): 
    return sorted(data, key = lambda d: d[key])
```

**Explicación técnica:**
- Usa `sorted()` con una función `lambda` como criterio de ordenamiento
- `lambda d: d[key]` extrae el valor de la clave especificada de cada diccionario
- Orden ascendente por defecto (lexicográfico para strings, numérico para números)

**Ejemplo de uso:**
```python
D = [
    {'make': 'Nokia', 'model': 216, 'color': 'Black'},
    {'make': 'Apple', 'model': 2, 'color': 'Silver'},
    {'make': 'Huawei', 'model': 50, 'color': 'Gold'},
    {'make': 'Samsung', 'model': 7, 'color': 'Blue'},
    {'make': 'Xiaomi', 'model': 931, 'color': 'Green'},
]
print(order_by_key(D, 'color'))
```

**Salida esperada (ordenado por 'color'):**
```
[{'make': 'Nokia', 'model': 216, 'color': 'Black'}, 
 {'make': 'Samsung', 'model': 7, 'color': 'Blue'}, 
 {'make': 'Huawei', 'model': 50, 'color': 'Gold'}, 
 {'make': 'Xiaomi', 'model': 931, 'color': 'Green'}, 
 {'make': 'Apple', 'model': 2, 'color': 'Silver'}]
```

---

### **Ejercicio 2: Potencias de Números**

#### Función: `calculate_n_powers(data, n)`
**¿Qué hace?**  
Calcula la potencia `n` de cada elemento en una lista de números.

**Parámetros:**
- `data`: Lista de números
- `n`: Exponente a aplicar (entero)

**Retorna:**  
Lista con cada elemento elevado a la potencia `n`

**Implementación:**
```python
def calculate_n_powers(data, n):
    return list(map(lambda x: x**n, data))
```

**Explicación técnica:**
- Usa `map()` para aplicar una transformación a cada elemento
- `lambda x: x**n` eleva cada número a la potencia `n`
- `list()` convierte el objeto map en una lista

**Ejemplo de uso:**
```python
E = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
print(calculate_n_powers(E, 2))  # Cuadrados
print(calculate_n_powers(E, 3))  # Cubos
```

**Salida esperada:**
```
[1, 4, 9, 16, 25, 36, 49, 64, 81, 100]  # n=2
[1, 8, 27, 64, 125, 216, 343, 512, 729, 1000]  # n=3
```

---

### **Ejercicio 3: Transposición de Matrices**

#### Función: `transpose_matrix(matrix)`
**¿Qué hace?**  
Transpone una matriz (intercambia filas por columnas).

**Parámetros:**
- `matrix`: Lista de listas representando una matriz

**Retorna:**  
Matriz transpuesta (filas → columnas, columnas → filas)

**Implementación:**
```python
def transpose_matrix(matrix):
    return list(map(lambda *row: list(row), zip(*matrix)))
```

**Explicación técnica:**
- `zip(*matrix)` desempaqueta las filas y las agrupa por columnas
- `map(lambda *row: list(row), ...)` convierte cada tupla en lista
- El operador `*` desempaqueta los argumentos

**Ejemplo de uso:**
```python
X = [
    [1, 2, 3, 1],
    [4, 5, 6, 0],
    [7, 8, 9, -1]
]
print(transpose_matrix(X))
```

**Salida esperada:**
```
[[1, 4, 7], [2, 5, 8], [3, 6, 9], [1, 0, -1]]
```

**Visualización:**
```
Original (3×4):          Transpuesta (4×3):
[1, 2, 3, 1]            [1, 4, 7]
[4, 5, 6, 0]     →      [2, 5, 8]
[7, 8, 9, -1]           [3, 6, 9]
                        [1, 0, -1]
```

---

### **Ejercicio 4: Filtrado de Elementos**

#### Función: `delete_elements(list1, list2)`
**¿Qué hace?**  
Elimina de `list1` todos los elementos que aparecen en `list2`.

**Parámetros:**
- `list1`: Lista original
- `list2`: Lista con elementos a eliminar

**Retorna:**  
Nueva lista con elementos de `list1` que NO están en `list2`

**Implementación:**
```python
def delete_elements(list1, list2):
    return list(filter(lambda x: x not in list2, list1))
```

**Explicación técnica:**
- Usa `filter()` para seleccionar elementos que cumplan una condición
- `lambda x: x not in list2` retorna `True` si el elemento NO está en `list2`
- Operación equivalente a diferencia de conjuntos, pero mantiene el orden

**Ejemplo de uso:**
```python
F = ['manzana', 'banana', 'amarillo', 'gris', 'rojo', 'naranja', 'morado']
G = ['manzana', 'banana', 'naranja', 'rojo', 'mango']
print(delete_elements(F, G))
```

**Salida esperada:**
```
['amarillo', 'gris', 'morado']
```

---

## 📁 Estructura del Proyecto

```
lab12/
├── README.md                 # Este archivo
└── src/
    └── main.py              # Código principal con los 4 ejercicios
```

## 🔍 Conceptos Clave de Programación Funcional

### 1. **Funciones Lambda**
Funciones anónimas de una sola expresión:
```python
lambda parametros: expresion
```

### 2. **map()**
Aplica una función a cada elemento de un iterable:
```python
map(funcion, iterable)  # Retorna iterador
```

### 3. **filter()**
Filtra elementos que cumplen una condición:
```python
filter(funcion_booleana, iterable)  # Retorna iterador
```

### 4. **sorted()**
Ordena un iterable usando una función clave:
```python
sorted(iterable, key=funcion)  # Retorna lista ordenada
```

---

**Autor:** Ernesto Ascencio 23009
**Fecha:** 2025  
**Lenguaje:** Python 3.6+
