package main

import (
	"fmt"
	"time"
)

const NMAX int = 10000

type userTest struct {
	date, year, depress, anxiety, stress int
	month, time, hasilD, hasilA, hasilS  string
}
type tabTest [NMAX]userTest

func main() {
	var arrTest tabTest
	var maxT int
	maxT = -1
	loginPage(&arrTest, maxT)
}

func loginPage(A *tabTest, m int) {
	//I.S.
	//F.S.
	var x, idx int

	fmt.Println("============   Dash Board   ===========")
	fmt.Println("|1. Start Test                        |")
	fmt.Println("|2. Riwayat                           |")
	fmt.Println("|3. Ubah Data                         |")
	fmt.Println("|4. Hapus Data                        |")
	fmt.Println("|5. Laporan                           |")
	fmt.Println("|6. Exit                              |")
	fmt.Println("=======================================")
	fmt.Println()
	fmt.Print("Pilih salah satu menu : ")
	fmt.Scan(&x)
	switch x {
	case 1:
		test(A, x, idx, &m)
	case 2:
		riwayat(A, x, m, &idx)
	case 3:
		changeData(A, x, m, idx)
	case 4:
		deleteData(A, x, idx, &m)
	case 5:
		laporan(A, m)
	case 6:
	}
}

func riwayat(A *tabTest, temp, x int, idx *int) {
	var find string = "null"
	var a, b, pick, d, y int
	if x < 0 {
		fmt.Println("=======================================")
		fmt.Println("           Data Not Yet Made           ")
		fmt.Println("=======================================")
		fmt.Println()
		fmt.Println()
		loginPage(A, x)
	} else {
		fmt.Println("===========================================   RIWAYAT dan List Data   ==========================================")
		for i := 0; i <= x; i++ {
			fmt.Printf("| %d. %d %s %d %s Depresi level : %12s, Kecemasan level : %12s, Stress level : %12s\n", i+1, A[i].date, A[i].month, A[i].year, A[i].time, A[i].hasilD, A[i].hasilA, A[i].hasilS)

		}
		if temp == 2 {
			fmt.Println("| ")
			fmt.Println("| Urutkan atau cari data ?  ")
			fmt.Println("| 1. Urutkan")
			fmt.Println("| 2. Cari data")
			fmt.Println("| 3. Kembali")
			fmt.Print("| Pilih angka (1/2/3) ")
			fmt.Scan(&pick)
			fmt.Println("| ")
			if pick == 1 {
				sortMenu(&a, &b)
				switch b {
				case 1:
					if a == 4 {
						sortDateAsc(*A, x, d, y, find)
					} else {
						sortSkorAsc(*A, a, x, find)
					}
				case 2:
					if a == 4 {
						sortDateDes(*A, x)
					} else {
						sortSkorDes(*A, a, x)
					}
				}
			} else if pick == 2 {
				searchMenu(&b)
				switch b {
				case 1:
					fmt.Print("| Tanggal : ")
					fmt.Scan(&d)
					fmt.Print("| Bulan (Depan wajib kapital) : ")
					fmt.Scan(&find)
					fmt.Print("| Tahun : ")
					fmt.Scan(&y)
					sortDateAsc(*A, x, d, y, find)
				case 2:
					fmt.Println("| Data yang akan di search :")
					fmt.Println("| 1. Depresi")
					fmt.Println("| 2. Cemas")
					fmt.Println("| 3. Stress")
					fmt.Print("| Pilih angka (1/2/3) ")
					fmt.Scan(&a)
					fmt.Println("|")
					fmt.Println("| Level data yang akan dicari (Normal/Ringan/Sedang/Berat/Sangat_Berat)")
					fmt.Println("| Note : Sesuaikan penulisan dengan contoh")
					fmt.Print("| Pilih level : ")
					fmt.Scan(&find)
					fmt.Println("|")
					sortSkorAsc(*A, a, x, find)
				}
			}
			loginPage(A, x)
		}
	}
	if temp == 3 || temp == 4 {
		fmt.Print("| Pilih data: ")
		fmt.Scan(idx)
		fmt.Println()
	}
}

