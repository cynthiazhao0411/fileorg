package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var categories = map[string]string{
	".jpg": "Images", ".jpeg": "Images", ".png": "Images", ".gif": "Images", ".webp": "Images",
	".pdf": "Documents", ".doc": "Documents", ".docx": "Documents", ".txt": "Documents", ".md": "Documents",
	".mp4": "Videos", ".mov": "Videos", ".avi": "Videos",
	".go": "Code", ".py": "Code", ".js": "Code", ".ts": "Code", ".java": "Code", ".c": "Code", ".cpp": "Code",
}

func main() {
	dryRun := flag.Bool("dry-run", true, "preview moves without changing files")
	apply := flag.Bool("apply", false, "actually move files")
	flag.Parse()

	dir := flag.Arg(0)
	if dir == "" {
		dir = "."
	}

	if *apply {
		*dryRun = false
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		os.Exit(1)
	}

	moved := 0
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		ext := strings.ToLower(filepath.Ext(name))
		cat, ok := categories[ext]
		if !ok {
			cat = "Others"
		}

		src := filepath.Join(dir, name)
		destDir := filepath.Join(dir, cat)
		dest := filepath.Join(destDir, name)

		if *dryRun {
			fmt.Printf("[dry-run] %s -> %s\n", src, dest)
			continue
		}

		if err := os.MkdirAll(destDir, 0755); err != nil {
			fmt.Println("Error creating folder:", err)
			continue
		}
		if err := os.Rename(src, dest); err != nil {
			fmt.Println("Error moving file:", err)
			continue
		}
		fmt.Printf("Moved %s -> %s\n", src, dest)
		moved++
	}

	if *dryRun {
		fmt.Println("\nDry run complete. Use --apply to actually move files.")
	} else {
		fmt.Printf("\nDone. Moved %d files.\n", moved)
	}
}