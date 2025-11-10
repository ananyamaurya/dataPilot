package config

import (
    "log"
    "os"
)

type Config struct {
    Port         string
    DBUrl        string
    MinioEndpoint string
    MinioKey     string
    MinioSecret  string
    MinioBucket  string
    JWTSecret    string
    AIServiceURL string
}

func Load() *Config {
    cfg := &Config{
        Port:          getenv("PORT", "8080"),
        DBUrl:         getenv("DATABASE_URL", "postgres://postgres:pass@postgres:5432/datapilot?sslmode=disable"),
        MinioEndpoint: getenv("MINIO_ENDPOINT", "minio:9000"),
        MinioKey:      getenv("MINIO_ACCESS_KEY", "minioadmin"),
        MinioSecret:   getenv("MINIO_SECRET_KEY", "minioadmin"),
        MinioBucket:   getenv("MINIO_BUCKET", "datasets"),
        JWTSecret:     getenv("JWT_SECRET", "change-me"),
        AIServiceURL:  getenv("AI_SERVICE_URL", "http://ai-service:8000"),
    }
    return cfg
}

func getenv(k, fallback string) string {
    if v := os.Getenv(k); v != "" {
        return v
    }
    return fallback
}
