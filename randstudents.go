package main
 
import (
        "encoding/json"
        "fmt"
        "github.com/hyperledger/fabric-contract-api-go/contractapi"
        "strconv"
	"time"
	"math/rand"
)

// 定义一个对象，继承合约对象
type Student struct {
        contractapi.Contract
}

// 上链信息（对象）
type StudentInfo struct {
        Id      string `json:"id"`
        Name    string `json:"name"`
        Age     string `json:"age"`
        College string `json:"college"`
        Major   string `json:"major"`
}

// 查询结果
type QueryResult struct {
        Key    string `json:"Key"`
        Record *StudentInfo
}

// 初始化账本
func (s *Student) InitLedger(ctx contractapi.TransactionContextInterface) error {
        StudentInfos := []StudentInfo{
                {Id: "1", Name: "张三", Age: "21", College: "机电", Major: "机械"},
                {Id: "2", Name: "李四", Age: "22", College: "计算", Major: "软件"},
		{Id: "3", Name: "王五", Age: "23", College: "材料", Major: "金属"},
		{Id: "4", Name: "赵六", Age: "23", College: "材料", Major: "金属"},
		{Id: "5", Name: "田七", Age: "23", College: "材料", Major: "金属"},
		{Id: "6", Name: "阿香", Age: "23", College: "材料", Major: "金属"},
		{Id: "4", Name: "阿凯", Age: "23", College: "材料", Major: "金属"},
		{Id: "8", Name: "阿兰", Age: "23", College: "材料", Major: "金属"},
		{Id: "9", Name: "旺财", Age: "23", College: "材料", Major: "金属"},
		}
        for _, StudentInfo := range StudentInfos {
                StudentInfoAsBytes, _ := json.Marshal(StudentInfo)
                err := ctx.GetStub().PutState(StudentInfo.Id, StudentInfoAsBytes)
                if err != nil {
                        return fmt.Errorf("Failed to put to world state. %s", err.Error())
                }
        }
        return nil
}

// 上链学生信息
func (s *Student) CreateStudentInfo(ctx contractapi.TransactionContextInterface, id string, name string, age string, college string, major string) error {
        StudentInfo := StudentInfo{
                Id:      id,
                Name:    name,
                Age:     age,
                College: college,
                Major:   major,
        }
        StudentInfoAsBytes, _ := json.Marshal(StudentInfo)
        return ctx.GetStub().PutState(StudentInfo.Id, StudentInfoAsBytes)
}

//随机写入学生信息
func (s *Student) RandCreateStudentInfo(ctx contractapi.TransactionContextInterface) error {
	rand.Seed(time.Now().Unix())
	id := strconv.Itoa(rand.Intn(100))
	name := strconv.Itoa(rand.Intn(9))
	age := strconv.Itoa(rand.Intn(9))
	college := strconv.Itoa(rand.Intn(9))
        major := strconv.Itoa(rand.Intn(9))
        StudentInfo := StudentInfo{
                Id:      id,
                Name:    name,
                Age:     age,
                College: college,
                Major:   major,
        }
        if id >= 10 {
                StudentInfoAsBytes, _ := json.Marshal(StudentInfo)
                return ctx.GetStub().PutState(StudentInfo.Id, StudentInfoAsBytes)
        }else {
                return error
        }
}

//随机查询学生信息
func (s *Student) RandQueryStudentInfo(ctx contractapi.TransactionContextInterface) (*StudentInfo, error) {
	rand.Seed(time.Now().Unix())
	StudentInfoId := strconv.Itoa(rand.Intn(8)+1)
	StudentInfoAsBytes, err := ctx.GetStub().GetState(StudentInfoId)
	if err != nil {
			return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}
	if StudentInfoAsBytes == nil {
			return nil, fmt.Errorf("%s does not exist", StudentInfoId)
	}
	stuInfo := new(StudentInfo)
	//注意： Unmarshal(data []byte, v interface{})的第二个参数为指针类型（结构体地址）
	err = json.Unmarshal(StudentInfoAsBytes, stuInfo) //stuInfo := new(StudentInfo)，stuInfo本身就是指针
	if err != nil {
			return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}
	return stuInfo, nil
}

//查询学生信息
func (s *Student) QueryStudentInfo(ctx contractapi.TransactionContextInterface, StudentInfoId string) (*StudentInfo, error) {
        StudentInfoAsBytes, err := ctx.GetStub().GetState(StudentInfoId)
        if err != nil {
                return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
        }
        if StudentInfoAsBytes == nil {
                return nil, fmt.Errorf("%s does not exist", StudentInfoId)
        }
        stuInfo := new(StudentInfo)
        //注意： Unmarshal(data []byte, v interface{})的第二个参数为指针类型（结构体地址）
        err = json.Unmarshal(StudentInfoAsBytes, stuInfo) //stuInfo := new(StudentInfo)，stuInfo本身就是指针
        if err != nil {
                return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
        }
        return stuInfo, nil
}
 
// 查询学生信息（查询的key末尾是数字，有对应的区间）
func (s *Student) QueryAllStudentInfos(ctx contractapi.TransactionContextInterface, startId, endId string) ([]QueryResult, error) {
        resultsIterator, err := ctx.GetStub().GetStateByRange(startId, endId)
        if err != nil {
                return nil, err
        }
        defer resultsIterator.Close()
        results := []QueryResult{}
        for resultsIterator.HasNext() {
                queryResponse, err := resultsIterator.Next()
 
                if err != nil {
                        return nil, err
                }
                StudentInfo := new(StudentInfo)
                _ = json.Unmarshal(queryResponse.Value, StudentInfo)
 
                queryResult := QueryResult{Key: queryResponse.Key, Record: StudentInfo}
                results = append(results, queryResult)
        }
        return results, nil
}

// 修改学生信息
func (s *Student) ChangeStudentInfo(ctx contractapi.TransactionContextInterface, id string, name string, age string, college string, major string) error {
        stuInfo, err := s.QueryStudentInfo(ctx, id)
        if err != nil {
                return err
        }
        stuInfo.Id = id
        stuInfo.Name = name
        stuInfo.Age = age
        stuInfo.College = college
        stuInfo.Major = major
        StudentInfoAsBytes, _ := json.Marshal(stuInfo)
        return ctx.GetStub().PutState(id, StudentInfoAsBytes)
}

func main() {
        chaincode, err := contractapi.NewChaincode(new(Student))
        if err != nil {
                fmt.Printf("Error create fabStudentInfo chaincode: %s", err.Error())
                return
        }
        if err := chaincode.Start(); err != nil {
                fmt.Printf("Error starting fabStudentInfo chaincode: %s", err.Error())
		}
}