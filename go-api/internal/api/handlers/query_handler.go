func QueryHandler(aiURL string, db *pgxpool.Pool) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var body struct {
            DatasetID string `json:"dataset_id"`
            Query     string `json:"query"`
        }
        json.NewDecoder(r.Body).Decode(&body)
        // fetch dataset object_url from DB
        datasetURL := "https://minio/..." // query DB

        // create query entry in DB (status queued)
        // call AI
        aiReq := services.AIRequest{DatasetURL: datasetURL, QueryText: body.Query, QueryID: "uuid"}
        resp, err := services.SendQueryToAI(aiURL, aiReq)
        if err != nil { http.Error(w, "ai error", 500); return }
        // pass through response to user
        w.Header().Set("Content-Type", "application/json")
        io.Copy(w, resp.Body)
    }
}
