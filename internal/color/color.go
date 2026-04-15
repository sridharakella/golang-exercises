package color

type Color struct {
	Value int
}

var (
	Red     = Color{Value: 1}
	Green   = Color{Value: 2}
	Blue    = Color{Value: 3}
	Magenta = Color{Value: 4}
	Yellow  = Color{Value: 5}
)

func Text(color Color, text string) string {
	
	return "text"
}
