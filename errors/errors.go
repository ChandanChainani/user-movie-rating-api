package errors

type AppError struct{
	err string
}

func (e *AppError) Error() string {
	return string(e.err)
}

var (
	NoDataFoundError = &AppError{ err: "no data found" }

	AuthFailed = map[string]string{
		"error": "user authentication failed",
	}
	MovieNameFieldMissing = map[string]string{
		"error": "name field missing",
	}
	MovieNotFound = map[string]string{
		"error": "movie not found",
	}
	NoDataFound = map[string]string{
		"error": "no data found",
	}
	AddMovieFailed = map[string]string{
		"error": "failed to rate the movie",
	}
	RateMovieFailed = map[string]string{
		"error": "failed to rate the movie",
	}
)
