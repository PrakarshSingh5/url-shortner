"use client";

import { useState } from "react";
import { motion } from "framer-motion";

const DEFAULT_API_BASE_URL = "http://localhost:8080";
const API_BASE_URL = (
  process.env.NEXT_PUBLIC_API_URL ?? DEFAULT_API_BASE_URL
).replace(/\/$/, "");
const SHORTEN_ENDPOINT = `${API_BASE_URL}/api/shorten`;

interface UrlFormProps {
  onSuccess: () => void;
}

export default function UrlForm({ onSuccess }: UrlFormProps) {
  const [url, setUrl] = useState("");
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");
  const [success, setSuccess] = useState("");

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError("");
    setSuccess("");
    setLoading(true);

    try {
      const response = await fetch(SHORTEN_ENDPOINT, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ url }),
      });

      const data = await response.json();

      if (!response.ok) {
        throw new Error(data.error || "Failed to shorten URL");
      }

      setSuccess(data.short_url);
      setUrl("");
      onSuccess();
    } catch (err) {
      setError(err instanceof Error ? err.message : "An error occurred");
    } finally {
      setLoading(false);
    }
  };

  return (
    <motion.div
      initial={{ opacity: 0, y: 20 }}
      animate={{ opacity: 1, y: 0 }}
      transition={{ duration: 0.5 }}
      className="w-full max-w-2xl"
    >
      <form onSubmit={handleSubmit} className="space-y-4">
        <motion.div
          whileFocus={{ scale: 1.02 }}
          transition={{ type: "spring", stiffness: 300 }}
        >
          <input
            type="text"
            value={url}
            onChange={(e) => setUrl(e.target.value)}
            placeholder="🔗 Enter URL to shorten..."
            className="w-full px-6 py-4 bg-[var(--card)] border border-[var(--border)] rounded-lg text-[var(--foreground)] placeholder-[var(--muted-foreground)] focus:outline-none focus:ring-2 focus:ring-[var(--primary)] focus:border-transparent transition-all"
            disabled={loading}
          />
        </motion.div>

        {error && (
          <motion.div
            initial={{ opacity: 0, y: -10 }}
            animate={{ opacity: 1, y: 0 }}
            className="px-4 py-2 bg-[var(--destructive)]/20 border border-[var(--destructive)] rounded-lg text-[var(--destructive)] text-sm"
          >
            ⚠️ {error}
          </motion.div>
        )}

        {success && (
          <motion.div
            initial={{ opacity: 0, y: -10 }}
            animate={{ opacity: 1, y: 0 }}
            className="px-4 py-2 bg-[var(--primary)]/20 border border-[var(--primary)] rounded-lg text-[var(--primary)] text-sm"
          >
            ✅ Shortened URL: <span className="font-mono">{success}</span>
          </motion.div>
        )}

        <motion.button
          type="submit"
          disabled={loading || !url.trim()}
          whileHover={{ scale: 1.02 }}
          whileTap={{ scale: 0.98 }}
          className="w-full px-6 py-4 bg-[var(--primary)] text-[var(--primary-foreground)] rounded-lg font-medium disabled:opacity-50 disabled:cursor-not-allowed transition-all hover:opacity-90"
        >
          {loading ? "⏳ Shortening..." : "✨ Shorten URL"}
        </motion.button>
      </form>
    </motion.div>
  );
}
