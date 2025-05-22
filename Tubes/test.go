package main

import (
	"fmt"
	"time"
)

func test(A *dataTes) {
	const layoutTime = "15:04"
	var answer int

	A.time = time.Now().Format(layoutTime)
	A.year = time.Now().Year()
	A.date = time.Now().Day()
	A.month = time.Now().Month().String()

	A.depress = 0
	A.anxiety = 0
	A.stress = 0
	fmt.Println("============   Question   ===========")
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
			A.depress += answer - 1
		} else if i == 2 || i == 4 || i == 7 || i == 9 || i == 15 || i == 19 || i == 20 {
			A.anxiety += answer - 1
		} else {
			A.stress += answer - 1
		}
	}

	A.depress *= 2
	A.anxiety *= 2
	A.stress *= 2

	if A.depress >= 0 && A.depress <= 9 {
		A.hasilD = "Normal"
	} else if A.depress >= 10 && A.depress <= 13 {
		A.hasilD = "Ringan"
	} else if A.depress >= 14 && A.depress <= 20 {
		A.hasilD = "Sedang"
	} else if A.depress >= 21 && A.depress <= 27 {
		A.hasilD = "Berat"
	} else {
		A.hasilD = "Sangat_Berat"
	}

	if A.anxiety >= 0 && A.anxiety <= 7 {
		A.hasilA = "Normal"
	} else if A.anxiety >= 8 && A.anxiety <= 9 {
		A.hasilA = "Ringan"
	} else if A.anxiety >= 10 && A.anxiety <= 14 {
		A.hasilA = "Sedang"
	} else if A.anxiety >= 15 && A.anxiety <= 19 {
		A.hasilA = "Berat"
	} else {
		A.hasilA = "Sangat_Berat"
	}

	if A.stress >= 0 && A.stress <= 14 {
		A.hasilS = "Normal"
	} else if A.stress >= 15 && A.stress <= 18 {
		A.hasilS = "Ringan"
	} else if A.stress >= 19 && A.stress <= 25 {
		A.hasilS = "Sedang"
	} else if A.stress >= 26 && A.stress <= 33 {
		A.hasilS = "Berat"
	} else {
		A.hasilS = "Sangat_Berat"
	}
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
		fmt.Println("| Saya merasa seperti akan panik.") //Anxiety
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
	option(x, a)
}

func option(x int, a *int) {
	fmt.Println("| 1. Tidak Pernah")
	fmt.Println("| 2. Kadang-kadang")
	fmt.Println("| 3. Cukup Sering")
	fmt.Println("| 4. Sangat Sering/selalu")
	fmt.Print("| Pilih angka (1/2/3/4) : ")
	fmt.Scan(a)
	fmt.Println("| ")
	if *a < 1 || *a > 4 {
		question(x, a)
	}
}
