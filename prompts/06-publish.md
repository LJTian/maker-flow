# Step 6 — publish (dialogue)

**English** · Agent loads this after step-5 approval. Skill: `skills/deploy.md`.

## Say to the human (adapt language)

MVP acceptance passed. Before going public, please confirm:

1. **What** should go live? (whole app / frontend only / API only)
2. **Where** should we publish?
   - Cloudflare Pages / GitHub Pages / Vercel — good for static or SPA
   - Your own VPS (Docker gateway) — good for APIs, workers, or full compose
   - Split is OK (e.g. frontend on Pages + API on VPS)
3. **Domain** preference? (platform default URL vs a hostname you already own)
4. Are platform logins / tokens ready on this machine?

I will run the publish steps myself after you choose — you do not need a deploy CLI command.

## After answers

Follow `release/publish/<target>.md`. Return public URL(s) and verification notes.
