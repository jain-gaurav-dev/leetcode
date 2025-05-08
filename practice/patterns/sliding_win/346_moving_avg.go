package sliding_win

import "container/list"

type MovingAverage struct {
	sum        float64
	count      float64
	window     *list.List
	maxWinSize int
}

func Constructor(size int) MovingAverage {
	return MovingAverage{
		sum:        0.0,
		count:      0,
		window:     list.New(),
		maxWinSize: size,
	}
}

func (this *MovingAverage) Next(val int) float64 {
	this.sum += float64(val)
	this.window.PushBack(val)

	if this.window.Len() > this.maxWinSize {
		firstElem := this.window.Front()

		this.sum -= float64(firstElem.Value.(int))
		this.window.Remove(firstElem)
	}

	return this.sum / float64(this.window.Len())
}
