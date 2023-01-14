package pp

func GetEditorData() {

	// for {
	// 	if memory.DynamicAddresses.IsReady == true {
	// 		if memory.MenuData.GameMode == 0 && memory.MenuData.OsuStatus == 1 {
	// 			err := readEditorData(&data, ezeditor, false)
	// 			if err != nil {
	// 				fmt.Println(err)
	// 			}
	// 			memory.GameplayData.PP.Pp = int32(data.Total)
	// 			memory.MenuData.Bm.Stats.BeatmapAR = float32(data.AR)
	// 			memory.MenuData.Bm.Stats.BeatmapCS = float32(data.CS)
	// 			memory.MenuData.Bm.Stats.BeatmapOD = float32(data.OD)
	// 			memory.MenuData.Bm.Stats.BeatmapHP = float32(data.HP)
	// 			memory.MenuData.Bm.Stats.BeatmapSR = cast.ToFloat32(fmt.Sprintf("%.2f", float32(data.StarRating)))

	// 			md5, err := hashFileMD5(memory.MenuData.Bm.Path.FullDotOsu)
	// 			if err != nil {
	// 				continue
	// 			}
	// 			if tempOsuMD5 != md5 {
	// 				tempOsuMD5 = md5
	// 				readEditorData(&data, ezeditor, true)
	// 				memory.MenuData.PP.PpStrains = data.Strain
	// 			}

	// 		}

	// 	}
	// 	time.Sleep(time.Duration(memory.UpdateTime) * time.Millisecond)
	// }

}
