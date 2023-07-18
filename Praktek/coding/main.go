package main

import (
	"fmt"
	"strconv"
)

type Pemain struct {
	Nama       string
	JumlahDadu int
	Poin       int
	Dadu       []int
}

func main() {
	var N int
	var M int
	fmt.Print("Masukkan jumlah pemain: ")
	fmt.Scan(&N)
	fmt.Print("Masukkan jumlah dadu: ")
	fmt.Scan(&M)

	pemain := make([]Pemain, N)

	for i := 1; i <= N; i++ {
		pemain[i-1] = Pemain{
			Nama:       fmt.Sprintf("Pemain %d", i),
			JumlahDadu: M,
			Poin:       0,
			Dadu:       make([]int, 10),
		}
	}

	round := 1
	for len(pemain) > 1 {
		fmt.Println("Giliran " + strconv.Itoa(round) + " lempar dadu")
		for i := range pemain {
				fmt.Println(pemain[i].Nama)
				for j := 0; j < pemain[i].JumlahDadu; j++ {
					fmt.Print("Dadu ke-" + strconv.Itoa(j+1) + " : ")
					fmt.Scan(&pemain[i].Dadu[j])
			}
		}

		for i := range pemain {
			if pemain[i].JumlahDadu > 0 {
				for j := 0; j <= pemain[i].JumlahDadu; j++ {
					if pemain[i].Dadu[j] == 6 {
						pemain[i].Poin++
						pemain[i].JumlahDadu--
					} else if pemain[i].Dadu[j] == 1 {
						pemain[i].JumlahDadu--
						if i == len(pemain)-1 {
							pemain[0].JumlahDadu++
						} else {
							pemain[i+1].JumlahDadu++
						}
					}
					pemain[i].Dadu[j] = 0
				}
			}
		}

		fmt.Println("Evaluasi")
		for i := range pemain {
			if pemain[i].JumlahDadu > 0 {
				fmt.Println("Poin " + pemain[i].Nama + " : " + strconv.Itoa(pemain[i].Poin) + ", sisa dadu " + strconv.Itoa(pemain[i].JumlahDadu))
			} else {
				fmt.Println("Poin " + pemain[i].Nama + " : " + strconv.Itoa(pemain[i].Poin) + ", Berhenti bermain karena tidak memiliki dadu")
			}
		}

		var pemainBaru []Pemain
		for i := range pemain {
			if pemain[i].JumlahDadu > 0 {
				pemainBaru = append(pemainBaru, pemain[i])
			}
		}
		pemain = pemainBaru
		round++
	}
	fmt.Println("Permainan berakhir, pemain dengan poin tertinggi menang")
}
