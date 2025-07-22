# Book Downloader

A web app to host and download textbooks by semester.

## Features
- Browse textbooks by semester
- Download textbooks
- Modern, responsive UI

## Project Structure

```
Book downloader/
├── backend/
│   ├── main.go
│   ├── books/           # Place PDF files here, organized by semester
│   │   └── Semester 1/
│   │       └── mathematics1.pdf, physics1.pdf
│   │   └── Semester 2/
│   │       └── mathematics2.pdf, chemistry.pdf
│   └── data/
│       └── books.json   # Metadata for books
├── frontend/
│   ├── index.html
│   ├── style.css
│   └── script.js
└── README.md
```

## Setup

1. **Install Go** (if not already): https://golang.org/dl/
2. **Install Gin**:
   ```sh
   go get -u github.com/gin-gonic/gin
   ```
3. **Add your PDF textbooks**:
   - Place files in `backend/books/<Semester Name>/` matching the filenames in `backend/data/books.json`.
4. **Run the server**:
   ```sh
   cd backend
   go run main.go
   ```
5. **Open the app**:
   - Visit [http://localhost:8080](http://localhost:8080) in your browser.

## Customization
- To add more semesters/books, update `backend/data/books.json` and add corresponding PDF files.

---

**Enjoy your Book Downloader app!** 