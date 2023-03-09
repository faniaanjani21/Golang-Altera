// 1. Kode berikut ini dituliskan tanpa mengikuti kebiasaan penulisan yang disarankan.
//     - Berapa banyak kekurangan dalam penulisan kode tersebut?
//     - Bagian mana saja terjadi kekurangan tersebut?
//     - Tuliskan alasan dari setiap kekurangan tersebut. Alasan bisa diberikan dalam bentuk komentar pada kode yang disediakan berikut.

package main

import (
	"errors"
	"fmt"
)

// definisikan struct user
type user struct {
	id       int
	namauser string // username diubah menjadi namauser
	kunci    string // password diubah menjadi kunci
}

// definisikan struct userservice
type userservice struct {
	data []*user // variabel t diubah menjadi data dengan tipe []*user
}

// implementasikan method getallusers() untuk struct userservice
func (us *userservice) getallusers() []*user {
	return us.data
}

// implementasikan method getuserbyid() untuk struct userservice
func (us *userservice) getuserbyid(id int) (*user, error) {
	for _, u := range us.data {
		if id == u.id {
			return u, nil
		}
	}
	return nil, errors.New("user tidak ditemukan")
}

// implementasikan method adduser() untuk struct userservice
func (us *userservice) adduser(namauser, kunci string) *user {
	// cek apakah data user sudah terdaftar
	for _, u := range us.data {
		if u.namauser == namauser {
			return nil
		}
	}

	// buat user baru
	u := &user{
		id:       len(us.data) + 1,
		namauser: namauser,
		kunci:    kunci,
	}

	// tambahkan user baru ke slice data
	us.data = append(us.data, u)

	return u
}

// implementasikan method deleteuser() untuk struct userservice
func (us *userservice) deleteuser(id int) (bool, error) {
	for i, u := range us.data {
		if id == u.id {
			// hapus user dari slice data
			us.data = append(us.data[:i], us.data[i+1:]...)
			return true, nil
		}
	}
	return false, errors.New("user tidak ditemukan")
}

// fungsi main
func main() {
	// inisialisasi data awal
	data := []*user{
		{id: 1, namauser: "user1", kunci: "pass1"},
		{id: 2, namauser: "user2", kunci: "pass2"},
	}

	// inisialisasi userservice dengan data awal
	us := userservice{data}

	// contoh penggunaan method getallusers()
	allUsers := us.getallusers()
	fmt.Println("Daftar user:")
	for _, u := range allUsers {
		// cetak data user
		fmt.Printf("ID: %d, Nama User: %s, Kunci: %s\n", u.id, u.namauser, u.kunci)
	}

	// contoh penggunaan method getuserbyid()
	u, err := us.getuserbyid(1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("User dengan ID %d: %s\n", u.id, u.namauser)
	}

	// contoh penggunaan method adduser()
	u = us.adduser("user3", "pass3")
	if u == nil {
		fmt.Println("User dengan nama user3 sudah terdaftar")
	} else {
		fmt.Printf("User baru dengan ID %d dan nama user %s telah ditambahkan\n", u.id, u.namauser)
	}

	// contoh penggunaan
	// method deleteuser()
	status, err := us.deleteuser(2)
	if err != nil {
		fmt.Println(err)
	} else if status {
		fmt.Println("User dengan ID 2 berhasil dihapus")
	} else {
		fmt.Println("User dengan ID 2 tidak ditemukan")
	}
}
