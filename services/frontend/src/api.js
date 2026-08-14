const PAYMENT_API_BASE_URL = import.meta.env.VITE_PAYMENT_API_URL ?? '';

async function parseResponse(response) {
  const contentType = response.headers.get('content-type') || '';

  if (contentType.includes('application/json')) {
    const payload = await response.json();

    if (!response.ok) {
      throw new Error(payload?.message || 'Request failed.');
    }

    return payload;
  }

  if (!response.ok) {
    throw new Error('Request failed.');
  }

  return response.text();
}

export async function createPayment(payload) {
  const response = await fetch(
    `${PAYMENT_API_BASE_URL}/payments`,
    {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Accept: 'application/json',
      },
      body: JSON.stringify(payload),
    },
  );

  return parseResponse(response);
}