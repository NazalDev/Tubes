package main

import (
	"fmt"
)

func laporan(A arrPengguna, userIdx int) {
	var i, j int = 0, 0
	var hasil, D, AH, S string
	fmt.Println("| Laporan :")
	j = A[userIdx].totTest

	for i < 5 && j > 0 {
		fmt.Printf("| %d. %d %s %d %s %s Depresi level : %12s, Kecemasan level : %12s, Stress level : %12s\n", i+1, A[userIdx].tabTest[i].date, A[userIdx].tabTest[i].month, A[userIdx].tabTest[i].year, A[userIdx].tabTest[i].time, A[userIdx].tabTest[i].userID, A[userIdx].tabTest[i].hasilD, A[userIdx].tabTest[i].hasilA, A[userIdx].tabTest[i].hasilS)
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

	fmt.Println("Depresi ", *n, *r, *s, *b, *sb) //nanti dihapus

}
func rataCemas(A arrPengguna, i, userIdx int, n, r, s, b, sb *int) {
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

	fmt.Println("anxiety ", *n, *r, *s, *b, *sb) //nanti dihapus
}
func rataStress(A arrPengguna, i, userIdx int, n, r, s, b, sb *int) {
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

	fmt.Println("stres ", *n, *r, *s, *b, *sb) //nanti dihapus

}
