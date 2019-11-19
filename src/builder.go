package builder

import "fmt"

type Director struct {
	houseBuilder HouseBuilder
	class string
}

func (d *Director) makeHouse(){
	d.houseBuilder.buildSwimmingPool()
	d.houseBuilder.buildGarage()
	d.houseBuilder.buildGarden()
}
func (d *Director) SetHouseBuilder(t string){
	switch t{
	case "Wood":
		d.houseBuilder = &WoodHouseBuilder{}
	case "Stone":
		d.houseBuilder = &StoneHouseBuilder{}
	}
	d.houseBuilder.buildHouse()
	d.class = t
}
func (d *Director) GetHouse() *House{
	switch d.class{
	case "Wood":
		return &d.houseBuilder.(*WoodHouseBuilder).House
	case "Stone":
		return &d.houseBuilder.(*StoneHouseBuilder).House
	}
	return nil
}
type HouseBuilder interface {
	buildHouse()
	buildSwimmingPool()
	buildGarden()
	buildGarage()
}
// 木头房施工队
type WoodHouseBuilder struct{
	House
}
func (wb *WoodHouseBuilder)buildHouse(){
	wb.Wall = "Wood Wall"
	wb.Door = "Wood Door"
	wb.Roof = "Wood Roof"
}
func (wb *WoodHouseBuilder) buildSwimmingPool(){
	wb.SwimmingPool = "Wood SwimmingPool"
}
func (wb *WoodHouseBuilder) buildGarden(){
	wb.Garden = "Wood Garden"
}
func (wb *WoodHouseBuilder) buildGarage(){
	wb.Garage = "Wood Garage"
}
// 石头房施工队
type StoneHouseBuilder struct{
	House
}
func (sb *StoneHouseBuilder)buildHouse(){
	sb.Wall = "Stone Wall"
	sb.Door = "Stone Door"
	sb.Roof = "Stone Roof"
}
func (sb *StoneHouseBuilder) buildSwimmingPool(){
	sb.SwimmingPool = "Stone SwimmingPool"
}
func (sb *StoneHouseBuilder) buildGarden(){
	sb.Garden = "Stone Garden"
}
func (sb *StoneHouseBuilder) buildGarage(){
	sb.Garage = "Stone Garage"
}
type House struct {
	Wall string
	Door string
	Roof string
	SwimmingPool string
	Garden string
	Garage string
}

func main() {
	director := Director{}
	director.SetHouseBuilder("Wood")
	director.makeHouse()
	house := director.GetHouse()
	fmt.Println("------------------------")
	fmt.Println(house.Wall)
	fmt.Println(house.Door)
	fmt.Println(house.Roof)
	fmt.Println(house.SwimmingPool)
	fmt.Println(house.Garage)
	fmt.Println(house.Garden)
	fmt.Println("------------------------")
	director.SetHouseBuilder("Stone")
	director.makeHouse()
	house = director.GetHouse()
	fmt.Println(house.Wall)
	fmt.Println(house.Door)
	fmt.Println(house.Roof)
	fmt.Println(house.SwimmingPool)
	fmt.Println(house.Garage)
	fmt.Println(house.Garden)
	fmt.Println("------------------------")
}

//output：
//------------------------
//Wood Wall
//Wood Door
//Wood Roof
//Wood SwimmingPool
//Wood Garage
//Wood Garden
//------------------------
//Stone Wall
//Stone Door
//Stone Roof
//Stone SwimmingPool
//Stone Garage
//Stone Garden
//------------------------

