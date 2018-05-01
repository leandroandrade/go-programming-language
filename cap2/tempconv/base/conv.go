package tempconv

// CToF converte a temperatura em Celsius para Fahrenheit
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC converte a temperatura de Fahrenheit para Celsius
func FToc(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
