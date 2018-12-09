package main

const (
	COLOR_RED  = 1  
	COLOR_GREEN = 2
)


func getColor(color int) string{
	var scolor string
	switch color {
	case COLOR_GREEN:
		scolor = "Green";
		break;
	default:
		scolor = "Red";
		break;
	}

	return scolor
}