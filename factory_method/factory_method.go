package factory_method

import "fmt"

type Product interface {
	use() string
}

type Creator interface {
	createProduct(owner string) Product
	registerProduct(product Product)
}

type Factory struct{}

func (*Factory) Create(creator Creator, owner string) Product {
	p := creator.createProduct(owner)
	creator.registerProduct(p)

	return p
}

type IDCard struct {
	owner string
}

func (c IDCard) use() string {
	return fmt.Sprintf("%sのカードを使います", c.owner)
}

func (c IDCard) getOwner() string {
	return c.owner
}

type IDCardFactory struct {
	*Factory
	owners []string
}

func (f *IDCardFactory) createProduct(owner string) Product {
	return IDCard{owner: owner}
}

func (f *IDCardFactory) registerProduct(product Product) {
	f.owners = append(f.owners, product.(IDCard).getOwner())
}
