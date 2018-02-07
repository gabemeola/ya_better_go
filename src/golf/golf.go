package golf

func shotsDescription(par int, shots int) string {
	shotsOverPar := shots - par
	switch shotsOverPar {
	case 0:
		return "par"
	case -1:
		return "birdie"
	case -2:
		return "eagle"
	default:
		return "no description"
	}
	// YOUR CODE HERE
	// If the shotsOverPar variable is 1, return "bogey".
	// If it's 0, return "par".
	// For -1 return "birdie".
	// For -2 return "eagle".
	// For all other cases, return "no description".
}
