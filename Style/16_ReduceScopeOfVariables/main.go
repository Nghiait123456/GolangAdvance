package ReduceScopeOfVariables

/**
Where possible, reduce scope of variables. Do not reduce the scope if it conflicts with Reduce Nesting.
*/

// this is bad
//err := os.WriteFile(name, data, 0644)
//if err != nil {
//return err
//}

// this is good
//if err := os.WriteFile(name, data, 0644); err != nil {
//return err
//}

/**
If you need a result of a function call outside of the if, then you should not try to reduce the scope.
*/

// this is bad
//func bad() {
//	if data, err := os.ReadFile(name); err == nil {
//		err = cfg.Decode(data)
//		if err != nil {
//			return err
//		}
//
//		fmt.Println(cfg)
//		return nil
//	} else {
//		return err
//	}
//}

// this is good
//func good() {
//	data, err := os.ReadFile(name)
//	if err != nil {
//		return err
//	}
//
//	if err := cfg.Decode(data); err != nil {
//		return err
//	}
//
//	fmt.Println(cfg)
//	return nil
//}
