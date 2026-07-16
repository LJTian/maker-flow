import { useEffect, useState } from "react";

const apiBase = import.meta.env.VITE_API_BASE_URL?.replace(/\/$/, "") ?? "";

export default function App() {
  const [apiHealth, setApiHealth] = useState<string>("");

  useEffect(() => {
    if (!apiBase) {
      return;
    }

    fetch(`${apiBase}/health`)
      .then((res) => res.json())
      .then((data) => setApiHealth(JSON.stringify(data)))
      .catch(() => setApiHealth("unreachable"));
  }, []);

  return (
    <main className="flex min-h-screen flex-col items-center justify-center bg-slate-950 p-8 text-slate-100">
      <h1 className="mb-2 text-3xl font-semibold">MVP Web UI</h1>
      <p className="mb-6 text-slate-400">Vite · React · TypeScript · Tailwind</p>
      <p className="mb-2 text-sm text-emerald-400">GET /health → ok</p>
      {apiBase ? (
        <p className="text-sm text-slate-300">
          Backend <code className="text-slate-400">{apiBase}</code>:{" "}
          {apiHealth || "…"}
        </p>
      ) : (
        <p className="text-sm text-slate-500">
          Set <code>VITE_API_BASE_URL</code> in <code>.env</code> to ping a
          backend API
        </p>
      )}
    </main>
  );
}
