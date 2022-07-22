/*
Стратегия (англ. strategy) — поведенческий шаблон проектирования,
предназначенный для определения семейства алгоритмов,
инкапсуляции каждого из них и обеспечения их взаимозаменяемости.
Это позволяет выбирать алгоритм путём определения соответствующего класса.
Шаблон Strategy позволяет менять выбранный алгоритм независимо от объектов-клиентов, которые его используют.

плюсы - инкапсуляция реализации различных алгоритмов.
	  - система становится независимой от возможных изменений.
	  - вызов всех алгоритмов одним стандартным образом.
минусы - создаение доп классов.
*/

package main

import "fmt"

type Strategy interface {
	Route(startPoint, endPoint int)
}

type Navigator struct {
	Strategy
}

func (n *Navigator) SetStrategy(str Strategy) {
	n.Strategy = str
}

type RoadStrategy struct {
}

func (r *RoadStrategy) Route(startPoint, endPoint int) {
	avgSpeed := 30
	trafficJam := 2
	total := endPoint - startPoint
	totalTime := total * 40 * trafficJam

	fmt.Printf("Road A - %d to B - %d, average speed = %d, traffic jam - %d, total long - %d and total time = %d min\n",
		startPoint, endPoint, avgSpeed, trafficJam, total, totalTime)
}

type PublicTransportStrategy struct {
}

func (p *PublicTransportStrategy) Route(startPoint, endPoint int) {
	avgSpeed := 40
	total := endPoint - startPoint
	totalTime := total * 40

	fmt.Printf("Public transport A - %d to B - %d, average speed = %d, total long - %d and total time = %d min\n",
		startPoint, endPoint, avgSpeed, total, totalTime)
}

type WalkStrategy struct {
}

func (w *WalkStrategy) Route(startPoint, endPoint int) {
	avgSpeed := 4
	total := endPoint - startPoint
	totalTime := total * 60

	fmt.Printf("Walk A - %d to B - %d, average speed = %d, total long - %d and total time = %d min\n",
		startPoint, endPoint, avgSpeed, total, totalTime)
}

var (
	start      = 0
	end        = 1000
	strategies = []Strategy{
		&PublicTransportStrategy{},
		&WalkStrategy{},
		&RoadStrategy{},
	}
)

func main() {
	navigator := Navigator{}

	for _, strat := range strategies {
		navigator.SetStrategy(strat)
		navigator.Route(start, end)

	}
}
