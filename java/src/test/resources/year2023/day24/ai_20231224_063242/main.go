
package main

import (
    "bufio"
    "fmt"
    "math/big"
    "os"
    "regexp"
    "strconv"
    "strings"
)

type Vector struct {
    x, y, z *big.Int
}

var pattern = regexp.MustCompile(`^(-?\d+),(-?\d+),(-?\d+)\s@\s(-?\d+),(-?\d+),(-?\d+)$`)

func parseVector(s string) Vector {
    matches := pattern.FindStringSubmatch(s)
    x, _ := new(big.Int).SetString(matches[1], 10)
    y, _ := new(big.Int).SetString(matches[2], 10)
    z, _ := new(big.Int).SetString(matches[3], 10)
    vx, _ := new(big.Int).SetString(matches[4], 10)
    vy, _ := new(big.Int).SetString(matches[5], 10)
    vz, _ := new(big.Int).SetString(matches[6], 10)

    return Vector{
        x: x,
        y: y,
        z: z,
        vx: vx,
        vy: vy,
        vz: vz,
    }
}

func gcd(a, b *big.Int) *big.Int {
    bigZero := big.NewInt(0)
    if b.Cmp(bigZero) == 0 {
        return a
    }
    return gcd(b, new(big.Int).Mod(a, b))
}

func lcm(numbers []*big.Int) *big.Int {
    result := big.NewInt(1)
    bigOne := big.NewInt(1)

    for _, n := range numbers {
        g := gcd(result, n)
        result.Div(result, g).Mul(result, n)
    }

    if result.Cmp(bigOne) == -1 {
        return bigOne
    }
    return result
}

func abs(x *big.Int) *big.Int {
    bigZero := big.NewInt(0)
    if x.Cmp(bigZero) == -1 {
        return new(big.Int).Neg(x)
    }
    return new(big.Int).Set(x)
}

func getCollisionTime(hailstone1, hailstone2 Vector) *big.Int {
    // Calculate time by using x and vx, y and vy

    tx := new(big.Int).Sub(hailstone2.x, hailstone1.x)
    tvx := new(big.Int).Sub(hailstone1.vx, hailstone2.vx)

    ty := new(big.Int).Sub(hailstone2.y, hailstone1.y)
    tvy := new(big.Int).Sub(hailstone1.vy, hailstone2.vy)

    tz := new(big.Int).Sub(hailstone2.z, hailstone1.z)
    tvz := new(big.Int).Sub(hailstone1.vz, hailstone2.vz)

    // Times must be the same
    gcdX := gcd(tx, tvx)
    gcdY := gcd(ty, tvy)
    gcdZ := gcd(tz, tvz)

    tx.Div(tx, gcdX)
    tvx.Div(tvx, gcdX)

    ty.Div(ty, gcdY)
    tvy.Div(tvy, gcdY)

    tz.Div(tz, gcdZ)
    tvz.Div(tvz, gcdZ)

    if tx.Cmp(ty) == 0 && tx.Cmp(tz) == 0 && (tvx.Cmp(tvy) == 0) && (tvx.Cmp(tvz) == 0) {
        return abs(tx)
    }

    return nil
}

func main() {
    inputFile, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Error opening input file:", err)
        return
    }
    defer inputFile.Close()

    scanner := bufio.NewScanner(inputFile)
    var hailstones []Vector

    for scanner.Scan() {
        parts := strings.SplitN(scanner.Text(), " @ ", 2)
        posVel := strings.Split(parts[0]+" "+parts[1], ", ")
        pos := make([]int64, 3)
        vel := make([]int64, 3)

        for i := 0; i < 3; i++ {
            pos[i], _ = strconv.ParseInt(posVel[i], 10, 64)
            vel[i], _ = strconv.ParseInt(posVel[i+3], 10, 64)
        }

        hailstones = append(hailstones, Vector{
            x:  big.NewInt(pos[0]),
            y:  big.NewInt(pos[1]),
            z:  big.NewInt(pos[2]),
            vx: big.NewInt(vel[0]),
            vy: big.NewInt(vel[1]),
            vz: big.NewInt(vel[2]),
        })
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading input:", err)
        return
    }

    times := make([]*big.Int, 0)
    for i := 0; i < len(hailstones); i++ {
        for j := i + 1; j < len(hailstones); j++ {
            time := getCollisionTime(hailstones[i], hailstones[j])
            if time != nil {
                times = append(times, time)
            }
        }
    }

    collisionTime := lcm(times)
    fmt.Println(collisionTime)
}
