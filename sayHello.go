package gosayhello

func SayHello(name string) string {
	return "Hello " + name
}

func CalculateSum(data []int) (results int) {
	for _, v := range data {
		results += v
	}

	return
}
