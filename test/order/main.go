package order

type Order struct {
	id int // 기본키
	orderNumber string // 주문번호
	orderId string // 주문자아이디
	orderRegDate string // 주문등록일자
	orderDelDate string // 주문삭제일자
	orderUpDate string // 주문수정일자
}