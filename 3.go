package main

import "fmt"

const NMAX = 1000000

// tipe bentukan partai
// tipe tabPartai: array of partai dengan kapasitas NMAX
type partai [NMAX]int

func main() {
	// deklarasi variabel
	var tabPartai partai
	var i, suara int
	var pass, idx_max, k, temp int

	// lakukan proses input suara secara berulang di sini, simpan ke dalam array
	// p, sehingga terdapat array p yang berisi hasil peroleh suara n partai.

	i = 0
	for i <= NMAX {
		fmt.Scan(&suara)
		if suara != -1 {
			tabPartai[suara] = tabPartai[suara] + 1
			i++
		}
		fmt.Scan(&suara)
	}

	// lakukan proses pengurutan dengan insertion sort descending berdasarkan jumlah suara yang diperoleh
	for pass = 1; pass <= NMAX-1; pass++ {
		idx_max = pass - 1
		for k = pass; k <= NMAX-1; k++ {
			if tabPartai[idx_max] > tabPartai[k] {
				idx_max = k
			}
		}
		temp = tabPartai[idx_max]
		tabPartai[idx_max] = tabPartai[pass-1]
		tabPartai[pass-1] = temp
	}
	// tampilkan array p
	for m := 0; m < NMAX; m++ {
		if tabPartai[m] != 0 {
			fmt.Print(m, "(", tabPartai[m], ")", "")
		}
	}
}

func posisi(t partai, n int, nama int) int {
	/* mengembalikan indeks partai yang memiliki nama yang dicari pada array t yang
	   berisi n partai atau -1 apabila tidak ditemukan , gunakan pencarian sekuensial
	*/
	var output int
	for i := 0; i < n; i++ {
		if t[i] == nama {
			output = i
		} else {
			output = -1
		}

	}
	return output
}
