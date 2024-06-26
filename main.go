package main

import (
    "movie-rating-service/config"
    "movie-rating-service/models"
    "movie-rating-service/routes"
)

func main() {
    // Загрузка переменных окружения
    config.LoadEnv()

    // Инициализация базы данных
    db := config.SetupDatabase()

    // Автоматическое создание таблиц на основе моделей
    db.AutoMigrate(&models.User{}, &models.Movie{}, &models.Rating{})

    // Настройка маршрутов
    r := routes.SetupRouter(db)

    // Запуск сервера
    r.Run(":8080")
}
