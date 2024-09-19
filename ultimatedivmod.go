package piscine

func UltimateDivMod(a *int, b *int) {
	quotient := *a / *b  // Calculate the quotient
	remainder := *a % *b // Calculate the remainder

	*a = quotient  // Store the quotient in the location pointed to by a
	*b = remainder // Store the remainder in the location pointed to by b
}
