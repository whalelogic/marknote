// Persist category open/close state across page loads
document.addEventListener('DOMContentLoaded', function () {
    const folders = document.querySelectorAll('details.category-folder');

    // Restore state from localStorage
    folders.forEach(folder => {
        const summary = folder.querySelector('.folder-name');
        const categoryName = summary ? summary.textContent.trim().split('(')[0].trim() : '';
        const savedState = localStorage.getItem('category-' + categoryName);

        if (savedState === 'open') {
            folder.setAttribute('open', '');
        }
    });

    // Save state on toggle
    folders.forEach(folder => {
        folder.addEventListener('toggle', function () {
            const summary = folder.querySelector('.folder-name');
            const categoryName = summary ? summary.textContent.trim().split('(')[0].trim() : '';

            if (folder.open) {
                localStorage.setItem('category-' + categoryName, 'open');
            } else {
                localStorage.removeItem('category-' + categoryName);
            }
        });
    });
});
