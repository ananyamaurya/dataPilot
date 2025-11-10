package services
import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
    "time"
)

type AIRequest struct {
    DatasetURL string `json:"dataset_url"`
    QueryText  string `json:"query"`
    QueryID    string `json:"query_id"`
    UserID     string `json:"user_id"`
}

func SendQueryToAI(aiURL string, req AIRequest) (*http.Response, error) {
    client := &http.Client{Timeout: 2 * time.Minute}
    b, _ := json.Marshal(req)
    resp, err := client.Post(fmt.Sprintf("%s/analyze", aiURL), "application/json", bytes.NewReader(b))
    return resp, err
}
