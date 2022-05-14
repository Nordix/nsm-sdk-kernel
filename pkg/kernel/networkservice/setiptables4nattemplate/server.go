// Copyright (c) 2022 Xored Software Inc and others.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build linux
// +build linux

// Package setiptables4nattemplate provides chain element for setup iptables rules template properties
package setiptables4nattemplate

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/networkservicemesh/api/pkg/api/networkservice"
	"github.com/networkservicemesh/api/pkg/api/networkservice/mechanisms/kernel"
	"github.com/networkservicemesh/sdk/pkg/networkservice/core/next"
)

type setIPTablesTemplateServer struct {
	rulesTemplate []string
}

// NewServer - returns a new networkservice.NetworkServiceServer that writes IP Tables rules template
// to kernel mechanism
func NewServer(rulesTemplate []string) networkservice.NetworkServiceServer {
	return &setIPTablesTemplateServer{
		rulesTemplate: rulesTemplate,
	}
}

func (s *setIPTablesTemplateServer) Request(ctx context.Context, request *networkservice.NetworkServiceRequest) (*networkservice.Connection, error) {
	mechanism := kernel.ToMechanism(request.GetConnection().GetMechanism())
	if mechanism != nil {
		mechanism.SetIPTables4NatTemplate(s.rulesTemplate...)
	}

	return next.Server(ctx).Request(ctx, request)
}

func (s *setIPTablesTemplateServer) Close(ctx context.Context, conn *networkservice.Connection) (*empty.Empty, error) {
	return next.Server(ctx).Close(ctx, conn)
}
