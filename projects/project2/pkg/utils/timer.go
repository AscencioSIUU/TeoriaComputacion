package utils

import (
	"fmt"
	"time"
)

// Start inicia un timer y retorna el tiempo inicial
func Start() time.Time {
	return time.Now()
}

// Elapsed retorna la duración desde el tiempo de inicio
func Elapsed(start time.Time) time.Duration {
	return time.Since(start)
}

// FormatDuration formatea una duración de manera legible
func FormatDuration(d time.Duration) string {
	// Microsegundos
	if d < time.Millisecond {
		return fmt.Sprintf("%.2fμs", float64(d.Microseconds()))
	}

	// Milisegundos
	if d < time.Second {
		return fmt.Sprintf("%.2fms", float64(d.Microseconds())/1000.0)
	}

	// Segundos
	if d < time.Minute {
		return fmt.Sprintf("%.2fs", d.Seconds())
	}

	// Minutos
	return fmt.Sprintf("%.2fm", d.Minutes())
}

// Measure ejecuta una función y mide su tiempo de ejecución
func Measure(fn func()) time.Duration {
	start := Start()
	fn()
	return Elapsed(start)
}

// MeasureWithResult ejecuta una función que retorna un valor y mide su tiempo
func MeasureWithResult[T any](fn func() T) (T, time.Duration) {
	start := Start()
	result := fn()
	duration := Elapsed(start)
	return result, duration
}
