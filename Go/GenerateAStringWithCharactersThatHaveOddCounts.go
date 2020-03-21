package Go

import "bytes"

func GenerateTheString(n int)  string {
	var ans bytes.Buffer
	if n % 2 == 0 {
		for i := 0; i < n-1; i++ {
			ans.WriteByte('a')
		}
		ans.WriteByte('b')
		return ans.String()
	} else {
		for i := 0; i < n; i++ {
			ans.WriteByte('a')
		}
		return ans.String()
	}
}
