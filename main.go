package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	readability "codeberg.org/readeck/go-readability/v2"
	htmltomarkdown "github.com/JohannesKaufmann/html-to-markdown/v2"
)

func sanitizeFilename(name string) string {
	name = strings.TrimSpace(name)
	replacer := strings.NewReplacer(
		"/", "-",
		"\\", "-",
		":", "-",
		"*", "-",
		"?", "-",
		`"`, "-",
		"<", "-",
		">", "-",
		"|", "-",
	)
	name = replacer.Replace(name)
	name = strings.TrimSpace(name)
	if name == "" {
		return "output"
	}
	return name
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: readify-cli <url> [output-filename]")
		os.Exit(1)
	}

	urlStr := os.Args[1]
	outputFile := ""
	if len(os.Args) >= 3 {
		outputFile = os.Args[2]
	}

	article, err := readability.FromURL(urlStr, 30*time.Second, func(r *http.Request) {
		r.Header.Set("User-Agent", "Mozilla/5.0 (compatible; readify-cli/1.0)")
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching URL: %v\n", err)
		os.Exit(1)
	}

	var buf bytes.Buffer
	if err := article.RenderHTML(&buf); err != nil {
		fmt.Fprintf(os.Stderr, "Error rendering article HTML: %v\n", err)
		os.Exit(1)
	}

	markdown, err := htmltomarkdown.ConvertString(buf.String())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error converting to Markdown: %v\n", err)
		os.Exit(1)
	}

	if outputFile == "" {
		outputFile = sanitizeFilename(article.Title()) + ".md"
	}

	if err := os.WriteFile(outputFile, []byte(markdown), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Saved: %s\n", outputFile)
}
