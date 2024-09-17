import React, { useEffect, useState } from 'react';

const UserListPage = () => {
    const [users, setUsers] = useState([]);
    const token = localStorage.getItem("token");

    useEffect(() => {
        const fetchUsers = async () => {
            const response = await fetch("http://localhost:8080/users", {
                headers: {
                    "Authorization": `Bearer ${token}`,
                },
            });

            if (response.ok) {
                const data = await response.json();
                setUsers(data);
            } else {
                alert("Failed to load users.");
            }
        };

        fetchUsers();
    }, [token]);

    return (
        <div>
            <h2>Users List</h2>
            <ul>
                {users.map((user) => (
                    <li key={user.id}>
                        {user.username}
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default UserListPage;