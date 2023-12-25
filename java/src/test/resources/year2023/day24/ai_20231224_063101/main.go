package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

type Hailstone struct {
	Position []int64
	Velocity []int64
}

func parseHailstone(line string) Hailstone {
	parts := strings.Split(strings.ReplaceAll(line, " @ ", ","), ",")
	posX, _ := strconv.ParseInt(parts[0], 10, 64)
	posY, _ := strconv.ParseInt(parts[1], 10, 64)
	posZ, _ := strconv.ParseInt(parts[2], 10, 64)

	velX, _ := strconv.ParseInt(parts[3], 10, 64)
	velY, _ := strconv.ParseInt(parts[4], 10, 64)
	velZ, _ := strconv.ParseInt(parts[5], 10, 64)

	return Hailstone{
		Position: []int64{posX, posY, posZ},
		Velocity: []int64{velX, velY, velZ},
	}
}

func gcd(a, b int64) int64 {
	var r, t int64
	if a < b {
		t = b
		b = a
		a = t
	}
	for b != 0 {
		r = a % b
		a = b
		b = r
	}
	return a
}

func extendedGCD(a, b int64) (int64, int64, int64) {
	var x, y int64
	x0, x1, y0, y1 := int64(1), int64(0), int64(0), int64(1)
	for b != 0 {
		q := a / b
		x, x0, x1 = x0-q*x1, x1, x0
		y, y0, y1 = y0-q*y1, y1, y0
		a, b = b, a%b
	}
	return a, x, y
}

func crt(a1, m1, a2, m2 int64) (int64, int64, error) {
	gcd, x1, _ := extendedGCD(m1, m2)
	if (a1-a2)%gcd != 0 {
		return 0, 0, fmt.Errorf("No solution")
	}
	return (a1 + x1*(a2-a1)/gcd*m1) % (m1 / gcd * m2), m1 / gcd * m2, nil
}

func findCollision(h []Hailstone) []int64 {
	cr := []int64{0, 1}
	var xval, xmod int64
	for i := 0; i < 3; i++ {
		for n := 0; n < len(h); n++ {
			for m := n + 1; m < len(h); m++ {
				A := h[m].Velocity[i] - h[n].Velocity[i]
				B := h[n].Position[i] - h[m].Position[i]

				// skip parallel paths
				if A == 0 {
					continue
				}

				AbsA := A
				if AbsA < 0 {
					AbsA = -AbsA
				}

				g := gcd(AbsA, B)
				// skip paths that are not going to collide
				if B%g != 0 {
					continue
				}

				if A < 0 {
					A = -A
					B = -B
				}

				xval = B / g
				xmod = A / g
				for xval < 0 {
					xval += xmod
				}
				xval %= xmod
				cr1, cr2, _ := crt(cr[0], cr[1], xval, xmod)
				cr = []int64{cr1, cr2}
				if cr[1] < 0 {
					cr[1] = -cr[1]
				}
			}
		}
	}
	return cr
}

func main() {
	file, err := os.Open("/Users/vkad2506/AdventOfCode/java/src/test/resources/year2023/day24/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var hailstones []Hailstone
	for scanner.Scan() {
		line := scanner.Text()
		hailstones = append(hailstones, parseHailstone(line))
	}

	collision := findCollision(hailstones)

	x, y, z := big.NewInt(0), big.NewInt(0), big.NewInt(0)
	x.GCD(nil, nil, big.NewInt(collision[0]), big.NewInt(hailstones[0].Velocity[0]))
	y.GCD(nil, nil, big.NewInt(collision[0]), big.NewInt(hailstones[0].Velocity[1]))
	z.GCD(nil, nil, big.NewInt(collision[0]), big.NewInt(hailstones[0].Velocity[2]))

	fmt.Println(x.Add(x, y).Add(x, z).Int64())
}
