package services

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockProofRepo struct {
	mock.Mock
}

func (m *MockProofRepo) SubmitProof(proofHash string) error {
	args := m.Called(proofHash)
	return args.Error(0)
}

type MockVerifier struct {
	mock.Mock
}

func (m *MockVerifier) VerifyProof(proof string) (bool, error) {
	args := m.Called(proof)
	return args.Bool(0), args.Error(1)
}

func TestSubmitProof_Success(t *testing.T) {
	mockRepo := new(MockProofRepo)
	mockVerifier := new(MockVerifier)
	usecase := NewSubmitProofUseCase(mockRepo, mockVerifier)

	proof := "proofData"
	hash := "0xabc123"

	mockVerifier.On("VerifyProof", proof).Return(true, nil)
	mockRepo.On("SubmitProof", hash).Return(nil)

	err := usecase.Submit(hash, proof)
	assert.NoError(t, err)
	mockVerifier.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestSubmitProof_InvalidProof(t *testing.T) {
	mockRepo := new(MockProofRepo)
	mockVerifier := new(MockVerifier)
	usecase := NewSubmitProofUseCase(mockRepo, mockVerifier)

	proof := "invalid"
	hash := "0xabc123"

	mockVerifier.On("VerifyProof", proof).Return(false, nil)

	err := usecase.Submit(hash, proof)
	assert.EqualError(t, err, "invalid zero-knowledge proof")
	mockVerifier.AssertExpectations(t)
}

func TestSubmitProof_VerifierError(t *testing.T) {
	mockRepo := new(MockProofRepo)
	mockVerifier := new(MockVerifier)
	usecase := NewSubmitProofUseCase(mockRepo, mockVerifier)

	proof := "error"
	hash := "0xabc123"

	mockVerifier.On("VerifyProof", proof).Return(false, errors.New("verifier failure"))

	err := usecase.Submit(hash, proof)
	assert.EqualError(t, err, "verifier failure")
	mockVerifier.AssertExpectations(t)
}
