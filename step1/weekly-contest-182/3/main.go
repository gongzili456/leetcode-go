package main

import (
	"fmt"
)

func main() {
	undergroundSystem := Constructor()
	undergroundSystem.CheckIn(45, "Leyton", 3)
	undergroundSystem.CheckIn(32, "Paradise", 8)
	undergroundSystem.CheckIn(27, "Leyton", 10)
	undergroundSystem.CheckOut(45, "Waterloo", 15)
	undergroundSystem.CheckOut(27, "Waterloo", 20)
	undergroundSystem.CheckOut(32, "Cambridge", 22)
	a1 := undergroundSystem.GetAverageTime("Paradise", "Cambridge") // 返回 14.0。从 "Paradise"（时刻 8）到 "Cambridge"(时刻 22)的行程只有一趟
	a2 := undergroundSystem.GetAverageTime("Leyton", "Waterloo")    // 返回 11.0。总共有 2 躺从 "Leyton" 到 "Waterloo" 的行程，编号为 id=45 的乘客出发于 time=3 到达于 time=15，编号为 id=27 的乘客于 time=10 出发于 time=20 到达。所以平均时间为 ( (15-3) + (20-10) ) / 2 = 11.0
	undergroundSystem.CheckIn(10, "Leyton", 24)
	a3 := undergroundSystem.GetAverageTime("Leyton", "Waterloo") // 返回 11.0
	undergroundSystem.CheckOut(10, "Waterloo", 38)
	a4 := undergroundSystem.GetAverageTime("Leyton", "Waterloo") // 返回 12.0

	fmt.Println("> ", a1, a2, a3, a4)
}

type state struct {
	begin string
	t     int
}

// UndergroundSystem is
type UndergroundSystem struct {
	Guests map[int]state
	Sum    map[string]int
	Count  map[string]int
}

// Constructor is
func Constructor() UndergroundSystem {
	return UndergroundSystem{
		Guests: map[int]state{},
		Sum:    map[string]int{},
		Count:  map[string]int{},
	}
}

// CheckIn is
func (u *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	u.Guests[id] = state{
		begin: stationName,
		t:     t,
	}
}

// CheckOut is
func (u *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	prev := u.Guests[id]
	delete(u.Guests, id)

	key := prev.begin + "_" + stationName
	dur := t - prev.t

	u.Sum[key] += dur
	u.Count[key]++
}

// GetAverageTime is
func (u *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	key := startStation + "_" + endStation

	return float64(u.Sum[key]) / float64(u.Count[key])
}
