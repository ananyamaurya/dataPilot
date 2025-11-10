func UploadHandler(minioClient *storage.MinioClient, db *pgxpool.Pool) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        r.ParseMultipartForm(50 << 20) // 50MB
        file, header, err := r.FormFile("file")
        if err != nil {
            http.Error(w, "file missing", 400); return
        }
        defer file.Close()
        // basic validation
        if !strings.HasSuffix(header.Filename, ".csv") {
            http.Error(w, "only CSV allowed", 400); return
        }
        objectKey := fmt.Sprintf("datasets/%s", header.Filename) // use uuid in real
        url, err := minioClient.Upload(r.Context(), objectKey, file, header.Size, header.Header.Get("Content-Type"))
        if err != nil {
            http.Error(w, "upload failed", 500); return
        }
        // insert metadata into DB...
        json.NewEncoder(w).Encode(map[string]string{"dataset_id": "uuid", "url": url})
    }
}
