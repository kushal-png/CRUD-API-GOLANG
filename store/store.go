package store

import (
	"CRUD_API/model"
	"errors"
)

var movies []model.Movie

func Movies() *[]model.Movie {
	return &movies
}

func InitializeMovies() {
	// Dummy data
	movies = append(movies, model.Movie{ID: "23", Title: "Horror", Author: &model.Author{Name: "Kushal", Age: 34}})
	movies = append(movies, model.Movie{ID: "26", Title: "Comedy", Author: &model.Author{Name: "Komal", Age: 24}})
}

func GetMovies() []model.Movie {
	return movies
}

func DeleteMovie(id string) error {
	for index, key := range movies {
		if key.ID == id {
			movies = append(movies[:index], movies[index+1:]...)
			return nil
		}
	}
	return errors.New("movie not found")
}

func GetMovie(id string) (*model.Movie, error) {
	for _, key := range movies {
		if key.ID == id {
			return &key, nil
		}
	}

	return nil, errors.New("movie not found")
}

func CreateMovie(m model.Movie) {
	movies = append(movies, m)
}

func UpdateMovie(id string, m model.Movie) error {
	for index, movie := range movies {
		if movie.ID == id {
			movies[index] = m
			return nil
		}
	}
	return errors.New("movie not found")
}
