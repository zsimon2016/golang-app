package car

import (
	// "car/herx/system/services"
	"context"
	"fmt"
)

type Car struct {
	// DB *mysql.Db

	Config *map[string]interface{}
}
type Args struct {
	R       map[string]interface{}
	Service string
	Method  string
}
type Reply struct {
	Result map[string]interface{}
}
type BySort []map[string]interface{}

func (a BySort) Len() int { return len(a) }
func (by BySort) Less(i, j int) bool {
	a := by[i]["time"].(string)
	b := by[j]["time"].(string)
	return a < b
}
func (a BySort) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (car *Car) Health(ctx context.Context, args *Args, reply *Reply) error {
	reply.Result = map[string]interface{}{
		"code": 200,
	}
	fmt.Println("1")
	return nil
}
