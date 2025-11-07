import Database from "better-sqlite3";
import path from "path";
import fs from "fs";

const dbPath = path.join(process.cwd(), "urls.db");

// Ensure the database file exists
if (!fs.existsSync(dbPath)) {
  fs.writeFileSync(dbPath, "");
}

const db = new Database(dbPath);

// Create table if it doesn't exist
db.exec(`
  CREATE TABLE IF NOT EXISTS urls (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    slug TEXT UNIQUE NOT NULL,
    original_url TEXT NOT NULL,
    created_at TEXT NOT NULL DEFAULT (datetime('now'))
  )
`);

export interface Url {
  id: number;
  slug: string;
  original_url: string;
  created_at: string;
}

export function createUrl(slug: string, originalUrl: string): Url {
  const stmt = db.prepare(
    "INSERT INTO urls (slug, original_url) VALUES (?, ?)"
  );
  const result = stmt.run(slug, originalUrl);

  const url = db
    .prepare("SELECT * FROM urls WHERE id = ?")
    .get(result.lastInsertRowid) as Url;
  return url;
}

export function getUrlBySlug(slug: string): Url | null {
  const stmt = db.prepare("SELECT * FROM urls WHERE slug = ?");
  const url = stmt.get(slug) as Url | undefined;
  return url || null;
}

export function getAllUrls(): Url[] {
  const stmt = db.prepare("SELECT * FROM urls ORDER BY created_at DESC");
  return stmt.all() as Url[];
}

export function slugExists(slug: string): boolean {
  const stmt = db.prepare("SELECT COUNT(*) as count FROM urls WHERE slug = ?");
  const result = stmt.get(slug) as { count: number };
  return result.count > 0;
}

export default db;

