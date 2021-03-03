package main
 
import (
    "fmt"
    // "math/rand"
	// "time"
	// "strconv"
	"encoding/json"
)

type StudentInfo struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Age     string `json:"age"`
	College string `json:"college"`
	Major   string `json:"major"`
}

func main() {

	// rand.Seed(time.Now().Unix())
	// id := strconv.Itoa(rand.Intn(100))
	// name := strconv.Itoa(rand.Intn(9))
	// age := strconv.Itoa(rand.Intn(9))
	// college := strconv.Itoa(rand.Intn(9))
	// major := strconv.Itoa(rand.Intn(9))
	// fmt.Println(id,name,age,college,major)

    StudentInfos := []StudentInfo{
		{Id: "1", Name: "张三", Age: "21", College: "机电", Major: "机械"},
		{Id: "2", Name: "李四", Age: "22", College: "计算", Major: "软件"},
		{Id: "3", Name: "王五", Age: "23", College: "材料", Major: "金属"},
	}
	for Num, Student := range StudentInfos {
		fmt.Println("Num:", Num, "Stident:", Student, "\n")
		StudentInfoAsBytes, _ := json.Marshal(Student)
		// fmt.Printf("StudentInfoAsBytes:%s", StudentInfoAsBytes)
		fmt.Println("编码后的Student:", StudentInfoAsBytes)

		stuInfo := new(StudentInfo)
		err := json.Unmarshal(StudentInfoAsBytes, stuInfo)
		if err != nil {
			fmt.Printf("Failed to read from world state. %s", err.Error())
		}
		fmt.Println("解码后的Student:", *stuInfo) 
	}
	// fmt.Println(StudentInfos)
}