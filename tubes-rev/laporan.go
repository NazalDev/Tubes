package main

import (
	"fmt"
	"time"
)

func laporan(A *tabTest, B arrPengguna, maxU, x int, userID string) {
	var i, j, k int = 0, 0, x
	var D, AH, S string
	fmt.Println("| Laporan :")
	j = x

	for i < 5 && j >= 0 && k != -1 {
		if userID == A[k].userID {
			fmt.Printf("| %d. %d %s %d %s Depresi level : %12s, Kecemasan level : %12s, Stress level : %12s\n", i+1, A[k].date, A[k].month, A[k].year, A[k].time, A[k].hasilD, A[k].hasilA, A[k].hasilS)
			j--
			i++
		}
		k--
	}
	for i = 0; i <= x; i++ {
		if userID == A[i].userID {
			rataDepresi(A, i, &D)
			rataCemas(A, i, &AH)
			rataStress(A, i, &S)
		}
	}
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
	loginPage(A, B, maxU, x, userID)
}

func rataDepresi(A *tabTest, i int, hasil *string) {
	var n, r, s, b, sb int = 0, 0, 0, 0, 0
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

	fmt.Println("Depresi ", n, r, s, b, sb) //nanti dihapus

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
func rataCemas(A *tabTest, i int, hasil *string) {
	var n, r, s, b, sb int = 0, 0, 0, 0, 0
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

	fmt.Println("anxiety ", n, r, s, b, sb) //nanti dihapus

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
func rataStress(A *tabTest, i int, hasil *string) {
	var n, r, s, b, sb int = 0, 0, 0, 0, 0
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

	fmt.Println("stres ", n, r, s, b, sb) //nanti dihapus

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
