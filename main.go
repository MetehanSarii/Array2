package main

import "fmt"

func arrayDemo3() {
	sayilar := [5]int{20, 30, 10, 12, 5}
	enBuyuk := sayilar[0]
	//lenght = Dizinin eleman Sayısı
	for i := 0; i < len(sayilar); i++ {
		if sayilar[i] > enBuyuk {
			enBuyuk = sayilar[i]
		}
	}
	fmt.Println("En Büyük: ", enBuyuk)
}
func main() {
	arrayDemo3()
}
