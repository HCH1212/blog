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

// 页面加载时获取所有标签
document.addEventListener('DOMContentLoaded', async () => {
    const tagList = document.getElementById('tag-list');
    try {
        const response = await fetch('http://localhost:8888/article/tags');
        if (!response.ok) {
            throw new Error('获取标签列表失败');
        }
        const data = await response.json();
        if (data.status === 2000) {
            const tags = data.data;
            tags.forEach(tag => {
                const tagLink = document.createElement('a');
                tagLink.href = '#';
                tagLink.textContent = tag;
                tagLink.addEventListener('click', (e) => {
                    e.preventDefault();
                    showArticlesByTag(tag);
                });
                tagList.appendChild(tagLink);
            });
        } else {
            console.error('获取标签列表失败:', data.message);
            const errorMsg = document.createElement('p');
            errorMsg.textContent = '获取标签列表失败，请稍后重试';
            tagList.appendChild(errorMsg);
        }
    } catch (error) {
        console.error('请求失败:', error);
        const errorMsg = document.createElement('p');
        errorMsg.textContent = '请求失败，请稍后重试';
        tagList.appendChild(errorMsg);
    }
});

// 显示特定标签的文章
async function showArticlesByTag(tag) {
    const articleList = document.getElementById('article-list');
    articleList.innerHTML = ''; // 清空之前的文章列表
    try {
        const response = await fetch(`http://localhost:8888/article/list/tag?tag=${tag}`);
        if (!response.ok) {
            throw new Error('获取文章列表失败');
        }
        const data = await response.json();
        if (data.status === 2000) {
            const articles = data.data;
            articles.forEach(article => {
                const articleDiv = document.createElement('div');
                articleDiv.classList.add('article-item');
                articleDiv.style.cursor = 'pointer'; // 设置鼠标指针为手型

                const title = document.createElement('h3');
                title.textContent = article.title;

                const content = document.createElement('p');
                content.textContent = article.content;

                articleDiv.appendChild(title);
                articleDiv.appendChild(content);

                // 为文章添加点击事件
                articleDiv.addEventListener('click', () => {
                    window.location.href = `article_one.html?id=${article.ID}`;
                });

                articleList.appendChild(articleDiv);
            });
        } else {
            console.error('获取文章列表失败:', data.message);
            const errorMsg = document.createElement('p');
            errorMsg.textContent = '获取文章列表失败，请稍后重试';
            articleList.appendChild(errorMsg);
        }
    } catch (error) {
        console.error('请求失败:', error);
        const errorMsg = document.createElement('p');
        errorMsg.textContent = '请求失败，请稍后重试';
        articleList.appendChild(errorMsg);
    }
}