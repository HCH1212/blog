// 主页初始化逻辑
document.addEventListener('DOMContentLoaded', () => {
    // 检查用户是否已登录
    const accessToken = localStorage.getItem('access_token');
    const refreshToken = localStorage.getItem('refresh_token');

    if (accessToken && refreshToken) {
        // 用户已登录，启动 Token 刷新定时器
        startTokenRefreshTimer(6 * 24 * 60 * 60 * 1000); // 每六天刷新一次
    }
});

// 启动 Token 刷新定时器
function startTokenRefreshTimer(interval) {
    // 设置定时器，每六天刷新一次 Token
    setInterval(refreshToken, interval);
}

// 刷新 Token
async function refreshToken() {
    const refreshToken = localStorage.getItem('refresh_token');
    if (!refreshToken) {
        console.error('未找到 refresh_token，请重新登录');
        return;
    }

    try {
        const response = await fetch('http://localhost:8888/user/refresh_token', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                refresh_token: refreshToken,
            }),
        });

        if (!response.ok) {
            throw new Error('刷新 Token 失败');
        }

        const data = await response.json();

        // 检查 status 字段是否为 2000
        if (data.status !== 2000) {
            throw new Error('刷新 Token 失败');
        }

        // 更新 Token
        localStorage.setItem('access_token', data.access_token);
        localStorage.setItem('refresh_token', data.refresh_token);

        console.log('Token 刷新成功');
    } catch (error) {
        console.error('刷新 Token 失败:', error.message);
        // 如果刷新失败，提示用户重新登录
        alert('登录已过期，请重新登录');
        localStorage.removeItem('access_token');
        localStorage.removeItem('refresh_token');
    }
}

// 检查登录状态
function checkLoginStatus() {
    const accessToken = localStorage.getItem('access_token');
    const loginLink = document.getElementById('login-link');
    const logoutLink = document.getElementById('logout-link');

    if (accessToken) {
        // 用户已登录，隐藏登录按钮，显示退出登录按钮
        loginLink.style.display = 'none';
        logoutLink.style.display = 'inline-block';
    } else {
        // 用户未登录，显示登录按钮，隐藏退出登录按钮
        loginLink.style.display = 'inline-block';
        logoutLink.style.display = 'none';
    }
}

// 页面加载时检查登录状态
document.addEventListener('DOMContentLoaded', checkLoginStatus);

// 退出登录功能
document.getElementById('logout-link').addEventListener('click', (e) => {
    e.preventDefault(); // 阻止默认跳转行为

    // 清除本地存储的 Token
    localStorage.removeItem('access_token');
    localStorage.removeItem('refresh_token');

    // 刷新页面
    window.location.reload();

    // 可选：跳转到登录页面
    // window.location.href = './login.html';
});

// 分页相关变量
let currentPage = 1;
let totalPages = 1;
const pageSize = 10;
let isLoading = false;

// 初始化分页功能
function initPagination() {
    // 绑定分页按钮事件
    document.getElementById('first-page').addEventListener('click', () => goToPage(1));
    document.getElementById('prev-page').addEventListener('click', () => goToPage(currentPage - 1));
    document.getElementById('next-page').addEventListener('click', () => goToPage(currentPage + 1));
    document.getElementById('last-page').addEventListener('click', () => goToPage(totalPages));

    // 初始加载第一页数据
    loadArticles(currentPage);
}

// 加载文章数据
function loadArticles(page) {
    if (isLoading || page < 1 || page > totalPages) return;

    isLoading = true;
    updatePaginationState();

    // 显示加载状态
    const articlesList = document.getElementById('articles-list');
    articlesList.innerHTML = '<div class="loading">加载中...</div>';

    fetch(`http://localhost:8888/article/list?page=${page}&pageSize=${pageSize}`)
        .then(response => {
            if (!response.ok) throw new Error('网络请求失败');
            return response.json();
        })
        .then(data => {
            if (data.status !== 2000) throw new Error(data.message || '获取数据失败');

            // 更新分页信息
            totalPages = Math.max(1, Math.ceil(data.data.total / pageSize));
            currentPage = Math.max(1, Math.min(page, totalPages));

            // 渲染文章列表
            renderArticles(data.data.articles);

            // 更新分页UI
            updatePaginationUI();
        })
        .catch(error => {
            console.error('请求失败:', error);
            document.getElementById('articles-list').innerHTML =
                `<div class="error">加载失败: ${error.message}</div>`;
        })
        .finally(() => {
            isLoading = false;
            updatePaginationState();
        });
}

