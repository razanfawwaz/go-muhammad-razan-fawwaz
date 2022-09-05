# (8) Recursive

## Pengertian Rekursif
Rekursif merupakan sebuah teknik pemrograman dimana sebuah fungsi memanggil dirinya sendiri. Fungsi rekursif biasanya memiliki kondisi dasar dan kondisi rekursif. Kondisi dasar merupakan kondisi dimana fungsi tidak memanggil dirinya sendiri lagi. Kondisi rekursif merupakan kondisi dimana fungsi memanggil dirinya sendiri.

### Contoh Rekursif
Berikut adalah contoh fungsi rekursif untuk perhitungan faktorial.

```go
package main

import "fmt"

func factorial(n int) int{
	if n == 1 {
        return 1
	} else {
        return n * factorial(n-1)
    }
}
```

## Number Theory
Merupakan percabangan dari matematika yang mengkaji penyelesaian persoalan matematika, pada mata kuliah biasanya disebut sebagai `Komputasi Numerik`.

## Sorting
Merupakan algoritma yang digunakan untuk mengurutkan data. Algoritma sorting biasanya digunakan untuk mengurutkan data sebelum dilakukan pencarian data.

### Jenis Sorting
Berikut adalah beberapa jenis sorting yang sering digunakan.
1. Bubble Sort
2. Insertion Sort
3. Selection Sort
4. Merge Sort
5. Quick Sort
6. Heap Sort
