# readify-cli

A CLI tool that fetches a URL, extracts the main article content (stripping ads, navigation, and clutter), and saves it as a Markdown file.

## How it works

1. Fetches the page at the given URL
2. Extracts the main readable content using [go-readability](https://codeberg.org/readeck/go-readability)
3. Converts the clean HTML to Markdown using [html-to-markdown](https://github.com/JohannesKaufmann/html-to-markdown)
4. Writes the result to a `.md` file

## Installation

```bash
go install github.com/adilkhash/readify-cli@latest
```

Or build from source:

```bash
git clone https://github.com/adilkhash/readify-cli
cd readify-cli
go build -o readify-cli .
```

## Usage

```
readify-cli <url> [output-filename]
```

**Auto-named output** (filename derived from the article title):

```bash
readify-cli https://en.wikipedia.org/wiki/Go_(programming_language)
# Saved: Go (programming language).md
```

**Explicit output filename:**

```bash
readify-cli https://en.wikipedia.org/wiki/Go_(programming_language) golang.md
# Saved: golang.md
```

## Requirements

- Go 1.22+
