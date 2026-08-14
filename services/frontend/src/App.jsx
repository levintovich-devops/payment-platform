import { useState } from 'react';
import { createPayment } from './api';

const defaultForm = {
  reference: 'INV-1001',
  amount: '150.00',
  currency: 'USD',
};

export default function App() {
  const [form, setForm] = useState(defaultForm);
  const [payment, setPayment] = useState(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');

  const handleChange = (event) => {
    const { name, value } = event.target;

    setForm((previous) => ({
      ...previous,
      [name]: value,
    }));
  };

  const handleCreatePayment = async () => {
    setLoading(true);
    setError('');

    try {
      const createdPayment = await createPayment(form);
      setPayment(createdPayment);
    } catch (requestError) {
      setError(requestError.message || 'Unable to create payment.');
    } finally {
      setLoading(false);
    }
  };

  return (
    <main className="app-shell">
      <section className="hero-card">
        <p className="eyebrow">Payment Platform</p>

        <h1>Enterprise Event-Driven Payment Platform</h1>

        <p className="description">
          Create a payment using the Payment Service REST API.
        </p>

        <div className="payment-form" aria-label="Create payment form">
          <label>
            Reference
            <input
              name="reference"
              value={form.reference}
              onChange={handleChange}
              placeholder="INV-1001"
            />
          </label>

          <label>
            Amount
            <input
              name="amount"
              value={form.amount}
              onChange={handleChange}
              placeholder="150.00"
            />
          </label>

          <label>
            Currency
            <input
              name="currency"
              value={form.currency}
              onChange={handleChange}
              placeholder="USD"
            />
          </label>
        </div>

        <div className="actions" aria-label="Primary actions">
          <button
            type="button"
            onClick={handleCreatePayment}
            disabled={loading}
          >
            {loading ? 'Creating...' : 'Create Payment'}
          </button>
        </div>

        {error ? (
          <p className="status error">{error}</p>
        ) : null}

        {payment ? (
          <div className="payment-card" aria-live="polite">
            <h2>Payment Result</h2>

            <dl>
              <div>
                <dt>ID</dt>
                <dd>{payment.id}</dd>
              </div>

              <div>
                <dt>Reference</dt>
                <dd>{payment.reference}</dd>
              </div>

              <div>
                <dt>Amount</dt>
                <dd>{payment.amount}</dd>
              </div>

              <div>
                <dt>Currency</dt>
                <dd>{payment.currency}</dd>
              </div>

              <div>
                <dt>Status</dt>
                <dd>{payment.status}</dd>
              </div>
            </dl>
          </div>
        ) : null}
      </section>
    </main>
  );
}