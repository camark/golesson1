package main;

import "fmt"
import "os"
import "strconv"

func main()  {
	fmt.Println("Hello calc")

	switch os.Args[1] {
	case "add":
		if len(os.Args)<=3 {
			fmt.Println("Usage add var1 var2");
			return;
		}

		i,err := strconv.Atoi(os.Args[2])

		if err!=nil {
			panic(err)
			
		}

		j,err := strconv.Atoi(os.Args[3])

		if err!=nil {
			panic(err)
		}

		sum := i+j

		fmt.Printf("%d+%d=%d",i,j,sum);
		break;
	case "three":
		i,err := strconv.Atoi(os.Args[2])

		if err!=nil {
			fmt.Println("Usage three argv");
			return
		}

		fmt.Printf("%d *3 = %d",i,i*3)
		break;

	default:
		fmt.Printf("Usage calc add v1 v2 or Usage three v1");
		break;

	}
}