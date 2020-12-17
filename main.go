package main

func main() {
	purchase := 3333_33
	percent := 1
	limit := 100
	bonus := (purchase / 100 * percent) / 100
	if bonus > limit {
		bonus = limit
	}
	println(bonus) // должно быть 33* (см. объяснение ниже)
}