func test(A *tabTest, temp, idx int, x *int) {
	const layoutTime = "15:04"
	var answer, tempMax int
	if temp != 3 {
		*x += 1
	} else {
		tempMax = *x
		*x = idx - 1
	}
	A[*x].time = time.Now().Format(layoutTime)
	A[*x].year = time.Now().Year()
	A[*x].date = time.Now().Day()
	A[*x].month = time.Now().Month().String()

	A[*x].depress = 0
	A[*x].anxiety = 0
	A[*x].stress = 0
	fmt.Println("============   Question   ===========")
	for i := 1; i <= 21; i++ {
		question(A, i, &answer)
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
			A[*x].depress += answer - 1
		} else if i == 2 || i == 4 || i == 7 || i == 9 || i == 15 || i == 19 || i == 20 {
			A[*x].anxiety += answer - 1
		} else {
			A[*x].stress += answer - 1
		}
	}

	A[*x].depress *= 2
	A[*x].anxiety *= 2
	A[*x].stress *= 2

	if A[*x].depress >= 0 && A[*x].depress <= 9 {
		A[*x].hasilD = "Normal"
	} else if A[*x].depress >= 10 && A[*x].depress <= 13 {
		A[*x].hasilD = "Ringan"
	} else if A[*x].depress >= 14 && A[*x].depress <= 20 {
		A[*x].hasilD = "Sedang"
	} else if A[*x].depress >= 21 && A[*x].depress <= 27 {
		A[*x].hasilD = "Berat"
	} else {
		A[*x].hasilD = "Sangat_Berat"
	}

	if A[*x].anxiety >= 0 && A[*x].anxiety <= 7 {
		A[*x].hasilA = "Normal"
	} else if A[*x].anxiety >= 8 && A[*x].anxiety <= 9 {
		A[*x].hasilA = "Ringan"
	} else if A[*x].anxiety >= 10 && A[*x].anxiety <= 14 {
		A[*x].hasilA = "Sedang"
	} else if A[*x].anxiety >= 15 && A[*x].anxiety <= 19 {
		A[*x].hasilA = "Berat"
	} else {
		A[*x].hasilA = "Sangat_Berat"
	}

	if A[*x].stress >= 0 && A[*x].stress <= 14 {
		A[*x].hasilS = "Normal"
	} else if A[*x].stress >= 15 && A[*x].stress <= 18 {
		A[*x].hasilS = "Ringan"
	} else if A[*x].stress >= 19 && A[*x].stress <= 25 {
		A[*x].hasilS = "Sedang"
	} else if A[*x].stress >= 26 && A[*x].stress <= 33 {
		A[*x].hasilS = "Berat"
	} else {
		A[*x].hasilS = "Sangat_Berat"
	}
	if temp == 3 {
		*x = tempMax
	}
	loginPage(A, *x)
}

func question(A *tabTest, x int, a *int) {
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
	option(A, x, a)
}

func option(A *tabTest, x int, a *int) {
	fmt.Println("| 1. Tidak Pernah")
	fmt.Println("| 2. Kadang-kadang")
	fmt.Println("| 3. Cukup Sering")
	fmt.Println("| 4. Sangat Sering/selalu")
	fmt.Print("| Pilih angka (1/2/3/4) : ")
	fmt.Scan(a)
	fmt.Println("| ")
	if *a < 1 || *a > 4 {
		question(A, x, a)
	}
}

func changeData(A *tabTest, temp, x, idx int) {
	riwayat(A, temp, x, &idx)
	test(A, temp, idx, &x)
	loginPage(A, x)
}

func deleteData(A *tabTest, temp, idx int, x *int) {
	riwayat(A, temp, *x, &idx)
	for i := idx - 1; i <= *x; i++ {
		A[i] = A[i+1]
	}
	*x -= 1
	riwayat(A, 1, *x, &idx)
	loginPage(A, *x)
}

