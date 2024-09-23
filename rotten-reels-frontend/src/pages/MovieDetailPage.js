import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';

const MovieDetailPage = ({ user }) => {
    const { id } = useParams();
    const [movie, setMovie] = useState(null);
    const [rating, setRating] = useState(0);
    const [review, setReview] = useState('');
    const [error, setError] = useState(null);
    const [message, setMessage] = useState(null);
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
                    setMovie(data); // Исправление: предполагаем, что данные фильма находятся прямо в ответе
                } else {
                    setError("Failed to load movie.");
                }
            } catch (error) {
                setError("An error occurred while fetching movie.");
            }
        };

        fetchMovie();
    }, [id, token]);

    const handleRatingSubmit = async (e) => {
        e.preventDefault();
        setMessage(null);

        try {
            const response = await fetch(`http://localhost:8080/movies/${id}/ratings`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${token}`,
                },
                body: JSON.stringify({ score: rating }),
            });

            if (response.ok) {
                setMessage('Rating submitted successfully.');
            } else {
                setError('Failed to submit rating.');
            }
        } catch (error) {
            setError('An error occurred while submitting rating.');
        }
    };

    const handleReviewSubmit = async (e) => {
        e.preventDefault();
        setMessage(null);

        try {
            const response = await fetch(`http://localhost:8080/movies/${id}/reviews`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${token}`,
                },
                body: JSON.stringify({
                    text: review,
                    user_id: Number(user.user_id), // Преобразуем user_id в число, если нужно
                    movie_id: Number(id),          // Преобразуем id фильма в число
                }),
            });

            if (response.ok) {
                setMessage('Review submitted successfully.');
                setReview(''); // Очищаем поле отзыва после успешной отправки
            } else {
                const errorData = await response.json(); // Логируем ошибку от сервера
                console.error('Error from server:', errorData);
                setError('Failed to submit review.');
            }
        } catch (error) {
            console.error('Error during fetch:', error); // Логируем ошибку сети
            setError('An error occurred while submitting review.');
        }
    };

    if (error) {
        return <div>{error}</div>;
    }

    return (
        <div>
            <h2>Movie Details</h2>
            {movie ? (
                <div>
                    <p>Title: {movie.title}</p>
                    <p>Description: {movie.description}</p>
                    <p>Average Rating: {movie.avg_rating}</p> {/* Показываем средний рейтинг фильма */}
                </div>
            ) : (
                <p>Loading movie...</p>
            )}

            {/* Форма для отправки оценки */}
            <form onSubmit={handleRatingSubmit}>
                <label>
                    Rate this movie:
                    <input
                        type="number"
                        value={rating}
                        onChange={(e) => setRating(e.target.value)}
                        min="1"
                        max="10"
                        required
                    />
                </label>
                <button type="submit">Submit Rating</button>
            </form>

            {/* Форма для отправки отзыва */}
            <form onSubmit={handleReviewSubmit}>
                <label>
                    Write a review:
                    <textarea
                        value={review}
                        onChange={(e) => setReview(e.target.value)}
                        required
                    />
                </label>
                <button type="submit">Submit Review</button>
            </form>

            {message && <p>{message}</p>}
        </div>
    );
};

export default MovieDetailPage;