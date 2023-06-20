package helpers

func ContainsPrimitive[K comparable](list []K, value K) bool {
	for _, j := range list {
		if j == value {
			return true
		}
	}
	return false
}
