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

// image
// 显示加载提示
function showLoading() {
    const previewDiv = document.getElementById('image-preview');
    previewDiv.innerHTML = '正在上传，请稍候...';
}

// 隐藏加载提示
function hideLoading() {
    const previewDiv = document.getElementById('image-preview');
    previewDiv.innerHTML = '';
}

// 处理成功上传结果
function handleUploadSuccess(data) {
    const previewDiv = document.getElementById('image-preview');
    if (!previewDiv) {
        console.error('未找到 image-preview 元素');
        return;
    }
    previewDiv.innerHTML = '上传成功，图片链接：<br>';
    data.data.forEach(url => {
        const link = document.createElement('a');
        link.href = url;
        link.target = '_blank';
        link.textContent = url;
        const br = document.createElement('br');
        previewDiv.appendChild(link);
        previewDiv.appendChild(br);
    });
}

// 处理上传失败结果
function handleUploadFailure(errorMessage) {
    alert(`上传失败：${errorMessage}`);
}

// 上传图片主函数
async function uploadImage() {
    const fileInput = document.getElementById('image-upload');
    if (fileInput.files.length === 0) {
        alert('请选择图片！');
        return;
    }

    const formData = new FormData();
    for (let i = 0; i < fileInput.files.length; i++) {
        formData.append('file', fileInput.files[i]);
    }

    // 从 localStorage 中获取 token
    const token = localStorage.getItem('access_token');
    if (!token) {
        alert('未找到有效的 token，请先登录');
        return;
    }

    showLoading();

    try {
        const response = await fetch('http://localhost:8888/image/upload', {
            method: 'POST',
            body: formData,
            headers: {
                // 添加 token 到请求头
                'Authorization': `Bearer ${token}`
            }
        });

        if (!response.ok) {
            const errorData = await response.json();
            handleUploadFailure(errorData.message);
            return;
        }

        const data = await response.json();
        console.log(data)
        if (typeof data === 'object' && data!== null && Array.isArray(data.data) && data.data.length > 0) {
            handleUploadSuccess(data);
        } else {
            handleUploadFailure('权限不足或到达次数限制');
        }
    } catch (error) {
        handleUploadFailure(`发生错误：${error.message}`);
    } finally {
        // hideLoading();
    }
}