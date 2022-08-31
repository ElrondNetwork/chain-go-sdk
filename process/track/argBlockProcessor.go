package track

import (
	"github.com/ElrondNetwork/elrond-go/process"
	"github.com/ElrondNetwork/elrond-go/sharding"
)

// ArgBlockProcessor holds all dependencies required to process tracked blocks in order to create new instances of
// block processor
type ArgBlockProcessor struct {
	HeaderValidator                       process.HeaderConstructionValidator
	RequestHandler                        process.RequestHandler
	ShardCoordinator                      sharding.Coordinator
	BlockTracker                          blockTrackerHandler
	SelfNotarizer                         blockNotarizerHandler
	SelfNotarizedHeadersNotifier          blockNotifierHandler
	RoundHandler                          process.RoundHandler
}
