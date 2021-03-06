package table

import (
	"log"
	"math/rand"
	"time"
)

type Bottom struct {
	tilePool []int
}

func (b *Bottom) GetInfo() {

}

func (b *Bottom) RemainingCount() {

}

func (b *Bottom) xiPai() {
	var single []int
	for i := 1; i < 10; i++ {
		single = append(single, i)
	}
	for i := 11; i < 20; i++ {
		single = append(single, i)
	}
	for i := 21; i < 30; i++ {
		single = append(single, i)
	}
	for i := 31; i < 38; i++ {
		single = append(single, i)
	}
	tilePool := make([]int, len(single)*4)
	copy(tilePool, single)
	copy(tilePool[len(single):], single)
	copy(tilePool[len(single)*2:], single)
	copy(tilePool[len(single)*3:], single)

	rand.Seed(time.Now().UnixNano())
	for i := range tilePool {
		j := rand.Intn(i + 1)
		tilePool[i], tilePool[j] = tilePool[j], tilePool[i]

	}

	b.tilePool = tilePool

	log.Println("init tile pool...", b.tilePool)
}

func (b *Bottom) ZhiSaiZi() {

}

func (b *Bottom) PopKaiPai() []int {
	log.Println("Bottom.PopKaiPai...", b.tilePool)
	tiles := b.tilePool[:13]
	b.tilePool = b.tilePool[13:]
	return tiles
}

func (b *Bottom) PopMoPai() (int, bool) {
	log.Println("Bottom.PopMoPai...", b.tilePool)
	if len(b.tilePool) == 0 {
		return -1, false
	}

	tile := b.tilePool[0]
	b.tilePool = b.tilePool[1:]
	return tile, true
}
