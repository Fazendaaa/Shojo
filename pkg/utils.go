package shojo

func numberOfCharacters(origin string) (count int) {
	count = 0

	for range origin {
		count++
	}

	return count
}

// unique based on this answer: https://www.golangprograms.com/remove-duplicate-values-from-slice.html
func unique(slice []interface{}) (list []interface{}) {
	keys := make(map[interface{}]bool)
	list = []interface{}{}

	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}

	return list
}

func packagesToInterface(packages []Package) (converted []interface{}) {
	converted = make([]interface{}, len(packages))

	for index, value := range packages {
		converted[index] = value
	}

	return converted
}
