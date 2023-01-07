package bw_95

func categorizeBox(length int, width int, height int, mass int) string {
	bulky := false
	heavy := false

	if length*width*height >= 1_000_000_000 {
		bulky = true
	}

	if length >= 10000 || width >= 10000 || height >= 10000 {
		bulky = true
	}

	if mass >= 100 {
		heavy = true
	}

	if bulky && heavy {
		return "Both"
	}

	if !bulky && !heavy {
		return "Neither"
	}

	if bulky && !heavy {
		return "Bulky"
	}

	if heavy && !bulky {
		return "Heavy"
	}

	return ""
}
