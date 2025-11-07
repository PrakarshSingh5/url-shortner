"use client";

import { useEffect, useState } from "react";
import { motion, AnimatePresence } from "framer-motion";
import UrlCard from "./UrlCard";

const DEFAULT_API_BASE_URL = "http://localhost:8080";
const API_BASE_URL = (
  process.env.NEXT_PUBLIC_API_URL ?? DEFAULT_API_BASE_URL
).replace(/\/$/, "");
const URLS_ENDPOINT = `${API_BASE_URL}/api/urls`;

interface Url {
  id: number;
  slug: string;
  original_url: string;
  short_url: string;
  created_at: string;
}

export default function UrlList() {
  const [urls, setUrls] = useState<Url[]>([]);
  const [loading, setLoading] = useState(true);

  const fetchUrls = async () => {
    try {
      const response = await fetch(URLS_ENDPOINT);
      if (response.ok) {
        const data = await response.json();
        setUrls(data);
      }
    } catch (error) {
      console.error("Error fetching URLs:", error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchUrls();

    // Listen for custom event to refresh list
    const handleRefresh = () => {
      fetchUrls();
    };

    window.addEventListener("url-shortened", handleRefresh);
    return () => {
      window.removeEventListener("url-shortened", handleRefresh);
    };
  }, []);

  if (loading) {
    return (
      <div className="w-full max-w-4xl">
        <div className="text-center text-[var(--muted-foreground)] py-8">
          ⏳ Loading...
        </div>
      </div>
    );
  }

  if (urls.length === 0) {
    return (
      <motion.div
        initial={{ opacity: 0 }}
        animate={{ opacity: 1 }}
        className="w-full max-w-4xl"
      >
        <div className="text-center text-[var(--muted-foreground)] py-8">
          📝 No shortened URLs yet. Create your first one above! 🚀
        </div>
      </motion.div>
    );
  }

  return (
    <div className="w-full max-w-4xl space-y-4">
      <h2 className="text-2xl font-semibold text-[var(--foreground)] mb-4 flex items-center gap-2">
        📋 Your Shortened URLs
      </h2>
      <AnimatePresence>
        {urls.map((url) => (
          <UrlCard
            key={url.id}
            id={url.id}
            originalUrl={url.original_url}
            shortUrl={url.short_url}
            createdAt={url.created_at}
          />
        ))}
      </AnimatePresence>
    </div>
  );
}
