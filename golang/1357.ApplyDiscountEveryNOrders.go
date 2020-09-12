type Cashier struct {
	n        int
	discount int
	price    map[int]int
	count    int
}

func Constructor(n int, discount int, products []int, prices []int) Cashier {
	var cashier Cashier
	cashier.n = n
	cashier.discount = discount
	cashier.price = make(map[int]int)
	for i, id := range products {
		cashier.price[id] = prices[i]
	}
	return cashier
}

func (this *Cashier) GetBill(product []int, amount []int) float64 {
	discount := float64(0)
	this.count++
	if this.count == this.n {
		discount = float64(this.discount) / 100
		this.count = 0
	}

	total := 0
	for i, id := range product {
		total += amount[i] * this.price[id]
	}

	return float64(total) * (float64(1) - discount)
}

