package main

import (
	"fmt"
)

const NMAX int = 10000
const MAXUSER int = 10

type dataTes struct {
	date, year, depress, anxiety, stress        int    //Test Variable
	month, time, hasilD, hasilA, hasilS, userID string //Test Variable
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
	//I.S.
	//F.S.

	var x int = 0
	var nUser, userIdx int
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
		register(A, &nUser, maxUser)
	case 2:
		login(A, *maxUser)
	case 3:
		if *maxUser > 0 {
			fmt.Println("==========  Log Data  ===========")
			for i := 0; i < *maxUser; i++ {
				fmt.Printf("| Username : %s\n", A[i].userId)
				for j := 0; j < A[i].totTest; j++ {
					fmt.Printf("| %d. %d %s %d %s\n", j+1, A[i].tabTest[j].date, A[i].tabTest[j].month, A[i].tabTest[j].year, A[i].tabTest[j].time)
				}
				fmt.Println("|")
			}
			fmt.Print("| Cari akun (Y/N) ?")
			fmt.Scan(pick)
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
						fmt.Printf("| %d. %d %s %d %s\n", i+1, A[userIdx].tabTest[i].date, A[userIdx].tabTest[i].month, A[userIdx].tabTest[i].year, A[userIdx].tabTest[i].time)
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
	default:
	}

}

func register(A *arrPengguna, n, maxU *int) {
	// I.S.
	// F.S.
	*n = *maxU
	fmt.Println("==========  Registrasi  =========")
	fmt.Print("| Masukkan Username : ")
	fmt.Scan(&A[*n].userId)
	fmt.Print("| Masukkan Password : ")
	fmt.Scan(&A[*n].pass)
	nameSecurity(*A, *n)
	*maxU += 1
	confirmLogin(*A, *maxU)
}

func confirmLogin(A arrPengguna, maxU int) {
	// I.S.
	// F.S.
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
	// I.S.
	// F.S.
	var tempUserID, tempPass string
	var i, idx int = 3, -1
	for i = 3; i >= 0 && idx == -1; i-- {
		if maxU > 0 {
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
		} else {
			fmt.Println("=== Belum ada akun yang dibuat ==")
		}
	}
	if idx == -1 {
		ui(A, &maxU)
	}
}

func searchtrue(A arrPengguna, ID, pass string) int {
	//I.S.
	//F.S.
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
