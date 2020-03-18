package main

import "fmt"

func similarRGB(color string) string {
	var rgb [3]int
	var s, t int
	var x, y, z int

	for i := 0; i < 3; i++ {
		m, n := color[1+2*i], color[2+2*i]
		if m >= 'a' {
			s = int(m-'a') + 10
		} else {
			s = int(m - '0')
		}
		if n >= 'a' {
			t = int(n-'a') + 10
		} else {
			t = int(n - '0')
		}

		if s == t {
			rgb[i] = s*16 + t
		} else {
			y = s*16 + t
			if s == 0 {
				x = 0
				z = 0x11
			} else {
				//r/g/b less than 0xXX
				x = (s-1)*16 + s - 1
				z = s*16 + s
			}
			//r/g/b bigger than 0xXX
			if y > z {
				x = z
				z = (s+1)*16 + s + 1
			}
			if y-x >= z-y {
				rgb[i] = int(z)
			} else {
				rgb[i] = int(x)
			}
		}
	}

	return fmt.Sprintf("#%02x%02x%02x", rgb[0], rgb[1], rgb[2])
}

func main() {
	fmt.Println(similarRGB("#09f166"))
	fmt.Println(similarRGB("#01f866"))
	fmt.Println(similarRGB("#71c986"))
	fmt.Println(similarRGB("#4e3fe1"))
}
