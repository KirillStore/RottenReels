import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom'; // Добавляем Link для перехода на страницу фильма

const MoviesListPage = () => {
    const [movies, setMovies] = useState([]);
    const [error, setError] = useState(null);
    const token = localStorage.getItem("token");

    useEffect(() => {
        const fetchMovies = async () => {
            try {
                const response = await fetch("http://localhost:8080/movies", {
                    headers: {
                        "Authorization": `Bearer ${token}`,
                    },
                });

                if (response.ok) {
                    const data = await response.json();
                    setMovies(data.movies); // Убедимся, что обращаемся к массиву фильмов
                } else {
                    setError("Failed to load movies.");
                }
            } catch (error) {
                setError("An error occurred while fetching movies.");
            }
        };

        fetchMovies();
    }, [token]);

    if (error) {
        return <div>{error}</div>;
    }

    return (
        <div>
            <h2>Movies List</h2>
            <ul>
                {Array.isArray(movies) && movies.length > 0 ? (
                    movies.map((movie) => (
                        <li key={movie.ID}>
                            <Link to={`/movies/${movie.ID}`}> {/* Переход на детали фильма */}
                                {movie.title} - {movie.description} - Average Rating: {movie.average_rating}
                            </Link>
                        </li>
                    ))
                ) : (
                    <li>No movies found.</li>
                )}
            </ul>
        </div>
    );
};

export default MoviesListPage;