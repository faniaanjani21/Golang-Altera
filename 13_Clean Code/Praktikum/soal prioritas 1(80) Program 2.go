// Kode berikut ini dituliskan tanpa mengikuti kebiasaan penulisan yang disarankan. Ubahlah penulisan kode berikut menjadi lebih terbaca dan lebih rapi!

package main

type Kendaraan struct {
	TotalRoda       int
	KecepatanPerJam int
}

type Mobil struct {
	Kendaraan
}

// Method Berjalan menambahkan kecepatan mobil sebanyak 10
func (m *Mobil) Berjalan() {
	m.TambahKecepatan(10)
}

// Method TambahKecepatan menambahkan kecepatan mobil sebesar kecepatanBaru ke kecepatanPerJam mobil
func (m *Mobil) TambahKecepatan(kecepatanBaru int) {
	m.KecepatanPerJam += kecepatanBaru
}

func main() {
	// Membuat instance Mobil bernama MobilCepat
	mobilCepat := Mobil{}
	mobilCepat.Berjalan()
	mobilCepat.Berjalan()
	mobilCepat.Berjalan()

	// Membuat instance Mobil bernama MobilLamban
	mobilLamban := Mobil{}
	mobilLamban.Berjalan()
}
