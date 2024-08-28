package pp

import (
	"time"

	"unsafe"

	"github.com/l3lackShark/gosumemory/memory"
	"github.com/spf13/cast"
)

import "C"

type Calculator struct {
	MapPath string
}

type ScoreParams struct {
	Mode          uint
	Mods          uint
	MaxCombo      uint
	Accuracy      float64
	MissCount     uint
	PassedObjects uint
}

type CalculatePerformanceResult struct {
	PP    float64
	Stars float64
}

func (calculator Calculator) Calculate(params ScoreParams) CalculatePerformanceResult {
	cMapPath := C.CString(calculator.MapPath)
	defer C.free(unsafe.Pointer(cMapPath))

	passedObjects := C.optionu32{t: C.uint(0), is_some: C.uchar(0)}
	if params.PassedObjects > 0 {
		passedObjects = C.optionu32{t: C.uint(params.PassedObjects), is_some: C.uchar(1)}
	}

	rawResult := C.calculate_score(
		cMapPath,
		C.uint(params.Mode),
		C.uint(params.Mods),
		C.uint(params.MaxCombo),
		C.double(params.Accuracy),
		C.uint(params.MissCount),
		passedObjects,
	)
	return CalculatePerformanceResult{PP: float64(rawResult.pp), Stars: float64(rawResult.stars)}
}

var strainArray []float64
var tempBeatmapFile string
var tempGameMode int32 = 4

var maniaSR float64
var maniaHitObjects float64
var tempMods string

func GetData() {
	for {

		if memory.DynamicAddresses.IsReady {
			if tempBeatmapFile != memory.MenuData.Bm.Path.BeatmapOsuFileString || memory.MenuData.Mods.PpMods != tempMods || memory.MenuData.GameMode != tempGameMode { //On map/mods change
				tempGameMode = memory.MenuData.GameMode // looks very ugly but will rewrite everything in 1.4.0
				tempBeatmapFile = memory.MenuData.Bm.Path.BeatmapOsuFileString
				tempMods = memory.MenuData.Mods.PpMods
				tempGameMode = memory.MenuData.GameMode
				mp3Time, err := calculateMP3Time()
				if err == nil {
					memory.MenuData.Bm.Time.Mp3Time = mp3Time
				}
			}

			switch memory.MenuData.OsuStatus {
			case 2:
				path := memory.MenuData.Bm.Path.FullDotOsu
				if memory.GameplayData.Combo.Max > 1 {
					calc := Calculator{MapPath: path}
					scoreParams := ScoreParams{
						Mode:          uint(memory.GameplayData.GameMode),
						Mods:          uint(memory.MenuData.Mods.AppliedMods),
						MaxCombo:      uint(memory.GameplayData.Combo.Max),
						Accuracy:      memory.GameplayData.Accuracy,
						MissCount:     uint(memory.GameplayData.Hits.H0),
						PassedObjects: 0, // TODO: fix
					}

					result := calc.Calculate(scoreParams)
					memory.GameplayData.PP.Pp = cast.ToInt32(result.PP)
				}
			case 7:
				//idle
			case 5:
				memory.GameplayData.PP.Pp = 0
			}
		}

		time.Sleep(time.Duration(memory.UpdateTime) * time.Millisecond)
	}
}
