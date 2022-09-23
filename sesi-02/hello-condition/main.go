package main

import "fmt"

func main() {
	var currentYear = 2021

	if age := currentYear - 1998; age < 17 {
		fmt.Println("Kamu belum boleh membuat kartu sim")
	} else {
		fmt.Println("Kamu sudah boleh membuat kartu sim")
	}

	var score = 8

	switch score {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	var dataScore = 6

	switch {
	case dataScore == 8:
		fmt.Println("perfect")
	case (dataScore < 8) && (dataScore > 3):
		fmt.Println("not bad")
	default:
		{
			fmt.Println("study harder")
			fmt.Println("you need to learn more")
		}
	}

	var score1 = 6

	switch {
	case score1 == 8:
		fmt.Println("perfect")
	case (score1 < 8) && (score1 > 3):
		fmt.Println("not bad")
		fallthrough
	case score1 < 5:
		fmt.Println("It is ok, but please study harder")
	default:
		{
			fmt.Println("study harder")
			fmt.Println("You don't have a good score yet")
		}
	}

	var score2 = 10

	if score2 > 7 {
		switch score2 {
		case 10:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice!")
		}
	} else {
		if score2 == 5 {
			fmt.Println("not bad")
		} else if score2 == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if score2 == 0 {
				fmt.Println("try harder!")
			}
		}
	}
}
