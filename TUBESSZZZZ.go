package main

import (
	"fmt"
)

// Define a struct to hold data for a user
type data struct {
	nama     string // The name of the user
	username string // The username of the user
	password string // The password of the user
	role     string // The role of the user (admin or student)
}

// Define a struct to hold an array of data and the number of elements
type arrData struct {
	info  [10]data // An array of user data
	nData int      // The number of elements in the array
}

// Define a struct to hold data for a student
type mahasiswa struct {
	nama    string // The name of the student
	nohp    int    // The phone number of the student
	nilai   int    // The score of the student
	jurusan string // The major of the student
	status  string // The status of the student (accepted/rejected/pending)
}

// Define a struct to hold an array of students and the number of elements
type arrMhs struct {
	info [10]mahasiswa // An array of student data
	nMhs int           // The number of elements in the array
}

func main() {
	fmt.Println("*** ------------------------------------------- ***")
	fmt.Println("***       Aplikasi pendaftaran mahasiswa        ***")
	fmt.Println("***         Created by ndul and payen           ***")
	fmt.Println("***        Algoritma Pemrograman 2024           ***")
	fmt.Println("*** ------------------------------------------- ***")
	var users arrData
	var students arrMhs
	menu(&users, &students)
}

// Function to display the main menu
func menu(users *arrData, students *arrMhs) {
	var pilihan int
	fmt.Println("*** Menu Utama ***")
	fmt.Println("1. Registrasi")
	fmt.Println("2. Login")
	fmt.Println("3. Keluar")

	fmt.Scan(&pilihan)
	for pilihan != 3 {
		switch pilihan {
		case 1:
			fmt.Println("*** Menu Registrasi ***")
			registrasi(users)
		case 2:
			fmt.Println("*** Menu Login ***")
			login(users, students)
		case 3:
			fmt.Println("*** Keluar ***")
			return
		default:
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}

		fmt.Println("*** Menu Utama ***")
		fmt.Println("1. Registrasi")
		fmt.Println("2. Login")
		fmt.Println("3. Keluar")
		fmt.Scan(&pilihan)
	}
}

// Function to handle user registration
func registrasi(users *arrData) {
	var user data
	var pil int
	fmt.Print("Masukkan Nama: ")
	fmt.Scan(&user.nama)
	fmt.Print("Masukkan Username: ")
	fmt.Scan(&user.username)
	fmt.Print("Masukkan Password: ")
	fmt.Scan(&user.password)
	fmt.Print("Masukkan Role (admin/student): ")
	fmt.Scan(&user.role)

	fmt.Println("1. Simpan, 2. Batal: ")
	fmt.Scan(&pil)
	switch pil {
	case 1:
		users.info[users.nData] = user
		users.nData++
		fmt.Println("Registrasi berhasil :)")
	case 2:
		fmt.Println("Registrasi dibatalkan :(")
	}
}

