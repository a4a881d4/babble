package proxy

import "github.com/babbleio/babble/hashgraph"

type AppProxy interface {
	SubmitCh() chan []byte
	CommitBlock(block hashgraph.Block) error
}

type BabbleProxy interface {
	CommitCh() chan hashgraph.Block
	SubmitTx(tx []byte) error
}
