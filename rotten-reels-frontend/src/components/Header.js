import React from 'react';
import { AppBar, Toolbar, Typography, Button } from '@mui/material';
import { Link, useNavigate } from 'react-router-dom';

const Header = ({ user }) => {
    const navigate = useNavigate();

    const handleLogout = () => {
        localStorage.removeItem("token");
        localStorage.removeItem("user");
        navigate("/login");
    };

    return (
        <AppBar position="static">
            <Toolbar>
                <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
                    Movie App
                </Typography>
                <Button color="inherit" component={Link} to="/movies">
                    Movies
                </Button>
                <Button color="inherit" component={Link} to="/users">
                    Users
                </Button>
                {user ? (
                    <>
                        <Typography variant="body1" component="div">
                            Logged in as: {user}
                        </Typography>
                        <Button color="inherit" onClick={handleLogout}>
                            Logout
                        </Button>
                    </>
                ) : (
                    <>
                        <Button color="inherit" component={Link} to="/login">
                            Login
                        </Button>
                        <Button color="inherit" component={Link} to="/register">
                            Register {/* Кнопка регистрации */}
                        </Button>
                    </>
                )}
            </Toolbar>
        </AppBar>
    );
};

export default Header;