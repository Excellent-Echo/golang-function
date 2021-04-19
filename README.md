## belajar membuat fungsi (helper)

## latihan 1

buatlah sebuah fungsi untuk melakukan perhitungan biaya PPN dan pengeluaran saat belanja dari sebuah struk dengan pajak PPN sebesar 10% : 

```js
Total(110000)  => 10.000 , 100.000 
Total(220000)  => 20.000, 200.000 
Total(500000)  => 45.454 , 454.545 
Total(1000000)  => 90.909 , 909.090 
```

## latihan 2
buatlah sebuah fungsi konversi dari nominal menjadi nilai rupiah

```js
changeToIdr(100000) => IDR 100,000,00
changeToIdr(1200000) => IDR 1.200,000,00
changeToIdr(140350000) => IDR 140.350,000,00
```

## latihan 3

buatlah sebuah fungsi melakukan pengecekan dari sebuah banyaknya angka yang dimasukkan, apakah angka terbanyak merupakan angka ganjil atau angka genap

```js
GenapGanjil(1,2,3,4,5) => angka terbanyak adalah ganjil
GenapGanjil(4,2) => angka terbanyak genap
GenapGanjil(10,20,30,13) => terbanyak genap
GenapGanjil(30,13,13,77,33,55,17,13) => terbanyak ganjil
GenapGanjil() => tidak ada angka
```

## latihan 4
buatlah sebuah fungsi yang dapat menentukan apakah orang tersebut layak mendapat bantuan dari pemerintah dilihat dari jarak rumah kurang dari 100 meter, kemudian orang tersebut berkeluarga, dan berumur lebih dari 20 tahun. 

```go
var data = map[string]string{
    "name" : "andi",
    "umur" : "30",
    "jarakRumah" : "50",
    "berkeluarga" : "ya",
}

var data2 = map[string]string{
    "name" : "santi",
    "umur" : "19",
    "jarakRumah" : "80",
    "berkeluarga" : "ya",
}

var data3 = map[string]string{
    "name" : "budi",
    "umur" : "45",
    "jarakRumah" : "120",
    "berkeluarga" : "ya",
}

// output
// andi layak mendapat bantuan dari pemerintah
// santi tidak layak mendapat bantuan dari pemerintah
// budi tidak layak mendapat bantuan dari pemerintah
```

## latihan 5

buatlah sebuah fungsi untuk mengubah dari konversi angka menjadi tipe data string sepeti berikut :

```js
ChangeNumtoStr(100000) => "seratus ribu"
ChangeNumtoStr(111000) => "seratus sebelas ribu"
ChangeNumtoStr(5124000) => "lima juta seratus dua puluh empat ribu"
ChangeNumtoStr(1543) => "seribu lima ratus empat puluh tiga"
ChangeNumtoStr(1234678) => "satu juta dua ratus tiga puluh empat ribu enam ratus tujuh puluh delapan"
ChangeNumtoStr(0) => "nol"
```

## latihan 6

Terdapat struk belanja yang berisi nama produk dan total uang yang dibayar termasuk didalamnya ada PPN. Tentukan besarnya total PPN yang diterima oleh customer selama berbelanja di sebuah supermarket, PPN yang ditentukan adalah sebesar 10% dari besaran tiap produk yang dibeli.

```go
var belanja1 = map[string]int{
    "sepatu" : 1100000, 
    "jaket" : 2200000, 
    "topi" : 594000, 
    "celana" : 803000,
    "sweater" : 330000, 
    "kausKaki" : 110000,
    "sabuk" : 55000,
}
// output 
//  total PPN yang diterima sebesar 472.000
```



