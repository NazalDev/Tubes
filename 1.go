package main

import (
	"fmt"
	"time"
)

const NMAX int = 1000

type userTest struct {
	date, year, depress, anxiety, stress, maxTest int
	month, time, hasilD, hasilA, hasilS           string
}
type tabUser [NMAX]userTest

var arrTest tabUser

func main() {
	var maxT int = -1
	loginPage(maxT)
}

func loginPage(m int) {
	//I.S.
	//F.S.
	var x, idx int

	fmt.Println("============Dash Board===========")
	fmt.Println("|1. Start Test                  |")
	fmt.Println("|2. Riwayat                     |")
	fmt.Println("|3. Ubah Data                   |")
	fmt.Println("|4. Hapus Data                  |")
	fmt.Println("|5. Exit                        |")
	fmt.Println("=================================")
	fmt.Println()
	fmt.Print("Pilih salah satu menu : ")
	fmt.Scan(&x)
	switch x {
	case 1:
		test(&m)
	case 2:
		riwayat(x, m, &idx)
	case 3:
		changeData(x, m, idx)
	case 4:
		deleteData()
	case 5:
	}
}

func riwayat(temp, x int, idx *int) {
	if x < 0 {
		fmt.Println("        Data Not Yet Made        ")
	} else {
		fmt.Println("====================================   RIWAYAT dan List Data   ====================================")
		for i := 0; i <= x; i++ {
			fmt.Printf("|%d. %d %s %d %s Depresi level : %s, Kecemasan level : %s, Stress level : %s %d\n", i+1, arrTest[i].date, arrTest[i].month, arrTest[i].year, arrTest[i].time, arrTest[i].hasilD, arrTest[i].hasilA, arrTest[i].hasilS, arrTest[i].depress)

		}
	}
	if temp == 3 || temp == 4 {
		fmt.Print("Pilih data: ")
		fmt.Scan(&*idx)
		fmt.Println()
	} else {
		loginPage(x)
	}
}

func test(x *int) {
	const layoutTime = "15:04"
	var answer int
	*x += 1
	arrTest[*x].time = time.Now().Format(layoutTime)
	arrTest[*x].year = time.Now().Year()
	arrTest[*x].month = time.Now().Month().String()
	arrTest[*x].date = time.Now().Day()

	arrTest[*x].depress = 0
	arrTest[*x].anxiety = 0
	arrTest[*x].stress = 0
	fmt.Println("============Question===========")
	for i := 1; i <= 21; i++ {
		question(i, &answer)
		if i == 3 || i == 5 || i == 10 || i == 13 || i == 17 || i == 16 || i == 21 {
			if i == 10 {
				switch answer {
				case 1:
					answer = 4
				case 2:
					answer = 3
				case 3:
					answer = 2
				case 4:
					answer = 1
				}
			}
			arrTest[*x].depress += answer - 1
		} else if i == 2 || i == 4 || i == 7 || i == 9 || i == 15 || i == 19 || i == 20 {
			arrTest[*x].anxiety += answer - 1
		} else {
			arrTest[*x].stress += answer - 1
		}
	}

	arrTest[*x].depress *= 2
	arrTest[*x].anxiety *= 2
	arrTest[*x].stress *= 2

	if arrTest[*x].depress >= 0 && arrTest[*x].depress <= 9 {
		arrTest[*x].hasilD = "Normal"
	} else if arrTest[*x].depress >= 10 && arrTest[*x].depress <= 13 {
		arrTest[*x].hasilD = "Ringan"
	} else if arrTest[*x].depress >= 14 && arrTest[*x].depress <= 20 {
		arrTest[*x].hasilD = "Sedang"
	} else if arrTest[*x].depress >= 21 && arrTest[*x].depress <= 27 {
		arrTest[*x].hasilD = "Berat"
	} else {
		arrTest[*x].hasilD = "Sangat Berat"
	}

	if arrTest[*x].anxiety >= 0 && arrTest[*x].anxiety <= 7 {
		arrTest[*x].hasilA = "Normal"
	} else if arrTest[*x].anxiety >= 8 && arrTest[*x].anxiety <= 9 {
		arrTest[*x].hasilA = "Ringan"
	} else if arrTest[*x].anxiety >= 10 && arrTest[*x].anxiety <= 14 {
		arrTest[*x].hasilA = "Sedang"
	} else if arrTest[*x].anxiety >= 15 && arrTest[*x].anxiety <= 19 {
		arrTest[*x].hasilA = "Berat"
	} else {
		arrTest[*x].hasilA = "Sangat Berat"
	}

	if arrTest[*x].stress >= 0 && arrTest[*x].stress <= 14 {
		arrTest[*x].hasilS = "Normal"
	} else if arrTest[*x].stress >= 15 && arrTest[*x].stress <= 18 {
		arrTest[*x].hasilS = "Ringan"
	} else if arrTest[*x].stress >= 19 && arrTest[*x].stress <= 25 {
		arrTest[*x].hasilS = "Sedang"
	} else if arrTest[*x].stress >= 26 && arrTest[*x].stress <= 33 {
		arrTest[*x].hasilS = "Berat"
	} else {
		arrTest[*x].hasilS = "Sangat Berat"
	}
	loginPage(*x)
}

