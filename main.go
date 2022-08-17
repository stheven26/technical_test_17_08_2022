package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	onlyOnce sync.Once
	dice     = []int{1, 2, 3, 4, 5, 6}
)

func rollDice() int {
	onlyOnce.Do(func() {
		rand.Seed(time.Now().UnixNano()) // only run once
	})

	return dice[rand.Intn(len(dice))]
}

func Dice(pemain, dadu int) {
	switch pemain {
	case 1:
		fmt.Println("Pemain Minimal 2")
	case 2:
		counterPemain1, counterPemain2 := 0, 0
		daduPemain1, daduPemain2 := []int{}, []int{}
		daduPemain1Lempar2, daduPemain2Lempar2 := []int{}, []int{}
		daduPemain1Lempar3, daduPemain2Lempar3 := []int{}, []int{}
		switch dadu {
		case 1:
			for i := 0; i < dadu; i++ {
				daduPemain1 = append(daduPemain1, rollDice())
				daduPemain2 = append(daduPemain2, rollDice())
			}

			fmt.Printf("Giliran %d lempar dadu:\n", 1)
			fmt.Println("Pemain 1:", daduPemain1)
			fmt.Println("Pemain 2:", daduPemain2)
			fmt.Println()

			i, j := 0, 0

			for i < len(daduPemain1) && j < len(daduPemain2) {
				if daduPemain1[i] == 1 {
					daduPemain2 = append(daduPemain2, 1)
					daduPemain1[i] = daduPemain1[len(daduPemain1)-1]
					daduPemain1[len(daduPemain1)-1] = 0
					daduPemain1 = daduPemain1[:len(daduPemain1)-1]
				}
				if daduPemain1[i] == 6 {
					counterPemain1++
					daduPemain1[i] = daduPemain1[len(daduPemain1)-1]
					daduPemain1[len(daduPemain1)-1] = 0
					daduPemain1 = daduPemain1[:len(daduPemain1)-1]
				}
				if daduPemain2[j] == 1 {
					daduPemain1 = append(daduPemain1, 1)
					daduPemain2[j] = daduPemain2[len(daduPemain2)-1]
					daduPemain2[len(daduPemain2)-1] = 0
					daduPemain2 = daduPemain2[:len(daduPemain2)-1]
				}
				if daduPemain2[j] == 6 {
					counterPemain2++
					daduPemain2[j] = daduPemain2[len(daduPemain2)-1]
					daduPemain2[len(daduPemain2)-1] = 0
					daduPemain2 = daduPemain2[:len(daduPemain2)-1]
				}
				i++
				j++
			}
			fmt.Println("Evaluasi")
			fmt.Printf("Pemain 1:(%d): %v\n", counterPemain1, daduPemain1)
			fmt.Printf("Pemain 2:(%d): %v\n", counterPemain2, daduPemain2)
			fmt.Println()
		case 2:
			for i := 0; i < dadu; i++ {
				daduPemain1 = append(daduPemain1, rollDice())
				daduPemain2 = append(daduPemain2, rollDice())
			}

			i, j, k, l := 0, 0, 0, 0

			fmt.Printf("Giliran %d lempar dadu:\n", 1)
			fmt.Println("Pemain 1:", daduPemain1)
			fmt.Println("Pemain 2:", daduPemain2)
			fmt.Println()
			for i < len(daduPemain1) && j < len(daduPemain2) {
				if daduPemain1[i] == 1 {
					daduPemain2 = append(daduPemain2, 1)
					daduPemain1[i] = daduPemain1[len(daduPemain1)-1]
					daduPemain1[len(daduPemain1)-1] = 0
					daduPemain1 = daduPemain1[:len(daduPemain1)-1]
				}
				if daduPemain1[i] == 6 {
					counterPemain1++
					daduPemain1[i] = daduPemain1[len(daduPemain1)-1]
					daduPemain1[len(daduPemain1)-1] = 0
					daduPemain1 = daduPemain1[:len(daduPemain1)-1]
				}
				if daduPemain2[j] == 1 {
					daduPemain1 = append(daduPemain1, 1)
					daduPemain2[j] = daduPemain2[len(daduPemain2)-1]
					daduPemain2[len(daduPemain2)-1] = 0
					daduPemain2 = daduPemain2[:len(daduPemain2)-1]
				}
				if daduPemain2[j] == 6 {
					counterPemain2++
					daduPemain2[j] = daduPemain2[len(daduPemain2)-1]
					daduPemain2[len(daduPemain2)-1] = 0
					daduPemain2 = daduPemain2[:len(daduPemain2)-1]
				}
				i++
				j++
			}
			fmt.Println("Evaluasi")
			fmt.Printf("Pemain 1:(%d): %v\n", counterPemain1, daduPemain1)
			fmt.Printf("Pemain 2:(%d): %v\n", counterPemain2, daduPemain2)
			fmt.Println()
			for i := range daduPemain1 {
				daduPemain1Lempar2 = append(daduPemain1Lempar2, rollDice())
				i++
			}

			for i := range daduPemain2 {
				daduPemain2Lempar2 = append(daduPemain2Lempar2, rollDice())
				i++
			}

			fmt.Printf("Giliran %d lempar dadu:\n", 2)
			fmt.Println("Pemain 1:", daduPemain1Lempar2)
			fmt.Println("Pemain 2:", daduPemain2Lempar2)
			fmt.Println()

			for k < len(daduPemain1Lempar2) && l < len(daduPemain2Lempar2) {
				if daduPemain1Lempar2[k] == 1 {
					daduPemain2Lempar2 = append(daduPemain2Lempar2, 1)
					daduPemain1Lempar2[k] = daduPemain1Lempar2[len(daduPemain1Lempar2)-1]
					daduPemain1Lempar2[len(daduPemain1Lempar2)-1] = 0
					daduPemain1Lempar2 = daduPemain1Lempar2[:len(daduPemain1Lempar2)-1]
				}
				if daduPemain1Lempar2[k] == 6 {
					counterPemain1++
					daduPemain1Lempar2[k] = daduPemain1Lempar2[len(daduPemain1Lempar2)-1]
					daduPemain1Lempar2[len(daduPemain1Lempar2)-1] = 0
					daduPemain1Lempar2 = daduPemain1Lempar2[:len(daduPemain1Lempar2)-1]
				}
				if daduPemain2Lempar2[l] == 1 {
					daduPemain1Lempar2 = append(daduPemain1Lempar2, 1)
					daduPemain2Lempar2[l] = daduPemain2Lempar2[len(daduPemain2Lempar2)-1]
					daduPemain2Lempar2[len(daduPemain2Lempar2)-1] = 0
					daduPemain2Lempar2 = daduPemain2Lempar2[:len(daduPemain2Lempar2)-1]
				}
				if daduPemain2Lempar2[l] == 6 {
					counterPemain2++
					daduPemain2Lempar2[l] = daduPemain2Lempar2[len(daduPemain2Lempar2)-1]
					daduPemain2Lempar2[len(daduPemain2Lempar2)-1] = 0
					daduPemain2Lempar2 = daduPemain2Lempar2[:len(daduPemain2Lempar2)-1]
				}
				k++
				l++
			}
			fmt.Println("Evaluasi")
			fmt.Printf("Pemain 1:(%d): %v\n", counterPemain1, daduPemain1Lempar2)
			fmt.Printf("Pemain 2:(%d): %v\n", counterPemain2, daduPemain2Lempar2)
			fmt.Println()
		case 3:
			for i := 0; i < dadu; i++ {
				daduPemain1 = append(daduPemain1, rollDice())
				daduPemain2 = append(daduPemain2, rollDice())
			}

			i, j, k, l, m, n := 0, 0, 0, 0, 0, 0

			fmt.Printf("Giliran %d lempar dadu:\n", 1)
			fmt.Println("Pemain 1:", daduPemain1)
			fmt.Println("Pemain 2:", daduPemain2)
			fmt.Println()
			for i < len(daduPemain1) && j < len(daduPemain2) {
				if daduPemain1[i] == 1 {
					daduPemain2 = append(daduPemain2, 1)
					daduPemain1[i] = daduPemain1[len(daduPemain1)-1]
					daduPemain1[len(daduPemain1)-1] = 0
					daduPemain1 = daduPemain1[:len(daduPemain1)-1]
				}
				if daduPemain1[i] == 6 {
					counterPemain1++
					daduPemain1[i] = daduPemain1[len(daduPemain1)-1]
					daduPemain1[len(daduPemain1)-1] = 0
					daduPemain1 = daduPemain1[:len(daduPemain1)-1]
				}
				if daduPemain2[j] == 1 {
					daduPemain1 = append(daduPemain1, 1)
					daduPemain2[j] = daduPemain2[len(daduPemain2)-1]
					daduPemain2[len(daduPemain2)-1] = 0
					daduPemain2 = daduPemain2[:len(daduPemain2)-1]
				}
				if daduPemain2[j] == 6 {
					counterPemain2++
					daduPemain2[j] = daduPemain2[len(daduPemain2)-1]
					daduPemain2[len(daduPemain2)-1] = 0
					daduPemain2 = daduPemain2[:len(daduPemain2)-1]
				}
				i++
				j++
			}
			fmt.Println("Evaluasi")
			fmt.Printf("Pemain 1:(%d): %v\n", counterPemain1, daduPemain1)
			fmt.Printf("Pemain 2:(%d): %v\n", counterPemain2, daduPemain2)
			fmt.Println()
			for i := range daduPemain1 {
				daduPemain1Lempar2 = append(daduPemain1Lempar2, rollDice())
				i++
			}

			for i := range daduPemain2 {
				daduPemain2Lempar2 = append(daduPemain2Lempar2, rollDice())
				i++
			}

			fmt.Printf("Giliran %d lempar dadu:\n", 2)
			fmt.Println("Pemain 1:", daduPemain1Lempar2)
			fmt.Println("Pemain 2:", daduPemain2Lempar2)
			fmt.Println()

			for k < len(daduPemain1Lempar2) && l < len(daduPemain2Lempar2) {
				if daduPemain1Lempar2[k] == 1 {
					daduPemain2Lempar2 = append(daduPemain2Lempar2, 1)
					daduPemain1Lempar2[k] = daduPemain1Lempar2[len(daduPemain1Lempar2)-1]
					daduPemain1Lempar2[len(daduPemain1Lempar2)-1] = 0
					daduPemain1Lempar2 = daduPemain1Lempar2[:len(daduPemain1Lempar2)-1]
				}
				if daduPemain1Lempar2[k] == 6 {
					counterPemain1++
					daduPemain1Lempar2[k] = daduPemain1Lempar2[len(daduPemain1Lempar2)-1]
					daduPemain1Lempar2[len(daduPemain1Lempar2)-1] = 0
					daduPemain1Lempar2 = daduPemain1Lempar2[:len(daduPemain1Lempar2)-1]
				}
				if daduPemain2Lempar2[l] == 1 {
					daduPemain1Lempar2 = append(daduPemain1Lempar2, 1)
					daduPemain2Lempar2[l] = daduPemain2Lempar2[len(daduPemain2Lempar2)-1]
					daduPemain2Lempar2[len(daduPemain2Lempar2)-1] = 0
					daduPemain2Lempar2 = daduPemain2Lempar2[:len(daduPemain2Lempar2)-1]
				}
				if daduPemain2Lempar2[l] == 6 {
					counterPemain2++
					daduPemain2Lempar2[l] = daduPemain2Lempar2[len(daduPemain2Lempar2)-1]
					daduPemain2Lempar2[len(daduPemain2Lempar2)-1] = 0
					daduPemain2Lempar2 = daduPemain2Lempar2[:len(daduPemain2Lempar2)-1]
				}
				k++
				l++
			}
			fmt.Println("Evaluasi")
			fmt.Printf("Pemain 1:(%d): %v\n", counterPemain1, daduPemain1Lempar2)
			fmt.Printf("Pemain 2:(%d): %v\n", counterPemain2, daduPemain2Lempar2)
			fmt.Println()

			for i := range daduPemain1Lempar2 {
				daduPemain1Lempar3 = append(daduPemain1Lempar3, rollDice())
				i++
			}

			for i := range daduPemain2Lempar2 {
				daduPemain2Lempar3 = append(daduPemain2Lempar3, rollDice())
				i++
			}

			fmt.Printf("Giliran %d lempar dadu:\n", 3)
			fmt.Println("Pemain 1:", daduPemain1Lempar3)
			fmt.Println("Pemain 2:", daduPemain2Lempar3)
			fmt.Println()

			for m < len(daduPemain1Lempar3) && n < len(daduPemain2Lempar3) {
				if daduPemain1Lempar3[m] == 1 {
					daduPemain2Lempar3 = append(daduPemain2Lempar3, 1)
					daduPemain1Lempar3[m] = daduPemain1Lempar3[len(daduPemain1Lempar3)-1]
					daduPemain1Lempar3[len(daduPemain1Lempar3)-1] = 0
					daduPemain1Lempar3 = daduPemain1Lempar3[:len(daduPemain1Lempar3)-1]
				}
				if daduPemain1Lempar3[m] == 6 {
					counterPemain1++
					daduPemain2Lempar3 = append(daduPemain2Lempar3, 1)
					daduPemain1Lempar3[m] = daduPemain1Lempar3[len(daduPemain1Lempar3)-1]
					daduPemain1Lempar3[len(daduPemain1Lempar3)-1] = 0
					daduPemain1Lempar3 = daduPemain1Lempar3[:len(daduPemain1Lempar3)-1]
				}
				if daduPemain2Lempar3[n] == 1 {
					daduPemain1Lempar3 = append(daduPemain1Lempar3, 1)
					daduPemain2Lempar3[n] = daduPemain2Lempar3[len(daduPemain2Lempar3)-1]
					daduPemain2Lempar3[len(daduPemain2Lempar3)-1] = 0
					daduPemain2Lempar3 = daduPemain2Lempar3[:len(daduPemain2Lempar3)-1]
				}
				if daduPemain2Lempar3[n] == 6 {
					counterPemain2++
					daduPemain2Lempar3[n] = daduPemain2Lempar3[len(daduPemain2Lempar3)-1]
					daduPemain2Lempar3[len(daduPemain2Lempar3)-1] = 0
					daduPemain2Lempar3 = daduPemain2Lempar3[:len(daduPemain2Lempar3)-1]
				}
				m++
				n++
			}
			fmt.Println("Evaluasi")
			fmt.Printf("Pemain 1:(%d): %v\n", counterPemain1, daduPemain1Lempar3)
			fmt.Printf("Pemain 2:(%d): %v\n", counterPemain2, daduPemain2Lempar3)
			fmt.Println()
		}
	case 3:
		counterPemain1, counterPemain2, counterPemain3 := 0, 0, 0
		daduPemain1, daduPemain2, daduPemain3 := []int{}, []int{}, []int{}
		daduPemain1Lempar2, daduPemain2Lempar2, daduPemain3Lempar2 := []int{}, []int{}, []int{}
		daduPemain1Lempar3, daduPemain2Lempar3, daduPemain3Lempar3 := []int{}, []int{}, []int{}
		switch dadu {
		case 1:
			for i := 0; i < dadu; i++ {
				daduPemain1 = append(daduPemain1, rollDice())
				daduPemain2 = append(daduPemain2, rollDice())
				daduPemain3 = append(daduPemain3, rollDice())
			}

			fmt.Printf("Giliran %d lempar dadu:\n", 1)
			fmt.Println("Pemain 1:", daduPemain1)
			fmt.Println("Pemain 2:", daduPemain2)
			fmt.Println("Pemain 3:", daduPemain3)
			fmt.Println()

			i, j, k := 0, 0, 0

			for i < len(daduPemain1) && j < len(daduPemain2) && k < len(daduPemain3) {
				if daduPemain1[i] == 1 {
					daduPemain2 = append(daduPemain2, 1)
					daduPemain1[i] = daduPemain1[len(daduPemain1)-1]
					daduPemain1[len(daduPemain1)-1] = 0
					daduPemain1 = daduPemain1[:len(daduPemain1)-1]
				}
				if daduPemain1[i] == 6 {
					counterPemain1++
					daduPemain1[i] = daduPemain1[len(daduPemain1)-1]
					daduPemain1[len(daduPemain1)-1] = 0
					daduPemain1 = daduPemain1[:len(daduPemain1)-1]
				}
				if daduPemain2[j] == 1 {
					daduPemain3 = append(daduPemain3, 1)
					daduPemain2[j] = daduPemain2[len(daduPemain2)-1]
					daduPemain2[len(daduPemain2)-1] = 0
					daduPemain2 = daduPemain2[:len(daduPemain2)-1]
				}
				if daduPemain2[j] == 6 {
					counterPemain2++
					daduPemain2[j] = daduPemain2[len(daduPemain2)-1]
					daduPemain2[len(daduPemain2)-1] = 0
					daduPemain2 = daduPemain2[:len(daduPemain2)-1]
				}
				if daduPemain3[k] == 1 {
					daduPemain1 = append(daduPemain1, 1)
					daduPemain3[k] = daduPemain3[len(daduPemain3)-1]
					daduPemain3[len(daduPemain3)-1] = 0
					daduPemain3 = daduPemain3[:len(daduPemain3)-1]
					if daduPemain3[k] == 6 {
						counterPemain3++
						daduPemain3[k] = daduPemain3[len(daduPemain3)-1]
						daduPemain3[len(daduPemain3)-1] = 0
						daduPemain3 = daduPemain3[:len(daduPemain3)-1]
					}
					i++
					j++
					k++
				}
				fmt.Println("Evaluasi")
				fmt.Printf("Pemain 1:(%d): %v\n", counterPemain1, daduPemain1)
				fmt.Printf("Pemain 2:(%d): %v\n", counterPemain2, daduPemain2)
				fmt.Printf("Pemain 3:(%d): %v\n", counterPemain3, daduPemain3)
				fmt.Println()
			}
		case 2:
			for i := 0; i < dadu; i++ {
				daduPemain1 = append(daduPemain1, rollDice())
				daduPemain2 = append(daduPemain2, rollDice())
				daduPemain3 = append(daduPemain3, rollDice())
			}

			fmt.Printf("Giliran %d lempar dadu:\n", 1)
			fmt.Println("Pemain 1:", daduPemain1)
			fmt.Println("Pemain 2:", daduPemain2)
			fmt.Println("Pemain 3:", daduPemain3)
			fmt.Println()

			i, j, k := 0, 0, 0

			for i < len(daduPemain1) && j < len(daduPemain2) && k < len(daduPemain3) {
				if daduPemain1[i] == 1 {
					daduPemain2 = append(daduPemain2, 1)
					daduPemain1[i] = daduPemain1[len(daduPemain1)-1]
					daduPemain1[len(daduPemain1)-1] = 0
					daduPemain1 = daduPemain1[:len(daduPemain1)-1]
				}
				if daduPemain1[i] == 6 {
					counterPemain1++
					daduPemain1[i] = daduPemain1[len(daduPemain1)-1]
					daduPemain1[len(daduPemain1)-1] = 0
					daduPemain1 = daduPemain1[:len(daduPemain1)-1]
				}
				if daduPemain2[j] == 1 {
					daduPemain3 = append(daduPemain3, 1)
					daduPemain2[j] = daduPemain2[len(daduPemain2)-1]
					daduPemain2[len(daduPemain2)-1] = 0
					daduPemain2 = daduPemain2[:len(daduPemain2)-1]
				}
				if daduPemain2[j] == 6 {
					counterPemain2++
					daduPemain2[j] = daduPemain2[len(daduPemain2)-1]
					daduPemain2[len(daduPemain2)-1] = 0
					daduPemain2 = daduPemain2[:len(daduPemain2)-1]
				}
				if daduPemain3[k] == 1 {
					daduPemain1 = append(daduPemain1, 1)
					daduPemain3[k] = daduPemain3[len(daduPemain3)-1]
					daduPemain3[len(daduPemain3)-1] = 0
					daduPemain3 = daduPemain3[:len(daduPemain3)-1]
				}
				if daduPemain3[k] == 6 {
					counterPemain3++
					daduPemain3[k] = daduPemain3[len(daduPemain3)-1]
					daduPemain3[len(daduPemain3)-1] = 0
					daduPemain3 = daduPemain3[:len(daduPemain3)-1]
				}
				i++
				j++
				k++
				fmt.Println("Evaluasi")
				fmt.Printf("Pemain 1:(%d): %v\n", counterPemain1, daduPemain1)
				fmt.Printf("Pemain 2:(%d): %v\n", counterPemain2, daduPemain2)
				fmt.Printf("Pemain 3:(%d): %v\n", counterPemain3, daduPemain3)
				fmt.Println()
			}
			l, m, n := 0, 0, 0

			for i := range daduPemain1 {
				daduPemain1Lempar2 = append(daduPemain1Lempar2, rollDice())
				i++
			}
			for i := range daduPemain2 {
				daduPemain2Lempar2 = append(daduPemain2Lempar2, rollDice())
				i++
			}
			for i := range daduPemain3 {
				daduPemain3Lempar2 = append(daduPemain3Lempar2, rollDice())
				i++
			}

			fmt.Printf("Giliran %d lempar dadu:\n", 2)
			fmt.Println("Pemain 1:", daduPemain1Lempar2)
			fmt.Println("Pemain 2:", daduPemain2Lempar2)
			fmt.Println("Pemain 3:", daduPemain3Lempar2)
			fmt.Println()

			for l < len(daduPemain1Lempar2) && m < len(daduPemain2Lempar2) && n < len(daduPemain3Lempar2) {
				if daduPemain1Lempar2[l] == 1 {
					daduPemain2Lempar2 = append(daduPemain2Lempar2, 1)
					daduPemain1Lempar2[l] = daduPemain1Lempar2[len(daduPemain1Lempar2)-1]
					daduPemain1Lempar2[len(daduPemain1Lempar2)-1] = 0
					daduPemain1Lempar2 = daduPemain1Lempar2[:len(daduPemain1Lempar2)-1]
				}
				if daduPemain1Lempar2[l] == 6 {
					counterPemain1++
					daduPemain1Lempar2[l] = daduPemain1Lempar2[len(daduPemain1Lempar2)-1]
					daduPemain1Lempar2[len(daduPemain1Lempar2)-1] = 0
					daduPemain1Lempar2 = daduPemain1Lempar2[:len(daduPemain1Lempar2)-1]
				}
				if daduPemain2Lempar2[m] == 1 {
					daduPemain3Lempar2 = append(daduPemain3Lempar2, 1)
					daduPemain2Lempar2[m] = daduPemain2Lempar2[len(daduPemain2Lempar2)-1]
					daduPemain2Lempar2[len(daduPemain2Lempar2)-1] = 0
					daduPemain2Lempar2 = daduPemain2Lempar2[:len(daduPemain2Lempar2)-1]
				}
				if daduPemain2Lempar2[m] == 6 {
					counterPemain2++
					daduPemain2Lempar2[m] = daduPemain2Lempar2[len(daduPemain2Lempar2)-1]
					daduPemain2Lempar2[len(daduPemain2Lempar2)-1] = 0
					daduPemain2Lempar2 = daduPemain2Lempar2[:len(daduPemain2Lempar2)-1]
				}
				if daduPemain3Lempar2[n] == 1 {
					daduPemain1Lempar2 = append(daduPemain1Lempar2, 1)
					daduPemain3Lempar2[n] = daduPemain3Lempar2[len(daduPemain3Lempar2)-1]
					daduPemain3Lempar2[len(daduPemain3Lempar2)-1] = 0
					daduPemain3Lempar2 = daduPemain3Lempar2[:len(daduPemain3Lempar2)-1]
				}
				if daduPemain3Lempar2[n] == 6 {
					counterPemain3++
					daduPemain3Lempar2[n] = daduPemain3Lempar2[len(daduPemain3Lempar2)-1]
					daduPemain3Lempar2[len(daduPemain3Lempar2)-1] = 0
					daduPemain3Lempar2 = daduPemain3Lempar2[:len(daduPemain3Lempar2)-1]
				}
				l++
				m++
				n++
			}
			fmt.Println("Evaluasi")
			fmt.Printf("Pemain 1:(%d): %v\n", counterPemain1, daduPemain1Lempar2)
			fmt.Printf("Pemain 2:(%d): %v\n", counterPemain2, daduPemain2Lempar2)
			fmt.Printf("Pemain 3:(%d): %v\n", counterPemain3, daduPemain3Lempar2)
			fmt.Println()
		case 3:
			for i := 0; i < dadu; i++ {
				daduPemain1 = append(daduPemain1, rollDice())
				daduPemain2 = append(daduPemain2, rollDice())
				daduPemain3 = append(daduPemain3, rollDice())
			}

			fmt.Printf("Giliran %d lempar dadu:\n", 1)
			fmt.Println("Pemain 1:", daduPemain1)
			fmt.Println("Pemain 2:", daduPemain2)
			fmt.Println("Pemain 3:", daduPemain3)
			fmt.Println()

			i, j, k := 0, 0, 0

			for i < len(daduPemain1) && j < len(daduPemain2) && k < len(daduPemain3) {
				if daduPemain1[i] == 1 {
					daduPemain2 = append(daduPemain2, 1)
					daduPemain1[i] = daduPemain1[len(daduPemain1)-1]
					daduPemain1[len(daduPemain1)-1] = 0
					daduPemain1 = daduPemain1[:len(daduPemain1)-1]
				}
				if daduPemain1[i] == 6 {
					counterPemain1++
					daduPemain1[i] = daduPemain1[len(daduPemain1)-1]
					daduPemain1[len(daduPemain1)-1] = 0
					daduPemain1 = daduPemain1[:len(daduPemain1)-1]
				}
				if daduPemain2[j] == 1 {
					daduPemain3 = append(daduPemain3, 1)
					daduPemain2[j] = daduPemain2[len(daduPemain2)-1]
					daduPemain2[len(daduPemain2)-1] = 0
					daduPemain2 = daduPemain2[:len(daduPemain2)-1]
				}
				if daduPemain2[j] == 6 {
					counterPemain2++
					daduPemain2[j] = daduPemain2[len(daduPemain2)-1]
					daduPemain2[len(daduPemain2)-1] = 0
					daduPemain2 = daduPemain2[:len(daduPemain2)-1]
				}
				if daduPemain3[k] == 1 {
					daduPemain1 = append(daduPemain1, 1)
					daduPemain3[k] = daduPemain3[len(daduPemain3)-1]
					daduPemain3[len(daduPemain3)-1] = 0
					daduPemain3 = daduPemain3[:len(daduPemain3)-1]
				}
				if daduPemain3[k] == 6 {
					counterPemain3++
					daduPemain3[k] = daduPemain3[len(daduPemain3)-1]
					daduPemain3[len(daduPemain3)-1] = 0
					daduPemain3 = daduPemain3[:len(daduPemain3)-1]
				}
				i++
				j++
				k++
			}
			fmt.Println("Evaluasi")
			fmt.Printf("Pemain 1:(%d): %v\n", counterPemain1, daduPemain1)
			fmt.Printf("Pemain 2:(%d): %v\n", counterPemain2, daduPemain2)
			fmt.Printf("Pemain 3:(%d): %v\n", counterPemain3, daduPemain3)
			fmt.Println()
			l, m, n := 0, 0, 0

			for i := range daduPemain1 {
				daduPemain1Lempar2 = append(daduPemain1Lempar2, rollDice())
				i++
			}
			for i := range daduPemain2 {
				daduPemain2Lempar2 = append(daduPemain2Lempar2, rollDice())
				i++
			}
			for i := range daduPemain3 {
				daduPemain3Lempar2 = append(daduPemain3Lempar2, rollDice())
				i++
			}

			fmt.Printf("Giliran %d lempar dadu:\n", 2)
			fmt.Println("Pemain 1:", daduPemain1Lempar2)
			fmt.Println("Pemain 2:", daduPemain2Lempar2)
			fmt.Println("Pemain 3:", daduPemain3Lempar2)
			fmt.Println()

			for l < len(daduPemain1Lempar2) && m < len(daduPemain2Lempar2) && n < len(daduPemain3Lempar2) {
				if daduPemain1Lempar2[l] == 1 {
					daduPemain2Lempar2 = append(daduPemain2Lempar2, 1)
					daduPemain1Lempar2[l] = daduPemain1Lempar2[len(daduPemain1Lempar2)-1]
					daduPemain1Lempar2[len(daduPemain1Lempar2)-1] = 0
					daduPemain1Lempar2 = daduPemain1Lempar2[:len(daduPemain1Lempar2)-1]
				}
				if daduPemain1Lempar2[l] == 6 {
					counterPemain1++
					daduPemain1Lempar2[l] = daduPemain1Lempar2[len(daduPemain1Lempar2)-1]
					daduPemain1Lempar2[len(daduPemain1Lempar2)-1] = 0
					daduPemain1Lempar2 = daduPemain1Lempar2[:len(daduPemain1Lempar2)-1]
				}
				if daduPemain2Lempar2[m] == 1 {
					daduPemain3Lempar2 = append(daduPemain3Lempar2, 1)
					daduPemain2Lempar2[m] = daduPemain2Lempar2[len(daduPemain2Lempar2)-1]
					daduPemain2Lempar2[len(daduPemain2Lempar2)-1] = 0
					daduPemain2Lempar2 = daduPemain2Lempar2[:len(daduPemain2Lempar2)-1]
				}
				if daduPemain2Lempar2[m] == 6 {
					counterPemain2++
					daduPemain2Lempar2[m] = daduPemain2Lempar2[len(daduPemain2Lempar2)-1]
					daduPemain2Lempar2[len(daduPemain2Lempar2)-1] = 0
					daduPemain2Lempar2 = daduPemain2Lempar2[:len(daduPemain2Lempar2)-1]
				}
				if daduPemain3Lempar2[n] == 1 {
					daduPemain1Lempar2 = append(daduPemain1Lempar2, 1)
					daduPemain3Lempar2[n] = daduPemain3Lempar2[len(daduPemain3Lempar2)-1]
					daduPemain3Lempar2[len(daduPemain3Lempar2)-1] = 0
					daduPemain3Lempar2 = daduPemain3Lempar2[:len(daduPemain3Lempar2)-1]
				}
				if daduPemain3Lempar2[n] == 6 {
					counterPemain3++
					daduPemain3Lempar2[n] = daduPemain3Lempar2[len(daduPemain3Lempar2)-1]
					daduPemain3Lempar2[len(daduPemain3Lempar2)-1] = 0
					daduPemain3Lempar2 = daduPemain3Lempar2[:len(daduPemain3Lempar2)-1]
				}
				l++
				m++
				n++
			}
			fmt.Println("Evaluasi")
			fmt.Printf("Pemain 1:(%d): %v\n", counterPemain1, daduPemain1Lempar2)
			fmt.Printf("Pemain 2:(%d): %v\n", counterPemain2, daduPemain2Lempar2)
			fmt.Printf("Pemain 3:(%d): %v\n", counterPemain3, daduPemain3Lempar2)
			fmt.Println()

			o, p, q := 0, 0, 0

			for i := range daduPemain1Lempar2 {
				daduPemain1Lempar3 = append(daduPemain1Lempar3, rollDice())
				i++
			}
			for i := range daduPemain2Lempar2 {
				daduPemain2Lempar3 = append(daduPemain2Lempar3, rollDice())
				i++
			}
			for i := range daduPemain3Lempar2 {
				daduPemain3Lempar3 = append(daduPemain3Lempar3, rollDice())
				i++
			}

			fmt.Printf("Giliran %d lempar dadu:\n", 3)
			fmt.Println("Pemain 1:", daduPemain1Lempar3)
			fmt.Println("Pemain 2:", daduPemain2Lempar3)
			fmt.Println("Pemain 3:", daduPemain3Lempar3)
			fmt.Println()

			for o < len(daduPemain1Lempar3) && p < len(daduPemain2Lempar3) && q < len(daduPemain3Lempar3) {
				if daduPemain1Lempar3[o] == 1 {
					daduPemain2Lempar3 = append(daduPemain2Lempar3, 1)
					daduPemain1Lempar3[o] = daduPemain1Lempar3[len(daduPemain1Lempar3)-1]
					daduPemain1Lempar3[len(daduPemain1Lempar3)-1] = 0
					daduPemain1Lempar3 = daduPemain1Lempar3[:len(daduPemain1Lempar3)-1]
				}
				if daduPemain1Lempar3[o] == 6 {
					counterPemain1++
					daduPemain1Lempar3[o] = daduPemain1Lempar3[len(daduPemain1Lempar3)-1]
					daduPemain1Lempar3[len(daduPemain1Lempar3)-1] = 0
					daduPemain1Lempar3 = daduPemain1Lempar3[:len(daduPemain1Lempar3)-1]
				}
				if daduPemain2Lempar3[p] == 1 {
					daduPemain3Lempar3 = append(daduPemain3Lempar3, 1)
					daduPemain2Lempar3[p] = daduPemain2Lempar3[len(daduPemain2Lempar3)-1]
					daduPemain2Lempar3[len(daduPemain2Lempar3)-1] = 0
					daduPemain2Lempar3 = daduPemain2Lempar3[:len(daduPemain2Lempar3)-1]
				}
				if daduPemain2Lempar3[p] == 6 {
					counterPemain2++
					daduPemain2Lempar3[p] = daduPemain2Lempar3[len(daduPemain2Lempar3)-1]
					daduPemain2Lempar3[len(daduPemain2Lempar3)-1] = 0
					daduPemain2Lempar3 = daduPemain2Lempar3[:len(daduPemain2Lempar3)-1]
				}
				if daduPemain3Lempar3[q] == 1 {
					daduPemain1Lempar3 = append(daduPemain1Lempar3, 1)
					daduPemain3Lempar3[q] = daduPemain3Lempar3[len(daduPemain3Lempar3)-1]
					daduPemain3Lempar3[len(daduPemain3Lempar3)-1] = 0
					daduPemain3Lempar3 = daduPemain3Lempar3[:len(daduPemain3Lempar3)-1]
				}
				if daduPemain3Lempar3[q] == 6 {
					counterPemain3++
					daduPemain3Lempar3[q] = daduPemain3Lempar3[len(daduPemain3Lempar3)-1]
					daduPemain3Lempar3[len(daduPemain3Lempar3)-1] = 0
					daduPemain3Lempar3 = daduPemain3Lempar3[:len(daduPemain3Lempar3)-1]
				}
				o++
				p++
				q++
			}
			fmt.Println("Evaluasi")
			fmt.Printf("Pemain 1:(%d): %v\n", counterPemain1, daduPemain1Lempar3)
			fmt.Printf("Pemain 2:(%d): %v\n", counterPemain2, daduPemain2Lempar3)
			fmt.Printf("Pemain 3:(%d): %v\n", counterPemain3, daduPemain3Lempar3)
			fmt.Println()
		}
	}
}

func main() {
	pemain, dadu := 3, 2
	Dice(pemain, dadu)
}
