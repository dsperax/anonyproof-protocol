---

## (EN)

# AnonyProof – Technical Architecture and Development Strategy


## 1. Technologies Used

### 🧠 Backend: Golang (Go) with Clean Architecture
Go is a compiled, safe, high-performance language created by Google.
We use Go to build the main AnonyProof API with a focus on scalability and testability.
Structure follows Clean Architecture with domain, ports, adapters separated.

### 🔒 Zero-Knowledge Proofs: Semaphore
Semaphore is a zk-SNARK library enabling private proof generation.
We use it to prove user legitimacy without revealing their identity.

### 🔐 Smart Contracts: Solidity + Hardhat
Solidity is used to write on-chain smart contracts. Hardhat handles deployment and local testing.

### 🧑‍💻 Frontend: Next.js + React
Next.js enables SEO-friendly SSR React interface for users to interact anonymously.

### 📦 Observability: OpenTelemetry + Grafana Cloud
OpenTelemetry captures logs, traces, and metrics. Exported to Grafana Cloud for visibility.

### 🌐 Blockchain Layer: zkSync / Polygon
Contracts are deployed on zkSync for privacy performance or Polygon for EVM compatibility.

## 2. Methodology

We adopt practices inspired by Google, Amazon, Microsoft:
- Domain-Driven Design (DDD)
- Ports and Adapters (Hexagonal)
- TDD (Test Driven Development)
- CI/CD (future)
- Rich Markdown documentation

## 3. Execution Plan

- Phase 1: Setup Go backend, Next.js frontend, Solidity contract
- Phase 2: Integrate Semaphore, proof endpoints
- Phase 3: QR generation, WalletConnect
- Phase 4: Observability setup
- Phase 5: IP, DDD, fingerprint tracking and ML layer

---

## (PT)

# AnonyProof – Arquitetura Técnica e Estratégia de Desenvolvimento


## 1. Tecnologias Utilizadas

### 🧠 Backend: Golang (Go) com Clean Architecture
Go é uma linguagem compilada e de alto desempenho criada pelo Google.
Utilizaremos Go para criar a API do AnonyProof com escalabilidade e testabilidade.
A arquitetura é limpa, com separação entre domínio, portas e adaptadores.

### 🔒 Zero-Knowledge Proofs: Semaphore
Semaphore é uma biblioteca de zk-SNARKs que permite provas privadas.
Utilizada para provar legitimidade sem revelar a identidade.

### 🔐 Smart Contracts: Solidity + Hardhat
Solidity será usada para contratos inteligentes. Hardhat para testes e deploys locais.

### 🧑‍💻 Frontend: Next.js + React
Next.js proporciona renderização SSR e uma interface moderna com React.

### 📦 Observabilidade: OpenTelemetry + Grafana Cloud
OpenTelemetry coleta logs, métricas e rastreamentos. Grafana Cloud exibe tudo com dashboards.

### 🌐 Blockchain Layer: zkSync / Polygon
Deploy em zkSync para privacidade e performance, ou Polygon para ampla compatibilidade EVM.

## 2. Metodologia

Inspirada em Google, Amazon e Microsoft:
- Domain-Driven Design (DDD)
- Ports and Adapters (Arquitetura Hexagonal)
- TDD (Desenvolvimento Orientado a Testes)
- CI/CD (futuro)
- Documentação com Markdown

## 3. Plano de Execução

- Etapa 1: Backend Go, Frontend Next.js, contrato Solidity
- Etapa 2: Integração com Semaphore, criação de endpoints
- Etapa 3: QR code, WalletConnect
- Etapa 4: Setup de observabilidade
- Etapa 5: Rastreio de IP, DDD e camada de ML

---

## (IT)

# AnonyProof – Architettura Tecnica e Strategia di Sviluppo


## 1. Tecnologie Utilizzate

### 🧠 Backend: Golang (Go) con Architettura Pulita
Go è un linguaggio compilato, sicuro e ad alte prestazioni creato da Google.
Utilizzato per l'API AnonyProof seguendo il modello Clean Architecture.

### 🔒 Zero-Knowledge Proofs: Semaphore
Semaphore è una libreria zk-SNARK per prove anonime e sicure.
Permette agli utenti di dimostrare la loro legittimità senza esporre l'identità.

### 🔐 Smart Contracts: Solidity + Hardhat
Solidity è usato per scrivere smart contract sulla blockchain.
Hardhat viene usato per test e distribuzioni locali.

### 🧑‍💻 Frontend: Next.js + React
Next.js offre SSR e un'interfaccia moderna per l'utente tramite React.

### 📦 Osservabilità: OpenTelemetry + Grafana Cloud
OpenTelemetry raccoglie log, metriche e tracce esportati su Grafana Cloud.

### 🌐 Blockchain Layer: zkSync / Polygon
Contratti deployati su zkSync per privacy o su Polygon per compatibilità EVM.

## 2. Metodologia

Ispirata a Google, Amazon, Microsoft:
- Domain-Driven Design (DDD)
- Ports and Adapters
- TDD (Test Driven Development)
- CI/CD (in futuro)
- Documentazione in Markdown

## 3. Piano Operativo

- Fase 1: backend Go, frontend Next.js, contratto Solidity
- Fase 2: Integrazione con Semaphore, endpoint
- Fase 3: QR code, WalletConnect
- Fase 4: osservabilità
- Fase 5: tracciamento IP, DDD, fingerprint e ML

---

## (RU)

# AnonyProof – Техническая Архитектура и Стратегия Разработки


## 1. Используемые Технологии

### 🧠 Backend: Golang (Go) и Чистая Архитектура
Go — это компилируемый язык с высокой производительностью, созданный Google.
Он используется для создания API AnonyProof по принципам чистой архитектуры.

### 🔒 Zero-Knowledge Proofs: Semaphore
Semaphore — это библиотека zk-SNARK, обеспечивающая создание анонимных доказательств.

### 🔐 Smart Contracts: Solidity + Hardhat
Solidity — язык для смарт-контрактов в Ethereum. Hardhat используется для тестирования и деплоя.

### 🧑‍💻 Frontend: Next.js + React
Next.js обеспечивает SSR и современный интерфейс на базе React.

### 📦 Наблюдаемость: OpenTelemetry + Grafana Cloud
OpenTelemetry собирает логи и метрики, отображаемые через Grafana Cloud.

### 🌐 Блокчейн: zkSync / Polygon
Контракты размещаются на zkSync (для конфиденциальности) или Polygon (для EVM-совместимости).

## 2. Методология

Основана на практике Google, Amazon и Microsoft:
- DDD (Проектирование, ориентированное на домен)
- Порты и адаптеры (гексагональная архитектура)
- TDD (разработка через тестирование)
- CI/CD (в будущем)
- Документация на Markdown

## 3. План Выполнения

- Этап 1: backend Go, frontend Next.js, контракт Solidity
- Этап 2: Интеграция Semaphore, эндпоинты
- Этап 3: QR-код и WalletConnect
- Этап 4: Наблюдаемость
- Этап 5: Отслеживание IP, DDD, отпечатков и ML

