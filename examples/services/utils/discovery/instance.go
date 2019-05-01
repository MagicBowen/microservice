package discovery

type instance struct {
	address string
}

func newInstance(address string) *instance {
	return &instance{address: address}
}

func (it *instance) getKey() string {
	return it.address
}
