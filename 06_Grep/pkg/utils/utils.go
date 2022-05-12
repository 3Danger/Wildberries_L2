package utils

func XOR(a, b bool) bool {
	return (a && !b) || (!a && b)
}
