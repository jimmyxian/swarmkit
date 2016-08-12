package manager

import (
	"github.com/docker/swarmkit/manager/orchestrator"
	"github.com/docker/swarmkit/manager/state/raft"
	"golang.org/x/net/context"
)

//
type Leader struct {
	replicatedOrchestrator *orchestrator.ReplicatedOrchestrator
}

func NewLeader(raftNode *raft.Node) *Leader {

	leader := &Leader{
		replicatedOrchestrator: orchestrator.NewReplicatedOrchestrator(raftNode.MemoryStore()),
	}

	return leader
}

func (l *Leader) Start(ctx context.Context) {
	l.replicatedOrchestrator.Start(ctx)
}

func (l *Leader) Stop(ctx context.Context) {
	l.replicatedOrchestrator.Stop(ctx)
}
