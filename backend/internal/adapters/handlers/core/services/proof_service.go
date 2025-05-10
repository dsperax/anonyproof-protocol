package services

import (
	"anonyproof-protocol/backend/internal/adapters/handlers/core/ports"
	"errors"
)

type SubmitProofUseCase struct {
	repo     ports.ProofRepository
	verifier ports.ZKVerifier
}

func NewSubmitProofUseCase(r ports.ProofRepository, v ports.ZKVerifier) *SubmitProofUseCase {
	return &SubmitProofUseCase{repo: r, verifier: v}
}

func (s *SubmitProofUseCase) Submit(proofHash string, proof string) error {
	ok, err := s.verifier.VerifyProof(proof)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("invalid zero-knowledge proof")
	}

	return s.repo.SubmitProof(proofHash)
}
