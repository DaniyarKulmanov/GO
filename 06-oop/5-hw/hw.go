package hw

import (
	"fmt"
	"math"
)

// Geom По условиям задачи, координаты не могут быть меньше 0.
type Geom struct {
	X1, Y1, X2, Y2 float64
}

// CalculateDistance Переименован Geom -> g рекомендации наименивания.
// CalculateDistance Geom -> *Geom меняем на указатель, напрямую работать с объектом (не создавать копию).
// distance -> d рекомендации наименивания более коротко.
func (g *Geom) CalculateDistance() (d float64) {
	if g.X1 < 0 || g.X2 < 0 || g.Y1 < 0 || g.Y2 < 0 {
		fmt.Println("Координаты не могут быть меньше нуля")
		return -1
	}
	//d этот код выполнится только если if вернет false.
	d = math.Sqrt(math.Pow(g.X2-g.X1, 2) + math.Pow(g.Y2-g.Y1, 2))
	// возврат расстояния между точками
	return d
}
