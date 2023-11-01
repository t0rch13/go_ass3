package main

import (
	"fmt"
	"sync"
)

type Restaurant struct { // Singleton pattern
	Orders []string
}

var instance *Restaurant
var once sync.Once

func GetRestaurant() *Restaurant {
	once.Do(func() {
		instance = &Restaurant{
			Orders: make([]string, 0),
		}
	})
	return instance
}

type Command interface { // Command pattern
	Execute()
}

type OrderCommand struct {
	Dish     string
	Receiver *Chef
}

func (c *OrderCommand) Execute() {
	c.Receiver.CookDish(c.Dish)
}

type Chef struct {
	Name string
}

func (chef *Chef) CookDish(dish string) {
	fmt.Printf("%s is preparing %s\n", chef.Name, dish)
	restaurant := GetRestaurant()
	restaurant.Orders = append(restaurant.Orders, dish)
}

type Waiter struct {
	Commands []Command
}

func (w *Waiter) TakeOrder(command Command) {
	w.Commands = append(w.Commands, command)
}

func (w *Waiter) PlaceOrders() {
	for _, command := range w.Commands {
		command.Execute()
	}
}

func main() {
	restaurant := GetRestaurant() // Singleton
	fmt.Println("Restaurant Orders:", restaurant.Orders)

	chef := &Chef{Name: "Master Chef"} // Command pattern
	waiter := &Waiter{}

	waiter.TakeOrder(&OrderCommand{Dish: "Pasta", Receiver: chef})
	waiter.TakeOrder(&OrderCommand{Dish: "Steak", Receiver: chef})
	waiter.TakeOrder(&OrderCommand{Dish: "Beshbarmak", Receiver: chef})

	waiter.PlaceOrders()

	fmt.Println("Restaurant Orders after placing orders:", restaurant.Orders)
}
