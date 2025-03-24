// 获取 DOM 元素
const loginForm = document.querySelector('.box .right');
const usernameInput = document.querySelector('.inputItem[placeholder="请输入用户名"]');
const passwordInput = document.querySelector('.inputItem[placeholder="请输入密码"]');
const loginButton = document.querySelector('.btn');
const errorMessage = document.createElement('div'); // 用于显示错误信息
errorMessage.style.color = '#ff4d4f'; // 错误信息颜色
errorMessage.style.marginTop = '10px'; // 错误信息间距
errorMessage.style.textAlign = 'center'; // 居中显示
loginForm.appendChild(errorMessage); // 将错误信息添加到表单中

// 登录按钮点击事件
loginButton.addEventListener('click', async (e) => {
    e.preventDefault(); // 阻止表单默认提交行为

    const username = usernameInput.value.trim();
    const password = passwordInput.value.trim();

    // 输入验证
    if (!username || !password) {
        errorMessage.textContent = '用户名和密码不能为空！';
        return;
    }

    // 显示加载状态
    loginButton.textContent = '登录中...';
    loginButton.disabled = true;

    try {
        // 发送登录请求
        const response = await fetch('http://localhost:8888/user/token', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                username: username,
                password: password,
            }),
        });

        // 处理响应
        if (!response.ok) {
            const errorData = await response.json();
            throw new Error(errorData.message || '登录失败，请检查用户名和密码');
        }

        const data = await response.json();

        // 检查 status 字段是否为 2000
        if (data.status !== 2000) {
            throw new Error('登录失败，请检查用户名和密码');
        }

        // 存储 token（示例使用 localStorage，生产环境建议使用更安全的方式）
        localStorage.setItem('access_token', data.data.token);
        localStorage.setItem('refresh_token', data.data.refresh_token);

        // 跳转到主页
        window.location.href = 'http://localhost:63342/blog/frontend/html/index.html';
    } catch (error) {
        // 显示错误信息
        errorMessage.textContent = error.message;
    } finally {
        // 恢复按钮状态
        loginButton.textContent = '登录';
        loginButton.disabled = false;
    }
});
