package main

import (
	"fmt"
	"time"
)

const NMAX int = 1024
const MAXUSER int = 8

type dataTes struct {
	date, year, depress, anxiety, stress int
	month, time, hasilD, hasilA, hasilS  string
}

type pengguna struct {
	userId, pass string
	totTest      int
	tabTest      [NMAX]dataTes
}
type arrPengguna [MAXUSER]pengguna

func main() {
	var user arrPengguna
	var maxUser int = 0
	ui(&user, &maxUser)
}

func ui(A *arrPengguna, maxUser *int) {
	//I.S. terdefinisi array A dengan tipe data arrPengguna dan maxUser
	//F.S. pemanggilan fungsi lain berdasarkan fungsinya masing-masing

	var x int = 0
	var userIdx int
	var pick string
	fmt.Println("========  Selamat Datang  =======")
	fmt.Println("|   Manajemen Kesehatan Mental  |")
	fmt.Println("|    2025 - Nazal Putra P. D.   |")
	fmt.Println("|===============================|")
	fmt.Println("|1. Registrasi                  |")
	fmt.Println("|2. Login                       |")
	fmt.Println("|3. Log                         |")
	fmt.Println("|4. Quit                        |")
	fmt.Println("=================================")
	fmt.Println()
	fmt.Print("Pilih salah satu menu : ")
	fmt.Scan(&x)
	switch x {
	case 1:
		register(A, maxUser)
	case 2:
		login(A, *maxUser)
	case 3:
		if *maxUser > 0 {
			fmt.Println("==========  Log Data  ===========")
			for i := 0; i < *maxUser; i++ {
				fmt.Printf("| Username : %s\n", A[i].userId)
				for j := 0; j < A[i].totTest; j++ {
					fmt.Printf("| %d. %2d %9s %d %s\n", j+1, A[i].tabTest[j].date, A[i].tabTest[j].month, A[i].tabTest[j].year, A[i].tabTest[j].time)
				}
				fmt.Println("|")
			}
			fmt.Print("| Cari akun (Y/N) ? ")
			fmt.Scan(&pick)
			fmt.Println("|")
			if pick == "y" || pick == "Y" {
				fmt.Print("| Masukkan username : ")
				fmt.Scan(&pick)
				userIdx = binarySearch(*A, *maxUser, pick)
				if userIdx == -1 {
					fmt.Println("|      Data tidak ditemukan     |")
				} else {
					fmt.Printf("| Username : %s\n", A[userIdx].userId)
					for i := 0; i < A[userIdx].totTest; i++ {
						fmt.Printf("| %d. %2d %9s %d %s\n", i+1, A[userIdx].tabTest[i].date, A[userIdx].tabTest[i].month, A[userIdx].tabTest[i].year, A[userIdx].tabTest[i].time)
					}
				}
			}
		} else {
			fmt.Println("=================================")
			fmt.Println("|       Data Not Yet Made       |")
			fmt.Println("=================================")
			fmt.Println("|                               |")
		}
		ui(A, maxUser)
	case 24:
		dataIn(A, maxUser)
		ui(A, maxUser)
	default:
	}

}

func register(A *arrPengguna, maxU *int) {
	// I.S. terdefinisi array pengguna A, n dan maxU
	// F.S. terisi array A dengan data user dan maxU bertambah 1
	fmt.Println("==========  Registrasi  =========")
	fmt.Print("| Masukkan Username : ")
	fmt.Scan(&A[*maxU].userId)
	fmt.Print("| Masukkan Password : ")
	fmt.Scan(&A[*maxU].pass)
	nameSecurity(*A, *maxU)
	*maxU++
	confirmLogin(*A, *maxU)
}

func confirmLogin(A arrPengguna, maxU int) {
	// I.S. terdefinisi array pengguna A dan maxU
	// F.S. pemanggilan fungsi login atau kembali ke menu utama
	var confirmInput string
	fmt.Println("=========  Confirmation  ========")
	fmt.Println("|                               |")
	fmt.Println("|        To login page?         |")
	fmt.Println("|              Y/N              |")
	fmt.Println("|                               |")
	fmt.Println("=================================")
	fmt.Scan(&confirmInput)
	if confirmInput == "y" || confirmInput == "Y" {
		login(&A, maxU)
	} else {
		ui(&A, &maxU)
	}
}

func login(A *arrPengguna, maxU int) {
	// I.S. terdefinisi array pengguna A dan maxU
	// F.S. pemanggilan fungsi tes utama atau kembali ke menu utama jika salah menginput akun
	var tempUserID, tempPass string
	var i, idx int = 3, -1
	for i = 3; i >= 0 && idx == -1 && maxU > 0; i-- {
		fmt.Println("============  Login  ============")
		fmt.Print("| Masukkan Username : ")
		fmt.Scan(&tempUserID)
		fmt.Print("| Masukkan Password : ")
		fmt.Scan(&tempPass)
		idx = searchtrue(*A, tempUserID, tempPass)
		if idx != -1 {
			loginPage(A, maxU, idx)
		} else if i != 0 {
			fmt.Println("=================================")
			fmt.Println("|                               |")
			fmt.Println("|   Password/username salah!    |")
			fmt.Printf("|     sisa attempt masuk %d      |\n", i)
			fmt.Println("=================================")
		}
	}

	if maxU < 1 {
		fmt.Println("=== Belum ada akun yang dibuat ==")
	}

	if idx == -1 {
		ui(A, &maxU)
	}
}

