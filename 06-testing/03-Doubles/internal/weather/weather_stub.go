package weather

func NewStub() *Stub {
	return &Stub{
		FuncGetTemperature: func(city string) (degrees float64, err error) {
			return
		},
	}
}

// stub.FuncGetTemperature = func(city string) (degrees float64, err error) {
// 	return 0, errors.New("something went wrong")
// }

type Stub struct {
	// // Temperature is the temperature that the stub will return.
	// Temperature float64
	// // Error is the error that the stub will return.
	// Error error

	// FuncGetTemperature is the function that will be called when GetTemperature is invoked.
	FuncGetTemperature func(city string) (degrees float64, err error)
}

func (s *Stub) GetTemperature(city string) (degrees float64, err error) {
	// static or concrete values
	// return 20, nil
	// return 0, errors.New("something went wrong")

	// dynamic or variable values (externalized to config params)
	// degrees = s.Temperature
	// err = s.Error
	// return

	degrees, err = s.FuncGetTemperature(city)
	return
}
