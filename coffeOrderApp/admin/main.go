package admin

// 관리자

var adminList = []Admin{
	{id: 1},
	{id: 2},
	{id: 3},
	{id: 4},
	{id: 5},
}

type Admin struct {
	id int
	name string
	userId string
	userPass string
	profilePicture string

	phone string
	email string
	address string
	snsIdList []string
}

type SearchKeyword struct {
	id int
	name string
	phone string
	email string

}

type AdminInterface interface {
	getAdminInfo(id int) Admin
	setAdminInfo(admin Admin) Admin
	removeAdminInfo(id int)
	editAdminInfo(admin Admin)
	getAdminInfoList(searchKeyword SearchKeyword) []Admin
}

func (admin *Admin) getAdminInfo(_id int) Admin {
	for _, a := range adminList {
		if a.id == _id {
			return a
		}
	}
	return Admin{}
}

func (admin *Admin) setAdminInfo(_a Admin) Admin {
	adminList = append(adminList, _a)
	return _a
}

func (admin *Admin) removeAdminInfo(_id int) {
	var newAmdinList = []Admin{}
	for _, a := range adminList {
		if a.id != _id {
			newAmdinList = append(newAmdinList, a)
		}
	}
}

func (admin *Admin) editAdminInfo(_a Admin) {
	for _, a := range adminList { // map로 바꾸는 방법은? 같은 코드를 계속 반복해야 한다.
		if a.id == _a.id {
			a = _a
		}
	}
}

func (admin *Admin) getAdminInfoList(_searchKeyword SearchKeyword) []Admin {
	var searchList = []Admin{}
	for _, a := range adminList {
		if a.name == _searchKeyword.name { // 정규식 어떻게 ?
			searchList = append(searchList, a)
		} else if a.id == _searchKeyword.id {
			searchList = append(searchList, a)
		} else if a.email == _searchKeyword.email {
			searchList = append(searchList, a)
		}
	}
	return searchList
}

func Main() {
	// 단위 테스트
	admin := Admin{}
	admin.setAdminInfo(Admin{id: 2, name: "Tonyspark"}) // 어드민 등록
	admin.getAdminInfo(2) // 어드민 조회
	admin.getAdminInfoList(SearchKeyword{name: "Tony"})
	admin.removeAdminInfo(2)
	admin.editAdminInfo(Admin{id:3, name: "원지은"})
}