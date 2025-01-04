package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	if min > max {
		return min
	} else {
		return rand.Intn(max-min) + min
	}
}

func calculateWeightCart(cart []Bag) int {
	sum := 0
	for i := range cart {
		sum += cart[i].weight
	}
	return sum
}

type Bag struct {
	weight int
}

func NewBag(weight int) Bag {
	return Bag{weight}
}

const maxWarehouse = 250 // требование скалада(кг)
const cartСapacity = 4   // емкость тележки

func main() {
	wg := sync.WaitGroup{}

	storageOfBags := make(chan Bag)
	go func() {
		for {
			time.Sleep(5 * time.Second)
			bag := NewBag(random(20, 35))
			storageOfBags <- bag
			fmt.Println("мещок готов")
		}
	}()

	warehouse := 0
	cart := make([]Bag, 3)
	wg.Add(1)
	go func() {
		for {
			for i := 0; i < cartСapacity; i++ {
				select {
				case bag := <-storageOfBags:
					cart[i] = bag
					time.Sleep(100 * time.Millisecond)
					fmt.Println("мешок погружен")
				}
			}
			fmt.Println("тележка полна, везу на склад")
			weightCart := calculateWeightCart(cart)
			transportationTime := 0.2 * float64(weightCart)
			time.Sleep(time.Duration(transportationTime) * time.Second)

			warehouse += weightCart
			fmt.Println("мешки на складе")
			if maxWarehouse <= warehouse {
				fmt.Println("Склад заполнен гуд ворк")
				wg.Done()
			}
		}
	}()
	wg.Wait()
}
