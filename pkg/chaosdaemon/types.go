// Copyright 2019 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package chaosdaemon

import (
	"github.com/vishvananda/netlink"

	pb "github.com/pingcap/chaos-mesh/pkg/chaosdaemon/pb"
)

func ToNetlinkNetemAttrs(netem *pb.Netem) netlink.NetemQdiscAttrs {
	return netlink.NetemQdiscAttrs{
		Latency:       netem.Time,
		DelayCorr:     netem.DelayCorr,
		Limit:         netem.Limit,
		Loss:          netem.Loss,
		LossCorr:      netem.LossCorr,
		Gap:           netem.Gap,
		Duplicate:     netem.Duplicate,
		DuplicateCorr: netem.DuplicateCorr,
		Jitter:        netem.Jitter,
		ReorderProb:   netem.Reorder,
		ReorderCorr:   netem.ReorderCorr,
		CorruptProb:   netem.Corrupt,
		CorruptCorr:   netem.CorruptCorr,
	}
}
