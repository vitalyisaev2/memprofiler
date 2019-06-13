package filesystem

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/golang/protobuf/ptypes"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"github.com/memprofiler/memprofiler/schema"
	"github.com/memprofiler/memprofiler/server/config"
)

// simple integration test for file-based storage
func TestDefaultStorage_Integration_Write_Read(t *testing.T) {
	fmt.Printf("Hello!")

	logger := logrus.New()
	logger.Out = os.Stdout

	// create new s in tmp dir
	dataDir, err := ioutil.TempDir("/tmp", "memprofiler")
	assert.NoError(t, err)

	defer func() {
		if err = os.RemoveAll(dataDir); err != nil {
			logger.WithError(err).Fatal("Failed to remove dir")
		}
	}()

	cfg := &config.FilesystemStorageConfig{
		DataDir:   dataDir,
		SyncWrite: false,
	}

	fmt.Printf("Hello!")
	s, err := NewStorage(logger, cfg)
	assert.NotNil(t, s)
	assert.NoError(t, err)

	// write some measurements
	serviceDesc := &schema.ServiceDescription{
		ServiceType:     "database",
		ServiceInstance: "localhost:8080",
	}
	saver, err := s.NewDataSaver(serviceDesc)
	assert.NoError(t, err)
	assert.NotNil(t, saver)

	input := []*schema.Measurement{
		{
			ObservedAt: ptypes.TimestampNow(),
			Locations: []*schema.Location{
				{
					MemoryUsage: &schema.MemoryUsage{AllocBytes: 1},
					Callstack: &schema.Callstack{
						Id:     "abcd",
						Frames: []*schema.StackFrame{{Name: "a", File: "b.go", Line: 1}},
					},
				},
			},
		},
		{
			ObservedAt: ptypes.TimestampNow(),
			Locations: []*schema.Location{
				{
					MemoryUsage: &schema.MemoryUsage{AllocBytes: 2},
					Callstack: &schema.Callstack{
						Id:     "edfg",
						Frames: []*schema.StackFrame{{Name: "b", File: "c.go", Line: 2}},
					},
				},
			},
		},
	}

	for _, mm := range input {
		err = saver.Save(mm)
		assert.NoError(t, err)
	}
	err = saver.Close()
	assert.NoError(t, err)

	// try to load data just written
	sessionDesc := &schema.SessionDescription{
		ServiceType:     serviceDesc.GetServiceType(),
		ServiceInstance: serviceDesc.GetServiceInstance(),
		SessionId:       saver.SessionID(),
	}
	loader, err := s.NewDataLoader(sessionDesc)
	assert.NotNil(t, loader)
	assert.NoError(t, err)

	outChan, err := loader.Load(context.Background())
	assert.NotNil(t, outChan)
	assert.NoError(t, err)

	output := make([]*schema.Measurement, 0, 2)
	for result := range outChan {
		assert.NotNil(t, result.Measurement)
		if !assert.NoError(t, result.Err) {
			assert.FailNow(t, "failed to read data")
		}
		output = append(output, result.Measurement)
	}

	err = loader.Close()
	assert.NoError(t, err)

	assert.Equal(t, len(input), len(output))
	assert.Equal(t, input, output)
}