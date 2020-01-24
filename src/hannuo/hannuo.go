package hannuo

import(
	"fmt"
)

func MoveHannuo(count byte){
	moveAction("A", "B", "C", count);
}

func moveAction(z1, z2, zb string, count byte){
	if count == 1{
		fmt.Printf("move Top from %s to %s.\n", z1, z2);
		return;
	}
	moveAction(z1, zb, z2, count - 1);
	moveAction(z1, z2, zb, 1);
	moveAction(zb, z2, z1, count - 1);
}
