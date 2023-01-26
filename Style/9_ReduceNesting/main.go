package __ReduceNesting

/**
Code should reduce nesting where possible by handling error cases/special conditions first and returning early or continuing the loop. Reduce the amount of code that is nested multiple levels.
*/

// this is bad
//func bad() {
//	for _, v := range data {
//		if v.F1 == 1 {
//			v = process(v)
//			if err := v.Call(); err == nil {
//				v.Send()
//			} else {
//				return err
//			}
//		} else {
//			log.Printf("Invalid v: %v", v)
//		}
//	}
//}

// this is good
//func good() {
//	for _, v := range data {
//		if v.F1 != 1 {
//			log.Printf("Invalid v: %v", v)
//			continue
//		}
//
//		v = process(v)
//		if err := v.Call(); err != nil {
//			return err
//		}
//		v.Send()
//	}
//}
