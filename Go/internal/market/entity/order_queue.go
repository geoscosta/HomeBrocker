package entity

type OrderQueue struct {
	Orders []*Order
}

// Função para comparar o preço de duas ordens.
func (oq *OrderQueue) Less(i, j int) bool {
	return oq.Orders[i].Price < oq.Orders[j].Price
}

// Função para inverter os valores das ordens
func (oq *OrderQueue) Swap(i, j int) {
	oq.Orders[i], oq.Orders[j] = oq.Orders[j], oq.Orders[i]
}

// Função para retornar a quantidade de ordens
func (oq *OrderQueue) Len() int {
	return len(oq.Orders)
}

// Função para adicionar as ordens
func (oq *OrderQueue) Push(x any) {
	oq.Orders = append(oq.Orders, x.(*Order))
}

// Função para retirar a ultima ordem da fila
func (oq *OrderQueue) Pop() any {
	oldOrders := oq.Orders
	qtdOrders := len(oldOrders)
	item := oldOrders[qtdOrders-1]
	oq.Orders = oldOrders[0 : qtdOrders-1]
	return item
}

func NewOrderQueue() *OrderQueue {
	return &OrderQueue{}
}
