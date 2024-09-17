import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';

const MovieDetailPage = () => {
    const { id } = useParams();
    const [movie, setMovie] = useState(null);
    const [error, setError] = useState(null);
    const token = localStorage.getItem("token");

    useEffect(() => {
        const fetchMovie = async () => {
            try {
                const response = await fetch(`http://localhost:8080/movies/${id}`, {
                    headers: {
                        "Authorization": `Bearer ${token}`,
                    },
                });

                if (response.ok) {
                    const data = await response.json();
                    setMovie(data.movie); // Подтягиваем данные фильма
                } else {
                    setError("Failed to load movie.");
                }
            } catch (error) {
                setError("An error occurred while fetching movie.");
            }
        };

        fetchMovie();
    }, [id, token]);

    if (error) {
        return <div>{error}</div>;
    }

    if (!movie) {
        return <div>Loading...</div>;
    }

    return (
        <div>
            <h2>{movie.title}</h2>
            <p>{movie.description}</p>
            <p>Average Rating: {movie.average_rating}</p>
            {/* Здесь можно добавить возможность оставить оценку */}
        </div>
    );
};

export default MovieDetailPage;