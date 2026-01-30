package database

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"imageUrl"`
}

var productList []Product

func Store(p Product) Product {
	p.ID = len(productList) + 1
	productList = append(productList, p)
	return p
}

func List() []Product {
	return productList
}

func Get(ProductId int) *Product {
	for _, product := range productList {

		if product.ID == ProductId {
			return &product
		}
	}

	return nil
}

func Update(product Product) {
	for idx, p := range productList {
		if p.ID == product.ID {
			productList[idx] = product
		}
	}
}

func Delete(productId int) {
	var tempList []Product = make([]Product, 0)

	for _, p := range productList {
		if productId != p.ID {
			tempList = append(tempList, p)
		}
	}

	productList = tempList
}

func init() {
	pd1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "Orange is orange. And that's why it's orange",
		Price:       230,
		ImgURL:      "https://upload.wikimedia.org/wikipedia/commons/thumb/4/43/Ambersweet_oranges.jpg/250px-Ambersweet_oranges.jpg",
	}

	pd2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "Apple is apple",
		Price:       340,
		ImgURL:      "https://hips.hearstapps.com/hmg-prod/images/apples-at-farmers-market-royalty-free-image-1627321463.jpg?crop=1xw:0.94466xh;center,top&resize=1200:*",
	}

	pd3 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "Banana is banana",
		Price:       40,
		ImgURL:      "https://www.allrecipes.com/thmb/jYmw-0Vijg1E_OuG2yGjEAcdQg4=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/ar-new-banana-adobe-ar-4x3-d8f0871e12214350be7ae5575eea4eed.jpg",
	}

	productList = append(productList, pd1)
	productList = append(productList, pd2)
	productList = append(productList, pd3)

}
