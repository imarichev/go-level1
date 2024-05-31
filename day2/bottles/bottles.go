package bottles

import "fmt"

func Humanize(count int) {
	var ending string
	lastNumber := count % 10
	secondNumber := (count % 100) / 10

	if(secondNumber == 1) {
		ending = "ок"
	}else {
		if lastNumber == 1 {
			ending = "ка"
		}else if lastNumber > 1 && lastNumber < 5 {
			ending = "ки"
		}else{
			ending = "ок"
		}
	}
	

	fmt.Printf("%d Бутыл%s\n", count, ending)
}