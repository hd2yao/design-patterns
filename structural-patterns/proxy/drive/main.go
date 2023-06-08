package main

func main() {
	car := NewCarProxy(&Driver{12})
	car.Drive()
	car2 := NewCarProxy(&Driver{22})
	car2.Drive()
}
