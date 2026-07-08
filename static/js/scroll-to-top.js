// "Go to top" floating button
document.addEventListener('DOMContentLoaded', function () {
    const btnScrollToTop = document.getElementById('btnScrollToTop');
    if (!btnScrollToTop) {
        return;
    }

    // Show the button once the page is scrolled down
    const toggleVisibility = function () {
        const scrolled = document.body.scrollTop > 300 || document.documentElement.scrollTop > 300;
        btnScrollToTop.style.display = scrolled ? 'block' : 'none';
    };

    window.addEventListener('scroll', toggleVisibility);
    toggleVisibility();

    // Smooth scroll to top on click
    btnScrollToTop.addEventListener('click', function () {
        window.scrollTo({ top: 0, left: 0, behavior: 'smooth' });
    });
});
