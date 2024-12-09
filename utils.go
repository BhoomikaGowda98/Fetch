package main

import (
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

func GenerateID() string {
	return uuid.New().String()
}

func CalculatePoints(receipt Receipt) int {
	points := 0

	// 1. Points for alphanumeric characters in retailer name
	points += len(regexp.MustCompile(`[a-zA-Z0-9]`).FindAllString(receipt.Retailer, -1))

	// 2. Points for round dollar total
	if strings.HasSuffix(receipt.Total, ".00") {
		points += 50
	}

	// 3. Points for total being a multiple of 0.25
	total, _ := strconv.ParseFloat(receipt.Total, 64)
	if math.Mod(total, 0.25) == 0 {
		points += 25
	}

	// 4. Points for every two items
	points += (len(receipt.Items) / 2) * 5

	// 5. Points for items with description length multiple of 3
	for _, item := range receipt.Items {
		desc := strings.TrimSpace(item.ShortDescription)
		if len(desc)%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(price * 0.2))
		}
	}

	// 6. Points if purchase date day is odd
	day, _ := strconv.Atoi(strings.Split(receipt.PurchaseDate, "-")[2])
	if day%2 != 0 {
		points += 6
	}

	// 7. Points if purchase time is between 2:00 PM and 4:00 PM
	t, _ := time.Parse("15:04", receipt.PurchaseTime)
	if t.Hour() == 14 {
		points += 10
	}

	return points
}
