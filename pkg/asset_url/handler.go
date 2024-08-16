package asset_url

import (
	"errors"
	"fmt"
	"net/url"
	"path"
)

// GenerateBlockChainLogoURL dynamically constructs the URL based on the blockchain name
func GenerateBlockChainLogoURL(blockchain string) (string, error) {
	// Check if the blockchain name is empty
	if blockchain == "" {
		return "", errors.New("blockchain name cannot be empty")
	}

	baseURL := "https://raw.githubusercontent.com/0xbase-Corp/assets/master/blockchains/"
	logoPath := "/info/logo.png"

	// Parse the base URL to ensure it's valid
	u, err := url.Parse(baseURL)
	if err != nil {
		return "", fmt.Errorf("invalid base URL: %v", err)
	}

	// Add blockchain to the path and then add the logoPath
	u.Path = path.Join(u.Path, blockchain, logoPath)

	return u.String(), nil
}

// GenerateAssetLogoURL dynamically constructs the URL based on the blockchain and asset ID
func GenerateAssetLogoURL(blockchain, id string) (string, error) {
	if blockchain == "" && id == "" {
		return "", errors.New("argument missing")
	}

	// Check if the blockchain name or asset ID is empty
	if blockchain == "" {
		return "", errors.New("blockchain name cannot be empty")
	}
	if id == "" {
		return "", errors.New("asset ID cannot be empty")
	}

	baseURL := "https://raw.githubusercontent.com/0xbase-Corp/assets/master/blockchains/"
	assetsPath := path.Join(blockchain, "assets", id, "logo.png")

	// Parse the base URL to ensure it's valid
	u, err := url.Parse(baseURL)
	if err != nil {
		return "", fmt.Errorf("invalid base URL: %v", err)
	}

	// Append the assets path to the base URL
	u.Path = path.Join(u.Path, assetsPath)

	return u.String(), nil
}
