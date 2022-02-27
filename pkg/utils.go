package shojo

func numberOfCharacters(origin string) (count int) {
	count = 0

	for range origin {
		count++
	}

	return count
}
