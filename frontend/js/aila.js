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

// ai
document.addEventListener('DOMContentLoaded', function () {
    const chatMessages = document.getElementById('chat-messages');
    const chatInput = document.getElementById('chat-input');
    const sendButton = document.getElementById('send-button');

    sendButton.addEventListener('click', async function () {
        const userMessage = chatInput.value.trim();
        if (userMessage) {
            // 显示用户消息
            const userMessageElement = document.createElement('div');
            userMessageElement.classList.add('message', 'user-message');
            userMessageElement.textContent = userMessage;
            chatMessages.appendChild(userMessageElement);

            // 滚动到最新消息
            chatMessages.scrollTop = chatMessages.scrollHeight;

            try {
                // 向后端 AI 发送请求
                const response = await fetch('http://localhost:8888/ai/chat', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/x-www-form-urlencoded'
                    },
                    body: `content=${encodeURIComponent(userMessage)}`
                });

                if (!response.ok) {
                    throw new Error('请求失败');
                }

                const data = await response.json();
                if (data.status === 2000) {
                    // 显示 AI 回复消息
                    const replyMessage = data.data;
                    const replyMessageElement = document.createElement('div');
                    replyMessageElement.classList.add('message', 'reply-message');
                    replyMessageElement.textContent = replyMessage;
                    chatMessages.appendChild(replyMessageElement);
                } else {
                    console.error('AI 回复出错:', data.message);
                }
            } catch (error) {
                console.error('请求出错:', error);
            }

            // 清空输入框
            chatInput.value = '';
        }
    });

    chatInput.addEventListener('keydown', function (event) {
        if (event.key === 'Enter') {
            sendButton.click();
        }
    });
});