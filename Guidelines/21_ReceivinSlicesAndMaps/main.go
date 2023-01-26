package main

import "fmt"

type Driver struct {
	trips []string
}

func (d *Driver) SetTripsBad(trips []string) {
	d.trips = trips
}

func (d *Driver) SetTripsGood(trips []string) {
	d.trips = make([]string, len(trips))
	copy(d.trips, trips)
}

func main() {
	d := &Driver{
		trips: []string{"HaNoi", "SaiGon"},
	}
	fmt.Println(d.trips)
	fmt.Println("----------------------------------change good, not same pointer---------------------------------------------------------")
	changeGood := []string{"HaNoi1", "SaiGon1"}
	fmt.Println(changeGood, &changeGood, d.trips)
	d.SetTripsGood(changeGood)
	fmt.Println(changeGood, &changeGood, d.trips)
	changeGood[0] = "Thanh Hoa"
	fmt.Println(changeGood, &changeGood, d.trips)
	fmt.Println(&changeGood[0], &d.trips[0])
	fmt.Println("----------------------------------end-----------------------------------------------------------------------------------")

	fmt.Println("----------------------------------change bad, same pointer--------------------------------------------------------------")
	changeBad := []string{"HaNoi1", "SaiGon1"}
	d.SetTripsBad(changeBad)
	fmt.Println(changeBad, &changeBad, d.trips)
	changeBad[0] = "HaiPhong"
	fmt.Println(changeBad, &changeBad, d.trips)
	fmt.Println(&changeBad[0], &d.trips[0])
	fmt.Println("----------------------------------end----------------------------------------------------------------------------------")

	fmt.Println("----------------------------------append and not change pointer---------------------------------------------------------")
	fmt.Println(changeBad, &changeBad, d.trips)
	changeBad = append(changeBad[:(len(changeBad)-1)], "Hue")
	fmt.Println(changeBad, &changeBad, d.trips)
	fmt.Println(&changeBad[0], &d.trips[0])
	fmt.Println("----------------------------------end------------------------------------------------------------------------------------")

	fmt.Println("----------------------------------append and change pointer--------------------------------------------------------------")
	fmt.Println(changeBad, &changeBad, d.trips)
	changeBad = append(changeBad, "PhuQuoc")
	fmt.Println(changeBad, &changeBad, d.trips)
	fmt.Println(&changeBad[0], &d.trips[0])
	fmt.Println("----------------------------------end------------------------------------------------------------------------------------")

}

//pointer and slice have many problem, we will explain step by step follow me:
/**
	1) Keep in mind that users can modify a map or slice you received as an argument if you store a reference to it.
    2) Slices and maps contain pointers to the underlying data so be wary of scenarios when they need to be copied.

you run code, and list result:
+) follow line 23 to 32: result:
----------------------------------change good, not same pointer---------------------------------------------------------
[HaNoi1 SaiGon1] &[HaNoi1 SaiGon1] [HaNoi SaiGon]
[HaNoi1 SaiGon1] &[HaNoi1 SaiGon1] [HaNoi1 SaiGon1]
[Thanh Hoa SaiGon1] &[Thanh Hoa SaiGon1] [HaNoi1 SaiGon1]
0xc00002a060 0xc00002a080
----------------------------------end-----------------------------------------------------------------------------------

 it is good change:
    in line 14: d.trips = make([]string, len(trips)), d.strip use make then  allocates and initializes an object, not same  trips
    in line 15, copy(d.trips, trips), copy will copy bytes, don't change memory area
    ==> d.strips and strip is independence
    in line 28: changeGood[0] = "Thanh Hoa",  changeGood will is [Thanh Hoa SaiGon1], and  d.strip not change, it is HaNoi1 SaiGon1
    in line 30: 	fmt.Println(&changeGood[0], &d.trips[0]),  0xc00002a060 0xc00002a080  ==>  d.strips and changeGood is independence
    ==>  code run and not confused when change changeGood

+) follow line 33 to 40: result:
----------------------------------change bad, same pointer--------------------------------------------------------------
[HaNoi1 SaiGon1] &[HaNoi1 SaiGon1] [HaNoi1 SaiGon1]
[HaiPhong SaiGon1] &[HaiPhong SaiGon1] [HaiPhong SaiGon1]
0xc00002a0a0 0xc00002a0a0
----------------------------------end----------------------------------------------------------------------------------
it is a bad change:
   in line 10: d.trips = trips , pointer to memory same copied.
   in line 37: changeBad[0] = "HaiPhong", when change element[0] of changeBad, d.strip[0] wil changed same. why it same pointer.
   in line 39: fmt.Println(&changeBad[0], &d.trips[0]) : 0xc00002a0a0 0xc00002a0a0 => we same pointer, very very risk when use.

+) follow line 42 to 47: result:
----------------------------------append and not change pointer---------------------------------------------------------
[HaiPhong SaiGon1] &[HaiPhong SaiGon1] [HaiPhong SaiGon1]
[HaiPhong Hue] &[HaiPhong Hue] [HaiPhong Hue]
0xc00002a0a0 0xc00002a0a0
----------------------------------end-----------------------------------------------------------------------------------
it is a bad change:
   in line 44: changeBad = append(changeBad[:(len(changeBad)-1)], "Hue") => change end element of badChange. With fc append(), it has sufficient capacity, the destination is resliced to accommodate the
   ==> pointer is same ==> result is wrong


+) fillow line 49 to 54: result:
----------------------------------append and change pointer--------------------------------------------------------------
[HaiPhong Hue] &[HaiPhong Hue] [HaiPhong Hue]
[HaiPhong Hue PhuQuoc] &[HaiPhong Hue PhuQuoc] [HaiPhong Hue]
0xc00007a040 0xc00002a0a0
----------------------------------end------------------------------------------------------------------------------------

it is a bad change but result not wronf, why?. Follow me:
   in line 51: changeBad = append(changeBad, "PhuQuoc"), add one element,  withc fc append(), it hasn't  sufficient, a new underlying array will be allocated
   => new memory, new pointer wil return
   in line fmt.Println(&changeBad[0], &d.trips[0]) : 0xc00007a040 0xc00002a0a0  => we not same pointer
   ==> when change changeBad, d.trip not change.

==================>>>>>> Summary >>>>>===================================================================================================
      when pointer have many rule in slice, in function of go,  to avoid risk, you should use good copy, never use bad copy
==================>>>>>> End >>>>>===================================================================================================
*/
