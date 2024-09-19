package piscine

type food struct {
	chipspreptime   int
	burgerpreptime  int
	nuggetspreptime int
}

func setOrder(ord *food) {
	ord.burgerpreptime = 15
	ord.chipspreptime = 10
	ord.nuggetspreptime = 12
}

func FoodDeliveryTime(order string) int {
	orders := &food{}

	setOrder(orders)

	if order == "burger" {
		return orders.burgerpreptime
	} else if order == "chips" {
		return orders.chipspreptime
	} else if order == "nuggets" {
		return orders.nuggetspreptime
	} else {
		return 404
	}
}
