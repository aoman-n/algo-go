package slice

func Exists[T comparable](list []T, target T) bool {
	for _, v := range list {
		if v == target {
			return true
		}
	}
	return false
}

func Remove[T comparable](list *[]T, i int) {
	*list = (*list)[:i+copy((*list)[i:], (*list)[i+1:])]
}
