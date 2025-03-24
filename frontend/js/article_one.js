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

document.addEventListener('DOMContentLoaded', () => {
    // 从URL中获取文章ID
    const urlParams = new URLSearchParams(window.location.search);
    const articleId = urlParams.get('id');

    if (!articleId) {
        console.error('未找到文章ID');
        alert('文章ID缺失，请返回首页');
        window.location.href = 'index.html';
        return;
    }

    // 获取文章详情
    fetch(`http://localhost:8888/article/get?id=${articleId}`)
        .then(response => response.json())
        .then(data => {
            if (data.status === 2000) {
                const article = data.data;

                // 渲染文章详情
                document.getElementById('article-title').textContent = article.title;
                document.getElementById('article-content').textContent = article.content;

                // 渲染标签
                const tagsContainer = document.getElementById('article-tags');
                tagsContainer.innerHTML = ''; // 清空原有标签
                const tags = article.tags? article.tags.split(',') : []; // 安全处理
                tags.forEach(tag => {
                    const tagElement = document.createElement('span');
                    tagElement.className = 'tag';
                    tagElement.textContent = tag.trim();
                    tagsContainer.appendChild(tagElement);
                });

                // 显示编辑和删除按钮
                const articleActions = document.getElementById('article-actions');
                articleActions.style.display = 'flex';

                // 编辑文章按钮点击事件
                const editArticleBtn = document.getElementById('edit-article-btn');
                let saveBtn = null;
                editArticleBtn.addEventListener('click', async () => {
                    if (saveBtn) return; // 如果保存按钮已经存在，直接返回
                    // 切换到编辑模式（这里简单地展示输入框，实际需要更复杂的逻辑）
                    const titleElement = document.getElementById('article-title');
                    const contentElement = document.getElementById('article-content');
                    const tagsElement = document.getElementById('article-tags');
                    titleElement.contentEditable = true;
                    contentElement.contentEditable = true;
                    tagsElement.contentEditable = true;

                    // 保存修改按钮（假设添加一个新按钮）
                    saveBtn = document.createElement('button');
                    saveBtn.textContent = '保存修改';
                    saveBtn.addEventListener('click', async () => {
                        const title = titleElement.textContent;
                        const content = contentElement.innerText;
                        const tags = Array.from(tagsElement.children).map(tag => tag.textContent).join(',');
                        const accessToken = localStorage.getItem('access_token');

                        try {
                            const response = await fetch(`http://localhost:8888/article/update?id=${articleId}`, {
                                method: 'POST',
                                headers: {
                                    'Content-Type': 'application/json',
                                    'Authorization': `Bearer ${accessToken}`
                                },
                                body: JSON.stringify({
                                    title: title,
                                    content: content,
                                    tags: tags
                                })
                            });

                            if (!response.ok) {
                                throw new Error('修改文章失败');
                            }

                            const data = await response.json();
                            if (data.status === 2000) {
                                alert('文章修改成功');
                                // 保存成功后退出编辑模式
                                titleElement.contentEditable = false;
                                contentElement.contentEditable = false;
                                tagsElement.contentEditable = false;
                                saveBtn.remove();
                                saveBtn = null;
                            } else {
                                alert('文章修改失败');
                                // 保存失败后移除保存按钮
                                titleElement.contentEditable = false;
                                contentElement.contentEditable = false;
                                tagsElement.contentEditable = false;
                                saveBtn.remove();
                                saveBtn = null;
                            }
                        } catch (error) {
                            console.error('修改文章失败:', error);
                            alert('文章修改失败');
                            // 保存失败后移除保存按钮
                            titleElement.contentEditable = false;
                            contentElement.contentEditable = false;
                            tagsElement.contentEditable = false;
                            saveBtn.remove();
                            saveBtn = null;
                        }
                    });

                    articleActions.appendChild(saveBtn);
                });

                // 删除文章按钮点击事件
                const deleteArticleBtn = document.getElementById('delete-article-btn');
                deleteArticleBtn.addEventListener('click', async () => {
                    const confirmDelete = confirm('确定要删除这篇文章吗？');
                    if (confirmDelete) {
                        const accessToken = localStorage.getItem('access_token');
                        try {
                            const response = await fetch(`http://localhost:8888/article/delete?id=${articleId}`, {
                                method: 'POST',
                                headers: {
                                    'Content-Type': 'application/json',
                                    'Authorization': `Bearer ${accessToken}`
                                },
                            });

                            if (!response.ok) {
                                throw new Error('删除文章失败');
                            }

                            const data = await response.json();
                            if (data.status === 2000) {
                                alert('文章删除成功');
                                // 跳转到文章列表页
                                window.location.href = 'index.html';
                            } else {
                                alert('文章删除失败');
                            }
                        } catch (error) {
                            console.error('删除文章失败:', error);
                            alert('文章删除失败');
                        }
                    }
                });
            } else {
                console.error('获取文章详情失败:', data.message);
                alert('获取文章详情失败，请稍后重试');
            }
        })
        .catch(error => {
            console.error('请求失败:', error);
            alert('请求失败，请稍后重试');
        });
});
    