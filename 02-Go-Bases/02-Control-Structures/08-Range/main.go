package main

import "fmt"

func main() {
	// songs
	songs := []string{
		"Imagine", // 0
		"Hotel California",
		"Bohemian Rhapsody",
		"Stairway to Heaven",
		"Hotel California",
		"Bohemian Rhapsody",
		"Back to Black",
		"Imagine",
		"New York",
	}

	// for i := 0; i < len(songs); i++ {
	// 	println(songs[i])
	// }

	// for ix, song := range songs {
	// 	fmt.Println("Index:", ix, "Song:", song)
	// }

	// frequency is a map to store the frequency of each song
	// key: song name, value: frequency
	frequency := make(map[string]int)
	for _, song := range songs {
		// if _, ok := frequency[song]; ok {
		// 	frequency[song]++
		// } else {
		// 	frequency[song] = 1
		// }
		frequency[song] += 2
	}

	fmt.Println("frequency:", frequency)
}