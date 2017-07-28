package filter

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
)

//MockConfig mock struct
type MockConfig struct {
	mock.Mock
}

func (m *MockConfig) processDirectory() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockConfig) isDir() bool {
	args := m.Called()
	return args.Bool(0)
}

func TestRunFilter(t *testing.T) {
	runFilterTestFactory := func(errorExpected bool, isDir bool, goodProcess bool) func(t *testing.T) {
		return func(t *testing.T) {
			// build mock
			mockFilterConfig := new(MockConfig)
			mockFilterConfig.On("isDir").Return(isDir)

			if isDir {
				if goodProcess {
					mockFilterConfig.On("processDirectory").Return(nil)
				} else {
					mockFilterConfig.On("processDirectory").Return(errors.New("test error"))
				}
			}

			// run RunFilter
			err := RunFilter(mockFilterConfig)
			if err != nil && !errorExpected {
				t.Error("Error thown when not expected")
			}

			mockFilterConfig.AssertExpectations(t)
		}
	}

	t.Run("No dir or process", runFilterTestFactory(true, false, false))
	t.Run("Is dir but process fails", runFilterTestFactory(true, true, false))
	t.Run("Is dir and process suceeds", runFilterTestFactory(true, true, true))
}
