package client

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// test client constuctor
func TestNewClient(t *testing.T) {
	type test struct {
		opts []ClientOpt
		want string
	}

	tests := []test{
		{opts: []ClientOpt{WithBaseUrl("https://api.dexscreener.com/latest")}, want: "https://api.dexscreener.com/latest"},
		{opts: []ClientOpt{}, want: "https://api.dexscreener.com/latest"},
	}

	c := &http.Client{}

	for _, tt := range tests {
		got := NewClient(c, tt.opts...)

		assert.NotNil(t, got)
		assert.NotNil(t, got.client)
		assert.Equal(t, tt.want, got.baseUrl)
	}
}

// test pairs request
func TestPairs(t *testing.T) {
	type test struct {
		chainId   string
		pairAddrs []string
		err       bool
		want      string
		desc      string
	}

	tests := []test{
		{
			chainId:   "solana",
			pairAddrs: []string{"2QrWsSWrGvoAqkDC5XSGqjS752RWLaopqAaGrbugSxBL"},
			err:       false,
			desc:      "jitosol/jup pair",
		},
		{
			chainId:   "",
			pairAddrs: []string{"0xA43fe16908251ee70EF74718545e4FE6C5cCEc9f"},
			err:       true,
			desc:      "no chain id",
		},
		{
			chainId:   "bsc",
			pairAddrs: []string{},
			err:       true,
			desc:      "no pairs",
		},
	}

	dc := NewClient(&http.Client{})

	for _, tt := range tests {
		_, err := dc.Pairs(tt.chainId, tt.pairAddrs)

		assert.Equal(t, tt.err, err != nil, tt.desc)
	}
}

// test tokens request
func TestTokens(t *testing.T) {
	type test struct {
		tokenAddrs []string
		err        bool
		want       string
		desc       string
	}

	tests := []test{
		{
			tokenAddrs: []string{"JUPyiwrYJFskUPiHa7hkeR8VUtAeFoSYbKedZNsDvCN"},
			err:        false,
			desc:       "jup",
		},
		{
			tokenAddrs: []string{"JUPyiwrYJFskUPiHa7hkeR8VUtAeFoSYbKedZNsDvCN,J1toso1uCk3RLmjorhTtrVwY9HJ7X8V9yYac6Y7kGCPn"},
			err:        false,
			desc:       "jup and jito",
		},
		{
			tokenAddrs: []string{},
			err:        true,
			desc:       "no tokens",
		},
	}

	dc := NewClient(&http.Client{})

	for _, tt := range tests {
		_, err := dc.Tokens(tt.tokenAddrs)

		assert.Equal(t, tt.err, err != nil, tt.desc)
	}
}

// test search request
func TestSearch(t *testing.T) {
	type test struct {
		query string
		err   bool
		want  string
		desc  string
	}

	tests := []test{
		{
			query: "JUPyiwrYJFskUPiHa7hkeR8VUtAeFoSYbKedZNsDvCN",
			err:   false,
			desc:  "jup token",
		},
		{
			query: "JUP/SOL",
			err:   false,
			desc:  "jup/sol pair",
		},
		{
			query: "",
			err:   true,
			desc:  "no query",
		},
		{
			query: "a",
			err:   true,
			desc:  "query len 1",
		},
	}

	dc := NewClient(&http.Client{})

	for _, tt := range tests {
		_, err := dc.Search(tt.query)

		assert.Equal(t, tt.err, err != nil, tt.desc, err)
	}
}
