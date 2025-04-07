// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

/// @title AnonyProof - Decentralized Identity Verification Contract
/// @notice Stores anonymous zero-knowledge proof events for validation

contract AnonyProof {
    struct ProofRecord {
        address user;
        bytes32 proofHash;
        uint256 timestamp;
    }

    mapping(bytes32 => ProofRecord) public proofs;

    event ProofSubmitted(address indexed user, bytes32 indexed proofHash, uint256 timestamp);

    /// @notice Submit a new zero-knowledge proof hash
    /// @param proofHash The hash of the zero-knowledge proof
    function submitProof(bytes32 proofHash) external {
        require(proofs[proofHash].timestamp == 0, "Proof already submitted.");
        proofs[proofHash] = ProofRecord(msg.sender, proofHash, block.timestamp);
        emit ProofSubmitted(msg.sender, proofHash, block.timestamp);
    }

    /// @notice Check if a proof exists
    /// @param proofHash The hash to check
    /// @return exists True if the proof exists
    function isProofValid(bytes32 proofHash) public view returns (bool exists) {
        return proofs[proofHash].timestamp != 0;
    }
}
