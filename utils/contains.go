
package utils

func Contains(s *[]string, e string) bool {
	for i := range *s {
		if (*s)[i] == e {
			return true
		}
	}
	return false
}