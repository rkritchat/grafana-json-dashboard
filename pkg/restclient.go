package pkg

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Resp struct {
	Amt float64 `json:"amt"`
}

func GetRemainingAmt() float64 {
	if r, err := http.Get(os.Getenv("target_url")); err!= nil {
		log.Printf("Exception while calling api : %v\n", err)
	}else{
		var resp Resp
		if err := json.NewDecoder(r.Body).Decode(&resp); err!= nil {
			log.Printf("Exception while decode json : %v\n", err)
		}
		return resp.Amt
	}
	return 0
}