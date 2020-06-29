package main
import "fmt"
func main() {
	var a, b, c, d, e, f int
	fmt.Scanf("%d %d %d\n", &a, &b, &c)
	fmt.Scanf("%d %d %d\n", &d, &e, &f)
	firLine := (a*100)+(b*10)+(c*1)
	secLine := (d*100)+(e*10)+(f*1)
	fmt.Printf("%d\n", firLine * f)
	fmt.Printf("%d\n", firLine * e)
	fmt.Printf("%d\n", firLine * d)
	fmt.Printf("%d\n", firLine * secLine)
}