func searchtrue(A arrPengguna, ID, pass string) int {
	//Fungsi ini akan mereturn index akun jika benar akun tersebut memang ada
	var i, idx int = 0, -1
	for i < MAXUSER && idx == -1 {
		if A[i].userId == ID && A[i].pass == pass {
			idx = i
		} else {
			i += 1
		}
	}
	return idx
}

func nameSecurity(A arrPengguna, idx int) {
	// I.S terdefinisi array pengguna A dan idx
	// F.S procedure ini akan mengecek apakah username pernah digunakan atau tidak dan akan terus mengulang input username sampai menginput username baru
	var i int = 0
	for i < idx {
		if A[idx].userId == A[i].userId {
			fmt.Println("| !!! Username telah digunakan !!!")
			fmt.Println()
			fmt.Print("| Masukkan Username : ")
			fmt.Scan(&A[idx].userId)
			fmt.Print("| Masukkan Password : ")
			fmt.Scan(&A[idx].pass)
			fmt.Println()
		} else {
			i++
		}
	}
}

func loginPage(A *arrPengguna, maxU, userIdx int) {
	//I.S. terdefinisi array pengguna A, maxU, dan useridx
	//F.S. memanggil fungsi lain dengan fungsi masing-masing
	var x, idx int
	for {
		fmt.Println("=============   Menu Test   ===========")
		fmt.Println("|1. Start Test                        |")
		fmt.Println("|2. Riwayat                           |")
		fmt.Println("|3. Ubah Data                         |")
		fmt.Println("|4. Hapus Data                        |")
		fmt.Println("|5. Laporan                           |")
		fmt.Println("|6. Back                              |")
		fmt.Println("|7. Exit                              |")
		fmt.Println("=======================================")
		fmt.Println()
		fmt.Print("Pilih salah satu menu : ")
		fmt.Scan(&x)
		switch x {
		case 1:
			test(&A[userIdx].tabTest[A[userIdx].totTest])
			A[userIdx].totTest += 1
		case 2:
			riwayat(*A, x, userIdx, &idx)
		case 3:
			changeData(A, userIdx, x, idx)
		case 4:
			deleteData(A, userIdx, x, idx)
		case 5:
			laporan(*A, userIdx)
		case 6:
			ui(A, &maxU)
		default:
		}
		if x == 7 || x == 6 {
			break
		}
	}
}

