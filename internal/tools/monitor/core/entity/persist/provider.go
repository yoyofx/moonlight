// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package persist

import (
	"context"
	"fmt"
	"time"

	"github.com/ping-cloudnative/moonlight-utils/base/logs"
	"github.com/ping-cloudnative/moonlight-utils/base/servicehub"
	"github.com/ping-cloudnative/moonlight/internal/tools/monitor/core/entity/storage"
	"github.com/ping-cloudnative/moonlight/internal/tools/monitor/core/storekit"
	"github.com/ping-cloudnative/moonlight/internal/tools/monitor/oap/collector/lib/kafka"
)

type config struct {
	Input                kafka.BatchReaderConfig `file:"input"`
	Parallelism          int                     `file:"parallelism" default:"1"`
	BufferSize           int                     `file:"buffer_size" default:"1024"`
	ReadTimeout          time.Duration           `file:"read_timeout" default:"5s"`
	PrintInvalidData     bool                    `file:"print_invalid_data" default:"false"`
	StorageWriterService string                  `file:"storage_writer_service" default:"entity-storage-elasticsearch-writer"`
}

type provider struct {
	Cfg           *config
	Log           logs.Logger
	Kafka         kafka.Interface `autowired:"kafkago"`
	StorageWriter storage.Storage

	r         storekit.BatchReader
	stats     Statistics
	validator Validator
}

func (p *provider) Init(ctx servicehub.Context) error {

	p.validator = newValidator(p.Cfg)
	if runner, ok := p.validator.(servicehub.ProviderRunnerWithContext); ok {
		ctx.AddTask(runner.Run, servicehub.WithTaskName("entity validator"))
	}

	p.stats = newStatistics()
	p.StorageWriter = ctx.Service(p.Cfg.StorageWriterService).(storage.Storage)
	r, err := p.Kafka.NewBatchReader(&p.Cfg.Input, kafka.WithReaderDecoder(p.decodeData))
	if err != nil {
		return err
	}
	p.r = r

	// add consumer task
	for i := 0; i < p.Cfg.Parallelism; i++ {
		ctx.AddTask(func(ctx context.Context) error {

			w, err := p.StorageWriter.NewWriter(ctx)
			if err != nil {
				return err
			}
			defer w.Close()
			return storekit.BatchConsume(ctx, r, w, &storekit.BatchConsumeOptions{
				BufferSize:          p.Cfg.BufferSize,
				ReadTimeout:         p.Cfg.ReadTimeout,
				ReadErrorHandler:    p.handleReadError,
				WriteErrorHandler:   p.handleWriteError,
				ConfirmErrorHandler: p.confirmErrorHandler,
				Statistics:          p.stats,
			})
		}, servicehub.WithTaskName(fmt.Sprintf("consumer(%d)", i)))
	}
	return nil
}

func (p *provider) Close() error {
	return p.r.Close()
}

func init() {
	servicehub.Register("entity-persist", &servicehub.Spec{
		ConfigFunc: func() interface{} { return &config{} },
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
