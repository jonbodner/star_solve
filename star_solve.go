package main

import (
	"fmt"
)

type ValSet struct {
	one, two, three, four *int
}

func (vs ValSet) sum() int {
	return *vs.one + *vs.two + *vs.three + *vs.four
}

func (vs ValSet) String() string {
	return fmt.Sprintf("%d + %d + %d + %d = %d", *vs.one, *vs.two, *vs.three, *vs.four, vs.sum())
}

var memo = make(map[int][][]int)

func perm(n int) <-chan []int {
	c := make(chan []int)
	go func() {
		defer close(c)
		if n == 1 {
			c <- []int{1}
			return
		}
		cache := memo[n-1]
		for i := 0; i < n; i++ {
			if len(cache) > 0 {
				//fmt.Println("in cache: ", n)
				for _, result := range cache {
					out := make([]int, n)
					out[i] = n
					popOut(out, result, i)
					c <- out
				}
			} else {
				c1 := perm(n - 1)
				for result := range c1 {
					out := make([]int, n)
					out[i] = n
					popOut(out, result, i)
					c <- out
				}
			}
		}
	}()
	return c
}

func popOut(out []int, result []int, pos int) {
	for k := 0; k < pos; k++ {
		out[k] = result[k]
	}
	for k := pos; k < len(result); k++ {
		out[k+1] = result[k]
	}
}

func main() {
	fillCache(1, 10)
	fmt.Println("doing solve")
	starSolve()
}

func fillCache(start, end int) {
	for maxCache := start; maxCache < end; maxCache++ {
		//fmt.Println("caching ", maxCache)
		//put stuff in the cache
		c := perm(maxCache)
		cache := make([][]int, 0)
		for r := range c {
			cache = append(cache, r)
		}
		//fmt.Println(cache)
		memo[maxCache] = cache
	}
	//fmt.Println(memo)
}

func starSolve() {

	var a [12]int
	bcde := ValSet{&a[1], &a[2], &a[3], &a[4]}
	bfil := ValSet{&a[1], &a[5], &a[8], &a[11]}
	acfh := ValSet{&a[0], &a[2], &a[5], &a[7]}
	adgk := ValSet{&a[0], &a[3], &a[6], &a[10]}
	hijk := ValSet{&a[7], &a[8], &a[9], &a[10]}
	egjl := ValSet{&a[4], &a[6], &a[9], &a[11]}

	//run through every permutation of 12P12 (which is 12!)
	c := perm(12)
	count := 0
	numFound := 0
	for r := range c {
		count++
		for k, v := range r {
			a[k] = v
		}
		if (bcde.sum() == bfil.sum()) && (bcde.sum() == acfh.sum()) && (bcde.sum() == adgk.sum()) && (bcde.sum() == hijk.sum()) && (bcde.sum() == egjl.sum()) {
			numFound++
			fmt.Println("Done!")
			fmt.Println(bcde)
			fmt.Println(bfil)
			fmt.Println(acfh)
			fmt.Println(adgk)
			fmt.Println(hijk)
			fmt.Println(egjl)

			fmt.Println("Answers are ", r)
			fmt.Println("Number of tries: ", count)
			//break

		}
	}
	fmt.Println()
	fmt.Println("Num checked: ", count)
	fmt.Println("num found: ", numFound)
}
