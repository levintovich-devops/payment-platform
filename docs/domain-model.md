# Domain Model (v1 Draft)

## Core Business Entities

### Payment
- Represents the complete payment lifecycle for v1.
- Key attributes: identifier, reference, amount, currency, status.

### NotificationRecord
- Represents a notification generated from a payment state change.
- Key attributes: identifier, payment reference, status, delivery state.

## Entity Relationships

- A Payment may generate zero or more NotificationRecords.
- NotificationRecords are associated with the related Payment.

## Entity Lifecycle / State Transitions

### Payment
- INITIATED -> CAPTURED
- INITIATED -> FAILED

### NotificationRecord
- Pending -> Sent
- Pending -> Failed

## Domain Events

- PaymentInitiated
- PaymentCaptured
- PaymentFailed
