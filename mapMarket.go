package main

import "fmt"

type Item struct {
	namaProduk  string
	hargaProduk int
	stokProduk  int
}

func main() {
	itemList := make(map[int]Item)
	itemList[1] = Item{namaProduk: "Coca-cola", hargaProduk: 1000, stokProduk: 8}
	itemList[2] = Item{namaProduk: "Fanta", hargaProduk: 1500, stokProduk: 18}
	itemList[3] = Item{namaProduk: "Sprite", hargaProduk: 1700, stokProduk: 9}
	itemList[4] = Item{namaProduk: "Pepsi", hargaProduk: 1500, stokProduk: 5}
	for _, value := range itemList {
		if value.stokProduk <= 10 {
			fmt.Println("Nama Barang :", value.namaProduk, "\nHarga Barang :", value.hargaProduk, "\nStok Produk :", value.stokProduk)
			fmt.Println("-----------------------")
		}
	}

}
