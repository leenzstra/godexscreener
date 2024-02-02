package client

type ClientOpt func(*DexscreenerClient)

func WithBaseUrl(baseUrl string) func(*DexscreenerClient) {
  return func(s *DexscreenerClient) {
    s.baseUrl = baseUrl
  }
}