package main

import (
	"fmt"
)

func main() {
	var width, height float64

	fmt.Println("=== PHẦN MỀM TÍNH TOÁN KÍCH THƯỚC CẮT CỬA NHÔM KÍNH ===")
	fmt.Print("Nhập chiều rộng phủ bì cửa (mm): ")
	fmt.Scanln(&width)
	fmt.Print("Nhập chiều cao phủ bì cửa (mm): ")
	fmt.Scanln(&height)

	if width <= 0 || height <= 0 {
		fmt.Println("Lỗi: Kích thước nhập vào phải lớn hơn 0!")
		return
	}

	// Công thức tính toán minh họa cho cửa đi 1 cánh mở quay hệ Xingfa 55
	trừKhungNgang := 100.0 // mm
	trừKhungDoc := 100.0   // mm
	kheHoNen := 10.0      // mm

	canhNgang := width - trừKhungNgang
	canhDoc := height - trừKhungDoc - kheHoNen
	kinhNgang := canhNgang - 160.0 // Khấu trừ cánh ngang ra kích thước kính
	kinhDoc := canhDoc - 160.0     // Khấu trừ cánh dọc ra kích thước kính

	fmt.Println("\n--- KẾT QUẢ TÍNH TOÁN KÍCH THƯỚC SẢN XUẤT ---")
	fmt.Printf("+ Chiều rộng cánh cửa: %.1f mm\n", canhNgang)
	fmt.Printf("+ Chiều cao cánh cửa:  %.1f mm\n", canhDoc)
	fmt.Printf("+ Kích thước kính cánh: %.1f x %.1f mm\n", kinhNgang, kinhDoc)
	fmt.Println("----------------------------------------------")
}
