func NewRouter(cfg *config.Config, db *pgxpool.Pool, minioClient *storage.MinioClient) *mux.Router {
    r := mux.NewRouter()
    // public
    r.HandleFunc("/healthz", healthHandler).Methods("GET")
    r.HandleFunc("/api/v1/auth/register", auth.RegisterHandler).Methods("POST")
    r.HandleFunc("/api/v1/auth/login", auth.LoginHandler).Methods("POST")

    // protected
    api := r.PathPrefix("/api/v1").Subrouter()
    api.Use(middlewares.JWTAuthMiddleware(cfg.JWTSecret))
    api.HandleFunc("/datasets", handlers.UploadHandler(minioClient, db)).Methods("POST")
    api.HandleFunc("/queries", handlers.QueryHandler(cfg.AIServiceURL, db)).Methods("POST")
    // more routes...
    return r
}
