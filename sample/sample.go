package sample

//here are some sample datasets 1-3 to load!
import (
	"fmt"

	"encoding/json"
)

type MyData struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Passion string `json:"passion"`
}




func ExampleData1() []MyData {

	var dummy_data []MyData

	dummy_data = append(dummy_data, MyData{ID:1, Name:"jimson",Passion:"food"})
	dummy_data = append(dummy_data, MyData{ID:2, Name:"janey",Passion:"gym"})
	dummy_data = append(dummy_data, MyData{ID:3, Name:"joey",Passion:"running"})

	//fmt.Println(dummy_data)
	test, _ := json.Marshal(dummy_data)
	test_str := string(test)
	fmt.Printf("test data - %s" ,test_str)
	
	return dummy_data
}


func ExampleData2() []MyData {

	var dummy_data []MyData

	dummy_data = append(dummy_data, MyData{ID:1, Name:"Liam", Passion:"Origami"})
	dummy_data = append(dummy_data, MyData{ID:2, Name:"Sophia", Passion:"Astronomy"})
	dummy_data = append(dummy_data, MyData{ID:3, Name:"Ethan", Passion:"Calligraphy"})
	dummy_data = append(dummy_data, MyData{ID:4, Name:"Ava", Passion:"Lockpicking"})
	dummy_data = append(dummy_data, MyData{ID:5, Name:"Noah", Passion:"Geocaching"})
	dummy_data = append(dummy_data, MyData{ID:6, Name:"Mia", Passion:"Soap Making"})
	dummy_data = append(dummy_data, MyData{ID:7, Name:"Jackson", Passion:"Parkour"})
	dummy_data = append(dummy_data, MyData{ID:8, Name:"Olivia", Passion:"Blacksmithing"})
	dummy_data = append(dummy_data, MyData{ID:9, Name:"Lucas", Passion:"Mushroom Foraging"})
	dummy_data = append(dummy_data, MyData{ID:10, Name:"Emma", Passion:"Sand Art"})

	//fmt.Println(dummy_data)
	test, _ := json.Marshal(dummy_data)
	test_str := string(test)
	fmt.Printf("test data - %s" ,test_str)
	
	return dummy_data
}

