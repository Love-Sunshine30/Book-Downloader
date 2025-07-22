document.addEventListener('DOMContentLoaded', () => {
    const semesterSelect = document.getElementById('semesterSelect');
    const booksContainer = document.getElementById('booksContainer');

    // Fetch semesters
    fetch('/api/semesters')
        .then(res => res.json())
        .then(data => {
            data.semesters.forEach(sem => {
                const opt = document.createElement('option');
                opt.value = sem;
                opt.textContent = sem;
                semesterSelect.appendChild(opt);
            });
            if (data.semesters.length > 0) {
                semesterSelect.value = data.semesters[0];
                loadBooks(data.semesters[0]);
            }
        });

    semesterSelect.addEventListener('change', () => {
        loadBooks(semesterSelect.value);
    });

    function loadBooks(semester) {
        booksContainer.innerHTML = '<div class="text-center">Loading...</div>';
        fetch(`/api/books/${encodeURIComponent(semester)}`)
            .then(res => res.json())
            .then(data => {
                if (!data.books || data.books.length === 0) {
                    booksContainer.innerHTML = '<div class="alert alert-info">No books found for this semester.</div>';
                    return;
                }
                booksContainer.innerHTML = '';
                data.books.forEach(book => {
                    const card = document.createElement('div');
                    card.className = 'card mb-3 book-card';
                    card.innerHTML = `
                        <div class="card-body d-flex justify-content-between align-items-center">
                            <span>${book.title}</span>
                            <a href="/download/${encodeURIComponent(semester)}/${encodeURIComponent(book.filename)}" class="btn btn-success" download>Download</a>
                        </div>
                    `;
                    booksContainer.appendChild(card);
                });
            });
    }
}); 