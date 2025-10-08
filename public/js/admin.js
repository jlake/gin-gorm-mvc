// Admin Panel JavaScript

// Admin initialization
document.addEventListener('DOMContentLoaded', function() {
    // Add active class to current navigation item
    const currentPath = window.location.pathname;
    document.querySelectorAll('nav a').forEach(link => {
        const href = link.getAttribute('href');
        if (href === currentPath || (currentPath.startsWith(href) && href !== '/')) {
            link.classList.add('active');
        }
    });

    // Initialize table sorting if table exists
    const tables = document.querySelectorAll('table');
    if (tables.length > 0) {
        initTableFeatures();
    }

    // Add confirmation for delete actions
    document.querySelectorAll('[data-confirm]').forEach(element => {
        element.addEventListener('click', function(e) {
            if (!confirm(this.getAttribute('data-confirm'))) {
                e.preventDefault();
            }
        });
    });
});

// Table features
function initTableFeatures() {
    // Add hover effect and row selection
    const rows = document.querySelectorAll('tbody tr');
    rows.forEach(row => {
        row.style.cursor = 'pointer';
        row.addEventListener('click', function(e) {
            // Don't trigger if clicking on a link or button
            if (e.target.tagName === 'A' || e.target.tagName === 'BUTTON') {
                return;
            }
            this.classList.toggle('selected');
        });
    });
}

// Admin API helper functions
const AdminAPI = {
    // User Management
    async getUsers(page = 1, limit = 100) {
        try {
            const response = await fetch(`/admin/api/users?page=${page}&limit=${limit}`);
            if (!response.ok) throw new Error('Failed to fetch users');
            return await response.json();
        } catch (error) {
            console.error('Error fetching users:', error);
            showNotification('ユーザーの取得に失敗しました', 'error');
            return null;
        }
    },

    async createUser(userData) {
        try {
            const response = await fetch('/admin/api/users', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(userData)
            });
            if (!response.ok) throw new Error('Failed to create user');
            showNotification('ユーザーを作成しました', 'success');
            return await response.json();
        } catch (error) {
            console.error('Error creating user:', error);
            showNotification('ユーザーの作成に失敗しました', 'error');
            return null;
        }
    },

    async updateUser(id, userData) {
        try {
            const response = await fetch(`/admin/api/users/${id}`, {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(userData)
            });
            if (!response.ok) throw new Error('Failed to update user');
            showNotification('ユーザーを更新しました', 'success');
            return await response.json();
        } catch (error) {
            console.error('Error updating user:', error);
            showNotification('ユーザーの更新に失敗しました', 'error');
            return null;
        }
    },

    async deleteUser(id) {
        try {
            const response = await fetch(`/admin/api/users/${id}`, {
                method: 'DELETE'
            });
            if (!response.ok) throw new Error('Failed to delete user');
            showNotification('ユーザーを削除しました', 'success');
            return true;
        } catch (error) {
            console.error('Error deleting user:', error);
            showNotification('ユーザーの削除に失敗しました', 'error');
            return false;
        }
    },

    // Article Management
    async getArticles(page = 1, limit = 100) {
        try {
            const response = await fetch(`/admin/api/articles?page=${page}&limit=${limit}`);
            if (!response.ok) throw new Error('Failed to fetch articles');
            return await response.json();
        } catch (error) {
            console.error('Error fetching articles:', error);
            showNotification('記事の取得に失敗しました', 'error');
            return null;
        }
    },

    async createArticle(articleData) {
        try {
            const response = await fetch('/admin/api/articles', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(articleData)
            });
            if (!response.ok) throw new Error('Failed to create article');
            showNotification('記事を作成しました', 'success');
            return await response.json();
        } catch (error) {
            console.error('Error creating article:', error);
            showNotification('記事の作成に失敗しました', 'error');
            return null;
        }
    },

    async updateArticle(id, articleData) {
        try {
            const response = await fetch(`/admin/api/articles/${id}`, {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(articleData)
            });
            if (!response.ok) throw new Error('Failed to update article');
            showNotification('記事を更新しました', 'success');
            return await response.json();
        } catch (error) {
            console.error('Error updating article:', error);
            showNotification('記事の更新に失敗しました', 'error');
            return null;
        }
    },

    async deleteArticle(id) {
        try {
            const response = await fetch(`/admin/api/articles/${id}`, {
                method: 'DELETE'
            });
            if (!response.ok) throw new Error('Failed to delete article');
            showNotification('記事を削除しました', 'success');
            return true;
        } catch (error) {
            console.error('Error deleting article:', error);
            showNotification('記事の削除に失敗しました', 'error');
            return false;
        }
    }
};

// Notification system
function showNotification(message, type = 'info') {
    // Create notification element
    const notification = document.createElement('div');
    notification.className = `notification notification-${type}`;
    notification.textContent = message;
    notification.style.cssText = `
        position: fixed;
        top: 20px;
        right: 20px;
        padding: 1rem 1.5rem;
        background: ${type === 'success' ? '#28a745' : type === 'error' ? '#dc3545' : '#17a2b8'};
        color: white;
        border-radius: 5px;
        box-shadow: 0 4px 12px rgba(0,0,0,0.15);
        z-index: 10000;
        animation: slideIn 0.3s ease;
    `;

    document.body.appendChild(notification);

    // Remove after 3 seconds
    setTimeout(() => {
        notification.style.animation = 'slideOut 0.3s ease';
        setTimeout(() => notification.remove(), 300);
    }, 3000);
}

// Format date helper
function formatDate(dateString) {
    const date = new Date(dateString);
    const options = { year: 'numeric', month: '2-digit', day: '2-digit', hour: '2-digit', minute: '2-digit' };
    return date.toLocaleDateString('ja-JP', options);
}

// Export for use in other scripts
if (typeof window !== 'undefined') {
    window.AdminAPI = AdminAPI;
    window.showNotification = showNotification;
    window.formatDate = formatDate;
}

// Add CSS animations
const style = document.createElement('style');
style.textContent = `
    @keyframes slideIn {
        from {
            transform: translateX(100%);
            opacity: 0;
        }
        to {
            transform: translateX(0);
            opacity: 1;
        }
    }
    @keyframes slideOut {
        from {
            transform: translateX(0);
            opacity: 1;
        }
        to {
            transform: translateX(100%);
            opacity: 0;
        }
    }
`;
document.head.appendChild(style);
