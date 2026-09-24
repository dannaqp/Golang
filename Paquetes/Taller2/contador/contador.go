package contador

func A(or string) int {
	cont := 0

	for _, char := range or {
		if char == 'a' || char == 'A' {
			cont++
		}
	}
	return cont
}

func E(or string) int {
	cont := 0

	for _, char := range or {
		if char == 'e' || char == 'E' {
			cont++
		}
	}
	return cont
}

func I(or string) int {
	cont := 0

	for _, char := range or {
		if char == 'i' || char == 'I' {
			cont++
		}
	}
	return cont
}

func O(or string) int {
	cont := 0

	for _, char := range or {
		if char == 'o' || char == 'O' {
			cont++
		}
	}
	return cont
}

func U(or string) int {
	cont := 0

	for _, char := range or {
		if char == 'u' || char == 'U' {
			cont++
		}
	}
	return cont
}
