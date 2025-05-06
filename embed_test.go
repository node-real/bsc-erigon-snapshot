package snapshothashes

import (
	"context"
	"testing"
)

func TestFetchSnapshotHashes(t *testing.T) {
	dat, err := fetchSnapshotHashes(context.Background(), Github, "https://raw.githubusercontent.com/node-real/bsc-erigon-snapshot/main/mainnet.toml")
	if err != nil {
		t.Errorf("fetchSnapshotHashes() failed: %v", err)
	}
	if len(dat) == 0 {
		t.Errorf("fetchSnapshotHashes() failed: empty data")
	}
}

func TestFetchSnapshotHashesAll(t *testing.T) {
	ok, err := LoadSnapshots(context.Background(), Github, "main")
	if err != nil {
		t.Errorf("LoadSnapshots() failed %s", err)
	}
	if !ok {
		t.Errorf("LoadSnapshots() failed")
	}
	if len(Bsc) == 0 {
		t.Errorf("Mainnet is empty")
	}
	if len(Chapel) == 0 {
		t.Errorf("Sepolia is empty")
	}
}
