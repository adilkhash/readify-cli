# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
go build ./...          # compile
go run main.go <url>    # run without building
go vet ./...            # static analysis
```

## Architecture

Single-file Go tool (`main.go`). No packages, no sub-directories.

**Pipeline:** URL → `readability.FromURL` (extracts article body) → `article.RenderHTML` → `htmltomarkdown.ConvertString` → write `.md` file.

**Key dependencies:**
- `codeberg.org/readeck/go-readability/v2` — strips boilerplate (ads, nav, etc.) from fetched HTML; used via `FromURL` which handles the HTTP request internally
- `github.com/JohannesKaufmann/html-to-markdown/v2` — converts the clean HTML to Markdown via `ConvertString`

**User-Agent:** A `User-Agent` header is set on every request via a `readability.RequestWith` modifier. Without it, many sites reject the request with "URL is not a HTML document".

**Filename fallback:** If no output filename is given, `sanitizeFilename(article.Title())` is used. Characters invalid in filenames (`/\:*?"<>|`) are replaced with `-`; empty result falls back to `"output"`.
