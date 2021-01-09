package utils

func  GetListOfStringAsSet(strings []string) map[string]bool{
	set := make(map[string]bool)
	for _, s := range strings {
		set[s] = true
	}
	return set
}
