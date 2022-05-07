package utils

func GetUniqueTags(tags []string) []string {
	items := make(map[string]bool)
	unique := []string{}
	for _, item := range tags {
		if _, val := items[item]; !val {
			items[item] = true
			unique = append(unique, item)
		}
	}
	return unique
}
