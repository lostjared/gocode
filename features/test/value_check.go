package value_check


func hasSpace(s string) bool {
	for i := range(s) {
		if(i == ' ' || i == '\t' || i =='\n') {
			return true
		}

	}
	return false
}