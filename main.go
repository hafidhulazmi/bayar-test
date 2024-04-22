package main

import (
	"errors"
	"fmt"
)

const (
	tax = 10
	app = 2000
)

func main() {
	fmt.Println(HitungHargaTotal(15000, 10000, 2))
	fmt.Println(PembayaranBarang(40000, "kredit", true))
}

func HitungHargaTotal(hargaItem float64, ongkir float64, qty int) (float64, error) {
	if hargaItem <= 0 {
		return 0, errors.New("harga barang tidak boleh nol")
	}
	if qty <= 0 {
		return 0, errors.New("jumlah barang tidak boleh nol")
	}
	if ongkir <= 0 {
		return 0, errors.New("ongkir tidak boleh nol")
	}

	hargaAkhirItem := hargaItem * float64(qty)
	hargaSetelahOngkir := hargaAkhirItem + ongkir
	pajak := hargaAkhirItem * tax / 100
	total := hargaSetelahOngkir + pajak + app

	return total, nil
}

func PembayaranBarang(hargaTotal float64, metodePembayaran string, dicicil bool) error {
	if hargaTotal <= 0 {
		return errors.New("harga total tidak bisa nol")
	}
	if metodePembayaran == "cod" || metodePembayaran == "transfer" || metodePembayaran == "debit" || metodePembayaran == "kredit" || metodePembayaran == "gerai" {

	} else {
		return errors.New("metode tidak dikenali")
	}
	if dicicil {
		if metodePembayaran != "kredit" {
			return errors.New("cicilan tidak memenuhi syarat(Metode pembayaran harus dicicil)")
		}
		if hargaTotal < 500000 {
			return errors.New("cicilan tidak memenuhi syarat(Total harga kurang dari 500.000)")
		}
	}
	return nil
}
