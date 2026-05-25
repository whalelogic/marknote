# MarkNote
### A simple, extensible markdown file viewer.

This web server serves all markdown files in the repository in a clean, browsable format.

<img width="1329" height="831" alt="image" src="https://github.com/user-attachments/assets/2120ed67-7ffc-477a-8910-f9eb8276edda" />


#### Why?

- 🗂️ Easily view and navigate your notes without leaving the browser
- 🚀 Fast, lightweight and minimal with no external dependencies. Perfect for extending with your own features.
- 📁 Automatically organizes files by directory structure
- 🔍 Search and filter capabilities (coming soon!)
- Secure and private - runs locally on your machine
- Just drop your `.md` files or entire folders in the `/static/` folder and they will be available on the web interface in an organized way.

## Features

- 📄 Automatically discovers all `.md` files in the repository
- 🎨 Clean, responsive UI with syntax highlighting
- 📂 Organized by directory/category
- 🔍 Easy navigation between files
- 📱 Mobile-friendly design

## Usage

From the repository root, run:

```bash
go run web/markdown-server.go
```

Then open your browser to: **http://localhost:8080**

## Structure

- `markdown-server.go` - Main server application
- `templates/` - HTML templates
  - `index.html` - Homepage listing all markdown files
  - `view.html` - Individual markdown file viewer
- `static/` - Static assets
  - `your_files/` - All your notes converted to HTML : )
  - `style.css` - Stylesheet for the entire site

## Dependencies

- `github.com/gomarkdown/markdown` - Markdown to HTML conversion

The server uses only Go standard library packages (`net/http`, `html/template`) plus the markdown parser.
