package config

import (
	"os"
)

type Config struct {
	BotToken     string
	PdfPath      string
	MergePdfPath string
}

func Load() (*Config, error) {
	// err := godotenv.Load()
	// if err != nil {
	// 	return nil, err
	// }

	return &Config{
		BotToken:     os.Getenv("BOT_TOKEN"),
		PdfPath:      os.Getenv("PDF_PATH"),
		MergePdfPath: os.Getenv("PDF_MERGE"),
	}, nil
}
