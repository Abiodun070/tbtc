package btc

import (
	"bytes"
	"encoding/hex"
	"fmt"
)

// LocalChain represents a local Bitcoin chain.
type LocalChain struct {
	headers         []*Header
	orphanedHeaders []*btc.Header
}

// ConnectLocal connects to the local Bitcoin chain and returns a chain handle.
func ConnectLocal() (Handle, error) {
	logger.Infof("connecting local Bitcoin chain")

	return &LocalChain{}, nil
}

// GetHeaderByHeight returns the block header from the longest block chain at
// the given block height.
func (lc *LocalChain) GetHeaderByHeight(height int64) (*Header, error) {
	for _, header := range lc.headers {
		if header.Height == height {
			return header, nil
		}
	}

	return nil, fmt.Errorf("no header with height [%v]", height)
}

// GetHeaderByDigest returns the block header for given digest (hash).
func (lc *LocalChain) GetHeaderByDigest(
	digest Digest,
) (*Header, error) {
	for _, header := range lc.headers {
		if bytes.Equal(header.Hash[:], digest[:]) {
			return header, nil
		}
	}

	for _, header := range lc.orphanedHeaders {
		if bytes.Equal(header.Hash[:], digest[:]) {
			return header, nil
		}
	}

	return nil, fmt.Errorf(
		"no header with digest [%v]",
		hex.EncodeToString(digest[:]),
	)
}

// GetBlockCount returns the number of blocks in the longest blockchain
func (lc *LocalChain) GetBlockCount() (int64, error) {
	return int64(len(lc.headers)), nil
}

// SetHeaders sets internal headers for testing purposes.
func (lc *LocalChain) SetHeaders(headers []*Header) {
	lc.headers = headers
}

// AppendHeader appends internal header for testing purposes.
func (c *Chain) AppendHeader(header *btc.Header) {
	c.headers = append(c.headers, header)
}

// SetOrphanedHeaders sets internal orphaned headers for testing purposes.
func (c *Chain) SetOrphanedHeaders(headers []*btc.Header) {
	c.orphanedHeaders = headers
}
