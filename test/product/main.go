package product

import (
	"fmt"
	"math/big"
)

// 제품
type Product struct {
	id int
	name string
	price big.Accuracy
	createDate string
	upDate string
	delDate string
	createUserId string
}

// 고객
type Client struct {
	userId string
	userName string
	userPhone int
	userAddress string
	userEmail string
}

// 주문
type Order struct {
	productId int // 제품 아이디
	clientUserId string // 고객 아이디
}

// 카트
type Cart struct {
	productId int // 제품 아이디
	clientUserId string // 고객 아이디
}

// 주문하다
func getOrder(order *Order) {

}

// 장바구니 생성
func setCart(card Cart) {

}

type ProductMethod interface {
	getProductDetail(productId int) Product
	//getProductList() []Product
	//insProduct(product Product) // 상품 추가
	//delProduct(productId int) // 상품 삭제
	//editProduct(product Product) // 상품 수정
}
type ClientMethod interface {

}
type CartMethod interface {

}

func (product Product) getProductDetail(productId int) Product {

	// 리펙토링 전: 코드가 길어지면 보기 힘들어짐.
	//if productId > 0 {
	//	return Product{id: 1}
	//} else {
	//	return Product{}
	//}

	// if문 역으로.
	fmt.Println("productId: ", productId)
	if productId != 0 {
		return Product{id: productId}
	}
	return Product{}
}

// 리시버는 product 안에 메서드가 있고, 메서드 구현은 언제 해야 하는가?
func productMng(productMethod ProductMethod) {
	fmt.Println("productMethod: ", productMethod.getProductDetail(1))
}


func Main() string {
	fmt.Println("Product ProductMain called")

	var product = Product{id: 1}
	productMng(product)

	return "product Main "
}
