package domain_test

import (
	"encode/domain"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestValidateIFVideoIsEmpty(t *testing.T) {
	video := domain.NewVideo()
	err := video.Validate()

	require.Error(t, err)
}

func TestVideoIdIsNotUuid(t *testing.T) {
	badID := "abc"
	video := domain.Video{
		ID:         badID,
		ResourceID: "a",
		FilePath:   "/path/to/file.mp4",
		CreatedAt:  time.Now(),
	}

	err := video.Validate()
	require.Error(t, err)
}

func TestVideoValidation(t *testing.T) {
	video := domain.Video{
		ID:         uuid.NewV4().String(),
		ResourceID: "a",
		FilePath:   "/path/to/file.mp4",
		CreatedAt:  time.Now(),
	}

	err := video.Validate()
	require.Nil(t, err)
}
