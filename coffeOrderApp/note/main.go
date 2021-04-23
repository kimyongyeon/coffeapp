package note

// 쪽지

type Note struct {
	id int // 기본키
	title string // 100자
	content string // 문자열제한? 500자
	nickName string // 별칭
	fromId string // 보내는 이
	toId string  // 받는 이

	regDate string // 등록일자
	upDate string // 수정일자
	register string // 등록자

	useYn bool // 블록된 쪽지
}

type SearchNote struct {

}

type NoteVO struct {

}

type NoteInterface interface {
	findOne(searchNote SearchNote) NoteVO
}