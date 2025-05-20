package main

import (
	"fmt"
)

const NMAX int = 10000
const MAXUSER int = 10

type dataTes struct {
	date, year, depress, anxiety, stress, testID int    //Test Variable
	month, time, hasilD, hasilA, hasilS, userID  string //Test Variable
}
type tabTest [NMAX]dataTes

type pengguna struct {
	userId, pass string
}
type arrPengguna [MAXUSER]pengguna

func main() {
	var user arrPengguna
	var test tabTest
	var maxUser int = 0
	var maxT int = -1
	ui(&test, &user, &maxUser, maxT)
}

func ui(A *tabTest, B *arrPengguna, maxUser *int, maxT int) {
	//I.S.
	//F.S.

	var x int = 0
	var nUser int

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
		register(B, *A, &nUser, maxUser, maxT)
	case 2:
		login(B, *A, maxUser, maxT)
	case 3:
		if maxT > -1 {
			fmt.Println("| Riwayat dan List Data :")
			for i := 0; i <= maxT; i++ {
				fmt.Printf("| %d. %d %s %d %s %s Depresi level : %12s, Kecemasan level : %12s, Stress level : %12s\n", i+1, A[i].date, A[i].month, A[i].year, A[i].time, A[i].userID, A[i].hasilD, A[i].hasilA, A[i].hasilS)
			}
		} else {
			fmt.Println("=================================")
			fmt.Println("|       Data Not Yet Made       |")
			fmt.Println("=================================")
			fmt.Println("|                               |")
		}
		ui(A, B, maxUser, maxT)
	default:
	}
}

func register(B *arrPengguna, A tabTest, n, maxU *int, maxT int) {
	// I.S.
	// F.S.
	*n = *maxU
	fmt.Println("==========  Registrasi  =========")
	fmt.Print("| Masukkan Username : ")
	fmt.Scan(&B[*n].userId)
	fmt.Print("| Masukkan Password : ")
	fmt.Scan(&B[*n].pass)
	nameSecurity(*B, *n)
	*maxU += 1
	confirmLogin(*B, A, maxT, maxU)
}

func confirmLogin(B arrPengguna, A tabTest, maxT int, maxU *int) {
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
		login(&B, A, maxU, maxT)
	} else {
		ui(&A, &B, maxU, maxT)
	}
}

func login(B *arrPengguna, A tabTest, maxU *int, maxT int) {
	// I.S.
	// F.S.
	var tempUserID, tempPass string
	var security bool
	var logged bool = false
	var i int
	for i = 3; i >= 0; i-- {
		if *maxU > 0 {
			fmt.Println("============  Login  ============")
			fmt.Print("| Masukkan Username : ")
			fmt.Scan(&tempUserID)
			fmt.Print("| Masukkan Password : ")
			fmt.Scan(&tempPass)
			searchtrue(B, tempUserID, tempPass, &security)
			if security {
				logged = true
				loginPage(&A, *B, *maxU, maxT, tempUserID)
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
	if !logged {
		ui(&A, B, maxU, maxT)
	}
}

func searchtrue(A *arrPengguna, ID, pass string, log *bool) {
	//I.S.
	//F.S.
	var i int
	for i < MAXUSER && !*log {
		if A[i].userId == ID && A[i].pass == pass {
			*log = true
		} else {
			i += 1
		}
	}
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
