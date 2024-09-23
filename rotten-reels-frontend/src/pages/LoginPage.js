import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';

const LoginPage = ({ setUser }) => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState(null);
    const navigate = useNavigate();

    const handleLogin = async (e) => {
        e.preventDefault();
        setError(null);

        try {
            const response = await fetch('http://localhost:8080/auth/login', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ username, password }),
            });

            if (response.ok) {
                // Получаем токен из ответа
                const { token } = await response.json();

                // Логируем токен для отладки
                console.log("Получен токен:", token);

                // Сохраняем токен в localStorage
                localStorage.setItem('token', token);

                // Парсим payload токена (вторая часть JWT)
                const base64Url = token.split('.')[1];
                const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
                const jsonPayload = decodeURIComponent(atob(base64).split('').map(function(c) {
                    return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
                }).join(''));

                const userInfo = JSON.parse(jsonPayload);

                // Логируем userInfo для отладки
                console.log("Декодированная информация о пользователе:", userInfo);

                // Обновляем состояние пользователя
                setUser(userInfo);

                // Перенаправляем на главную страницу
                navigate('/movies');
            } else {
                setError('Invalid credentials');
            }
        } catch (error) {
            console.error("Ошибка при логине:", error);
            setError('An error occurred during login.');
        }
    };

    return (
        <div>
            <h2>Login</h2>
            <form onSubmit={handleLogin}>
                <div>
                    <label>Username:</label>
                    <input
                        type="text"
                        value={username}
                        onChange={(e) => setUsername(e.target.value)}
                    />
                </div>
                <div>
                    <label>Password:</label>
                    <input
                        type="password"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                    />
                </div>
                {error && <p>{error}</p>}
                <button type="submit">Login</button>
            </form>
        </div>
    );
};

export default LoginPage;