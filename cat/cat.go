package cat

type Cat struct {
    Name       string　// Public
    tailLength int // private
}

// getter
func (c *Cat) GetTailLength() int {
    return c.tailLength
}

// setter
func (c *Cat) SetTailLength(length int) {
    c.tailLength = length
}