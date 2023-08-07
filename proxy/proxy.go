package proxy

import "fmt"

type Pay interface {
	Pay(int) error
}

type AliPay struct {
}

func (a *AliPay) Pay(i int) error {
	fmt.Println("this is Alipay")
	i = 1
	return nil
}

type User struct {
	AliPay
}

func (u User) Pay(i int) error {
	fmt.Println("this is user's pay")
	err := u.AliPay.Pay(i)
	return err
}
