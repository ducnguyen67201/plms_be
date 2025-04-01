package problem_domain

type Problem struct {
	ProblemID       int64  `json:"problem_id"`
	ContestID       *int64 `json:"contest_id"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	DifficultyLevel string `json:"difficulty_level"`
	RepeatedTimes   int64  `json:"repeated_times"`
	Type            string `json:"type"`
}