package server

func checkPlayedSlice(s [][] int) bool {
	if len(s) != 13 {
		return false
	} else {
		for _, i := range s {
			if len(i) != 5 {
				return false
			} else {
				for _, i2 := range i {
					if i2 > 6 || i2 < 1 {
						return false
					}
				}
			}
		}
		return true
	}
}
