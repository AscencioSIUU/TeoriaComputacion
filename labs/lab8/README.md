# Laboratorio 8 — Guía rápida (estructura + comandos)

> Proyecto en Go con **un solo `main.go`** que ejecuta los Ej. 1–3 y además **genera gráficas** desde los CSV.
> El código fuente está en **inglés**; esta guía está en español.

---

## Estructura de carpetas utilizada

```
lab8/
├─ go.mod
├─ main.go
├─ config/
│  ├─ types.go          # Runner, Result
│  ├─ bench.go          # TimeN (medición) + AppendCSV
│  └─ plotcsv.go        # PlotCSV / PlotCSVWithOpts (graficado PNG)
├─ ex1/
│  └─ ex1.go            # Ex1(n int) uint64    -> O(n^2 log n)
├─ ex2/
│  └─ ex2.go            # Ex2(n int) uint64    -> O(n)
├─ ex3/
│  └─ ex3.go            # Ex3(n int) uint64    -> O(n^2)
├─ results/             # CSVs generados (ex01.csv, ex02.csv, ex03.csv)
└─ plots/               # PNGs generados (ex01.png, ex02.png, ex03.png)
```

**Notas**

* Import path interno: usa el de tu `go.mod` (por ejemplo `example.com/lab8/...` o `lab8/...` si así lo tienes).
* `main.go` acepta **dos modos**:

  * `-mode=run` → ejecuta un ejercicio y guarda CSV.
  * `-mode=plot` → lee un CSV y genera un PNG.

---

## Complejidades (resumen para el informe)

* **Ejercicio 1:** `i: n/2..n`, `j: 1..n/2`, `k: 1,2,4,...≤n` → **O(n² log n)**; espacio **O(1)**.
* **Ejercicio 2:** doble `for` pero `break` inmediato en el interno → **O(n)** (si `n≤1`, **O(1)**); espacio **O(1)**.
* **Ejercicio 3:** `i: 1..n/3`, `j: 1..n` con paso 4 → **O(n²)**; espacio **O(1)**.

---

Video: 

---

## Preparación del entorno

```bash
# 1) (una sola vez) inicializa el módulo
go mod init <tu-modulo>     # p.ej. example.com/lab8

# 2) (si graficarás) instala la librería de gráficos
go get gonum.org/v1/plot/...
go mod tidy
```

---

## Uso — modo ejecución (generar CSV)

Flags principales en **modo run**:

* `-exercise=1|2|3`  → elige el ejercicio.
* `-n=<int>`         → un tamaño (si no usas -ns).
* `-ns="a,b,c"`      → lista de tamaños (override de `-n`).
* `-runs=<int>`      → repeticiones para promediar (≥1).
* `-out=<path.csv>`  → salida CSV (si se omite: `results/ex0X.csv` según ejercicio).

### Ejercicio 1

```bash
# Set sugerido (1e6 suele ser prohibitivo en ex1)
go run . -exercise=1 -ns="1,10,100,1000,10000,100000" -runs=3
# CSV por defecto: results/ex01.csv
```

### Ejercicio 2

```bash
go run . -exercise=2 -ns="1,10,100,1000,10000,100000,1000000" -runs=3
# CSV por defecto: results/ex02.csv
```

### Ejercicio 3

```bash
go run . -exercise=3 -ns="1,10,100,1000,10000,100000" -runs=3
# CSV por defecto: results/ex03.csv
```

### Ejercicio 4

```bash
go run . -mode=gen-linear -ns="1,10,100,1000,10000,100000" -p=0.5 -outdir="results/ex04"
```


---

## Uso — modo gráfico (generar PNG desde CSV)

Flags en **modo plot**:

* `-inplot=<csv>`    → CSV de entrada (obligatorio).
* `-outplot=<png>`   → PNG de salida (por defecto `plots/out.png`).
* `-title="..."`     → título del gráfico.
* `-logx`, `-logy`   → ejes logarítmicos.
* `-ymin=<float>`    → mínimo del eje Y (opcional; útil para evitar “pared” en 0).
* `-nolines`         → solo puntos (sin línea de conexión).

