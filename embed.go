package snapshothashes

import (
	"context"
	_ "embed"
	"fmt"
	"io"
	"net/http"

	_ "github.com/erigontech/erigon-snapshot/webseed"
)

//go:embed mainnet.toml
var Mainnet []byte

//go:embed sepolia.toml
var Sepolia []byte

//go:embed amoy.toml
var Amoy []byte

//go:embed bor-mainnet.toml
var BorMainnet []byte

//go:embed gnosis.toml
var Gnosis []byte

//go:embed chiado.toml
var Chiado []byte

//go:embed holesky.toml
var Holesky []byte

//go:embed bsc.toml
var Bsc []byte

//go:embed chapel.toml
var Chapel []byte

type SnapshotSource int

const (
	Github SnapshotSource = 0
	R2     SnapshotSource = 1
)

func getURLByChain(source SnapshotSource, chain, branch string) string {
	if source == Github {
		return fmt.Sprintf("https://raw.githubusercontent.com/node-real/bsc-erigon-snapshot/%s/%s.toml", branch, chain)
	} else if source == R2 {
		return fmt.Sprintf("https://download.snapshots.bnbchain.world/%s/%s.toml", branch, chain)
	}

	panic(fmt.Sprintf("unknown snapshot source: %d", source))
}

func LoadSnapshots(ctx context.Context, source SnapshotSource, branch string) (fetched bool, err error) {
	var (
		bscUrl    = getURLByChain(source, "bsc", branch)
		chapelUrl = getURLByChain(source, "chapel", branch)
	)
	var hashes []byte
	// Try to fetch the latest snapshot hashes from the web

	if hashes, err = fetchSnapshotHashes(ctx, source, bscUrl); err != nil {
		fetched = false
		return
	}
	Bsc = hashes

	if hashes, err = fetchSnapshotHashes(ctx, source, chapelUrl); err != nil {
		fetched = false
		return
	}
	Chapel = hashes
	fetched = true
	return fetched, nil
}

func fetchSnapshotHashes(ctx context.Context, source SnapshotSource, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	if source == R2 {
		insertCloudflareHeaders(req)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch snapshot hashes by %q: status code %d %s", url, resp.StatusCode, resp.Status)
	}
	res, err := io.ReadAll(resp.Body)
	if len(res) == 0 {
		return nil, fmt.Errorf("empty response from %s", url)
	}
	return res, err
}

// TODO: this was taken originally from erigon repo, downloader.go; we need to decide on a unique place to store such headers
var cloudflareHeaders = http.Header{
	"lsjdjwcush6jbnjj3jnjscoscisoc5s": []string{"I%OSJDNFKE783DDHHJD873EFSIVNI7384R78SSJBJBCCJBC32JABBJCBJK45"},
}

func insertCloudflareHeaders(req *http.Request) {
	for key, value := range cloudflareHeaders {
		req.Header[key] = value
	}
}