func sortSkorAsc(A tabTest, n, x int, find string) {
	var temp userTest
	var i, idxMin int
	switch n {
	case 1:
		for i = 0; i < x; i++ {
			idxMin = i
			for j := i + 1; j <= x; j++ {
				if A[idxMin].depress > A[j].depress {
					idxMin = j
				}
			}
			temp = A[idxMin]
			A[idxMin] = A[i]
			A[i] = temp
		}
	case 2:
		for i = 0; i < x; i++ {
			idxMin = i
			for j := i + 1; j <= x; j++ {
				if A[idxMin].anxiety > A[j].anxiety {
					idxMin = j
				}
			}
			temp = A[idxMin]
			A[idxMin] = A[i]
			A[i] = temp
		}
	case 3:
		for i = 0; i < x; i++ {
			idxMin = i
			for j := i + 1; j <= x; j++ {
				if A[idxMin].stress > A[j].stress {
					idxMin = j
				}
			}
			temp = A[idxMin]
			A[idxMin] = A[i]
			A[i] = temp
		}
	}
	if find == "null" {
		fmt.Println("=============================================   Hasil Mengurutkan   ============================================")
		for i = 0; i <= x; i++ {
			fmt.Printf("| %d. %d %s %d %s Depresi level : %12s, Kecemasan level : %12s, Stress level : %12s\n", i+1, A[i].date, A[i].month, A[i].year, A[i].time, A[i].hasilD, A[i].hasilA, A[i].hasilS)
		}
	} else {
		sequentialSearch(A, n, x, find)
	}
}
func sortSkorDes(A tabTest, n, x int) {
	var temp userTest
	var i, idxMin int
	switch n {
	case 1:
		for i = 0; i < x; i++ {
			idxMin = i
			for j := i + 1; j <= x; j++ {
				if A[idxMin].depress < A[j].depress {
					idxMin = j
				}
			}
			temp = A[idxMin]
			A[idxMin] = A[i]
			A[i] = temp
		}
	case 2:
		for i = 0; i < x; i++ {
			idxMin = i
			for j := i + 1; j <= x; j++ {
				if A[idxMin].anxiety < A[j].anxiety {
					idxMin = j
				}
			}
			temp = A[idxMin]
			A[idxMin] = A[i]
			A[i] = temp
		}
	case 3:
		for i = 0; i < x; i++ {
			idxMin = i
			for j := i + 1; j <= x; j++ {
				if A[idxMin].stress < A[j].stress {
					idxMin = j
				}
			}
			temp = A[idxMin]
			A[idxMin] = A[i]
			A[i] = temp
		}
	}
	fmt.Println("=============================================   Hasil Mengurutkan   ============================================")
	for i = 0; i <= x; i++ {
		fmt.Printf("| %d. %d %s %d %s Depresi level : %12s, Kecemasan level : %12s, Stress level : %12s\n", i+1, A[i].date, A[i].month, A[i].year, A[i].time, A[i].hasilD, A[i].hasilA, A[i].hasilS)
	}
}

func sortDateAsc(A tabTest, x, date, year int, month string) {
	var i, j int
	var temp userTest
	i = 1
	for i <= x {
		j = i
		temp = A[j]
		for j > 0 && temp.date < A[j-1].date {
			A[j] = A[j-1]
			j -= 1
		}
		A[j] = temp
		i += 1
	}
	if month == "null" {
		fmt.Println("=============================================   Hasil Mengurutkan   ============================================")
		for i = 0; i <= x; i++ {
			fmt.Printf("| %d. %d %s %d %s Depresi level : %12s, Kecemasan level : %12s, Stress level : %12s\n", i+1, A[i].date, A[i].month, A[i].year, A[i].time, A[i].hasilD, A[i].hasilA, A[i].hasilS)
		}
	} else {
		binarySearch(A, x, date, year, month)
	}
}

func sortDateDes(A tabTest, x int) {
	var i, j int
	var temp userTest
	i = 1
	for i <= x {
		j = i
		temp = A[j]
		for j > 0 && temp.date > A[j-1].date {
			A[j] = A[j-1]
			j -= 1
		}
		A[j] = temp
		i += 1
	}
	fmt.Println("=============================================   Hasil Mengurutkan   ============================================")
	for i = 0; i <= x; i++ {
		fmt.Printf("| %d. %d %s %d %s Depresi level : %12s, Kecemasan level : %12s, Stress level : %12s\n", i+1, A[i].date, A[i].month, A[i].year, A[i].time, A[i].hasilD, A[i].hasilA, A[i].hasilS)
	}
}

func sortMenu(a, b *int) {
	fmt.Println("| Urutkan sesuai :")
	fmt.Println("| 1. Level Depresi")
	fmt.Println("| 2. Level Kecemasan")
	fmt.Println("| 3. Level Stress")
	fmt.Println("| 4. Tanggal pengerjaan")
	fmt.Print("| Pilih angka (1/2/3/4)  ")
	fmt.Scan(a)
	fmt.Println("| ")
	fmt.Println("| Menaik/menurun? ")
	fmt.Println("| 1. Menaik")
	fmt.Println("| 2. Menurun")
	fmt.Print("| Pilih angka (1/2)  ")
	fmt.Scan(b)
	fmt.Println()
}

