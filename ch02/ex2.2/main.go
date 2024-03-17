package main

import (
	"fmt"
	"os"
	"strconv"
)

type meter float64
type foot float64

func (m meter) String() string { return fmt.Sprintf("%g m", m) }
func (f foot) String() string  { return fmt.Sprintf("%g f", f) }

func MeterToFoot(m meter) foot { return foot(m * 3.281) }
func FootToMeter(f foot) meter { return meter(f / 3.281) }

type kilogram float64
type pound float64

func (kg kilogram) String() string { return fmt.Sprintf("%g kg", kg) }
func (lb pound) String() string    { return fmt.Sprintf("%g lb", lb) }

func KilogramToPound(kg kilogram) pound { return pound(kg * 2.205) }
func PoundToKilogram(lb pound) kilogram { return kilogram(lb / 2.205) }

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		m := meter(t)
		f := foot(t)

		kg := kilogram(t)
		lb := pound(t)

		fmt.Println("== distance ==")
		fmt.Printf("%s = %s\n", m, MeterToFoot(m))
		fmt.Printf("%s = %s\n", f, FootToMeter(f))
		fmt.Println("== mass ==")
		fmt.Printf("%s = %s\n", kg, KilogramToPound(kg))
		fmt.Printf("%s = %s\n\n", lb, PoundToKilogram(lb))

	}
}
