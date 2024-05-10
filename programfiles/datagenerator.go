package mathskills

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func DataGenerate(filename string, count int) error {
	// create a new instance of rand.Rand with a unique seed
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	// create data file
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
	}
	// generate random numbers

	for i := 0; i < count; i++ {
		num := (r.Intn(1000))
		_, err := file.WriteString(fmt.Sprintf("%d\n", num))
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("successfully created data.txt")
	return nil
}
