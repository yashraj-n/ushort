package services

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log/slog"
	"math/big"
	"time"

	"github.com/yashraj-n/ushort/repository"
)

func CreateRedirectURL(ctx context.Context, original string, ipAddr string) (string, error) {
	hash := generateHash(original)
	config := &repository.Links{
		Redirect:  original,
		Hash:      hash,
		CreatedAt: time.Now(),
		IpAddr:    ipAddr,
	}

	jsonData, err := json.Marshal(config)
	if err != nil {
		slog.Error("Failed to serialize link data", "error", err)
		return "", err
	}

	err = repository.GetRedisInstance().Set(ctx, hash, string(jsonData), 0).Err()
	if err != nil {
		slog.Error("Failed to Create Redirect URL", "original", original, "hash", hash, "error", err)
		return "", err
	}
	return hash, nil
}
func GetRedirectURL(ctx context.Context, hash string) (string, error) {
	link, ok := repository.GetLruCache().Get(hash)
	if ok {
		return link.Redirect, nil
	}

	jsonData, err := repository.GetRedisInstance().Get(ctx, hash).Result()
	if err != nil {
		slog.Error("Failed to get redirect URL", "hash", hash, "error", err)
		return "", err
	}

	var linkData repository.Links
	if err := json.Unmarshal([]byte(jsonData), &linkData); err != nil {
		slog.Error("Failed to deserialize link data", "hash", hash, "error", err)
		return "", err
	}

	repository.GetLruCache().Add(hash, linkData)
	return linkData.Redirect, nil
}

func generateHash(url string) string {
	now := time.Now().UnixNano()

	randomLength, _ := rand.Int(rand.Reader, big.NewInt(3))
	length := int(randomLength.Int64()) + 3 // Results in lengths between 3-5

	seed := fmt.Sprintf("%d%x", now, url)
	seedLen := len(seed)

	charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	result := ""
	for i := 0; i < length; i++ {
		// Mix in the seed value to affect the random selection
		seedValue := int64(seed[i%seedLen])
		charIdx, _ := rand.Int(rand.Reader, big.NewInt(62))
		idx := (charIdx.Int64() + seedValue) % 62
		result += string(charSet[idx])
	}

	return result
}
