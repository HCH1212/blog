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