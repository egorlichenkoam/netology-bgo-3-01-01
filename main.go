package main

/**
Домашняя работа bgo-3-01-01

https://github.com/netology-code/bgo-homeworks/tree/master/01_std
*/
func main() {
	purchase := 3333_33
	percent := 1
	limit := 100
	bonus := (purchase / 100 * percent) / 100
	if bonus > limit {
		bonus = limit
	}
	println(bonus)
}
