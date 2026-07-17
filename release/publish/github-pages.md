# Publish target: github-pages

**English** · Agent-only. Human confirmation required first.

## When

Static or SPA (`web-vite` `dist/`). Repo must be on GitHub; human controls the repo settings.

## Prerequisites

- Human approved step 5 and chose GitHub Pages
- `gh` authenticated (`gh auth status`) or `GITHUB_TOKEN` with Pages write
- Agree: user/org site vs project site; branch (`gh-pages` / `main`) vs Actions

## Build

Same as Cloudflare Pages — produce `dist/`:

```bash
docker run --rm -v "$PWD:/app" -w /app node:22-alpine \
  sh -c "npm ci && npm run build"
```

For project Pages with a non-root base, ensure Vite `base` matches (e.g. `/<repo>/`) per PRO — confirm with the human if unsure.

## Publish (simple branch upload)

```bash
# Example: publish dist to branch gh-pages
rm -rf .gh-pages-tmp && mkdir .gh-pages-tmp
cp -R dist/. .gh-pages-tmp/
cd .gh-pages-tmp
git init
git checkout -b gh-pages
git add .
git -c user.email="agent@maker-flow.local" -c user.name="maker-flow-agent" commit -m "pages"
git remote add origin "https://github.com/<owner>/<repo>.git"
git push -f origin gh-pages
cd .. && rm -rf .gh-pages-tmp
```

Then ensure repo **Settings → Pages** source is branch `gh-pages` `/` (human may need to click once if never enabled).

Prefer GitHub Actions workflow if the human already has one — do not fight an existing pipeline.

## Verify

`https://<owner>.github.io/<repo>/` (or user site URL). Open `/`.

## Rollback

Force-push a previous known-good `dist` to `gh-pages`, or revert the Actions deploy.
