package  main

import "fmt"

//实现商品结构体：
type good struct{
	name string
	price int
    amount  int
	nature  string
}



//实现接口1

type  Controller interface {
	Check()
	Update()
	Print()
}


type Goods struct{
	amount  int
	message  string
}

func (g*Goods) check_amount(new_amount  int )  {
	g.amount = new_amount
}


func (g Goods  ) Check (){
   	var g1  Goods
	   g1.check_amount(30 )
	fmt.Println(g1.amount )
}

func (g Goods  ) Update (){
	var g1  Goods
	g1.check_amount(30 )
	g1.check_amount(35)
	fmt.Println(g1.amount )
}

func (g Goods) Print() {
	var g1 Goods
	fmt.Println(g1.message)
}
//实现电子产品结构体：
type attribute struct {
     brand	string
	 model  int
}

func main(){
	 good1 := good{
		name :
			price :
		amount :
			nature:attribute{
			brand : "brand"
			model :"model"
	 },
	}
	fmt.Println("good1=%#v\n",good1)
}

//实现接口2：
type Print interface {
   Print ()
}


type Function struct{
	xinghao int
}

func (f Function  ) print()  {
	var f1  Function
	fmt.Println(f1.xinghao)
}





















