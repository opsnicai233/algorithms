package dynamicPlanning

import (
	"log"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	log.Println(MaxProfit(prices))
}