// 渲染文章列表
function renderArticles(articles) {
    const articlesList = document.getElementById('articles-list');
    articlesList.innerHTML = '';

    if (!articles || articles.length === 0) {
        articlesList.innerHTML = '<div class="no-data">暂无文章</div>';
        return;
    }

    // 渲染每篇文章
    articles.forEach(article => {
        const articleCard = document.createElement('div');
        articleCard.className = 'article-card';

        const articleLink = document.createElement('a');
        articleLink.href = `article_one.html?id=${article.ID}`;
        articleLink.className = 'article-link';

        // 文章标题
        const title = document.createElement('h2');
        title.className = 'article-title';
        title.textContent = article.title || '无标题';

        // 文章内容
        const content = document.createElement('div');
        content.className = 'article-content';
        content.innerHTML = formatContent(article.content);

        // 文章元信息
        const meta = document.createElement('div');
        meta.className = 'article-meta';
        meta.textContent = formatDate(article.CreatedAt);

        // 文章标签
        const tags = document.createElement('div');
        tags.className = 'article-tags';
        tags.innerHTML = formatTags(article.tags);

        // 组合元素
        articleLink.appendChild(title);
        articleLink.appendChild(content);
        articleLink.appendChild(meta);
        articleLink.appendChild(tags);
        articleCard.appendChild(articleLink);
        articlesList.appendChild(articleCard);
    });
}

// 辅助函数：格式化内容
function formatContent(text) {
    if (!text) return '无内容';
    const content = text.length > 50 ? text.substring(0, 50) + '...' : text;
    return content.replace(/\n/g, ' ');
}

// 辅助函数：格式化标签
function formatTags(tagString) {
    if (!tagString || tagString.trim() === '') return '';
    return tagString.split(',').map(tag => `<span>${tag.trim()}</span>`).join('');
}

// 辅助函数：格式化日期
function formatDate(dateString) {
    if (!dateString) return '发布于: 未知时间';
    const date = new Date(dateString);
    return `发布于: ${date.toLocaleDateString()} ${date.toLocaleTimeString()}`;
}

// 更新分页UI
function updatePaginationUI() {
    // 更新页码信息
    document.getElementById('current-page').textContent = currentPage;
    document.getElementById('total-pages').textContent = totalPages;

    // 更新页码按钮
    const pageNumbers = document.getElementById('page-numbers');
    pageNumbers.innerHTML = '';

    // 计算显示的页码范围（当前页居中，最多显示5个页码）
    let startPage = Math.max(1, currentPage - 2);
    let endPage = Math.min(totalPages, currentPage + 2);

    // 调整页码范围确保显示足够数量的页码
    while (endPage - startPage < 4 && endPage < totalPages) endPage++;
    while (endPage - startPage < 4 && startPage > 1) startPage--;

    // 生成页码按钮
    for (let i = startPage; i <= endPage; i++) {
        const pageBtn = document.createElement('button');
        pageBtn.textContent = i;
        pageBtn.className = 'page-number';
        if (i === currentPage) {
            pageBtn.classList.add('active');
            pageBtn.disabled = true;
        }
        pageBtn.addEventListener('click', () => goToPage(i));
        pageNumbers.appendChild(pageBtn);
    }
}

// 更新分页按钮状态
function updatePaginationState() {
    const controls = ['first-page', 'prev-page', 'next-page', 'last-page'];
    const states = {
        'first-page': currentPage === 1,
        'prev-page': currentPage === 1,
        'next-page': currentPage === totalPages,
        'last-page': currentPage === totalPages
    };

    controls.forEach(id => {
        const btn = document.getElementById(id);
        if (btn) btn.disabled = isLoading || states[id];
    });
}

// 跳转到指定页
function goToPage(page) {
    if (!isLoading && page >= 1 && page <= totalPages && page !== currentPage) {
        currentPage = page;
        loadArticles(page);
    }
}

// 在页面加载完成后调用 initPagination 函数
document.addEventListener('DOMContentLoaded', initPagination);