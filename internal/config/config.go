package config

import (
	"log"  // logs errors and fatal messages
	"os"   //reads env variables
	"time" // handles time durations

	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	MongoURI   string
	MongoDB    string
	JWTSecret  string
	AccessTTL  time.Duration
	RefreshTTL time.Duration
	CORSOrigin string
	Env        string
}

func Load() *Config {
	_ = godotenv.Load()
	acc, err := time.ParseDuration(getEnv("ACCESS_TTL", "15m"))
	if err != nil {
		log.Fatalf("invalid ACCESS_TTL")
	}
	ref, err := time.ParseDuration(getEnv("REFRESH_TTL", "168h"))
	if err != nil {
		log.Fatalf("invalid REFRESH_TTL")
	}

	return &Config{
		Port:       getEnv("PORT", "8080"),
		MongoURI:   mustEnv("MONGO_URI"),
		MongoDB:    getEnv("MONGO_DB", "todo_db"),
		JWTSecret:  mustEnv("JWT_SECRET"),
		AccessTTL:  acc,
		RefreshTTL: ref,
		CORSOrigin: getEnv("CORS_ORIGIN", "http://localhost:3000"),
		Env:        getEnv("ENV", "dev"),
	}
}

func mustEnv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("missing env %s", k)
	}
	return v
}
func getEnv(k, d string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return d
}
