package finnhub

import  (
	finnhub_go "github.com/Finnhub-Stock-API/finnhub-go/v2"
)

type finnhubSession struct {
  client *finnhub_go.APIClient
  symbols []string
}

func NewFinnhubSession(apiKey string) *finnhubSession {
  fs := &finnhubSession{}
  cfg := finnhub_go.NewConfiguration()
  cfg.AddDefaultHeader("X-Finnhub-Token", apiKey)
  fs.client = finnhub_go.NewAPIClient(cfg)
  return fs
}
