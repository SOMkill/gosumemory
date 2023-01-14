package pp

func GetFCData() {
	// ezfc = C.ezpp_new()
	// C.ezpp_set_autocalc(ezfc, 1)
	// for {

	// 	if memory.DynamicAddresses.IsReady == true {

	// 		switch memory.GameplayData.GameMode {
	// 		case 0, 1:

	// 			if memory.MenuData.OsuStatus == 2 && memory.GameplayData.Combo.Max > 0 {
	// 				var data PPfc
	// 				readFCData(&data, ezfc, C.float(memory.GameplayData.Accuracy))
	// 				res, err := wiekuCalcCrutch(memory.MenuData.Bm.Path.FullDotOsu, int16(memory.MenuData.Bm.Stats.BeatmapMaxCombo), int16(C.ezpp_nobjects(ezfc)-1)-memory.GameplayData.Hits.H100-memory.GameplayData.Hits.H50, memory.GameplayData.Hits.H100, memory.GameplayData.Hits.H50, 0)
	// 				if err != nil {
	// 					pp.Println(err)
	// 					memory.GameplayData.PP.PPifFC = cast.ToInt32(float64(data.RestSS))
	// 				} else {
	// 					memory.GameplayData.PP.PPifFC = res
	// 				}

	// 				memory.GameplayData.Hits.Grade.Current = data.GradeCurrent
	// 				memory.GameplayData.Hits.Grade.Expected = data.GradeExpected
	// 			}
	// 			switch memory.MenuData.OsuStatus {
	// 			case 1, 4, 5, 13, 2:
	// 				if memory.MenuData.OsuStatus == 2 && memory.MenuData.Bm.Time.PlayTime > 150 { //To catch up with the F2-->Enter
	// 					time.Sleep(250 * time.Millisecond)
	// 					continue
	// 				}
	// 				//TODO: figure out how to calc %% pp on the new rework
	// 				// if memory.GameplayData.GameMode == 0 {
	// 				// 	wiekuCalcCrutch(memory.MenuData.Bm.Path.FullDotOsu, int16(memory.MenuData.Bm.Stats.BeatmapMaxCombo), desired300Hits())
	// 				// }
	// 				var data PPfc
	// 				readFCData(&data, ezfc, 100.0)
	// 				memory.MenuData.PP.PpSS = cast.ToInt32(float64(data.Acc))
	// 				readFCData(&data, ezfc, 99.0)
	// 				memory.MenuData.PP.Pp99 = cast.ToInt32(float64(data.Acc))
	// 				readFCData(&data, ezfc, 98.0)
	// 				memory.MenuData.PP.Pp98 = cast.ToInt32(float64(data.Acc))
	// 				readFCData(&data, ezfc, 97.0)
	// 				memory.MenuData.PP.Pp97 = cast.ToInt32(float64(data.Acc))
	// 				readFCData(&data, ezfc, 96.0)
	// 				memory.MenuData.PP.Pp96 = cast.ToInt32(float64(data.Acc))
	// 				readFCData(&data, ezfc, 95.0)
	// 				memory.MenuData.PP.Pp95 = cast.ToInt32(float64(data.Acc))
	// 			}
	// 		}

	// 	}

	// 	time.Sleep(250 * time.Millisecond)
	// }
}