func laporan(A *tabTest, x int) {
	var i, j int = 0, 0
	var D, AH, S string
	fmt.Println("==================================================   Laporan   =================================================")
	j = x

	for i < 5 && j >= 0 {
		fmt.Printf("| %d. %d %s %d %s Depresi level : %12s, Kecemasan level : %12s, Stress level : %12s\n", i+1, A[i].date, A[i].month, A[i].year, A[i].time, A[i].hasilD, A[i].hasilA, A[i].hasilS)
		j -= 1
		i += 1
	}
	rataDepresi(A, x, &D)
	rataCemas(A, x, &AH)
	rataStress(A, x, &S)
	fmt.Println("| Rata-rata depresi : ", D)
	fmt.Println("| Rata-rata cemas   : ", AH)
	fmt.Println("| Rata-rata stress  : ", S)

	if D == "Sangat_Berat" || AH == "Sangat_Berat" || S == "Sangat_Berat" {
		fmt.Println("| Berkonsultasilah dengan psikolog atau psikiater")
	} else if D == "Berat" || AH == "Berat" || S == "Berat" {
		fmt.Println("| Sangat direkomendasikan konsultasi dengan psikolog atau psikiater")
	} else if D == "Sedang" || AH == "Sedang" || S == "Sedang" {
		fmt.Println("| Tetaplah tenang dan cobalah berbicara kepada orang lain atau psikolog atau psikiater")
	} else if D == "Ringan" || AH == "Ringan" || S == "Ringan" {
		fmt.Println("| Jangan berfikir kamu itu buruk, teruslah melangkah pelan-pelan")
	} else {
		fmt.Println("| Pertahankan kondisimu yang sekarang :D")
	}
	fmt.Println("|")
	if D > AH && D > S {
		fmt.Println("| \"Sometimes when you're in a dark place, you think you've been buried, but actually you've been planted.\" - Christine Caine")
	} else if AH > D && AH > S {
		fmt.Println("| \"Trust yourself. You've survived a lot, and you'll survive whatever is coming.\" - Robert Tew")
	} else if S > D && S > AH {
		fmt.Println("| \"The greatest weapon against stress is our ability to choose one thought over another\" - William James")
	} else {
		fmt.Println("| \"The best revenge is massive success.\" - Frank Sinatra")
	}
	loginPage(A, x)
}

