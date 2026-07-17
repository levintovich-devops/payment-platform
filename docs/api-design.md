# API Design (v1 Draft)

## Public REST API Surface

### POST /payments
- Purpose: Initiate a new payment.
- Request responsibility: Accept the payment initiation request and trigger the payment lifecycle transition to INITIATED.
- Response responsibility: Return the created payment resource and its current status.

### POST /payments/{paymentId}/capture
- Purpose: Capture an existing initiated payment.
- Request responsibility: Accept the capture request for the specified payment and trigger the payment lifecycle transition to CAPTURED.
- Response responsibility: Return the updated payment resource and its current status.

### GET /payments/{paymentId}
- Purpose: Retrieve an existing payment and its current status.
- Request responsibility: Identify the specified payment and retrieve its current state.
- Response responsibility: Return the payment resource and its current status.

### GET /payments
- Purpose: List payment history with pagination.
- Request responsibility: Accept pagination parameters and retrieve the requested payment history page.
- Response responsibility: Return the page of payment resources and pagination metadata.
