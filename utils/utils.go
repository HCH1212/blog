package utils

func InList(key string, list []string) bool {
	for _, v := range list {
		if v == key {
			return true
		}
	}
	return false
}