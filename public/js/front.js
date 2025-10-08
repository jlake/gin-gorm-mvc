// Front-end JavaScript

// Smooth scrolling and navigation
document.addEventListener('DOMContentLoaded', function() {
    // Add active class to current navigation item
    const currentPath = window.location.pathname;
    document.querySelectorAll('nav a').forEach(link => {
        if (link.getAttribute('href') === currentPath) {
            link.classList.add('active');
        }
    });

    // Smooth scroll for anchor links
    document.querySelectorAll('a[href^="#"]').forEach(anchor => {
        anchor.addEventListener('click', function (e) {
            e.preventDefault();
            const target = document.querySelector(this.getAttribute('href'));
            if (target) {
                target.scrollIntoView({
                    behavior: 'smooth',
                    block: 'start'
                });
            }
        });
    });

    // Add animation to cards on scroll
    const cards = document.querySelectorAll('.card');
    const observerOptions = {
        threshold: 0.1,
        rootMargin: '0px 0px -50px 0px'
    };

    const cardObserver = new IntersectionObserver(function(entries) {
        entries.forEach(entry => {
            if (entry.isIntersecting) {
                entry.target.style.opacity = '0';
                entry.target.style.transform = 'translateY(20px)';
                setTimeout(() => {
                    entry.target.style.transition = 'all 0.6s ease';
                    entry.target.style.opacity = '1';
                    entry.target.style.transform = 'translateY(0)';
                }, 100);
                cardObserver.unobserve(entry.target);
            }
        });
    }, observerOptions);

    cards.forEach(card => {
        cardObserver.observe(card);
    });
});

// API helper functions for front-end
const FrontAPI = {
    // Fetch all articles
    async getArticles(page = 1, limit = 100) {
        try {
            const response = await fetch(`/front/api/articles?page=${page}&limit=${limit}`);
            if (!response.ok) throw new Error('Failed to fetch articles');
            return await response.json();
        } catch (error) {
            console.error('Error fetching articles:', error);
            return null;
        }
    },

    // Fetch single article
    async getArticle(id) {
        try {
            const response = await fetch(`/front/api/articles/${id}`);
            if (!response.ok) throw new Error('Failed to fetch article');
            return await response.json();
        } catch (error) {
            console.error('Error fetching article:', error);
            return null;
        }
    },

    // Fetch user
    async getUser(id) {
        try {
            const response = await fetch(`/front/api/users/${id}`);
            if (!response.ok) throw new Error('Failed to fetch user');
            return await response.json();
        } catch (error) {
            console.error('Error fetching user:', error);
            return null;
        }
    }
};

// Format date helper
function formatDate(dateString) {
    const date = new Date(dateString);
    const options = { year: 'numeric', month: 'long', day: 'numeric', hour: '2-digit', minute: '2-digit' };
    return date.toLocaleDateString('ja-JP', options);
}

// Export for use in other scripts
if (typeof window !== 'undefined') {
    window.FrontAPI = FrontAPI;
    window.formatDate = formatDate;
}
