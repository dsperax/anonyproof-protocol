# Project Overview: AnonyProof

## What is AnonyProof?

AnonyProof is a privacy-first, decentralized protocol that enables users to verify their identity or trustworthiness without exposing sensitive personal information (e.g., phone numbers, CPF, address).

It leverages:
- **Zero-Knowledge Proofs (ZKPs)** for privacy-preserving validation
- **Smart Contracts on zkSync/Polygon** for trustless, tamper-proof verification
- **GhostID Wallet Interface** for user interaction via QR codes
- **Clean Architecture (Backend in Go)** for scalability and maintainability
- **OpenTelemetry Observability Stack** to ensure transparency and traceability
- **Fraud Detection Layer** using IP, DDD, fingerprinting, and (soon) ML-based behavior analysis

## Key Principles

- No PII is ever stored or transmitted
- Identity proofs are anonymous, decentralized, and verifiable
- Backend is modular with domain-driven design and port/adapters separation
- Designed to support high security, auditability, and global interoperability

## Target Use Cases

- Secure onboarding for fintech platforms
- Anonymous login for marketplaces and forums
- Pre-authenticated access to government or health portals
