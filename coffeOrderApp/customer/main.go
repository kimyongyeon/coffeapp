package customer

// 고객정보

type Customer struct {
	id int
	name string
	email string
	userId string
	userPass string
	address string

	coffePoint int64
	snsUserId []string // jwt, 카카오, 네이버페이 연동시 사용

	regDate string  // 가입일자
	editDate string  // 수정일자

	useYn bool // 휴면유무
}

// 일촌, 친구사이
type OneConnCustomer struct {
	masterId string // 나의 아이디
	oneConnIds []string // 친구들 아이디

	regDate string  // 가입일자
	editDate string  // 수정일자

	useYn bool // 휴면유무
}

// 직장, 동료사이
type GroupCustomer struct {
	masterId string // 나의 아이디
	groupId []string  // 그룹 아이디

	regDate string  // 가입일자
	editDate string  // 수정일자

	useYn bool // 휴면유무
}

type OneConnCustomerVO struct {
	masterId string
	oneConnId string
	oneConnCustomer OneConnCustomer
}

type GroupCustomerVO struct {
	masterId string
	groupId string
	groupCustomere GroupCustomer
}

type PointVO struct {
	point int
	customer Customer
}

type CustomerVO struct {
	userId string // 로그인용
	userPass string  // 로그인용
	snsId string // 로그인용
	customer Customer // 회원가입, 회원수정, 휴면 설정용
}

type CustomerInterface interface {

	// 고객 관리
	setJoinCustomer(_customer CustomerVO) CustomerVO // 회원가입
	isCustomerId(_customer CustomerVO) bool // 회원아이디 존재여부 체크
	setSleepCustomer(_customer CustomerVO) CustomerVO // 휴면계정 설정
	isSleepCustomer(_customer CustomerVO) bool // 휴면계정 여부
	login(_customer CustomerVO) bool // 로그인
	editCustomer(_customer CustomerVO) CustomerVO // 회원정보 수정
	snsLogin(_customer CustomerVO) CustomerVO // SNS 회원 로그인
	snsJoin(_customer CustomerVO) CustomerVO // SNS 회원가입

	// 포인트
	incPoint(_p PointVO) // 포인트 증가
	decPoint(_p PointVO) // 포인트 감소
	getTotalPoint(_p PointVO) int // 전체 포인트 조회

	// 친구 관리
	getOneConnIds(_masterId string) []string // 친구 조회
	isOneConnIds(_searchOneConn OneConnCustomerVO) bool // 친구 유무 체크
	addOneConnId(_searchOneConn OneConnCustomerVO) OneConnCustomer // 친구 추가
	removeOneConnId(_searchOneConn OneConnCustomerVO) OneConnCustomer // 친구 삭제

	// 그룹 관리
	getGroupCustomerId(_masterId string) []string // 그룹 친구 조회
	isGroupCustomerId(_searchGroupCustomer GroupCustomerVO) // 그룹 유무 체크
	addGroupId(_searchGroupCustomer GroupCustomerVO) // 그룹 추가
	removeGroupId(_searchGroupCustomer GroupCustomerVO) // 그룹 삭제
}
