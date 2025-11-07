"use client";

import { motion } from "framer-motion";
import { useState } from "react";

interface UrlCardProps {
  id: number;
  originalUrl: string;
  shortUrl: string;
  createdAt: string;
}

export default function UrlCard({
  originalUrl,
  shortUrl,
  createdAt,
}: UrlCardProps) {
  const [copied, setCopied] = useState(false);

  const handleCopy = async () => {
    try {
      await navigator.clipboard.writeText(shortUrl);
      setCopied(true);
      setTimeout(() => setCopied(false), 2000);
    } catch (err) {
      console.error("Failed to copy:", err);
    }
  };

  const formatDate = (dateString: string) => {
    const date = new Date(dateString);
    return date.toLocaleDateString("en-US", {
      year: "numeric",
      month: "short",
      day: "numeric",
      hour: "2-digit",
      minute: "2-digit",
    });
  };

  return (
    <motion.div
      initial={{ opacity: 0, y: 20 }}
      animate={{ opacity: 1, y: 0 }}
      exit={{ opacity: 0, y: -20 }}
      whileHover={{ scale: 1.02 }}
      className="bg-[var(--card)] border border-[var(--border)] rounded-lg p-6 space-y-4 transition-all hover:border-[var(--accent)]"
    >
      <div className="space-y-2">
        <div className="flex items-center justify-between">
          <p className="text-sm text-[var(--muted-foreground)] flex items-center gap-1">
            🌐 Original URL
          </p>
          <span className="text-xs text-[var(--muted-foreground)]">
            🕒 {formatDate(createdAt)}
          </span>
        </div>
        <a
          href={originalUrl}
          target="_blank"
          rel="noopener noreferrer"
          className="text-[var(--foreground)] hover:text-[var(--primary)] transition-colors break-all line-clamp-2"
        >
          {originalUrl}
        </a>
      </div>

      <div className="space-y-2">
        <p className="text-sm text-[var(--muted-foreground)] flex items-center gap-1">
          🔗 Short URL
        </p>
        <div className="flex items-center gap-2">
          <a
            href={shortUrl}
            target="_blank"
            rel="noopener noreferrer"
            className="text-[var(--primary)] hover:opacity-80 transition-colors font-mono flex-1 break-all"
          >
            {shortUrl}
          </a>
          <motion.button
            onClick={handleCopy}
            whileHover={{ scale: 1.1 }}
            whileTap={{ scale: 0.9 }}
            className="px-4 py-2 bg-[var(--secondary)] hover:bg-[var(--accent)] rounded-lg transition-colors text-sm"
          >
            {copied ? "✅ Copied!" : "📋 Copy"}
          </motion.button>
        </div>
      </div>
    </motion.div>
  );
}
