# Publish target: vercel

**English** · Agent-only. Human confirmation required first.

## When

Static or SPA (`web-vite`). Optional later: Vercel-native SSR — only if the PRO and stack match; do not invent Next.js.

## Prerequisites

- Human approved step 5 and chose Vercel
- `vercel` CLI available (`npx vercel`) and human confirmed login (`vercel whoami`) **or** `VERCEL_TOKEN` in env
- Link/create project name agreed in chat

## Build & publish

From the web app directory:

```bash
# Non-interactive when token present
npx vercel pull --yes --environment=production ${VERCEL_TOKEN:+--token "$VERCEL_TOKEN"}
npx vercel build ${VERCEL_TOKEN:+--token "$VERCEL_TOKEN"}
npx vercel deploy --prebuilt --prod ${VERCEL_TOKEN:+--token "$VERCEL_TOKEN"}
```

Or one-shot (CLI builds):

```bash
npx vercel --prod --yes ${VERCEL_TOKEN:+--token "$VERCEL_TOKEN"}
```

Framework preset: Vite. Output `dist`. Set root directory if the app lives under `web/`.

Custom domain: add in Vercel project Domains after human confirms the hostname.

## Verify

Open the production URL from the CLI output. Check `/`.

## Rollback

`vercel rollback` or promote a previous deployment in the Vercel dashboard per human request.
