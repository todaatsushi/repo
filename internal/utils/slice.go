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

func TagInItem(tag *string, tags *[]string) bool {
	for _, tag2 := range *tags {
		if *tag == tag2 {
			return true
		}
	}
	return false
}
