# fileorg

A small CLI that sorts files in a folder by extension.

## Usage

```bash
go run . <directory> --dry-run
go run . <directory> --apply
```bash


`--dry-run` prints what would move. `--apply` actually moves files.

Files go into Images, Documents, Videos, Code, and Others.

## Why

My downloads folder gets messy. I wrote this to practice Go and to have something that previews before it touches anything.

## Notes

- Dry-run is the default.
- Only top-level files, subfolders are skipped.
- Uses os.Rename, so it stays on the same disk.
