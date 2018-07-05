package linked_list

import (
	"strconv"
	"strings"
	"testing"
)

func TestLinkedList_Iterate(t *testing.T) {
	raw := []int{1, 2, 5, 56, 7, 5, 8, 23, 7, 8, 3, 3, 5, 7, 87}

	ll := LinkedList{}
	for _, i := range raw {
		ll.Add(i)
	}

	var result []string
	for ll.Next() {
		result = append(result, strconv.Itoa(ll.Current()))
	}

	if strings.Join(result, " ") != intArrayToString(raw) {
		t.Errorf("не совпадает:\nожидается %v\n передано %v", intArrayToString(raw), result)
	}
}

func intArrayToString(intValues []int) string {
	var stringValues []string
	for _, v := range intValues {
		stringValues = append(stringValues, strconv.Itoa(v))
	}
	return strings.Join(stringValues, " ")
}
