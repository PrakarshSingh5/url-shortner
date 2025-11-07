"use client";

import { motion } from "framer-motion";
import UrlForm from "@/components/UrlForm";
import UrlList from "@/components/UrlList";
import ThemeToggle from "@/components/ThemeToggle";

export default function Home() {
  const handleUrlShortened = () => {
    // Trigger refresh of URL list
    window.dispatchEvent(new Event("url-shortened"));
  };

  const features = [
    {
      icon: "⚡",
      title: "Lightning Fast",
      description: "Instant URL shortening",
    },
    { icon: "🔒", title: "Secure", description: "Safe and reliable links" },
    { icon: "📊", title: "Simple", description: "Easy to use interface" },
    { icon: "🎨", title: "Beautiful", description: "Clean and modern design" },
  ];

  return (
    <main className="min-h-screen bg-[var(--background)] px-4 py-12 md:py-20 transition-colors">
      <ThemeToggle />

      <div className="max-w-6xl mx-auto space-y-16">
        {/* Hero Section */}
        <motion.div
          initial={{ opacity: 0, y: -20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.5 }}
          className="text-center space-y-6"
        >
          <div className="flex justify-center items-center gap-3 mb-4">
            <span className="text-6xl md:text-8xl">🔗</span>
          </div>
          <h1 className="text-4xl md:text-6xl font-bold text-[var(--foreground)]">
            URL Shortener
          </h1>
          <p className="text-lg md:text-xl text-[var(--muted-foreground)] max-w-2xl mx-auto">
            Transform long URLs into short, shareable links. Clean, fast, and
            simple. ✨
          </p>
        </motion.div>

        {/* Features Section */}
        <motion.div
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.5, delay: 0.2 }}
          className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6"
        >
          {features.map((feature, index) => (
            <motion.div
              key={feature.title}
              initial={{ opacity: 0, y: 20 }}
              animate={{ opacity: 1, y: 0 }}
              transition={{ duration: 0.5, delay: 0.3 + index * 0.1 }}
              whileHover={{ scale: 1.05, y: -5 }}
              className="bg-[var(--card)] border border-[var(--border)] rounded-lg p-6 text-center space-y-3 transition-all hover:border-[var(--primary)]"
            >
              <div className="text-4xl">{feature.icon}</div>
              <h3 className="text-lg font-semibold text-[var(--foreground)]">
                {feature.title}
              </h3>
              <p className="text-sm text-[var(--muted-foreground)]">
                {feature.description}
              </p>
            </motion.div>
          ))}
        </motion.div>

        {/* Form Section */}
        <motion.div
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.5, delay: 0.4 }}
          className="flex justify-center"
        >
          <div className="w-full max-w-2xl">
            <div className="text-center mb-6">
              <h2 className="text-2xl md:text-3xl font-bold text-[var(--foreground)] mb-2">
                🚀 Get Started
              </h2>
              <p className="text-[var(--muted-foreground)]">
                Paste your long URL below and get a short link instantly!
              </p>
            </div>
            <UrlForm onSuccess={handleUrlShortened} />
          </div>
        </motion.div>

        {/* URL List Section */}
        <motion.div
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.5, delay: 0.5 }}
          className="flex justify-center"
        >
          <UrlList />
        </motion.div>

        {/* Footer Section */}
        <motion.div
          initial={{ opacity: 0 }}
          animate={{ opacity: 1 }}
          transition={{ duration: 0.5, delay: 0.6 }}
          className="text-center py-8 border-t border-[var(--border)]"
        >
          <p className="text-[var(--muted-foreground)]">
            Made with ❤️ using Next.js, TypeScript, and Tailwind CSS
          </p>
        </motion.div>
      </div>
    </main>
  );
}
