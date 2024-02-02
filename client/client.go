package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/leenzstra/godexscreener/types"
)

const defaultBaseUrl = "https://api.dexscreener.com/latest"

type DexscreenerClient struct {
	baseUrl string
	client  *http.Client
}

// New Dexcreener http client with options
func NewClient(client *http.Client, opts ...ClientOpt) *DexscreenerClient {
	d := &DexscreenerClient{
		client: client,
		baseUrl: defaultBaseUrl,
	}

	for _, opt := range opts {
		opt(d)
	}

	return d

}

// get request wrapper
func (d *DexscreenerClient) get(url string, result interface{}) error {
	resp, err := d.client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, result)
	if err != nil {
		return err
	}

	return nil
}

// Get one or multiple pairs by chain and pair address
// 
// One or multiple, comma-separated pair addresses (up to 30 addresses).
// Example: 0xAbc1 or 0xAbc1,0xAbc2,0xAbc3
func (d *DexscreenerClient) Pairs(chainId string, pairAddrs []string) (*types.PairsResponse, error) {
	url, err := url.JoinPath(d.baseUrl, "dex", "pairs", chainId, strings.Join(pairAddrs, ","))
	if err != nil {
		return nil, err
	}

	pairs := &types.PairsResponse{}
	if err := d.get(url, pairs); err != nil {
		return nil, err
	}

	return pairs, nil
}

// Get one or multiple pairs by token address.
// 
// One or multiple, comma-separated token addresses (up to 30 addresses).
// Example: 0xAbc1 or 0xAbc1,0xAbc2,0xAbc3
func (d *DexscreenerClient) Tokens(tokenAddrs []string) (*types.PairsResponse, error) {
	url, err := url.JoinPath(d.baseUrl, "dex", "tokens", strings.Join(tokenAddrs, ","))
	if err != nil {
		return nil, err
	}

	pairs := &types.PairsResponse{}
	if err := d.get(url, pairs); err != nil {
		return nil, err
	}

	return pairs, nil
}

// Search for pairs matching query. Query may include pair address, token address, token name or token symbol.
// 
// Example: WBTC or WBTC/USDC or 0xAbc01
func (d *DexscreenerClient) Search(query string) (*types.PairsResponse, error) {
	// query must have at least 2 chars
	if len(query) < 2 {
		return nil, fmt.Errorf("query must have at least 2 chars")
	}

	u, err := url.JoinPath(d.baseUrl, "dex", "search")
	if err != nil {
		return nil, err
	}

	url, err := url.Parse(u)
	if err != nil {
		return nil, err
	}

	// add ?q parameter
	q := url.Query()
	q.Add("q", query)
	url.RawQuery = q.Encode()

	pairs := &types.PairsResponse{}
	if err := d.get(url.String(), pairs); err != nil {
		return nil, err
	}

	return pairs, nil
}


