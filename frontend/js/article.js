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

// 添加文章功能
const addArticle = async (articleData) => {
    try {
        const accessToken = localStorage.getItem('access_token');
        if (!accessToken) {
            throw new Error('未登录，请先登录');
        }

        const response = await fetch('http://localhost:8888/article/add', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': 'Bearer ' + accessToken // 管理员token
            },
            body: JSON.stringify(articleData)
        });

        if (!response.ok) {
            const errorData = await response.json();
            throw new Error(errorData.message || '添加文章失败');
        }

        const data = await response.json();
        // 检查 status 字段是否为 2000
        if (data.status !== 2000) {
            throw new Error('添加文章失败');
        }
        console.log('文章添加成功:', data);
        alert('文章发布成功！');
        window.location.href = 'index.html'; // 发布成功后跳转到首页
        return data;
    } catch (error) {
        console.error('添加文章失败:', error);
        alert(error.message || '添加文章失败，请重试');
        if (error.message === '未登录，请先登录') {
            window.location.href = 'login.html'; // 跳转到登录页面
        }
        throw error;
    }
};

// 表单提交事件
document.getElementById('add-article-form').addEventListener('submit', async (e) => {
    e.preventDefault(); // 阻止表单默认提交行为

    // 获取表单数据
    const title = document.getElementById('article-title').value;
    const content = document.getElementById('article-content').value;
    const tags = document.getElementById('article-tags').value;

    // 构造请求体
    const articleData = {
        title: title,
        content: content,
        tags: tags,
    };

    // 调用后端接口
    try {
        await addArticle(articleData);
    } catch (error) {
        console.error('发布文章失败:', error);
    }
});