package ch04

type Product struct {
	ID            int
	InStock       bool
	InSomeoneCart bool
}

func (p *Product) CanBeBought() bool {
	return p.InStock && !p.InSomeoneCart
}

type ShoppingCart struct {
	ID          int
	Products    []Product
	IsFull      bool
	MaxCartSize int
}

func (s *ShoppingCart) AddToCart(p Product) bool {
	if s.IsFull {
		return false
	}
	if p.CanBeBought() {
		s.Products = append(s.Products, p)
		return true
	}
	if s.MaxCartSize == len(s.Products) {
		s.IsFull = true
	}
	return true
}