func ExampleData3() []MyData {

	var dummy_data []MyData

	dummy_data = append(dummy_data, MyData{ID:1, Name:"Liam", Passion:"Origami"})
	dummy_data = append(dummy_data, MyData{ID:2, Name:"Sophia", Passion:"Astronomy"})
	dummy_data = append(dummy_data, MyData{ID:3, Name:"Ethan", Passion:"Calligraphy"})
	dummy_data = append(dummy_data, MyData{ID:4, Name:"Ava", Passion:"Lockpicking"})
	dummy_data = append(dummy_data, MyData{ID:5, Name:"Noah", Passion:"Geocaching"})
	dummy_data = append(dummy_data, MyData{ID:6, Name:"Mia", Passion:"Soap Making"})
	dummy_data = append(dummy_data, MyData{ID:7, Name:"Jackson", Passion:"Parkour"})
	dummy_data = append(dummy_data, MyData{ID:8, Name:"Olivia", Passion:"Blacksmithing"})
	dummy_data = append(dummy_data, MyData{ID:9, Name:"Lucas", Passion:"Mushroom Foraging"})
	dummy_data = append(dummy_data, MyData{ID:10, Name:"Emma", Passion:"Sand Art"})
	dummy_data = append(dummy_data, MyData{ID:11, Name:"Henry", Passion:"Wood Carving"})
	dummy_data = append(dummy_data, MyData{ID:12, Name:"Charlotte", Passion:"Pottery"})
	dummy_data = append(dummy_data, MyData{ID:13, Name:"Benjamin", Passion:"Juggling"})
	dummy_data = append(dummy_data, MyData{ID:14, Name:"Amelia", Passion:"Metal Detecting"})
	dummy_data = append(dummy_data, MyData{ID:15, Name:"Daniel", Passion:"Candle Making"})
	dummy_data = append(dummy_data, MyData{ID:16, Name:"Harper", Passion:"Bird Watching"})
	dummy_data = append(dummy_data, MyData{ID:17, Name:"William", Passion:"Lego Building"})
	dummy_data = append(dummy_data, MyData{ID:18, Name:"Ella", Passion:"Knife Throwing"})
	dummy_data = append(dummy_data, MyData{ID:19, Name:"Matthew", Passion:"Cosplaying"})
	dummy_data = append(dummy_data, MyData{ID:20, Name:"Scarlett", Passion:"Tattoo Designing"})
	dummy_data = append(dummy_data, MyData{ID:21, Name:"James", Passion:"Knife Making"})
	dummy_data = append(dummy_data, MyData{ID:22, Name:"Grace", Passion:"Beekeeping"})
	dummy_data = append(dummy_data, MyData{ID:23, Name:"Andrew", Passion:"Puzzle Solving"})
	dummy_data = append(dummy_data, MyData{ID:24, Name:"Victoria", Passion:"Cactus Collecting"})
	dummy_data = append(dummy_data, MyData{ID:25, Name:"Michael", Passion:"Guitar Playing"})
	dummy_data = append(dummy_data, MyData{ID:26, Name:"Zoe", Passion:"Quilling"})
	dummy_data = append(dummy_data, MyData{ID:27, Name:"Nathan", Passion:"Telescope Making"})
	dummy_data = append(dummy_data, MyData{ID:28, Name:"Lily", Passion:"Chess Playing"})
	dummy_data = append(dummy_data, MyData{ID:29, Name:"Ryan", Passion:"Kite Flying"})
	dummy_data = append(dummy_data, MyData{ID:30, Name:"Isabelle", Passion:"Canoeing"})
	dummy_data = append(dummy_data, MyData{ID:31, Name:"Joseph", Passion:"Baking"})
	dummy_data = append(dummy_data, MyData{ID:32, Name:"Madison", Passion:"Yoga"})
	dummy_data = append(dummy_data, MyData{ID:33, Name:"Samuel", Passion:"Gaming"})
	dummy_data = append(dummy_data, MyData{ID:34, Name:"Penelope", Passion:"Magic Tricks"})
	dummy_data = append(dummy_data, MyData{ID:35, Name:"Logan", Passion:"Leather Crafting"})
	dummy_data = append(dummy_data, MyData{ID:36, Name:"Evelyn", Passion:"Dancing"})
	dummy_data = append(dummy_data, MyData{ID:37, Name:"David", Passion:"Skateboarding"})
	dummy_data = append(dummy_data, MyData{ID:38, Name:"Chloe", Passion:"Fencing"})
	dummy_data = append(dummy_data, MyData{ID:39, Name:"Caleb", Passion:"RC Car Racing"})
	dummy_data = append(dummy_data, MyData{ID:40, Name:"Stella", Passion:"Rock Climbing"})
	dummy_data = append(dummy_data, MyData{ID:41, Name:"Adam", Passion:"Piano Playing"})
	dummy_data = append(dummy_data, MyData{ID:42, Name:"Hannah", Passion:"Soap Carving"})
	dummy_data = append(dummy_data, MyData{ID:43, Name:"Isaac", Passion:"Swimming"})
	dummy_data = append(dummy_data, MyData{ID:44, Name:"Natalie", Passion:"Embroidery"})
	dummy_data = append(dummy_data, MyData{ID:45, Name:"Sebastian", Passion:"Fishing"})
	dummy_data = append(dummy_data, MyData{ID:46, Name:"Aria", Passion:"Calligraphy"})
	dummy_data = append(dummy_data, MyData{ID:47, Name:"Jack", Passion:"Sculpting"})
	dummy_data = append(dummy_data, MyData{ID:48, Name:"Brooklyn", Passion:"Parkour"})
	dummy_data = append(dummy_data, MyData{ID:49, Name:"Jonathan", Passion:"Drone Flying"})
	dummy_data = append(dummy_data, MyData{ID:50, Name:"Leah", Passion:"Hiking"})
	
	//fmt.Println(dummy_data)
	test, _ := json.Marshal(dummy_data)
	test_str := string(test)
	fmt.Printf("test data - %s" ,test_str)
	
	return dummy_data
}
 



