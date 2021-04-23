package feedback

// 피드백

type Feedback struct {
	id int
	writer string
	title string
	content string
	feedbackType FeedbackType

	regDate string
	upDate string
	useYn bool
}

type FeedbackType struct {
	code string
	codeName string
}

type SearchFeedback struct {
	currentPage int
	totalPageCount int
	keyword string
	feedback Feedback
}

type FeedbackVO struct {
	currentPage int
	totalPageCount int
	feedback []Feedback
}

type FeedbackInterface interface {
	write(feedback Feedback) bool // 피드백 쓰기
	edit(feedback Feedback) Feedback // 피드백 수정
	remove(feedback Feedback) bool // 피드백 삭제
	findOne(searchFeedback SearchFeedback) Feedback // 피드백 상세 조회
	findPaging(searchFeedback SearchFeedback) []FeedbackVO // 피드백 페이징 목록 조회
}