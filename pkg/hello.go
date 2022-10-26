package target


import (
)

// mockgen requirement +
type Math interface {
    Add(x, y int) int
    Subtract(x, y int) int
}
// mockgen requirement -

// interface can be call by test +
func GaoAdd(math Math, a, b int) int {
    v := math.Add(a, b)
    return v
}

func GaoSub(math Math, a, b int) int {
    v := math.Subtract(a, b)
    return v
}
// interface can be call by test -


func Add(x ,y int) int {
    return x + y
}

func Subtract(x, y int) int {
    return x - y
}
