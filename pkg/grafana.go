package pkg

import (
	"time"
)

//Grafana search response
type SearchResp struct {
	Type string `json:"type"`
	Target string `json:"target"`
}

//Gafana query response
type QueryResp struct {
	Target     string `json:"target"`
	Datapoints []details  `json:"datapoints"`
}

//store value and time in UnixNano
type details []interface{}

type RemainingAmt struct {
	Target string
	Datapoint []details
}

//memory storage
var AMT RemainingAmt

func init(){
	AMT.Target = "amount"
	//comment if need to use graph
	AMT.Datapoint = make([]details,1)
}

func InitRemainingAMT(){
	for {
		time.Sleep(3 * time.Second)
		remainingAMT := GetRemainingAmt()

		//AMT.Datapoint = append(AMT.Datapoint, details{remainingAMT, time.Now().UnixNano() / 1000000}) this for graph
		AMT.Datapoint[0]=details{remainingAMT, time.Now().UnixNano() / 1000000}
	}
}

func Query() []QueryResp {
	resp := make([]QueryResp, 1)
	resp[0] = QueryResp{Target: AMT.Target,Datapoints: AMT.Datapoint}
	return resp
}

func InitSearch() *SearchResp {
	return &SearchResp{Type: "timeseries",Target: "amount"}
}