package gosoln

func MinOperations(logs []string) int {
	awayFromRoot := 0
	for _, v := range logs {
		switch v {
		case "../":
			awayFromRoot--
			if awayFromRoot < 0 {
				awayFromRoot = 0
			}
		case "./":
			continue
		default:
			awayFromRoot++
		}
	}
	return awayFromRoot
}
