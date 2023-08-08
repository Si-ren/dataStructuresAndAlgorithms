package main

import (
	"fmt"
	"time"
)

// 订单状态枚举
type OrderStatus int

const (
	OrderPlaced OrderStatus = iota
	OrderShipped
	OrderDelivered
)

// 订单结构体
type Order struct {
	ID        int
	Status    OrderStatus
	CreatedAt time.Time
}

// 观察者接口
type Observer interface {
	Update(order *Order)
}

// 主题接口
type Subject interface {
	Register(observer Observer)
	Unregister(observer Observer)
	Notify(order *Order)
}

// 订单管理器
type OrderManager struct {
	observers []Observer
}

func (om *OrderManager) Register(observer Observer) {
	om.observers = append(om.observers, observer)
}

func (om *OrderManager) Unregister(observer Observer) {
	for i, o := range om.observers {
		if o == observer {
			om.observers = append(om.observers[:i], om.observers[i+1:]...)
			break
		}
	}
}

func (om *OrderManager) Notify(order *Order) {
	for _, observer := range om.observers {
		observer.Update(order)
	}
}

// 订单状态观察者
type OrderStatusObserver struct {
	Name string
}

func (oso *OrderStatusObserver) Update(order *Order) {
	fmt.Printf("Observer %s received order status update. Order ID: %d, Status: %d\n", oso.Name, order.ID, order.Status)
}

func main() {
	orderManager := &OrderManager{}

	observer1 := &OrderStatusObserver{Name: "Observer 1"}
	observer2 := &OrderStatusObserver{Name: "Observer 2"}
	observer3 := &OrderStatusObserver{Name: "Observer 3"}

	orderManager.Register(observer1)
	orderManager.Register(observer2)
	orderManager.Register(observer3)

	order := &Order{
		ID:        1,
		Status:    OrderPlaced,
		CreatedAt: time.Now(),
	}

	orderManager.Notify(order)

	order.Status = OrderShipped
	orderManager.Notify(order)

	order.Status = OrderDelivered
	orderManager.Notify(order)

	orderManager.Unregister(observer2)

	order.Status = OrderPlaced
	orderManager.Notify(order)
}
