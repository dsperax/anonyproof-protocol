package ports

// ZKVerifier abstracts the validation of zero-knowledge proofs
type ZKVerifier interface {
    VerifyProof(proof string) (bool, error)
}
