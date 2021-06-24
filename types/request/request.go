package request

type SearchUser struct {
	ID string `json:"id"`
}

type SearchMovie struct {
	Name string `json:"name"`
}

type User struct {
	Email string `json:"email"`
}

type Movie struct {
	Name string `json:"name"`
	Description string `json:"description"`
}

type UserMovieRating struct {
	Rating float32  `json:"rating"`
	UserID string   `json:"user_id"`
	MovieID string  `json:"movie_id"`
}

type UserMovieComment struct {
	Comment string  `json:"comment"`
	UserID string   `json:"user_id"`
	MovieID string  `json:"movie_id"`
}
