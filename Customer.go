package banking

type Customer struct {
	Name string
}

func NewCustomer(name string) *Customer {
	return &Customer{Name: name}
}

func (c *Customer) String() string {
	return c.Name
}
