
package slice_operations

func Contains(ele string, sl []string) bool {
	for _, i := range sl {
		if i == ele {
			return true
		}
	}
	return false
}

// To find all the elements from a slice of s1 that are not present in slice s2
func Difference(s1 []string, s2 []string) (keys []string) {
	var unique map[string]bool = make(map[string]bool)
	for _, i := range s1 {
		if !Contains(i, s2) {
			unique[i] = true
		}
	}

	for key := range unique {
		keys = append(keys, key)
	}
	return
}

func Unique(sl []string) (keys []string) {
	var unique map[string]bool = make(map[string]bool)
	for _, ele := range sl {
		unique[ele] = true
	}
	for key := range unique {
		keys = append(keys, key)
	}
	return
}



