package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

var allNumbers []string
var initialized bool

func initAllNumber() {
	if initialized {
		return
	}
	defer func() {
		initialized = true
	}()

	names := make([]string, 0, 26*26*1000)
	for c1 := 'A'; c1 <= 'Z'; c1++ {
		for c2 := 'A'; c2 <= 'Z'; c2++ {
			for n := 0; n <= 999; n++ {
				name := fmt.Sprintf("%c%c%03d", c1, c2, n)
				names = append(names, name)
			}
		}
	}

	r := rand.New(rand.NewSource(time.Now().UnixMicro()))
	r.Shuffle(len(names), func(i, j int) {
		names[i], names[j] = names[j], names[i]
	})

	allNumbers = names
}

// Define the Robot type here.
type Robot struct {
	name string
}

func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}

	if !initialized {
		initAllNumber()
	}

	if len(allNumbers) == 0 {
		return "", fmt.Errorf("no more robot names available")
	}
	r.name = allNumbers[0]
	allNumbers = allNumbers[1:]
	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = ""
}
