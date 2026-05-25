package main

import (
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/whalelogic/marknote/internal/render"
)

type MarkdownFile struct {
	Name     string
	Path     string
	Category string
}

type PageData struct {
	Title          string
	Content        template.HTML
	Files          []MarkdownFile
	CurrentPath    string
	CategoryCounts map[string]int
}

var notesRoot string
var templates *template.Template

func mustGlob(pattern string) []string {
	matches, _ := filepath.Glob(pattern)
	return matches
}

func init() {
	var err error
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	notesRoot = filepath.Join(cwd, "static")
	if _, err := os.Stat(notesRoot); os.IsNotExist(err) {
		notesRoot = filepath.Join(cwd, "web/static")
	}
}

func main() {
	templatePath := "templates/*.html"
	if _, err := filepath.Glob(templatePath); err != nil || len(mustGlob(templatePath)) == 0 {
		templatePath = "web/templates/*.html"
	}
	templates = template.Must(template.New("").ParseGlob(templatePath))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/view/", viewHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(notesRoot))))

	log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	files := getAllMarkdownFiles()
	data := PageData{
		Title:          "Markdown Documentation Platform",
		Files:          files,
		CategoryCounts: getCategoryCounts(files),
	}
	templates.ExecuteTemplate(w, "index.html", data)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/view/")
	if path == "" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	// Prevent directory traversal
	cleanedPath := filepath.Clean(path)
	if strings.Contains(cleanedPath, "..") {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	fullPath := filepath.Join(notesRoot, cleanedPath)
	content, err := os.ReadFile(fullPath)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	htmlContent := render.MarkdownToHTML(content)
	files := getAllMarkdownFiles()

	data := PageData{
		Title:          prettifyName(filepath.Base(path)),
		Content:        template.HTML(htmlContent),
		Files:          files,
		CurrentPath:    cleanedPath,
		CategoryCounts: getCategoryCounts(files),
	}

	templates.ExecuteTemplate(w, "view.html", data)
}

func getAllMarkdownFiles() []MarkdownFile {
	var files []MarkdownFile

	filepath.WalkDir(notesRoot, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		// Skip the root notes directory itself
		if path == notesRoot {
			return nil
		}

		if d.IsDir() {
			// Skip internal folders
			name := d.Name()
			if name == ".git" || name == "bin" || name == "out" || name == "node_modules" {
				return filepath.SkipDir
			}
			return nil
		}

		if strings.HasSuffix(d.Name(), ".md") {
			relPath, _ := filepath.Rel(notesRoot, path)
			category := getCategoryFromPath(relPath)

			files = append(files, MarkdownFile{
				Name:     prettifyName(d.Name()),
				Path:     relPath,
				Category: category,
			})
		}

		return nil
	})

	sort.Slice(files, func(i, j int) bool {
		if files[i].Category != files[j].Category {
			return files[i].Category < files[j].Category
		}
		return files[i].Name < files[j].Name
	})

	return files
}

func getCategoryFromPath(path string) string {
	dir := filepath.Dir(path)
	if dir == "." {
		return "General"
	}
	parts := strings.Split(dir, string(filepath.Separator))
	return strings.Title(parts[0])
}

func prettifyName(filename string) string {
	name := strings.TrimSuffix(filename, ".md")
	name = strings.ReplaceAll(name, "-", " ")
	name = strings.ReplaceAll(name, "_", " ")
	return strings.Title(name)
}

func getCategoryCounts(files []MarkdownFile) map[string]int {
	counts := make(map[string]int)
	for _, file := range files {
		counts[file.Category]++
	}
	return counts
}
