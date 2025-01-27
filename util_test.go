package kbucket

import (
	"testing"
	"time"

	"github.com/libp2p/go-libp2p-kbucket/peerdiversity"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/peerstore"
	"github.com/libp2p/go-libp2p/core/test"
	"github.com/stretchr/testify/require"
)

// TODO: think, should I make this function public??
func newEmptyFeaturesRT(bucketsize int, localID ID, latency time.Duration, m peerstore.Metrics, 
	usefulnessGracePeriod time.Duration, df *peerdiversity.Filter) (*RoutingTable, error) {
	return NewRoutingTable(bucketsize, localID, peer.NewFeatureSet(), nil, latency, m, usefulnessGracePeriod, df)
}

func TestCloser(t *testing.T) {
	Pa := test.RandPeerIDFatal(t)
	Pb := test.RandPeerIDFatal(t)
	var X string

	// returns true if d(Pa, X) < d(Pb, X)
	for {
		X = string(test.RandPeerIDFatal(t))
		if xor(ConvertPeerID(Pa), ConvertKey(X)).less(xor(ConvertPeerID(Pb), ConvertKey(X))) {
			break
		}
	}

	require.True(t, Closer(Pa, Pb, X))

	// returns false if d(Pa,X) > d(Pb, X)
	for {
		X = string(test.RandPeerIDFatal(t))
		if xor(ConvertPeerID(Pb), ConvertKey(X)).less(xor(ConvertPeerID(Pa), ConvertKey(X))) {
			break
		}

	}
	require.False(t, Closer(Pa, Pb, X))
}
