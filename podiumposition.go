package piscine

func PodiumPosition(podium [][]string) [][]string {
	for i := 0; i < len(podium); i++ {
		for j := i + 1; j < len(podium); j++ {
			if len(podium[i]) == len(podium[j]) {
				podium[i], podium[j] = podium[j], podium[i]
			}
		}
	}
	return podium
}

// func main() {
// 	p4 := []string{"4th Place"}
// 	p3 := []string{"3rd Place"}
// 	p2 := []string{"2nd Place"}
// 	p1 := []string{"1st Place"}

// 	position := [][]string{p4, p3, p2, p1}
// 	fmt.Println(PodiumPosition(position))
// }