func question(x int, a *int) {
	switch x {
	case 1:
		fmt.Println("| Saya merasa kesulitan untuk tenang.") //Stressed
	case 2:
		fmt.Println("| Saya merasa mulut saya kering.") //Anxiety
	case 3:
		fmt.Println("| Saya merasa tidak bisa mengalami perasaan positif sama sekali.") //Depressed
	case 4:
		fmt.Println("| Saya mengalami kesulitan bernafas (misalnya, napas cepat, terengah-engah walaupun tanpa aktivitas fisik).") //Anxiety
	case 5:
		fmt.Println("| Saya merasa tidak bisa memotivasi diri untuk melakukan aktivitas.") //Depressed
	case 6:
		fmt.Println("| Saya cenderung bereaksi secara berlebihan terhadap situasi.") //Stressed
	case 7:
		fmt.Println("| Saya mengalami gemetar (misalnya, tangan gemetar).") //Anxiety
	case 8:
		fmt.Println("| Saya merasa kesulitan untuk bersantai.") //Stressed
	case 9:
		fmt.Println("| Saya merasa gugup tanpa alasan yang jelas.") //Anxiety
	case 10:
		fmt.Println("| Saya merasa bahwa saya memiliki banyak energi positif.") //Depressed skor dibalik
	case 11:
		fmt.Println("| Saya merasa mudah tersinggung.") //Stressed
	case 12:
		fmt.Println("| Saya merasa tidak sabaran ketika tertunda dalam melakukan sesuatu.") //Stressed
	case 13:
		fmt.Println("| Saya merasa sedih dan muram.") //Depressed
	case 14:
		fmt.Println("| Saya merasa sulit untuk tidak bereaksi secara emosional terhadap situasi.") //Stressed
	case 15:
		fmt.Println("| Saya merasa seperti panik.") //Anxiety
	case 16:
		fmt.Println("| Saya merasa bahwa tidak ada harapan sama sekali.") //Depressed
	case 17:
		fmt.Println("| Saya merasa saya tidak layak menjadi seseorang.") //Depressed
	case 18:
		fmt.Println("| Saya merasa sangat mudah terganggu.") //Stressed
	case 19:
		fmt.Println("| Saya merasa berkeringat tanpa melakukan aktivitas fisik (misalnya, tangan berkeringat).") //Anxiety
	case 20:
		fmt.Println("| Saya merasa takut tanpa alasan jelas.") //Anxiety
	case 21:
		fmt.Println("| Saya merasa hidup saya tidak berarti. ") //Depressed
	}
	option(x, &*a)
}

func option(x int, a *int) {
	fmt.Println("| 1. Tidak Pernah")
	fmt.Println("| 2. Kadang-kadang")
	fmt.Println("| 3. Cukup Sering")
	fmt.Println("| 4. Sangat Sering/selalu")
	fmt.Print("| Pilih angka (1/2/3/4) : ")
	fmt.Scan(&*a)
	fmt.Println("| ")
	if *a < 1 || *a > 4 {
		question(x, &*a)
	}
}

func changeData(temp, x, idx int) {
	riwayat(temp, x, &idx)
	fmt.Print("it works!")
}

func deleteData() {

}
