package main

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
    set map[int]int
	keys []int
}


func Constructor() RandomizedSet {
    return RandomizedSet{
		set: make(map[int]int),
		keys: make([]int,0),
	}
}


func (this *RandomizedSet) Insert(val int) bool {
    _, ok := this.set[val]

	if !ok {
		this.keys = append(this.keys, val)
		this.set[val] = len(this.keys)-1
		return true
	}

	return false
}


func (this *RandomizedSet) Remove(val int) bool {
    _, ok := this.set[val]

	if !ok {
		return ok
	}

	delete(this.set, val)

	index := this.set[val]
	lastIndex := len(this.keys) - 1
	this.keys[lastIndex], this.keys[index] = this.keys[index], this.keys[lastIndex]

	this.keys = this.keys[:lastIndex]

	return ok
}


func (this *RandomizedSet) GetRandom() int {
	return this.keys[rand.Intn(len(this.keys))]
}

func main(){ 
	randomizedSet := Constructor()
	fmt.Println(randomizedSet.Insert(1))
	fmt.Println(randomizedSet)
	fmt.Println(randomizedSet.Remove(2))
	fmt.Println(randomizedSet.Insert(2))
	fmt.Println(randomizedSet)
	fmt.Println(randomizedSet.GetRandom())
	fmt.Println(randomizedSet.Remove(1))
	fmt.Println(randomizedSet.Insert(2))
	fmt.Println(randomizedSet.GetRandom())
	fmt.Println(randomizedSet.Remove(2))
	fmt.Println(randomizedSet)

	randomizedSet2 := Constructor()
	fmt.Println(randomizedSet2.Remove(0))
	fmt.Println(randomizedSet2.Remove(0))
	fmt.Println(randomizedSet2.Insert(0))
	fmt.Println(randomizedSet2.GetRandom())
	fmt.Println(randomizedSet2.Remove(0))
	fmt.Println(randomizedSet)
	fmt.Println(randomizedSet2.Insert(0))
}
