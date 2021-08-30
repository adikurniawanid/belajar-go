package main

import "fmt"

func logging() {
	fmt.Println("Selesai menjalankan function")
}

func endApp() {
	// recover untuk menangkap data panic, panic akan terhenti, program tetap berjalan
	message := recover()
	if message != nil {
		fmt.Println("Terjadi error", message)
	}
	fmt.Println("Aplikasi Selesai")
}

func runApp(value int) {
	// defer mirip __construct php
	// menjalankan walau fungsi runApp berhasil atau gagal
	defer logging()
	result := 10 / value
	fmt.Println("Run Application", result)
}

func runAppBoolean(error bool) {
	defer endApp()
	// panic untuk menghentikan program ketika error tetapi defer tetap dijalankan
	if error {
		panic("APLIKASI ERROR")
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp(2)
	runAppBoolean(false)
}
