package ports

type ProofRepository interface {
    SubmitProof(proofHash string) error
}
