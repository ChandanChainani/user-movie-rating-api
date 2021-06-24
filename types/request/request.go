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
	Rating float32  `json:"rating"   validate:"numeric"`
	UserID string   `json:"user_id"  validate:"required"`
	MovieID string  `json:"movie_id" validate:"required"`
}

type UserMovieComment struct {
	Comment string  `json:"comment"  validate:"required"`
	UserID string   `json:"user_id"  validate:"required"`
	MovieID string  `json:"movie_id" validate:"required"`
}
