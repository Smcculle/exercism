package allergies

var allergens = []string{"eggs", "peanuts", "shellfish", "strawberries",
	"tomatoes", "chocolate", "pollen", "cats"}

// Allergies reports the list of allergens
// BenchmarkAllergies-4  149 ns/op	      64 B/op	       3 allocs/op
func Allergies(score uint) (result []string) {

	for i := 0; score > 0 && i < len(allergens); i, score = i+1, score>>1 {
		if score&1 == 1 {
			result = append(result, allergens[i])
		}
	}
	return
}

// AllergicTo returns true if allergic to substance.
// BenchmarkAllergicTo-4   59.4 ns/op	       0 B/op	       0 allocs/op
func AllergicTo(score uint, substance string) bool {
	for i := 0; i < len(allergens); i++ {
		if allergens[i] == substance {
			var flag uint = 1 << uint(i)
			return (score & flag) > 0
		}
	}
	return false
}
