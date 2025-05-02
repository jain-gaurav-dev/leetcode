package misc

import "container/list"

type hit struct {
	timestamp int
	count     int
}
type HitCounter struct {
	q         *list.List
	winSize   int
	totalHits int
}

func Constructor() HitCounter {
	return HitCounter{
		q:       list.New(),
		winSize: 300,
	}
}

func (this *HitCounter) Hit(timestamp int) {
	this.prune(timestamp)
	last := this.q.Back()
	if last != nil && last.Value.(hit).timestamp == timestamp {
		newHit := &hit{
			timestamp: timestamp,
			count:     last.Value.(hit).count + 1,
		}

		this.q.Remove(last)
		this.q.PushBack(newHit)
	} else {
		this.q.PushBack(hit{timestamp: timestamp, count: 1})
	}

	this.totalHits += 1
}

func (this *HitCounter) prune(timestamp int) {
	front := this.q.Front()
	oldestAllowedTs := timestamp - this.winSize + 1
	for front != nil && front.Value.(hit).timestamp < oldestAllowedTs {
		hitsToRemove := front.Value.(hit).count
		this.totalHits -= hitsToRemove

		this.q.Remove(front)
		front = this.q.Front()
	}
}

func (this *HitCounter) GetHits(timestamp int) int {
	this.prune(timestamp)

	return this.totalHits
}

/**
 * Your HitCounter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Hit(timestamp);
 * param_2 := obj.GetHits(timestamp);
 */
