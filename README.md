# AnonyProof Protocol

> Secure, decentralized identity verification using Zero-Knowledge Proofs — no CPF, phone, or address ever exposed.

**AnonyProof** is a privacy-first identity verification protocol that allows users to prove they are trustworthy **without revealing any personal information**.  
Built on **Zero-Knowledge Proofs (ZKPs)** and **blockchain-based trust**, AnonyProof enables a safer and anonymous digital identity.

---

## 🚀 What Problem Does It Solve?

Millions of people are targeted by scams, extortion, and social engineering daily — all enabled by easy access to personal data.

AnonyProof solves this by:
- Removing the need to expose personal identifiers (CPF, phone, etc.)
- Enabling private and verifiable proof of trust
- Providing a safe way to interact with apps and services

---

## 🧠 Core Technologies

- **Backend**: Go + Fiber
- **Frontend**: Next.js (React)
- **ZKP Engine**: Semaphore (zk-SNARKs)
- **Smart Contracts**: Solidity + Hardhat
- **Blockchain Layer**: zkSync or Polygon
- **Wallet Integration**: WalletConnect + QRCode
- **Observability**: OpenTelemetry, Grafana Cloud, Loki, Tempo
- **Hosting**:
  - MVP: Railway
  - Production: AWS

---

## 🧱 Architecture Overview

                        ┌────────────────────────────┐
                        │             End User               │
                        └────────────┬───────────────┘
                                         │
                               Scan QR / Connect Wallet
                                         │
                        ┌────────────▼──────────────┐
                        │          GhostID Frontend         │
                        │      (Next.js + WalletConnect)    │
                        └────────────┬──────────────┘
                                         │
                               Request proof via API
                                         │
                        ┌────────────▼──────────────┐
                        │           AnonyProof API          │
                        │            (Go + Fiber)           │
                        └───────┬──────────┬────────┘
                                  │              │
             Generate/Verify ZKP  │              │  Send proof txn
                                  │              │
                     ┌─────────▼┐       ┌─ ─▼───────────┐
                     │  ZKP Engine  │      │   Smart Contracts  │
                     │  Semaphore   │      │   Solidity + ZK    │
                     └───────────┘      └─────┬─────────┘
                                                    │
                                       Record on-chain ZKP proof
                                                    │
                         ┌────────────────────▼──────────────────┐
                         │           Blockchain Layer (zkSync)              │
                         └───────────────────────────────────────┘

---

## 🕵️ Key Privacy Features

- ✅ Zero-Knowledge Proof-based identity validation  
- 🔒 No personal data exposed or stored  
- 📍 Fraud detection via IP, DDD, and browser fingerprinting  
- 👁️ Full system observability (tracing, metrics, logs)  
- 🧠 Future ML scoring to detect abuse patterns  

---

## 👻 GhostID Interface

**GhostID** is the user-friendly identity wallet for interacting with AnonyProof.  
It allows people to validate themselves using QR codes and cryptographic proofs — with zero exposure.

---

## 📊 Observability & Monitoring

- Structured logging with contextual metadata  
- Real-time distributed tracing (OpenTelemetry)  
- Custom metrics and dashboards (Grafana)  
- Fraud detection hooks (IP / DDD tracking)  

---

## 🧩 Roadmap

- [x] Define architecture and tech stack  
- [ ] MVP: Backend API with Go + ZKP validation  
- [ ] Deploy smart contracts on zkSync testnet  
- [ ] Build GhostID frontend with QR-based login  
- [ ] Add full observability (Loki + Tempo)  
- [ ] Implement ML module for fraud scoring (v2)  

---

## 🤝 Contributing

If you care about privacy and a better internet, join us.  
Open a PR, suggest features, or report an issue.


## ▶️ Running the Project Locally

### 🧑‍💻 Frontend (GhostID)

1. Open a terminal and navigate to the frontend folder:

   ```bash
   cd frontend
   ```

2. Install dependencies:

   ```bash
   npm install
   ```

3. Start the development server:

   ```bash
   npm run dev
   ```

4. Open your browser at:

   ```
   http://localhost:3000
   ```

---

### 🧠 Backend (AnonyProof API)

1. Open a terminal and navigate to the backend folder:

   ```bash
   cd backend
   ```

2. Run the application:

   ```bash
   go run cmd/api/main.go
   ```

3. Test the API health check:

   ```
   GET http://localhost:3000/health
   ```

   Expected response:

   ```json
   {
     "status": "ok"
   }
   ```

---

---

## 🛡️ License

MIT License

---

**AnonyProof** – *Privacy is a right, not a privilege.*
