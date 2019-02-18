package space

func Age(seconds float64, planet string) float64 {
	switch planet {
	case "Earth":
		return seconds / 31557600
	case "Mercury":
		return seconds / 7600387.75
	case "Venus":
		return seconds / 19413750.40
	case "Mars":
		return seconds / 59352813.92
	case "Jupiter":
		return seconds / 374347972.1
	case "Saturn":
		return seconds / 929273280.9
	case "Uranus":
		return seconds / 2649555255.46
	case "Neptune":
		return seconds / 5196859067.52
	default:
		return 0.00
	}
}
