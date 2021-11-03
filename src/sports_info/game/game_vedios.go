package game

import (
	"context"
	"sports_info/utils"
	"sync"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type DBGameVedioInfo struct {
	EventGroupIdCust int32
	//PTP              int32
	SportIdCust int32
	//FEType           int32
	//SrcTP            int32
	//EventDate        time.Time
	EventDateCust time.Time
	//HomeTeam         string
	//AwayTeam         string
	//HomeTeamEn       string
	//AwayTeamEn       string
	//EventGroupId     int32
	//SportId          int32
	HomeTeamCust   string
	AwayTeamCust   string
	AwayTeamCustEn string
	HomeTeamCustEn string
	//CompName       string
	//CompNameEn     string
	CompNameCust string
	//CompNameCustEn string
	//ExVDataFlag    int32
	//HomeTeamId     int32
	//AwayTeamId     int32
	//CompId         int32
	HomeTeamIdCust int32
	AwayTeamIdCust int32
	CompIdCust     int32
	Vdata          []map[string]interface{}
	Gifdata        []map[string]interface{}
}

type DBGameVedioInfoSt struct {
	infos *[]*DBGameVedioInfo
	tm    int64
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
	inst._loadDB(DB_MATCH_HGUAN_IM, PLATFORM_IM)
	inst._loadDB(DB_MATCH_HGUAN_SABA, PLATFORM_SABA)
}

func (inst *GameVediosInfoManager) _loadDB(tab string, ptp int32) {
	cur := G_DBCollects.Find(tab, bson.M{"ExVDataFlag": 1}, nil)
	//cur := G_DBCollects.Find(DB_MATCH_HGUAN_IM, bson.M{}, nil)
	if cur == nil {
		G_log.Errorf("vedios info tab[%s] ptp[%d]  find err:%s", tab, ptp, G_DBCollects.Get_last_error())
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
		newPtpSt := DBGameVedioInfoSt{&newAry, curtm}
		inst.m_infos[ptp] = &newPtpSt
		infosAry = newPtpSt.infos
	}
	for cur.Next(context.Background()) {
		var dd DBGameVedioInfo
		err := cur.Decode(&dd)
		if err != nil {
			G_log.Errorf("vedios info tab[%s] ptp[%d]  decode object err", tab, ptp)
			continue
		}
		*infosAry = append(*infosAry, &dd)
	}
	inst.m_lock.Unlock()
}

func (inst *GameVediosInfoManager) initVedios() {
	inst.loadDB()
}

func (inst *GameVediosInfoManager) update(tm_sec int32) {
	inst.m_intervalTM += tm_sec
	if inst.m_intervalTM >= VEDIOS_INFO_UPDATE_TM {
		inst.m_intervalTM = 0
		inst.loadDB()
	}
}

func (inst *GameVediosInfoManager) getVediosInfo(ptp int32) interface{} {
	inst.m_lock.Lock()
	defer inst.m_lock.Unlock()
	info, ex := inst.m_infos[ptp]
	if !ex {
		return nil
	}

	return info.infos
}
