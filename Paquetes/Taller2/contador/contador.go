package contador

func A(or string) int {
	cont := 0

	for _, char := range or {
		if char == 'a' || char == 'A' || char == 'á' || char == 'Á' || char == 'ä' || char == 'Ä' {
			cont++
		}
	}
	return cont
}

func E(or string) int {
	cont := 0

	for _, char := range or {
		if char == 'e' || char == 'E' || char == 'é' || char == 'É' || char == 'ë' || char == 'Ë' {
			cont++
		}
	}
	return cont
}

func I(or string) int {
	cont := 0

	for _, char := range or {
		if char == 'i' || char == 'I' || char == 'í' || char == 'Í' || char == 'ï' || char == 'Ï' {
			cont++
		}
	}
	return cont
}

func O(or string) int {
	cont := 0

	for _, char := range or {
		if char == 'o' || char == 'O' || char == 'ó' || char == 'Ó' || char == 'ö' || char == 'Ö' {
			cont++
		}
	}
	return cont
}

func U(or string) int {
	cont := 0

	for _, char := range or {
		if char == 'u' || char == 'U' || char == 'ú' || char == 'Ú' || char == 'ü' || char == 'Ü' {
			cont++
		}
	}
	return cont
}
