package gosoln

type ProductOfNumbers struct {
	prefixProduct []int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{
		prefixProduct: []int{1},
	}
}

func (p *ProductOfNumbers) Add(num int) {
	if num > 0 {
		p.prefixProduct = append(p.prefixProduct, p.prefixProduct[len(p.prefixProduct)-1]*num)
	} else {
		p.prefixProduct = []int{1}
	}
}

func (p *ProductOfNumbers) GetProduct(k int) int {
	n := len(p.prefixProduct)
	if k >= n {
		return 0
	} else {
		return p.prefixProduct[n-1] / p.prefixProduct[n-k-1]
	}
}
