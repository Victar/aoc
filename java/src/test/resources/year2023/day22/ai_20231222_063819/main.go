
type Brick struct {
    x1, y1, z1 int // One endpoint of the brick
    x2, y2, z2 int // The other endpoint of the brick
    supportCount int // Number of bricks directly below
}
