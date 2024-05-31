package main

import "fmt"

func main() {
	value1, value2, operation := readTask()
	var norm int
	array := [2]interface{}{value1, value2}
	for _, znach := range array {
		switch v := znach.(type) {
		case float64:
			norm += 1
		case int:
			fmt.Printf("value=%d: %T", v, v)
			break
		case bool:
			fmt.Printf("value=%t: %T", v, v)
			break
		case string:
			fmt.Printf("value=%s: %T", v, v)
			break
		}
	}
	if norm == 2 {
		switch operation {
		case "+":
			result := value1.(float64) + value2.(float64)
			fmt.Printf("%.4f", result)
		case "-":
			result := value1.(float64) - value2.(float64)
			fmt.Printf("%.4f", result)
		case "*":
			result := value1.(float64) * value2.(float64)
			fmt.Printf("%.4f", result)
		case "/":
			result := value1.(float64) / value2.(float64)
			fmt.Printf("%.4f", result)
		default:
			fmt.Println("Неизвестная операция")
		}
	}

}
func readTask() (back1, back2, back3 interface{}) {
	return 10.0, 2.0, "/"
}