// Function to handle user login
func login(users *arrData, students *arrMhs) {
	var username, password string
	var found bool

	fmt.Print("Username: ")
	fmt.Scan(&username)
	fmt.Print("Password: ")
	fmt.Scan(&password)

	for i := 0; i < users.nData; i++ {
		if users.info[i].username == username && users.info[i].password == password {
			fmt.Println("Login berhasil!")
			if users.info[i].role == "admin" {
				adminMenu(students)
			} else {
				studentMenu(students)
			}
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Username atau password salah.")
	}
}

// Function to display admin menu
func adminMenu(students *arrMhs) {
	var pilihan int
	fmt.Println("*** Menu Admin ***")
	fmt.Println("1. Tambah Mahasiswa")
	fmt.Println("2. Edit Mahasiswa")
	fmt.Println("3. Hapus Mahasiswa")
	fmt.Println("4. Tambah/Edit Nilai")
	fmt.Println("5. Hapus Nilai")
	fmt.Println("6. Tampilkan Data Mahasiswa Berdasarkan Jurusan")
	fmt.Println("7. Tampilkan Data Mahasiswa Diterima/Ditolak")
	fmt.Println("8. Tampilkan Data Mahasiswa Terurut")
	fmt.Println("9. Kembali ke Menu Utama")

	fmt.Scan(&pilihan)
	for pilihan != 9 {
		switch pilihan {
		case 1:
			tambahMahasiswa(students)
		case 2:
			editMahasiswa(students)
		case 3:
			hapusMahasiswa(students)
		case 4:
			tambahEditNilai(students)
		case 5:
			hapusNilai(students)
		case 6:
			tampilkanMahasiswaJurusan(students)
		case 7:
			tampilkanMahasiswaStatus(students)
		case 8:
			tampilkanMahasiswaTerurut(students)
		case 9:
			return
		default:
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}

		fmt.Println("*** Menu Admin ***")
		fmt.Println("1. Tambah Mahasiswa")
		fmt.Println("2. Edit Mahasiswa")
		fmt.Println("3. Hapus Mahasiswa")
		fmt.Println("4. Tambah/Edit Nilai")
		fmt.Println("5. Hapus Nilai")
		fmt.Println("6. Tampilkan Data Mahasiswa Berdasarkan Jurusan")
		fmt.Println("7. Tampilkan Data Mahasiswa Diterima/Ditolak")
		fmt.Println("8. Tampilkan Data Mahasiswa Terurut")
		fmt.Println("9. Kembali ke Menu Utama")
		fmt.Scan(&pilihan)
	}
}

// Function to display student menu
func studentMenu(students *arrMhs) {
	var pilihan int
	fmt.Println("*** Menu Mahasiswa ***")
	fmt.Println("1. Tampilkan Data Mahasiswa Berdasarkan Jurusan")
	fmt.Println("2. Tampilkan Data Mahasiswa Diterima/Ditolak")
	fmt.Println("3. Tampilkan Data Mahasiswa Terurut")
	fmt.Println("4. Kembali ke Menu Utama")

	fmt.Scan(&pilihan)
	for pilihan != 4 {
		switch pilihan {
		case 1:
			tampilkanMahasiswaJurusan(students)
		case 2:
			tampilkanMahasiswaStatus(students)
		case 3:
			tampilkanMahasiswaTerurut(students)
		case 4:
			return
		default:
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}

		fmt.Println("*** Menu Mahasiswa ***")
		fmt.Println("1. Tampilkan Data Mahasiswa Berdasarkan Jurusan")
		fmt.Println("2. Tampilkan Data Mahasiswa Diterima/Ditolak")
		fmt.Println("3. Tampilkan Data Mahasiswa Terurut")
		fmt.Println("4. Kembali ke Menu Utama")
		fmt.Scan(&pilihan)
	}
}

// Function to add a student
func tambahMahasiswa(students *arrMhs) {
	fmt.Println("*** Tambah Mahasiswa ***")
	fmt.Print("Masukkan Nama: ")
	fmt.Scan(&students.info[students.nMhs].nama)
	fmt.Print("Masukkan Nomor HP : +62-")

	for {
		var nohp int
		_, err := fmt.Scan(&nohp)
		length := len(fmt.Sprintf("%d", nohp))
		if err == nil && length >= 8 && length <= 13 {
			students.info[students.nMhs].nohp = nohp
			break
		}
		fmt.Println("Nomor HP tidak valid, coba lagi (minimal 8 angka dan maksimal 13 angka): ")
	}

	fmt.Print("Masukkan Jurusan (Gunakan underscore ( _ ) untuk pengganti spasi ): ")
	fmt.Scan(&students.info[students.nMhs].jurusan)
	students.info[students.nMhs].status = "pending"
	students.nMhs++
	fmt.Println("Mahasiswa berhasil ditambahkan.")
}

// Function to edit student data
func editMahasiswa(students *arrMhs) {
	var nama string
	fmt.Print("Nama Mahasiswa yang akan diubah: ")
	fmt.Scan(&nama)

	index := cariMahasiswa(nama, *students)
	if index == -1 {
		fmt.Println("Mahasiswa tidak ditemukan.")
		return
	}

	fmt.Print("Nama Baru: ")
	fmt.Scan(&students.info[index].nama)
	fmt.Print("Nomor HP Baru: ")

	for {
		var nohp int
		_, err := fmt.Scan(&nohp)
		length := len(fmt.Sprintf("%d", nohp))
		if err == nil && length >= 8 && length <= 13 {
			students.info[index].nohp = nohp
			break
		}
		fmt.Println("Nomor HP tidak valid, coba lagi (minimal 8 angka dan maksimal 13 angka): ")
	}

	fmt.Print("Jurusan Baru: ")
	fmt.Scan(&students.info[index].jurusan)
	fmt.Println("Data mahasiswa berhasil diubah.")
}

// Function to delete a student
func hapusMahasiswa(students *arrMhs) {
	var nama string
	fmt.Print("Nama Mahasiswa yang akan dihapus: ")
	fmt.Scan(&nama)

	index := cariMahasiswa(nama, *students)
	if index == -1 {
		fmt.Println("Mahasiswa tidak ditemukan.")
		return
	}

	students.info[index] = students.info[students.nMhs-1]
	students.nMhs--
	fmt.Println("Mahasiswa berhasil dihapus.")
}

// Function to add or edit student scores
func tambahEditNilai(students *arrMhs) {
	var nama string
	var nilai int
	fmt.Print("Nama Mahasiswa: ")
	fmt.Scan(&nama)

	index := cariMahasiswa(nama, *students)
	if index == -1 {
		fmt.Println("Mahasiswa tidak ditemukan.")
		return
	}

	fmt.Print("Nilai: ")
	fmt.Scan(&nilai)
	students.info[index].nilai = nilai
	if nilai >= 75 && nilai <= 100 {
		students.info[index].status = "accepted"
	} else if nilai <= 74 && nilai >= 50 {
		students.info[index].status = "pending"
	} else {
		students.info[index].status = "rejected"
	}
	fmt.Println("Nilai mahasiswa berhasil ditambahkan/diedit.")
}

// Function to delete student scores
func hapusNilai(students *arrMhs) {
	var nama string
	fmt.Print("Nama Mahasiswa: ")
	fmt.Scan(&nama)

	index := cariMahasiswa(nama, *students)
	if index == -1 {
		fmt.Println("Mahasiswa tidak ditemukan.")
		return
	}

	students.info[index].nilai = 0
	fmt.Println("Nilai mahasiswa berhasil dihapus.")
}

// Function to display students by major
func tampilkanMahasiswaJurusan(students *arrMhs) {
	var jurusan string
	fmt.Print("Masukkan Jurusan: ")
	fmt.Scan(&jurusan)

	fmt.Printf("Mahasiswa di jurusan %s:\n", jurusan)
	for i := 0; i < students.nMhs; i++ {
		if students.info[i].jurusan == jurusan {
			fmt.Printf("Nama: %s, No. HP: %d, Nilai: %d, Status: %s\n", students.info[i].nama, students.info[i].nohp, students.info[i].nilai, students.info[i].status)
		}
	}
}

// Function to display students by status
func tampilkanMahasiswaStatus(students *arrMhs) {
	var status string
	fmt.Print("Masukkan Status (accepted/rejected/pending): ")
	fmt.Scan(&status)

	fmt.Printf("Mahasiswa dengan status %s:\n", status)
	for i := 0; i < students.nMhs; i++ {
		if students.info[i].status == status {
			fmt.Printf("Nama: %s, No. HP: %d, Nilai: %d, Jurusan: %s\n", students.info[i].nama, students.info[i].nohp, students.info[i].nilai, students.info[i].jurusan)
		}
	}
}

// Function to display sorted student data
func tampilkanMahasiswaTerurut(students *arrMhs) {
	fmt.Println("Data mahasiswa terurut berdasarkan nilai:")
	for i := 0; i < students.nMhs-1; i++ {
		for j := 0; j < students.nMhs-i-1; j++ {
			if students.info[j].nilai < students.info[j+1].nilai {
				students.info[j], students.info[j+1] = students.info[j+1], students.info[j]
			}
		}
	}

	for i := 0; i < students.nMhs; i++ {
		fmt.Printf("Nama: %s, No. HP: %d, Nilai: %d, Jurusan: %s, Status: %s\n", students.info[i].nama, students.info[i].nohp, students.info[i].nilai, students.info[i].jurusan, students.info[i].status)
	}
}

// Function to search for a student by name
func cariMahasiswa(nama string, students arrMhs) int {
	for i := 0; i < students.nMhs; i++ {
		if students.info[i].nama == nama {
			return i
		}
	}
	return -1
}
