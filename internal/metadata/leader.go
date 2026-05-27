// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package metadata

import (
	"context"
	"database/sql"
	"sync/atomic"
	"time"

	"github.com/openrundev/openrun/internal/types"
)

type LeaderElection struct {
	*types.Logger
	metadata              *Metadata
	db                    *sql.DB
	nodeId                string
	hostname              string
	heartbeatLeaseSecs    int
	heartbeatIntervalSecs int
	isLeader              atomic.Bool
	cancel                context.CancelFunc
}

type LeaderState struct {
	LeaderID               string
	Hostname               string
	LastHeartbeatAt        time.Time
	LastLeadershipChangeAt time.Time
}

func NewLeaderElection(logger *types.Logger, metadata *Metadata, config *types.ServerConfig, nodeId string, hostname string) *LeaderElection {
	_ = "STUB: not implemented"
	return nil
}

func (l *LeaderElection) CreateTables(ctx context.Context, tx types.Transaction) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *LeaderElection) IsLeader() bool { _ = "STUB: not implemented"; return false }

func (l *LeaderElection) tryAcquire(ctx context.Context) (*LeaderState, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

// No rows returned => did not acquire (WHERE condition failed)

func (l *LeaderElection) heartbeat(ctx context.Context) (time.Time, bool, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), false, nil
}

func (l *LeaderElection) Stop() { _ = "STUB: not implemented"; return }

func (l *LeaderElection) StartLoop(parentCtx context.Context) { _ = "STUB: not implemented"; return }

// Try to acquire leadership immediately on startup

// Lost leadership
