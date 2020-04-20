package minmeetingrooms

import "github.com/shomali11/go-interview/datastructures/sets/hashmultisets"

// MinMeetingRooms returns minimum meeting rooms required
func MinMeetingRooms(intervals [][]int) int {
	multiset := hashmultisets.New()
	for _, interval := range intervals {
		for i := interval[0]; i <= interval[1]; i++ {
			multiset.Add(i)
		}
	}

	multiSetPairs := multiset.GetTopValues()
	if len(multiSetPairs) == 0 {
		return 0
	}
	return multiSetPairs[0].Count
}
