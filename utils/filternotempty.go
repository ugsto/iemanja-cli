package utils

func FilterNotEmpty(array []string) []string {
	var filteredArray []string = []string{}
	for _, item := range array {
		if item != "" {
			filteredArray = append(filteredArray, item)
		}
	}
	return filteredArray
}
