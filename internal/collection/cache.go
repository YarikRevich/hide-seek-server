package collection

import "time"

//Runs cache loop which checks if collection
//data is expired to exist
func RunCacheLoop(){
	tick := time.NewTicker(time.Second)
	for range tick.C{
		for k, v := range World{
			if time.Since(v.Cache) == 0{
				delete(World, k)	
			}
		}
		for k, v := range User{
			if time.Since(v.Cache) == 0{
				delete(World, k)	
			}
		}
	}
}