export default function App() {
  return (
    <main className="app-shell">
      <section className="hero-card">
        <p className="eyebrow">Payment Platform</p>

        <h1>Enterprise Event-Driven Payment Platform</h1>

        <p className="description">
          Modern cloud-native payment platform built with Go, React, Kafka,
          PostgreSQL, Docker, Kubernetes and Jenkins.
        </p>

        <div className="actions" aria-label="Primary actions">
          <button type="button">
            Create Payment
          </button>

          <button
            type="button"
            className="secondary"
          >
            View Payments
          </button>
        </div>
      </section>
    </main>
  );
}