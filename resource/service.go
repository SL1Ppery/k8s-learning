package resource

import (
	"fmt"
	"sync"
	"time"
)

type service struct {
	Name      string
	NameSpace string
	Type      string
	Port      string
}

func CreateService(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i < 5; i++ {
		fmt.Printf("Creating Service%d\n", i)
		time.Sleep(time.Second)
	}
	fmt.Printf("Creating End")
}
