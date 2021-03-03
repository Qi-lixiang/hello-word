package main
 
import (
        "encoding/json"
        "fmt"
        "github.com/hyperledger/fabric-contract-api-go/contractapi"
        "strconv"
)

type Demo struct {
	contractapi.Contract
}

// 上链信息（对象）
type DemoInfo struct {
		Number		  string `json:"number"`
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


func (d *Demo) Init(ctx contractapi.TransactionContextInterface) error {
	DemoInfos := []DemoInfo{
                {
						Number:	"1",
                        Repo_url: "https://39.108.179.25/sszdal/ipark_pda.git",
                        ProjectID: "2oij490j8",
                        ProjectName: "ipark_pda",
                        CommitID: "8bc47f",
                        Branch: "master",
                        CommitDate: "20210129",
                        CommitTime: "1626"},{
						Number:	"2",
                        Repo_url: "https://39.108.179.25/sszdal/ipark_pda.git",
                        ProjectID: "2oij490j8",
                        ProjectName: "ipark_pda",
                        CommitID: "66a5d4",
                        Branch: "master",
                        CommitDate: "20200508",
                        CommitTime: "1525"},{
						Number:	"3",
                        Repo_url: "https://39.108.179.25/sszdal/ipark_pda.git",
                        ProjectID: "2oij490j8",
                        ProjectName: "ipark_pda",
                        CommitID: "467d1e",
                        Branch: "master",
                        CommitDate: "20191217",
                        CommitTime: "1048"},{
						Number:	"4",
                        Repo_url: "https://39.108.179.25/sszdal/ipark_pda.git",
                        ProjectID: "2oij490j8",
                        ProjectName: "ipark_pda",
                        CommitID: "cdcefe",
                        Branch: "master",
                        CommitDate: "20191216",
                        CommitTime: "1107"},{
						Number:	"5",
                        Repo_url: "https://39.108.179.25/sszdal/ipark_pda.git",
                        ProjectID: "2oij490j8",
                        ProjectName: "ipark_pda",
                        CommitID: "013230",
                        Branch: "master",
                        CommitDate: "20191213",
                        CommitTime: "1636"},{
						Number:	"6",
                        Repo_url: "https://39.108.179.25/sszdal/ipark_pda.git",
                        ProjectID: "2oij490j8",
                        ProjectName: "ipark_pda",
                        CommitID: "3a0e2e",
                        Branch: "master",
                        CommitDate: "20191212",
                        CommitTime: "1903"},{
						Number:	"7",
                        Repo_url: "https://39.108.179.25/sszdal/ipark_pda.git",
                        ProjectID: "2oij490j8",
                        ProjectName: "ipark_pda",
                        CommitID: "510dfd",
                        Branch: "master",
                        CommitDate: "20191206",
                        CommitTime: "1935"},{
						Number:	"8",
                        Repo_url: "https://39.108.179.25/sszdal/ipark_pda.git",
                        ProjectID: "2oij490j8",
                        ProjectName: "ipark_pda",
                        CommitID: "f01ac2",
                        Branch: "master",
                        CommitDate: "20191120",
						CommitTime: "1617"},{
						Number:	"9",
                        Repo_url: "https://39.108.179.25/sszdal/ipark_pda.git",
                        ProjectID: "2oij490j8",
                        ProjectName: "ipark_pda",
                        CommitID: "06543c",
                        Branch: "master",
                        CommitDate: "20191120",
                        CommitTime: "1514"},{
						Number:	"10",
                        Repo_url: "https://39.108.179.25/sszdal/ipark_pda.git",
                        ProjectID: "2oij490j8",
                        ProjectName: "ipark_pda",
                        CommitID: "7e5983",
                        Branch: "master",
                        CommitDate: "20191015",
                        CommitTime: "1506"},{
						Number:	"11",
                        Repo_url: "https://39.108.179.25/sszdal/ipark_pda.git",
                        ProjectID: "2oij490j8",
                        ProjectName: "ipark_pda",
                        CommitID: "77d831",
                        Branch: "master",
                        CommitDate: "20191012",
                        CommitTime: "1820"},
        }
		for _, DemoInfo := range DemoInfos {
                DemoInfoAsBytes, _ := json.Marshal(DemoInfo)
                err := ctx.GetStub().PutState(DemoInfo.Number, DemoInfoAsBytes)
                if err != nil {
                        return fmt.Errorf("Failed to put to world state. %s", err.Error())
                }
        }
		return nil

}

func (d *Demo) Query(ctx contractapi.TransactionContextInterface) (*DemoInfo, error) {
	startId := ""
	endId := ""
	resultsIterator, err := ctx.GetStub().GetStateByRange(startId, endId)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()
	DemoInfos := []DemoInfo{}
	lsDate := []DemoInfo{}
        for resultsIterator.HasNext() {
                queryResponse, err := resultsIterator.Next()

                if err != nil {
                        return nil, err
                }
                DemoInfo := new(DemoInfo)
                _ = json.Unmarshal(queryResponse.Value, DemoInfo)
                DemoInfos = append(DemoInfos, *DemoInfo)
        }

		maxDate := 0
		maxTime := 0
        for _, DemoInfo := range DemoInfos {
			num, _ := strconv.Atoi(DemoInfo.CommitDate)
			if maxDate < num {
				maxDate = num
			}
        }
		for _, DemoInfo := range DemoInfos {
			num, _ := strconv.Atoi(DemoInfo.CommitDate)
			if maxDate == num {
				lsDate = append(lsDate, DemoInfo)
			}
			// lsDate = append(lsDate, DemoInfo.CommitDate)
        }
		if len(lsDate) !=1 {
			// fmt.Println(len(lsDate))
			for _, info := range lsDate {
				num, _ := strconv.Atoi(info.CommitTime)
				if maxTime < num {
					maxTime = num
				}
			}
			for _, info := range lsDate {
				num, _ := strconv.Atoi(info.CommitTime)
				if maxTime == num {
					return &info, nil
				}
			}
			
		}
		return &lsDate[0], nil
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
