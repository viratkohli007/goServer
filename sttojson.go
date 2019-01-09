package main
import ("fmt"
        "encoding/json"
        "net/http"
        "log"
)

type Emp struct{
	Name string `json: "Empname"`
	Num int `json: "Empid"`
}

func main() {
	

	http.HandleFunc("/", jarvis)
	err := http.ListenAndServe(":8080", nil)
	if err != nil{
		fmt.Println(err)
	}
}
func jarvis(w http.ResponseWriter, r *http.Request) {
	
	emp := &Emp{Name: "James Bond", Num: 007}
	e , err := json.Marshal(emp)
	if err != nil{
		log.Fatal(err)
	}
	w.Write(e)
	fmt.Printf("%T", e)
}