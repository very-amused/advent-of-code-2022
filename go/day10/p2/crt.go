package p2

import "fmt"

type CRT struct {
	cpu    *CPU
	pixels []bool
}

func (c *CRT) draw() {
	c.pixels = make([]bool, 240)
	for cycles := 0; true; cycles++ {
		// Horizontal position of sprite
		sprite := c.cpu.Registers["X"]
		hPosition := (cycles % 40)
		if hPosition == sprite || hPosition == sprite-1 || hPosition == sprite+1 {
			c.pixels[cycles] = true
		}
		if done := c.cpu.cycle(); done {
			break
		}
	}
}

func (c CRT) render() {
	for i, pixel := range c.pixels {
		if i > 0 && i%40 == 0 {
			fmt.Println()
		}
		if pixel {
			fmt.Printf("#")
		} else {
			fmt.Printf(".")
		}
	}
	fmt.Println()
}
