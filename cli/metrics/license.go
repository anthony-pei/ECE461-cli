package metrics

import "fmt"

type license_metric struct {
}

func (l license_metric) calculate(m module) float64 {
	// Object l of type license matrix and m of type module with function get_url()\
	fmt.Println("Calculating license metric")
	return 0.0
}
