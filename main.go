package main

func SortIntegerTable(table []int) []int {
	n := len(table)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if table[j] > table[j+1] {
				table[j], table[j+1] = table[j+1], table[j]
			}
		}
	}
	return table
}
func main() {
	s := []int{9, 2, 3, 3, 1, 5, 5, 8, 7, 6, 7}
	SortIntegerTable(s)
}
