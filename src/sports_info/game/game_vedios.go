package game

import (
	"context"
	"sports_info/utils"
	"sync"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type DBGameVedioInfoNormal struct {
	//EventGroupIdCust int32
	//PTP              int32
	//SportIdCust int32
	////FEType           int32
	////SrcTP            int32
	////EventDate        time.Time
	//EventDateCust time.Time
	//HomeTeam string
	//AwayTeam string
	////HomeTeamEn       string
	////AwayTeamEn       string
	EventGroupId int32
	SportId      int32
	//HomeTeamCust string
	//AwayTeamCust string
	//AwayTeamCustEn string
	//HomeTeamCustEn string
	////CompName       string
	////CompNameEn     string
	//CompNameCust string
	////CompNameCustEn string
	////ExVDataFlag    int32
	////HomeTeamId     int32
	////AwayTeamId     int32
	////CompId         int32
	//HomeTeamIdCust int32
	//AwayTeamIdCust int32
	//CompIdCust     int32
	Vdata []map[string]interface{}
	//Gifdata []map[string]interface{}
}

type DBGameVedioInfo struct {
	EventGroupIdCust int32
	//PTP              int32
	SportIdCust int32
	////FEType           int32
	////SrcTP            int32
	////EventDate        time.Time
	//EventDateCust time.Time
	////HomeTeam         string
	////AwayTeam         string
	////HomeTeamEn       string
	////AwayTeamEn       string
	////EventGroupId     int32
	////SportId          int32
	//HomeTeamCust string
	//AwayTeamCust string
	//AwayTeamCustEn string
	//HomeTeamCustEn string
	////CompName       string
	////CompNameEn     string
	//CompNameCust string
	////CompNameCustEn string
	////ExVDataFlag    int32
	////HomeTeamId     int32
	////AwayTeamId     int32
	////CompId         int32
	//HomeTeamIdCust int32
	//AwayTeamIdCust int32
	//CompIdCust     int32
	Vdata []map[string]interface{}
	//Gifdata []map[string]interface{}
}

type DBGameVedioInfoSt struct {
	infosmap *map[int32]*[]*DBGameVedioInfo
	infos    *[]*DBGameVedioInfo
	tm       int64
}
type GameVediosInfoManager struct {
	m_lock       sync.Mutex
	m_infos      map[int32]*DBGameVedioInfoSt
	m_intervalTM int32
	//m_data map[string]interface{}
}

func (inst *GameVediosInfoManager) init() {
	utils.Once(func() {
		inst.initVedios()
	})
}

func (inst *GameVediosInfoManager) loadDB() {
	inst.m_infos = map[int32]*DBGameVedioInfoSt{}

	//im
	inst._loadDB(DB_REAL_TIME_EVENT, PLATFORM_YABOIM, PLATFORM_IM, true)
	inst._loadDB(DB_REAL_TIME_EVENT_MATCH_HGUAN_IM, 0, PLATFORM_IM, false)
	inst._loadDB(DB_REAL_TIME_EVENT_MATCH_M602_IM, 0, PLATFORM_IM, false)
	inst._loadDB(DB_REAL_TIME_EVENT_MATCH_M603_IM, 0, PLATFORM_IM, false)
	inst._loadDB(DB_REAL_TIME_EVENT_MATCH_SABA_IM, 0, PLATFORM_IM, false)

	//saba
	inst._loadDB(DB_REAL_TIME_EVENT_MATCH_HGUAN_SABA, 0, PLATFORM_SABA, false)
	inst._loadDB(DB_REAL_TIME_EVENT_MATCH_YABOIM_SABA, 0, PLATFORM_SABA, false)
	inst._loadDB(DB_REAL_TIME_EVENT_MATCH_M602_SABA, 0, PLATFORM_SABA, false)
	inst._loadDB(DB_REAL_TIME_EVENT_MATCH_M603_SABA, 0, PLATFORM_SABA, false)
}

func (inst *GameVediosInfoManager) _loadDB(tab string, srctp int32, ptp int32, normal bool) {
	tt := time.Now().Unix()
	ttt := time.Unix(tt-4*60*60, 0)

	var filter *bson.M
	if normal {
		filter = &bson.M{
			"ExVDataFlag": 1,
			"srcTP":       srctp,
			"EventDate": bson.M{
				"$gt": ttt,
			}}
	} else {

		filter = &bson.M{
			"ExVDataFlag": 1,
			"EventDateCust": bson.M{
				"$gt": ttt,
			}}
	}

	cur := G_DBCollects.Find(tab, *filter, nil)

	//	cur := G_DBCollects.Find(tab, bson.M{
	//		"ExVDataFlag": 1,
	//		"EventDateCust": bson.M{
	//			"$gt": ttt,
	//		}}, nil)
	//
	//cur := G_DBCollects.Find(DB_MATCH_HGUAN_IM, bson.M{}, nil)
	if cur == nil {
		G_log.Errorf("vedios info tab[%s] srctp[%d] ptp[%d]  find err:%s", tab, srctp, ptp, G_DBCollects.Get_last_error())
		return
	}

	inst.m_lock.Lock()
	var infosAry *[]*DBGameVedioInfo
	ptpSt, ex := inst.m_infos[ptp]
	if ex {
		infosAry = ptpSt.infos
	} else {
		curtm := time.Now().Unix()
		newAry := []*DBGameVedioInfo{}

		newPtpSt := DBGameVedioInfoSt{&map[int32]*[]*DBGameVedioInfo{}, &newAry, curtm}
		inst.m_infos[ptp] = &newPtpSt
		infosAry = newPtpSt.infos
	}
	for cur.Next(context.Background()) {
		var ddd DBGameVedioInfo
		if normal {
			var dd DBGameVedioInfoNormal
			err := cur.Decode(&dd)
			if err != nil {
				G_log.Errorf("vedios info tab[%s] ptp[%d]  decode object err", tab, ptp)
				continue
			}
			ddd.EventGroupIdCust = dd.EventGroupId
			ddd.SportIdCust = dd.SportId
			//ddd.Gifdata = dd.Gifdata
			//ddd.HomeTeamCust = dd.HomeTeam
			//ddd.AwayTeamCust = dd.AwayTeam

			Vdata := []map[string]interface{}{}
			vdataobj := map[string]interface{}{}
			vdataobj["srctp"] = srctp
			v_ary := []map[string]interface{}{}
			for _, av := range dd.Vdata {
				v_ary = append(v_ary, av)
			}
			vdataobj["vdata"] = v_ary
			Vdata = append(Vdata, vdataobj)
			ddd.Vdata = Vdata
		} else {

			err := cur.Decode(&ddd)
			if err != nil {
				G_log.Errorf("vedios info tab[%s] ptp[%d]  decode object err", tab, ptp)
				continue
			}
		}
		*infosAry = append(*infosAry, &ddd)
		ary, emx := (*inst.m_infos[ptp].infosmap)[ddd.SportIdCust]
		if !emx {

			(*inst.m_infos[ptp].infosmap)[ddd.SportIdCust] = &[]*DBGameVedioInfo{}
			ary = (*inst.m_infos[ptp].infosmap)[ddd.SportIdCust]
		}
		*ary = append(*ary, &ddd)
	}
	inst.m_lock.Unlock()
}

func (inst *GameVediosInfoManager) initVedios() {

	//ary := []map[string]interface{}{}
	//tar_ary := []map[string]interface{}{}
	//for _, av := range ary {
	//	tar_ary = append(tar_ary, av)
	//	//fmt.Println(i)
	//}
	//fmt.Println(tar_ary)

	inst.loadDB()
}

func (inst *GameVediosInfoManager) update(tm_sec int32) {
	inst.m_intervalTM += tm_sec
	if inst.m_intervalTM >= VEDIOS_INFO_UPDATE_TM {
		inst.m_intervalTM = 0
		inst.loadDB()
	}
}

func (inst *GameVediosInfoManager) getVediosInfo(ptp int32, spid int32) interface{} {
	inst.m_lock.Lock()
	defer inst.m_lock.Unlock()
	info, ex := inst.m_infos[ptp]
	if !ex {
		return nil
	}
	if spid > 0 {
		v := (*info.infosmap)[spid]
		return v
	} else {
		return info.infos
	}

}
