package domain_test

import (
	"encode/domain"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestNewJob(t *testing.T) {
	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path/to/video.mp4"
	video.CreatedAt = time.Now()

	job, err := domain.NewJob("path", "CONVERTED", video)
	require.NotNil(t, job)
	require.Nil(t, err)
}