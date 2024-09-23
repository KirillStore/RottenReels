import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import LoginPage from './pages/LoginPage';
import RegisterPage from './pages/RegisterPage';
import MoviesListPage from './pages/MoviesListPage';
import MovieDetailPage from './pages/MovieDetailPage';
import UserListPage from './pages/UserListPage';
import Header from './components/Header';

const App = () => {
    const user = localStorage.getItem("user");

    return (
        <Router>
            <Header user={user} />
            <Routes>
                <Route path="/login" element={<LoginPage />} />
                <Route path="/register" element={<RegisterPage />} />
                <Route path="/movies" element={<MoviesListPage />} />
                <Route path="/" element={<MoviesListPage />} />
                <Route path="/movies/:id" element={<MovieDetailPage />} />
                <Route path="/users" element={<UserListPage />} />
            </Routes>
        </Router>
    );
};

export default App;