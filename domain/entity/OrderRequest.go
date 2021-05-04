package entity

type OrderRequest struct {
	RestaurantID        *int64               `json:"restaurantID"`
	OrderNumberID       *int64               `json:"orderNumberID"`
	WaiterID            *string              `json:"waiterID"`
	OrderRequestRecords []OrderRequestRecord `json:"data"`
}

type OrderRequestRecord struct {
	ID        *int64 `json:"id"`
	TableID   int64  `json:"tableID"`
	ProductID int64  `json:"productID"`
	Comments  string `json:"comments"`
}

func (l *OrderRequest) GetDistinctProductIDsAndTableIDs() ([]int64, []int64) {
	productKeys := make(map[int64]bool)
	tableKeys := make(map[int64]bool)
	productResult := make([]int64, 0)
	tableResult := make([]int64, 0)
	for _, u := range l.OrderRequestRecords {
		if _, value := productKeys[u.ProductID]; !value {
			productResult = append(productResult, u.ProductID)
			productKeys[u.ProductID] = true
		}
		if _, value := tableKeys[u.TableID]; !value {
			tableResult = append(tableResult, u.TableID)
			tableKeys[u.TableID] = true
		}
	}
	return productResult, tableResult
}

func (l *OrderRequest) Discover() ([]OrderRequestRecord, []OrderRequestRecord, []int64) {
	keys := make(map[int64]bool)
	nonNullIDresult := make([]int64, 0)
	nonNullresult := make([]OrderRequestRecord, 0)
	nullResult := make([]OrderRequestRecord, 0)
	for _, u := range l.OrderRequestRecords {
		if u.ID == nil {
			nullResult = append(nullResult, u)
			continue
		}
		if _, duplicate := keys[*u.ID]; !duplicate {
			nonNullresult = append(nonNullresult, u)
			nonNullIDresult = append(nonNullIDresult, *u.ID)
			keys[*u.ID] = true
		}
	}
	return nullResult, nonNullresult, nonNullIDresult
}

func (l *OrderRequest) HasOrder(orderId int64) bool {
	for _, u := range l.OrderRequestRecords {
		if u.ID == nil {
			continue
		}
		if *u.ID == orderId {
			return true
		}
	}
	return false
}

func (l *OrderRequest) IsEmpty() bool {
	return len(l.OrderRequestRecords) == 0
}