### Ejemplos

```bash
# EX1: tiempos crecen mucho → eje Y log
go run . -mode=plot \
  -inplot=results/ex01.csv \
  -outplot=plots/ex01.png \
  -title="Exercise 1: n vs time" \
  -logy=true

# EX2: lineal
go run . -mode=plot \
  -inplot=results/ex02.csv \
  -outplot=plots/ex02.png \
  -title="Exercise 2: n vs time"

# EX3: log-log para ver tendencia ~cuadrática (recta con pendiente ~2)
go run . -mode=plot \
  -inplot=results/ex03.csv \
  -outplot=plots/ex03.png \
  -title="Exercise 3: n vs time" \
  -logx=true -logy=true -ymin=0.001

# EX4:
# Best (Θ(1)) – escala lineal está bien
go run . -mode=plot -inplot=results/ex04/ex04_best.csv \
  -outplot=plots/ex04_best.png -title="Linear Search — Best case" -ymin=0.8

# Avg (éxito uniforme) – O(n)
go run . -mode=plot -inplot=results/ex04/ex04_avg_success.csv \
  -outplot=plots/ex04_avg_success.png -title="Linear Search — Average (success uniform)" -ymin=0.8

# Avg (mixto p=0.5) – O(n)
go run . -mode=plot -inplot=results/ex04/ex04_avg_mixed_p0.50.csv \
  -outplot=plots/ex04_avg_mixed.png -title="Linear Search — Average (p=0.5)" -ymin=0.8

# Worst – O(n)
go run . -mode=plot -inplot=results/ex04/ex04_worst.csv \
  -outplot=plots/ex04_worst.png -title="Linear Search — Worst case" -ymin=0.8
```

---

## Notas de medición y buenas prácticas

* **Precisión**: en `TimeN` se usa `Duration.Nanoseconds()` convertido a ms con decimales para evitar `0.000 ms` en tamaños pequeños.
* **Evitar I/O** en bucles: no uses `fmt.Printf` dentro de los loops; se simula trabajo con un contador (`uint64`) para impedir que el compilador elimine el cuerpo.
* **Persistencia incremental**: cada fila se escribe al CSV inmediatamente; si cancelas (Ctrl+C), lo ya medido queda guardado.
* **Tamaños grandes**: documenta en el CSV/README si ciertos `n` fueron “too slow on my machine”, especialmente en **Ex1** y **Ex3**.

---

## Referencias oficiales de Go

* `flag` (parsing de CLI): [https://pkg.go.dev/flag](https://pkg.go.dev/flag)
* `time` (durations, `Nanoseconds`/`Since`): [https://pkg.go.dev/time](https://pkg.go.dev/time)
* `encoding/csv` (lectura/escritura CSV): [https://pkg.go.dev/encoding/csv](https://pkg.go.dev/encoding/csv)
* `image/color` (colores en gráficos): [https://pkg.go.dev/image/color](https://pkg.go.dev/image/color)
* Módulos e imports (`go.mod`, `go get`, `go mod tidy`): [https://go.dev/doc/modules/managing](https://go.dev/doc/modules/managing)
* Tutorial `go run` / `go build`: [https://go.dev/doc/tutorial/](https://go.dev/doc/tutorial/)

**Librería de gráficos (Gonum)**

* `gonum/plot`: [https://pkg.go.dev/gonum.org/v1/plot](https://pkg.go.dev/gonum.org/v1/plot)
* `plotter`: [https://pkg.go.dev/gonum.org/v1/plot/plotter](https://pkg.go.dev/gonum.org/v1/plot/plotter)
* `vg`: [https://pkg.go.dev/gonum.org/v1/plot/vg](https://pkg.go.dev/gonum.org/v1/plot/vg)

Si quieres, te armo un `Makefile` con targets `ex1`, `ex2`, `ex3`, `plot1`, `plot2`, `plot3` para automatizar todo.