func rataDepresi(A *tabTest, x int, hasil *string) {
	var n, r, s, b, sb int = 0, 0, 0, 0, 0
	for i := 0; i <= x; i++ {
		if time.Now().Month().String() == A[i].month {
			if A[i].hasilD == "Normal" {
				n += 1
			} else if A[i].hasilD == "Ringan" {
				r += 1
			} else if A[i].hasilD == "Sedang" {
				s += 1
			} else if A[i].hasilD == "Berat" {
				b += 1
			} else if A[i].hasilD == "Sangat_Berat" {
				sb += 1
			}
		}
	}

	if n > r && n > s && n > b && n > sb {
		*hasil = "Normal"
	} else if r > n && r > s && r > b && r > sb {
		*hasil = "Ringan"
	} else if s > n && s > r && s > b && s > sb {
		*hasil = "Sedang"
	} else if b > n && b > s && b > r && b > sb {
		*hasil = "Berat"
	} else if sb > n && sb > s && sb > b && sb > r {
		*hasil = "Sangat_Berat"
	} else {
		*hasil = "Data tidak cukup"
	}
}
func rataCemas(A *tabTest, x int, hasil *string) {
	var n, r, s, b, sb int = 0, 0, 0, 0, 0
	for i := 0; i <= x; i++ {
		if time.Now().Month().String() == A[i].month {
			if A[i].hasilD == "Normal" {
				n += 1
			} else if A[i].hasilA == "Ringan" {
				r += 1
			} else if A[i].hasilA == "Sedang" {
				s += 1
			} else if A[i].hasilA == "Berat" {
				b += 1
			} else if A[i].hasilA == "Sangat_Berat" {
				sb += 1
			}
		}
	}
	if n > r && n > s && n > b && n > sb {
		*hasil = "Normal"
	} else if r > n && r > s && r > b && r > sb {
		*hasil = "Ringan"
	} else if s > n && s > r && s > b && s > sb {
		*hasil = "Sedang"
	} else if b > n && b > s && b > r && b > sb {
		*hasil = "Berat"
	} else if sb > n && sb > s && sb > b && sb > r {
		*hasil = "Sangat_Berat"
	} else {
		*hasil = "Data tidak cukup"
	}
}
func rataStress(A *tabTest, x int, hasil *string) {
	var n, r, s, b, sb int = 0, 0, 0, 0, 0
	for i := 0; i <= x; i++ {
		if time.Now().Month().String() == A[i].month {
			if A[i].hasilD == "Normal" {
				n += 1
			} else if A[i].hasilS == "Ringan" {
				r += 1
			} else if A[i].hasilS == "Sedang" {
				s += 1
			} else if A[i].hasilS == "Berat" {
				b += 1
			} else if A[i].hasilS == "Sangat_Berat" {
				sb += 1
			}
		}
	}
	if n > r && n > s && n > b && n > sb {
		*hasil = "Normal"
	} else if r > n && r > s && r > b && r > sb {
		*hasil = "Ringan"
	} else if s > n && s > r && s > b && s > sb {
		*hasil = "Sedang"
	} else if b > n && b > s && b > r && b > sb {
		*hasil = "Berat"
	} else if sb > n && sb > s && sb > b && sb > r {
		*hasil = "Sangat_Berat"
	} else {
		*hasil = "Data tidak cukup"
	}
}

func searchMenu(pick *int) {
	fmt.Println("=========== Pencarian Data  ===========")
	fmt.Println("| 1. Berdasarkan tanggal")
	fmt.Println("| 2. Berdasarkan skor")
	fmt.Print("| Pilih angka (1/2) ")
	fmt.Scan(pick)
	fmt.Println("| ")
}

func sequentialSearch(A tabTest, n, x int, find string) {
	var found bool = false
	var j int = 1
	switch n {
	case 1:
		for i := 0; i <= x; i++ {
			if find == A[i].hasilD {
				fmt.Printf("| %d. %d %s %d %s Depresi level : %12s, Kecemasan level : %12s, Stress level : %12s\n", i+1, A[i].date, A[i].month, A[i].year, A[i].time, A[i].hasilD, A[i].hasilA, A[i].hasilS)
				found = true
				j++
			}
		}
	case 2:
		for i := 0; i <= x; i++ {
			if find == A[i].hasilA {
				fmt.Printf("| %d. %d %s %d %s Depresi level : %12s, Kecemasan level : %12s, Stress level : %12s\n", i+1, A[i].date, A[i].month, A[i].year, A[i].time, A[i].hasilD, A[i].hasilA, A[i].hasilS)
				found = true
				j++
			}
		}
	case 3:
		for i := 0; i <= x; i++ {
			if find == A[i].hasilS {
				fmt.Printf("| %d. %d %s %d %s Depresi level : %12s, Kecemasan level : %12s, Stress level : %12s\n", i+1, A[i].date, A[i].month, A[i].year, A[i].time, A[i].hasilD, A[i].hasilA, A[i].hasilS)
				found = true
				j++
			}
		}
	}
	if !found {
		fmt.Println("|         Data tidak ditemukan        |")
	}
}

func binarySearch(A tabTest, x, date, year int, month string) {
	var med, kr, kn int = 0, 0, x - 1
	var j int = 1
	var found bool
	for i := 0; i <= x; i++ {
		found = false
		for kr <= kn && found == false {
			med = (kr + kn) / 2
			if date > A[med].date {
				kn = med - 1
			} else if date < A[med].date {
				kr = med + 1
			} else if A[med].month == month && A[med].year == year {
				fmt.Printf("| %d. %d %s %d %s Depresi level : %12s, Kecemasan level : %12s, Stress level : %12s\n", i+1, A[i].date, A[i].month, A[i].year, A[i].time, A[i].hasilD, A[i].hasilA, A[i].hasilS)
				found = true
				j++
			}
		}
	}
	if !found {
		fmt.Println("|         Data tidak ditemukan        |")
	}

}
