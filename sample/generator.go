package sample

import pb "github.com/aleroawani/grpc/pb/proto"

// NewKeyboard returns a new sample keyboard
func NewKeyboard() *pb.Keyboard {
	keyboard := &pb.Keyboard{
		Layout: randomKeyboardLayout(),
		Backlit: randomBool(),


	}
	return keyboard

}

// NewCPU returns a new sample CPU 

func NewCPU() *pb.CPU {
	brand := randomCPUBrand()
	name := randomCPUName(brand)

	numberCores := randomInt(2,8)
	numberThreads := randomInt(numberCores, 12)

	minGhz := randomFloat64(2.0, 3.5)
	maxGhz := randomFloat64(minGhz, 5.0)




	cpu := &pb.CPU{
		Brand: brand,
		Name: name,
		NumberCores: uint32(numberCores),
		NumberThreads: uint32(numberThreads),
		MinGhz: minGhz,
		MaxGhz: maxGhz,

	}

	return cpu
}

// NewGPU returns a new sample GPU
func NewGPU() *pb.GPU {
	brand := randomGPUBrand()
	name := randomGPUName(brand)

	minGhz := randomFloat64(1.0, 1.5)
	mazGhz := randomFloat64(minGhz, 2.0)
	
	gpu := &pb.GPU{
		Brand: brand,
		Name: name,
	}

	return gpu 
}