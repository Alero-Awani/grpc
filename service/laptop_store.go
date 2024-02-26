package service

// LaptopStore is an interface to store laptop
type LaptopStore interface {
	// Save saves the laptop to the store
	Save (laptop *pb.Laptop) error
}


// InMemoryLaptopStore stores laptop in memory
type InMemoryLaptopStore struct {
	data map[string]*pb.Laptop
}
