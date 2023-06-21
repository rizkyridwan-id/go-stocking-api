package helpers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"stockingapi/app/configs"
	"time"
)

func ComputeSignature(apiKey, apiSecret, timestamp string) string {
	data := apiKey + apiSecret + timestamp
	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}

func ValidateSignature(signature, timestamp string, tolerance int) bool {
	parsedTime, err := time.Parse(time.RFC3339, timestamp)
	if err != nil {
		return false
	}
	fmt.Println("[DEBUG] (ValidateSignature) timestamp:", timestamp)
	fmt.Println("[DEBUG] (ValidateSignature) parsedTime:", parsedTime)

	currentTime := time.Now()
	fmt.Println("[DEBUG] (ValidateSignature) currentTime:", timestamp)

	timeDiff := currentTime.Sub(parsedTime).Seconds()
	fmt.Println("[DEBUG] (ValidateSignature) timeDiff:", timeDiff)

	if timeDiff < -float64(tolerance) || timeDiff > float64(tolerance) {
		fmt.Println("[DEBUG] (ValidateSignature) Expired Signature.")
		fmt.Println("[DEBUG] (ValidateSignature) - timeDiff:", timeDiff)
		return false
	}

	config := configs.LoadConfig()

	expectedSignature := ComputeSignature(config.API_KEY, config.SECRET_KEY, timestamp)

	return signature == expectedSignature
}
