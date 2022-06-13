package generator

import "math/rand"

func ShortLinkGenerate(length int) string {
	var symbols = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	link := make([]rune, length)
	for i := range link {
		link[i] = symbols[rand.Intn(len(symbols))]
	}

	return string(link)
}