func test(A *dataTes) {
	// I.S terdefinisi array test A
	// F.S terisi data tes untuk pengguna tersebut
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
		answer = question(i)
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

func question(x int) int {
	// Fungsi ini akan mereturn jawaban dari soal yang telah di output
	var a int
	for {

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
		fmt.Println("| 1. Tidak Pernah")
		fmt.Println("| 2. Kadang-kadang")
		fmt.Println("| 3. Cukup Sering")
		fmt.Println("| 4. Sangat Sering/selalu")
		fmt.Print("| Pilih angka (1/2/3/4) : ")
		fmt.Scan(&a)
		fmt.Println("| ")
		if a >= 1 && a <= 4 {
			break
		}
	}
	return a
}

func riwayat(A arrPengguna, temp, userIdx int, idx *int) {
	// I.S terdefinisi array pengguna A, temp, userIdx, idx
	// F.S mengoutput hasil riwayat tes pengguna tersebut
	var a, b, pick, d, y int
	var find string = "null"
	if A[userIdx].totTest < 0 {
		fmt.Println("=======================================")
		fmt.Println("           Data Not Yet Made           ")
		fmt.Println("=======================================")
		fmt.Println("|                                     |")
	} else {
		fmt.Println("| Riwayat dan List Data :")
		for i := 0; i < A[userIdx].totTest; i++ {
			fmt.Printf("| %d. %2d %9s %d %s Depresi level : %12s, Kecemasan level : %12s, Stress level : %12s\n", i+1, A[userIdx].tabTest[i].date, A[userIdx].tabTest[i].month, A[userIdx].tabTest[i].year, A[userIdx].tabTest[i].time, A[userIdx].tabTest[i].hasilD, A[userIdx].tabTest[i].hasilA, A[userIdx].tabTest[i].hasilS)
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
				if a == 4 {
					sortDateDes(A, userIdx)
				} else {
					switch b {
					case 1:
						sortSkorAsc(A, userIdx, a, find)
					case 2:
						sortSkorDes(A, userIdx, a)
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
					sortDateAsc(&A, userIdx, d, y, find)
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
					sortSkorAsc(A, userIdx, a, find)
				}
			}
		}
	}
	if temp == 3 || temp == 4 {
		fmt.Print("| Pilih data: ")
		fmt.Scan(idx)
		fmt.Println()
	}
}

func monthToInt(month string) int {
	// Fungsi ini akan mengubah nama bulan menjadi angka
	switch month {
	case "January":
		return 1
	case "February":
		return 2
	case "March":
		return 3
	case "April":
		return 4
	case "May":
		return 5
	case "June":
		return 6
	case "July":
		return 7
	case "August":
		return 8
	case "September":
		return 9
	case "October":
		return 10
	case "November":
		return 11
	case "December":
		return 12
	case "Januari":
		return 1
	case "Februari":
		return 2
	case "Maret":
		return 3
	case "Mei":
		return 5
	case "Juni":
		return 6
	case "Juli":
		return 7
	case "Agustus":
		return 8
	case "Desember":
		return 12
	default:
		return -1 // Invalid month
	}
}

func changeData(A *arrPengguna, userIdx, temp, idx int) {
	// I.S. terdefinisi array pengguna A, userIdx, temp, idx
	// F.S. mengubah data tes pengguna tersebut untuk data tertentu
	var a string
	if idx < A[userIdx].totTest && idx >= 0 {
		riwayat(*A, temp, userIdx, &idx)
		idx--
		confirm(&a)
		if a == "y" || a == "Y" {
			test(&A[userIdx].tabTest[idx])
			sortDateAsc(A, userIdx, 0, 0, "null")
		}
	}
}

func deleteData(A *arrPengguna, userIdx, temp, idx int) {
	// I.S. terdefinisi array pengguna A, userIdx, temp, idx
	// F.S. menghapus salah satu data tes pengguna
	var a string
	riwayat(*A, temp, userIdx, &idx)
	confirm(&a)
	if a == "y" || a == "Y" {
		for i := idx - 1; i < A[userIdx].totTest; i++ {
			A[userIdx].tabTest[i] = A[userIdx].tabTest[i+1]
		}
		A[userIdx].totTest--
		riwayat(*A, 2, userIdx, &idx)
	}
}

func confirm(answer *string) {
	// I.S. terdefinisi answer
	// F.S. mengonfirmasi perilaku pengguna dalam menghapus atau mengubah data tes
	fmt.Println("=======================================")
	fmt.Println("|            Yakin ? Y / N            |")
	fmt.Println("=======================================")
	fmt.Scan(answer)
}

func laporan(A arrPengguna, userIdx int) {
	// I.S. terdefinisi array pengguna A dan userIdx
	// F.S. akan mengoutput 5 tes terakhir pengguna dan rata-rata hasil tes
	var i, j int = 0, 0
	var hasil, D, AH, S string
	fmt.Println("| Laporan :")
	j = A[userIdx].totTest - 1

	for i < 5 && j > 0 {
		fmt.Printf("| %d. %2d %9s %d %s Depresi level : %12s, Kecemasan level : %12s, Stress level : %12s\n", i+1, A[userIdx].tabTest[j].date, A[userIdx].tabTest[j].month, A[userIdx].tabTest[j].year, A[userIdx].tabTest[j].time, A[userIdx].tabTest[j].hasilD, A[userIdx].tabTest[j].hasilA, A[userIdx].tabTest[j].hasilS)
		j--
		i++

	}
	for j = 1; j <= 3; j++ {
		var n, r, s, b, sb int = 0, 0, 0, 0, 0
		for i = 0; i < A[userIdx].totTest && i < 30; i++ {
			switch j {
			case 1:
				rataDepresi(A, i, userIdx, &n, &r, &s, &b, &sb)
			case 2:
				rataCemas(A, i, userIdx, &n, &r, &s, &b, &sb)
			case 3:
				rataStress(A, i, userIdx, &n, &r, &s, &b, &sb)
			}
		}
		if n > r && n > s && n > b && n > sb {
			hasil = "Normal"
		} else if r > n && r > s && r > b && r > sb {
			hasil = "Ringan"
		} else if s > n && s > r && s > b && s > sb {
			hasil = "Sedang"
		} else if b > n && b > s && b > r && b > sb {
			hasil = "Berat"
		} else if sb > n && sb > s && sb > b && sb > r {
			hasil = "Sangat_Berat"
		} else {
			hasil = "Data tidak cukup"
		}
		switch j {
		case 1:
			fmt.Println("| Rata-rata depresi : ", hasil)
			D = hasil
		case 2:
			fmt.Println("| Rata-rata cemas   : ", hasil)
			AH = hasil
		case 3:
			fmt.Println("| Rata-rata stress  : ", hasil)
			S = hasil
		}
	}

	if D == "Sangat_Berat" || AH == "Sangat_Berat" || S == "Sangat_Berat" {
		fmt.Println("| Berkonsultasilah dengan psikolog atau psikiater")
	} else if D == "Berat" || AH == "Berat" || S == "Berat" {
		fmt.Println("| Sangat direkomendasikan konsultasi dengan psikolog atau psikiater")
	} else if D == "Sedang" || AH == "Sedang" || S == "Sedang" {
		fmt.Println("| Tetaplah tenang dan cobalah berbicara kepada orang lain atau psikolog atau psikiater")
	} else if D == "Ringan" || AH == "Ringan" || S == "Ringan" {
		fmt.Println("| Jangan berfikir kamu itu buruk, teruslah melangkah pelan-pelan")
	} else if D == "Data tidak cukup" || AH == "Data tidak cukup" || S == "Data tidak cukup" {
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
}

func rataDepresi(A arrPengguna, i, userIdx int, n, r, s, b, sb *int) {
	// I.S. terdefinisi array pengguna A, i, userIdx
	// F.S. akan menghitung total hasil jawaban depresi pengguna
	if A[userIdx].tabTest[i].hasilD == "Normal" {
		*n += 1
	} else if A[userIdx].tabTest[i].hasilD == "Ringan" {
		*r += 1
	} else if A[userIdx].tabTest[i].hasilD == "Sedang" {
		*s += 1
	} else if A[userIdx].tabTest[i].hasilD == "Berat" {
		*b += 1
	} else if A[userIdx].tabTest[i].hasilD == "Sangat_Berat" {
		*sb += 1
	}
}

func rataCemas(A arrPengguna, i, userIdx int, n, r, s, b, sb *int) {
	// I.S. terdefinisi array pengguna A, i, userIdx
	// F.S. akan menghitung total hasil jawaban kecemasan pengguna
	if A[userIdx].tabTest[i].hasilA == "Normal" {
		*n += 1
	} else if A[userIdx].tabTest[i].hasilA == "Ringan" {
		*r += 1
	} else if A[userIdx].tabTest[i].hasilA == "Sedang" {
		*s += 1
	} else if A[userIdx].tabTest[i].hasilA == "Berat" {
		*b += 1
	} else if A[userIdx].tabTest[i].hasilA == "Sangat_Berat" {
		*sb += 1
	}
}

func rataStress(A arrPengguna, i, userIdx int, n, r, s, b, sb *int) {
	// I.S. terdefinisi array pengguna A, i, userIdx
	// F.S. akan menghitung total hasil jawaban stress pengguna
	if A[userIdx].tabTest[i].hasilS == "Normal" {
		*n += 1
	} else if A[userIdx].tabTest[i].hasilS == "Ringan" {
		*r += 1
	} else if A[userIdx].tabTest[i].hasilS == "Sedang" {
		*s += 1
	} else if A[userIdx].tabTest[i].hasilS == "Berat" {
		*b += 1
	} else if A[userIdx].tabTest[i].hasilS == "Sangat_Berat" {
		*sb += 1
	}
}

func searchMenu(pick *int) {
	// I.S. terdefinisi pick
	// F.S. hasil pilihan fitur search
	fmt.Println("=========== Pencarian Data  ===========")
	fmt.Println("| 1. Berdasarkan tanggal")
	fmt.Println("| 2. Berdasarkan skor")
	fmt.Print("| Pilih angka (1/2) ")
	fmt.Scan(pick)
	fmt.Println("|")
}

func sequentialSearch(A arrPengguna, userIdx, n int, find string) {
	// I.S. terdefinisi array pengguna A, userIdx, n, find
	// F.S. search secara sequential mencari data tes berdasarkan hasil per kategori
	var found bool = false
	var j int = 1
	switch n {
	case 1:
		for i := 0; i < A[userIdx].totTest; i++ {
			if find == A[userIdx].tabTest[i].hasilD {
				fmt.Printf("| %d. %2d %9s %d %s Depresi level : %12s, Kecemasan level : %12s, Stress level : %12s\n", j, A[userIdx].tabTest[i].date, A[userIdx].tabTest[i].month, A[userIdx].tabTest[i].year, A[userIdx].tabTest[i].time, A[userIdx].tabTest[i].hasilD, A[userIdx].tabTest[i].hasilA, A[userIdx].tabTest[i].hasilS)
				found = true
				j++
			}
		}
	case 2:
		for i := 0; i < A[userIdx].totTest; i++ {
			if find == A[userIdx].tabTest[i].hasilA {
				fmt.Printf("| %d. %2d %9s %d %s Depresi level : %12s, Kecemasan level : %12s, Stress level : %12s\n", j, A[userIdx].tabTest[i].date, A[userIdx].tabTest[i].month, A[userIdx].tabTest[i].year, A[userIdx].tabTest[i].time, A[userIdx].tabTest[i].hasilD, A[userIdx].tabTest[i].hasilA, A[userIdx].tabTest[i].hasilS)
				found = true
				j++
			}
		}
	case 3:
		for i := 0; i < A[userIdx].totTest; i++ {
			if find == A[userIdx].tabTest[i].hasilS {
				fmt.Printf("| %d. %2d %9s %d %s Depresi level : %12s, Kecemasan level : %12s, Stress level : %12s\n", j, A[userIdx].tabTest[i].date, A[userIdx].tabTest[i].month, A[userIdx].tabTest[i].year, A[userIdx].tabTest[i].time, A[userIdx].tabTest[i].hasilD, A[userIdx].tabTest[i].hasilA, A[userIdx].tabTest[i].hasilS)
				found = true
				j++
			}
		}
	}
	if !found {
		fmt.Println("|         Data tidak ditemukan        |")
	}
}

func dateSearch(A arrPengguna, userIdx, date, year int, month string) {
	// I.S. terdefinisi array pengguna A, userIdx, date, year, month
	// F.S. search secara sequential mencari data tes berdasarkan tanggal pengerjaan
	var found bool = false
	var j int = 1
	for i := 0; i < A[userIdx].totTest; i++ {
		if date == A[userIdx].tabTest[i].date {
			if month == A[userIdx].tabTest[i].month && year == A[userIdx].tabTest[i].year {
				fmt.Printf("| %d. %2d %9s %d %s Depresi level : %12s, Kecemasan level : %12s, Stress level : %12s\n", j, A[userIdx].tabTest[i].date, A[userIdx].tabTest[i].month, A[userIdx].tabTest[i].year, A[userIdx].tabTest[i].time, A[userIdx].tabTest[i].hasilD, A[userIdx].tabTest[i].hasilA, A[userIdx].tabTest[i].hasilS)
				found = true
				j++
			}
		}
	}
	if !found {
		fmt.Println("|         Data tidak ditemukan        |")
	}
}

func binarySearch(A arrPengguna, n int, find string) int {
	// Fungsi ini akan mengeluarkan index user yang ingin dicari secara binary
	var found int = -1
	var med, kr, kn int = 0, 0, n - 1
	for kr <= kn && found == -1 {
		med = (kr + kn) / 2
		if A[med].userId == find {
			found = med
		} else if A[med].userId < find {
			kr = med + 1
		} else {
			kn = med - 1
		}
	}
	return found
}

func sortSkorAsc(A arrPengguna, userIdx, n int, find string) {
	// I.S. terdefinisi array pengguna A, userIdx, n, find
	// F.S. mengurutkan dan mengoutput data tes pengguna secara menaik menggunakan selection sort berdasarkan hasil tes per kategori
	var temp dataTes
	var i, idxMin int
	switch n {
	case 1:
		for i = 0; i < A[userIdx].totTest-1; i++ {
			idxMin = i
			for j := i + 1; j < A[userIdx].totTest; j++ {
				if A[userIdx].tabTest[idxMin].depress > A[userIdx].tabTest[j].depress {
					idxMin = j
				}
			}
			temp = A[userIdx].tabTest[idxMin]
			A[userIdx].tabTest[idxMin] = A[userIdx].tabTest[i]
			A[userIdx].tabTest[i] = temp
		}
	case 2:
		for i = 0; i < A[userIdx].totTest-1; i++ {
			idxMin = i
			for j := i + 1; j < A[userIdx].totTest; j++ {
				if A[userIdx].tabTest[idxMin].anxiety > A[userIdx].tabTest[j].anxiety {
					idxMin = j
				}
			}
			temp = A[userIdx].tabTest[idxMin]
			A[userIdx].tabTest[idxMin] = A[userIdx].tabTest[i]
			A[userIdx].tabTest[i] = temp
		}
	case 3:
		for i = 0; i < A[userIdx].totTest-1; i++ {
			idxMin = i
			for j := i + 1; j < A[userIdx].totTest; j++ {
				if A[userIdx].tabTest[idxMin].stress > A[userIdx].tabTest[j].stress {
					idxMin = j
				}
			}
			temp = A[userIdx].tabTest[idxMin]
			A[userIdx].tabTest[idxMin] = A[userIdx].tabTest[i]
			A[userIdx].tabTest[i] = temp
		}
	}
	if find == "null" {
		fmt.Println("| Hasil Mengurutkan :")
		for i = 0; i < A[userIdx].totTest; i++ {
			fmt.Printf("| %d. %2d %9s %d %s Depresi level : %12s, Kecemasan level : %12s, Stress level : %12s\n", i+1, A[userIdx].tabTest[i].date, A[userIdx].tabTest[i].month, A[userIdx].tabTest[i].year, A[userIdx].tabTest[i].time, A[userIdx].tabTest[i].hasilD, A[userIdx].tabTest[i].hasilA, A[userIdx].tabTest[i].hasilS)
		}
	} else {
		sequentialSearch(A, userIdx, n, find)
	}
}

func sortSkorDes(A arrPengguna, userIdx, n int) {
	// I.S. terdefinisi array pengguna A, userIdx, n
	// F.S. mengurutkan dan mengoutput data tes pengguna secara menurun menggunakan selection sort berdasarkan hasil tes per kategori
	var temp dataTes
	var i, idxMin int
	switch n {
	case 1:
		for i = 0; i < A[userIdx].totTest-1; i++ {
			idxMin = i
			for j := i + 1; j < A[userIdx].totTest; j++ {
				if A[userIdx].tabTest[idxMin].depress < A[userIdx].tabTest[j].depress {
					idxMin = j
				}
			}
			temp = A[userIdx].tabTest[idxMin]
			A[userIdx].tabTest[idxMin] = A[userIdx].tabTest[i]
			A[userIdx].tabTest[i] = temp
		}
	case 2:
		for i = 0; i < A[userIdx].totTest-1; i++ {
			idxMin = i
			for j := i + 1; j < A[userIdx].totTest; j++ {
				if A[userIdx].tabTest[idxMin].anxiety < A[userIdx].tabTest[j].anxiety {
					idxMin = j
				}
			}
			temp = A[userIdx].tabTest[idxMin]
			A[userIdx].tabTest[idxMin] = A[userIdx].tabTest[i]
			A[userIdx].tabTest[i] = temp
		}
	case 3:
		for i = 0; i < A[userIdx].totTest-1; i++ {
			idxMin = i
			for j := i + 1; j < A[userIdx].totTest; j++ {
				if A[userIdx].tabTest[idxMin].stress < A[userIdx].tabTest[j].stress {
					idxMin = j
				}
			}
			temp = A[userIdx].tabTest[idxMin]
			A[userIdx].tabTest[idxMin] = A[userIdx].tabTest[i]
			A[userIdx].tabTest[i] = temp
		}
	}
	fmt.Println("| Hasil Mengurutkan :")
	for i = 0; i < A[userIdx].totTest; i++ {
		fmt.Printf("| %d. %2d %9s %d %s Depresi level : %12s, Kecemasan level : %12s, Stress level : %12s\n", i+1, A[userIdx].tabTest[i].date, A[userIdx].tabTest[i].month, A[userIdx].tabTest[i].year, A[userIdx].tabTest[i].time, A[userIdx].tabTest[i].hasilD, A[userIdx].tabTest[i].hasilA, A[userIdx].tabTest[i].hasilS)
	}
}

func sortDateAsc(A *arrPengguna, userIdx, date, year int, month string) {
	// I.S. terdefinisi array pengguna A, userIdx, date, year, month
	// F.S. mengurutkan dan mengoutput data tes pengguna secara menaik menggunakan insertion sort berdasarkan tanggal pengerjaan
	var i, j int
	var temp dataTes
	i = 1
	for i < A[userIdx].totTest {
		j = i
		temp = A[userIdx].tabTest[j]
		for j > 0 && ((temp.year < A[userIdx].tabTest[j-1].year) ||
			(temp.year == A[userIdx].tabTest[j-1].year && monthToInt(temp.month) < monthToInt(A[userIdx].tabTest[j-1].month)) ||
			(temp.year == A[userIdx].tabTest[j-1].year && monthToInt(temp.month) == monthToInt(A[userIdx].tabTest[j-1].month) && temp.date < A[userIdx].tabTest[j-1].date)) {
			A[userIdx].tabTest[j] = A[userIdx].tabTest[j-1]
			j -= 1
		}
		A[userIdx].tabTest[j] = temp
		i += 1
	}
	if month == "null" {
		fmt.Println("| Hasil Mengurutkan :")
		for i = 0; i < A[userIdx].totTest; i++ {
			fmt.Printf("| %d. %2d %9s %d %s Depresi level : %12s, Kecemasan level : %12s, Stress level : %12s\n", i+1, A[userIdx].tabTest[i].date, A[userIdx].tabTest[i].month, A[userIdx].tabTest[i].year, A[userIdx].tabTest[i].time, A[userIdx].tabTest[i].hasilD, A[userIdx].tabTest[i].hasilA, A[userIdx].tabTest[i].hasilS)
		}
	} else {
		dateSearch(*A, userIdx, date, year, month)
	}
}

func sortDateDes(A arrPengguna, userIdx int) {
	// I.S. terdefinisi array pengguna A, userIdx, date, year, month
	// F.S. mengurutkan dan mengoutput data tes pengguna secara menurun menggunakan insertion sort berdasarkan tanggal pengerjaan
	var i, j int
	var temp dataTes
	i = 1
	for i < A[userIdx].totTest {
		j = i
		temp = A[userIdx].tabTest[j]
		for j > 0 && ((temp.year > A[userIdx].tabTest[j-1].year) ||
			(temp.year == A[userIdx].tabTest[j-1].year && monthToInt(temp.month) > monthToInt(A[userIdx].tabTest[j-1].month)) ||
			(temp.year == A[userIdx].tabTest[j-1].year && monthToInt(temp.month) == monthToInt(A[userIdx].tabTest[j-1].month) && temp.date > A[userIdx].tabTest[j-1].date)) {
			A[userIdx].tabTest[j] = A[userIdx].tabTest[j-1]
			j -= 1
		}
		A[userIdx].tabTest[j] = temp
		i += 1
	}
	fmt.Println("| Hasil Mengurutkan :")
	for i = 0; i < A[userIdx].totTest; i++ {
		fmt.Printf("| %d. %2d %9s %d %s Depresi level : %12s, Kecemasan level : %12s, Stress level : %12s\n", i+1, A[userIdx].tabTest[i].date, A[userIdx].tabTest[i].month, A[userIdx].tabTest[i].year, A[userIdx].tabTest[i].time, A[userIdx].tabTest[i].hasilD, A[userIdx].tabTest[i].hasilA, A[userIdx].tabTest[i].hasilS)
	}
}

func sortMenu(a, b *int) {
	fmt.Println("| Urutkan sesuai :")
	fmt.Println("| 1. Level Depresi")
	fmt.Println("| 2. Level Kecemasan")
	fmt.Println("| 3. Level Stress")
	fmt.Println("| 4. Tanggal pengerjaan (menurun)")
	fmt.Print("| Pilih angka (1/2/3/4)  ")
	fmt.Scan(a)
	if *a != 4 {
		fmt.Println("| ")
		fmt.Println("| Menaik/menurun? ")
		fmt.Println("| 1. Menaik")
		fmt.Println("| 2. Menurun")
		fmt.Print("| Pilih angka (1/2)  ")
		fmt.Scan(b)
		fmt.Println()
	}
}

// Dev mode type 24 in account menu
func dataIn(A *arrPengguna, maxU *int) {
	//user 1
	A[0].userId = "name1"
	A[0].pass = "pass1"
	A[0].totTest = 8
	//test 1
	A[0].tabTest[0].date = 3
	A[0].tabTest[0].month = "May"
	A[0].tabTest[0].year = 2012
	A[0].tabTest[0].time = "11:50"
	A[0].tabTest[0].depress = 20
	A[0].tabTest[0].anxiety = 16
	A[0].tabTest[0].stress = 24
	A[0].tabTest[0].hasilD = "Sedang"
	A[0].tabTest[0].hasilA = "Berat"
	A[0].tabTest[0].hasilS = "Sedang"

	// test 2
	A[0].tabTest[1].date = 3
	A[0].tabTest[1].month = "May"
	A[0].tabTest[1].year = 2012
	A[0].tabTest[1].time = "23:11"
	A[0].tabTest[1].depress = 24
	A[0].tabTest[1].anxiety = 2
	A[0].tabTest[1].stress = 10
	A[0].tabTest[1].hasilD = "Berat"
	A[0].tabTest[1].hasilA = "Normal"
	A[0].tabTest[1].hasilS = "Normal"

	//test 3
	A[0].tabTest[2].date = 4
	A[0].tabTest[2].month = "May"
	A[0].tabTest[2].year = 2012
	A[0].tabTest[2].time = "06:12"
	A[0].tabTest[2].depress = 12
	A[0].tabTest[2].anxiety = 20
	A[0].tabTest[2].stress = 26
	A[0].tabTest[2].hasilD = "Ringan"
	A[0].tabTest[2].hasilA = "Sangat_Berat"
	A[0].tabTest[2].hasilS = "Berat"

	//test 4
	A[0].tabTest[3].date = 4
	A[0].tabTest[3].month = "May"
	A[0].tabTest[3].year = 2012
	A[0].tabTest[3].time = "21:02"
	A[0].tabTest[3].depress = 28
	A[0].tabTest[3].anxiety = 2
	A[0].tabTest[3].stress = 14
	A[0].tabTest[3].hasilD = "Sangat_Berat"
	A[0].tabTest[3].hasilA = "Normal"
	A[0].tabTest[3].hasilS = "Normal"

	//test 5
	A[0].tabTest[4].date = 1
	A[0].tabTest[4].month = "June"
	A[0].tabTest[4].year = 2012
	A[0].tabTest[4].time = "08:00"
	A[0].tabTest[4].depress = 4
	A[0].tabTest[4].anxiety = 18
	A[0].tabTest[4].stress = 30
	A[0].tabTest[4].hasilD = "Normal"
	A[0].tabTest[4].hasilA = "Berat"
	A[0].tabTest[4].hasilS = "Berat"

	//test 6
	A[0].tabTest[5].date = 5
	A[0].tabTest[5].month = "August"
	A[0].tabTest[5].year = 2012
	A[0].tabTest[5].time = "01:12"
	A[0].tabTest[5].depress = 10
	A[0].tabTest[5].anxiety = 12
	A[0].tabTest[5].stress = 22
	A[0].tabTest[5].hasilD = "Ringan"
	A[0].tabTest[5].hasilA = "Sedang"
	A[0].tabTest[5].hasilS = "Sedang"

	//test 7
	A[0].tabTest[6].date = 2
	A[0].tabTest[6].month = "January"
	A[0].tabTest[6].year = 2013
	A[0].tabTest[6].time = "23:59"
	A[0].tabTest[6].depress = 26
	A[0].tabTest[6].anxiety = 16
	A[0].tabTest[6].stress = 8
	A[0].tabTest[6].hasilD = "Berat"
	A[0].tabTest[6].hasilA = "Berat"
	A[0].tabTest[6].hasilS = "Normal"

	//test 8
	A[0].tabTest[7].date = 21
	A[0].tabTest[7].month = "January"
	A[0].tabTest[7].year = 2013
	A[0].tabTest[7].time = "08:29"
	A[0].tabTest[7].depress = 20
	A[0].tabTest[7].anxiety = 8
	A[0].tabTest[7].stress = 8
	A[0].tabTest[7].hasilD = "Sedang"
	A[0].tabTest[7].hasilA = "Ringan"
	A[0].tabTest[7].hasilS = "Normal"

	//user 2
	A[1].userId = "name2"
	A[1].pass = "pass2"
	A[1].totTest = 3

	//test 1
	A[1].tabTest[0].date = 1
	A[1].tabTest[0].month = "May"
	A[1].tabTest[0].year = 2012
	A[1].tabTest[0].time = "11:50"
	A[1].tabTest[0].depress = 28
	A[1].tabTest[0].anxiety = 20
	A[1].tabTest[0].stress = 32
	A[1].tabTest[0].hasilD = "Sangat_Berat"
	A[1].tabTest[0].hasilA = "sangat_Berat"
	A[1].tabTest[0].hasilS = "Berat"
	//test 2
	A[1].tabTest[1].date = 5
	A[1].tabTest[1].month = "June"
	A[1].tabTest[1].year = 2012
	A[1].tabTest[1].time = "23:11"
	A[1].tabTest[1].depress = 4
	A[1].tabTest[1].anxiety = 4
	A[1].tabTest[1].stress = 12
	A[1].tabTest[1].hasilD = "Normal"
	A[1].tabTest[1].hasilA = "Normal"
	A[1].tabTest[1].hasilS = "Normal"

	//test 3
	A[1].tabTest[2].date = 2
	A[1].tabTest[2].month = "May"
	A[1].tabTest[2].year = 2013
	A[1].tabTest[2].time = "06:12"
	A[1].tabTest[2].depress = 16
	A[1].tabTest[2].anxiety = 10
	A[1].tabTest[2].stress = 2
	A[1].tabTest[2].hasilD = "Sedang"
	A[1].tabTest[2].hasilA = "Sedang"
	A[1].tabTest[2].hasilS = "Normal"

	//user 3
	A[2].userId = "name3"
	A[2].pass = "pass3"
	A[2].totTest = 5

	//test 1
	A[2].tabTest[0].date = 2
	A[2].tabTest[0].month = "May"
	A[2].tabTest[0].year = 2012
	A[2].tabTest[0].time = "11:50"
	A[2].tabTest[0].depress = 14
	A[2].tabTest[0].anxiety = 18
	A[2].tabTest[0].stress = 6
	A[2].tabTest[0].hasilD = "Sedang"
	A[2].tabTest[0].hasilA = "Berat"
	A[2].tabTest[0].hasilS = "Normal"

	//test 2
	A[2].tabTest[1].date = 3
	A[2].tabTest[1].month = "May"
	A[2].tabTest[1].year = 2012
	A[2].tabTest[1].time = "23:11"
	A[2].tabTest[1].depress = 28
	A[2].tabTest[1].anxiety = 12
	A[2].tabTest[1].stress = 6
	A[2].tabTest[1].hasilD = "Sangat_Berat"
	A[2].tabTest[1].hasilA = "Sedang"
	A[2].tabTest[1].hasilS = "Normal"

	//test 3
	A[2].tabTest[2].date = 5
	A[2].tabTest[2].month = "July"
	A[2].tabTest[2].year = 2012
	A[2].tabTest[2].time = "06:12"
	A[2].tabTest[2].depress = 14
	A[2].tabTest[2].anxiety = 8
	A[2].tabTest[2].stress = 30
	A[2].tabTest[2].hasilD = "Sedang"
	A[2].tabTest[2].hasilA = "Ringan"
	A[2].tabTest[2].hasilS = "Berat"

	//test 4
	A[2].tabTest[3].date = 3
	A[2].tabTest[3].month = "September"
	A[2].tabTest[3].year = 2012
	A[2].tabTest[3].time = "21:02"
	A[2].tabTest[3].depress = 24
	A[2].tabTest[3].anxiety = 20
	A[2].tabTest[3].stress = 22
	A[2].tabTest[3].hasilD = "Berat"
	A[2].tabTest[3].hasilA = "Sangat_Berat"
	A[2].tabTest[3].hasilS = "Sedang"

	//test 5
	A[2].tabTest[4].date = 1
	A[2].tabTest[4].month = "June"
	A[2].tabTest[4].year = 2013
	A[2].tabTest[4].time = "08:00"
	A[2].tabTest[4].depress = 16
	A[2].tabTest[4].anxiety = 6
	A[2].tabTest[4].stress = 2
	A[2].tabTest[4].hasilD = "Sedang"
	A[2].tabTest[4].hasilA = "Normal"
	A[2].tabTest[4].hasilS = "Normal"

	*maxU = 3

}
