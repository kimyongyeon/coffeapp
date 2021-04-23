package board

import "fmt"

// 전역변수
var boardList = []Board{
	{1, "사나", "트와이스", "JYP", Comment{}},
	{2, "태연", "소녀시대", "SM", Comment{}},
	{3, "효리", "핑클", "빅엔터", Comment{}},
	{4, "공민지", "2ne1", "YG", Comment{}},
}

// 목록, 상세, 글쓰기, 글수정, 글삭제
type Board struct {
	boardId int
	title   string
	content string
	writer  string
	comment Comment
}

type Comment struct {
	commentId int // 댓글 아이디
	boardId   int // 게시글 아이디
	title     string
	content   string
	writer    string
}

type SearchItem struct {
	title   string
	content string
	writer  string
}

type BoardInterface interface {
	list() []Board                         // 목록보기
	detail() Board                         // 상세보기
	write(board *Board) Board              // 글쓰기
	remove(board *Board) bool              // 삭제
	edit(board *Board) Board               // 글수정
	search(searchItem *SearchItem) []Board // 글 검색
}

func (board *Board) search(item *SearchItem) []Board {
	return []Board{}
}

func (board *Board) edit(pBoard *Board) Board {
	return Board{}
}

func (board *Board) remove(pBoard *Board) bool {
	return true
}

func (board *Board) list() []Board {
	return []Board{}
}

func (board *Board) detail() Board {
	return Board{}
}

func (board *Board) write(pBoard *Board) Board {
	return Board{
		pBoard.boardId,
		pBoard.title,
		pBoard.content,
		pBoard.writer,
		pBoard.comment,
	}
}

func init() {
	fmt.Println(boardList)
}

func boardManager(boardInter BoardInterface) {
	fmt.Println("누가 진짜 인가?", boardInter.list())
}

func Main() {
	b := Board{}
	boardManager(&b)
}
