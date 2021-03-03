package main
 
import (
        "encoding/json"
        "fmt"
        "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// 定义一个对象，继承合约对象
type Demo struct {
        contractapi.Contract
}

// 上链信息（对象）
type DemoInfo struct {
        Repo_url      string `json:"repo_url"`
        ProjectID     string `json:"projectID"`
        ProjectName   string `json:"projectName"`
        CommitID      string `json:"commitID"`
        Branch        string `json:"branch"`
        CommitDate    string `json:"commitDate"`
        CommitTime    string `json:"commitTime"`
}

// 查询结果
type QueryResult struct {
        Key    string `json:"Key"`
        Record *DemoInfo
}

// 初始化账本
func (s *Demo) InitLedger(ctx contractapi.TransactionContextInterface) error {
        DemoInfos := []DemoInfo{
                {
                        Repo_url: "https://39.108.179.25/sszdal/ipark_pda.git",
                        ProjectID: "2oij490j8",
                        ProjectName: "ipark_pda",
                        CommitID: "8bc47f",
                        Branch: "master",
                        CommitDate: "20210129",
                        CommitTime: "1626"
                },
                {
                        Repo_url: "https://39.108.179.25/sszdal/ipark_pda.git",
                        ProjectID: "2oij490j8",
                        ProjectName: "ipark_pda",
                        CommitID: "66a5d4",
                        Branch: "master",
                        CommitDate: "20200508",
                        CommitTime: "1525"},
                {
                        Repo_url: "https://39.108.179.25/sszdal/ipark_pda.git",
                        ProjectID: "2oij490j8",
                        ProjectName: "ipark_pda",
                        CommitID: "467d1e",
                        Branch: "master",
                        CommitDate: "20191217",
                        CommitTime: "1048"
                },
                {
                        Repo_url: "https://39.108.179.25/sszdal/ipark_pda.git",
                        ProjectID: "2oij490j8",
                        ProjectName: "ipark_pda",
                        CommitID: "cdcefe",
                        Branch: "master",
                        CommitDate: "20191216",
                        CommitTime: "1107"
                },
                {
                        Repo_url: "https://39.108.179.25/sszdal/ipark_pda.git",
                        ProjectID: "2oij490j8",
                        ProjectName: "ipark_pda",
                        CommitID: "013230",
                        Branch: "master",
                        CommitDate: "20191213",
                        CommitTime: "1636"
                },
                {
                        Repo_url: "https://39.108.179.25/sszdal/ipark_pda.git",
                        ProjectID: "2oij490j8",
                        ProjectName: "ipark_pda",
                        CommitID: "3a0e2e",
                        Branch: "master",
                        CommitDate: "20191212",
                        CommitTime: "1903"
                },
                {
                        Repo_url: "https://39.108.179.25/sszdal/ipark_pda.git",
                        ProjectID: "2oij490j8",
                        ProjectName: "ipark_pda",
                        CommitID: "510dfd",
                        Branch: "master",
                        CommitDate: "20191206",
                        CommitTime: "1935"
                },
                {
                        Repo_url: "https://39.108.179.25/sszdal/ipark_pda.git",
                        ProjectID: "2oij490j8",
                        ProjectName: "ipark_pda",
                        CommitID: "f01ac2",
                        Branch: "master",
                        CommitDate: "20191120",
                        CommitTime: "1617"
                },
                {
                        Repo_url: "https://39.108.179.25/sszdal/ipark_pda.git",
                        ProjectID: "2oij490j8",
                        ProjectName: "ipark_pda",
                        CommitID: "06543c",
                        Branch: "master",
                        CommitDate: "20191120",
                        CommitTime: "1514"
                },
                {
                        Repo_url: "https://39.108.179.25/sszdal/ipark_pda.git",
                        ProjectID: "2oij490j8",
                        ProjectName: "ipark_pda",
                        CommitID: "7e5983",
                        Branch: "master",
                        CommitDate: "20191015",
                        CommitTime: "1506"
                },
                {
                        Repo_url: "https://39.108.179.25/sszdal/ipark_pda.git",
                        ProjectID: "2oij490j8",
                        ProjectName: "ipark_pda",
                        CommitID: "77d831",
                        Branch: "master",
                        CommitDate: "20191012",
                        CommitTime: "1820"
                },
        }
        for _, DemoInfo := range DemoInfos {
                DemoInfoAsBytes, _ := json.Marshal(DemoInfo)
                err := ctx.GetStub().PutState(DemoInfo.Repo_url, DemoInfoAsBytes)
                if err != nil {
                        return fmt.Errorf("Failed to put to world state. %s", err.Error())
                }
        }
        return nil
}

// 上链信息
// func (s *Demo) CreateDemoInfo(ctx contractapi.TransactionContextInterface, repo_url string, projectID string, projectName string, commitID string, branch string) error {
//         DemoInfo := DemoInfo{
//                 Repo_url:      repo_url,
//                 ProjectID:     projectID,
//                 ProjectName:   projectName,
//                 CommitID:      commitID,
//                 Branch:        branch,
//                 CommitDate:    commitDate,
//                 CommitTime:    commitTime
//         }
//         DemoInfoAsBytes, _ := json.Marshal(DemoInfo)
//         return ctx.GetStub().PutState(DemoInfo.Repo_url, DemoInfoAsBytes)
// }

//查询信息
// func (s *Demo) QueryDemoInfo(ctx contractapi.TransactionContextInterface) (*DemoInfo, error) {
//         DemoInfoAsBytes, err := ctx.GetStub().GetState(DemoInfoRepo_url)
//         if err != nil {
//                 return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
//         }
//         demoInfo := new(DemoInfo)
//         //注意： Unmarshal(data []byte, v interface{})的第二个参数为指针类型（结构体地址）
//         err = json.Unmarshal(DemoInfoAsBytes, demoInfo) //demoInfo := new(DemoInfo)，demoInfo本身就是指针
//         if err != nil {
//                 return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
//         }
//         return demoInfo, nil
// }

func (s *Demo) QueryAllDemo(ctx contractapi.TransactionContextInterface) ([]QueryResult, error) {
	startKey := ""
	endKey := ""

	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

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

		demoInfo := new(DemoInfo)
		_ = json.Unmarshal(queryResponse.Value, demoInfo)

		queryResult := QueryResult{Key: queryResponse.Key, Record: demoInfo}
		results = append(results, queryResult)
	}

	return results, nil
}

func main() {
        chaincode, err := contractapi.NewChaincode(new(Demo))
        if err != nil {
                fmt.Printf("Error create fabDemoInfo chaincode: %s", err.Error())
                return
        }
        if err := chaincode.Start(); err != nil {
                fmt.Printf("Error starting fabDemoInfo chaincode: %s", err.Error())
        }
}