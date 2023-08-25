package main

import (
	"fmt"
	"sync"
)

type BankAccount struct {
	balance float64
	mutex   sync.Mutex
}

func (acc*BankAccount) Deposit(amount float64){
	acc.mutex.Lock()
	defer acc.mutex.Unlock()
	fmt.Printf("Depositing %.2f\n",amount)
	acc.balance+=amount
	fmt.Printf("New balance :%.2f\n",acc.balance)
}

func (acc*BankAccount) Withdraw(amt float64){
	acc.mutex.Lock()
	defer acc.mutex.Unlock()

	if amt>=acc.balance{
		acc.balance-=amt
		fmt.Printf("%.2f amount withdraw successfully",amt)
		fmt.Printf("New balance :%.2f\n",acc.balance)

	} else {
		fmt.Println("Insufficient fund")
	}
}

func main(){
	cash:=&BankAccount{
		balance: 5000,
	}
	var wg sync.WaitGroup

	for i:=0;i<5;i++{
		wg.Add(1)
		go func(index int){
			defer wg.Done()
			if index%2==0{
				cash.Deposit(2000.0)
			} else {
				cash.Withdraw(1500.0)
			}
		}(i)
	}
	wg.Wait()
	fmt.Printf("final Balance:%.2f",cash.balance)

}