package sort

import (
	"log"
	"testing"
)

func TestPopupSortFunc(t *testing.T) {
	arr := []int{
		2, 354, 65, 2, 3, 26, 0, 687, 9, 4, 214, 35, 7,
	}

	log.Println(PopupSortFunc(arr))
}
