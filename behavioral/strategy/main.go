package main

import "fmt"

type Point struct {
	X, Y int
}

// Общий интерфейс для всех способов построения маршрута.
// Navigator будет работать только с ним и не будет знать,
// какой конкретно алгоритм сейчас используется.
type RouteStrategy interface {
	BuildRoute(from, to Point) []Point
}

// Отдельная реализация алгоритма.
// Здесь могла бы находиться сложная логика поиска маршрута
// по дорогам
type CarStrategy struct{}

func (CarStrategy) BuildRoute(from, to Point) []Point {
	fmt.Println("Building route by car")
	return []Point{from, to}
}

// Еще один вариант стратегии - пешком.
type WalkingStrategy struct{}

func (WalkingStrategy) BuildRoute(from, to Point) []Point {
	fmt.Println("Building route by walking")
	return []Point{from, to}
}

// Еще один вариант стратегии - общественный транспорт.
// Чтобы добавить новый способ построения маршрута,
// нам не нужно изменять Navigator
type PublicTransportStrategy struct{}

func (PublicTransportStrategy) BuildRoute(from, to Point) []Point {
	fmt.Println("Building route by public transport")
	return []Point{from, to}
}

// Navigator - основной объект, который использует стратегию.
//
// Важно: здесь нет CarStrategy, WalkingStrategy etc.,
// Есть только интерфейс RouteStrategy.
type Navigator struct {
	strategy RouteStrategy
}

// Позволяет заменить алгоритм во время работы программы.
func (n *Navigator) SetStrategy(strategy RouteStrategy) {
	n.strategy = strategy
}

// Navigator сам маршрут не строит.
// Он просто передаёт эту работу текущей стратегии.
func (n *Navigator) BuildRoute(from, to Point) []Point {
	return n.strategy.BuildRoute(from, to)
}

func main() {
	from := Point{0, 0}
	to := Point{10, 10}

	// Выбираем конкретный алгоритм
	navigator := Navigator{
		strategy: CarStrategy{},
	}

	navigator.BuildRoute(from, to)

	// Меняем алгоритм - сам Navigator при этом не изменяется
	navigator.SetStrategy(WalkingStrategy{})
	navigator.BuildRoute(from, to)

	// И снова можем поменять стратегию
	navigator.SetStrategy(PublicTransportStrategy{})
	navigator.BuildRoute(from, to)
}
