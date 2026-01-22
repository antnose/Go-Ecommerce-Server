package database

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"imageUrl"`
}

var ProductList []Product

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

	ProductList = append(ProductList, pd1)
	ProductList = append(ProductList, pd2)
	ProductList = append(ProductList, pd3)

}
