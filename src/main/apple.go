package main
import "fmt"

type apple struct {
	address string
	color int
	yy float32
}


func (a apple) msg() string{
	
	scolor := getColor(a.color)

	ret := fmt.Sprintf("Apple 产地 ->%s,颜色->%s,营业指数->%f",a.address,scolor,a.yy)

	return ret;
}

func (a *apple) setColor(c int){

	origin_color := a.color
	a.color = c

	fmt.Printf("Color change from %s to %s\n",getColor(origin_color),getColor(c))
}

func (a apple) printInformation(){
	fmt.Println(a.msg())
}
