package game

import (
	"context"
	"sync"
	"test_gin/utils"
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
	//EventDateCust    time.Time
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
	infos []DBGameVedioInfo
	tm    int64
}
type GameVediosInfoManager struct {
	m_lock       sync.Mutex
	m_infos      map[int32]DBGameVedioInfoSt
	m_intervalTM int32
	//m_data map[string]interface{}
}

func (inst *GameVediosInfoManager) init() {
	utils.Once(func() {
		inst.m_infos = map[int32]DBGameVedioInfoSt{}
		inst.initVedios()
	})
}

func (inst *GameVediosInfoManager) loadDB() {
	inst._loadDB(DB_MATCH_HGUAN_IM, PLATFORM_IM)
}

func (inst *GameVediosInfoManager) _loadDB(tab string, ptp int32) {
	cur := G_DBCollects.Find(tab, bson.M{"ExVDataFlag": 1, "PTP": ptp}, nil)
	//cur := G_DBCollects.Find(DB_MATCH_HGUAN_IM, bson.M{}, nil)
	if cur == nil {
		return
	}
	infos := []DBGameVedioInfo{}
	curtm := time.Now().Unix()
	for cur.Next(context.Background()) {
		var dd DBGameVedioInfo
		err := cur.Decode(&dd)
		if err != nil {
			G_log.Errorf("vedios info tab[%s] ptp[%d]  decode object err", tab, ptp)
			continue
		}
		infos = append(infos, dd)
	}
	//jj := utils.ToJson(infos[0])
	//if len(jj) <= 0 {
	//	G_log.Errorf("vedios info tab[%s] ptp[%d]  encode json err", tab, ptp)
	//	return
	//}
	infosst := DBGameVedioInfoSt{infos, curtm}
	//json 字符串
	//jj, ee := json.Marshal(infos)
	//if ee == nil {
	//	fmt.Println(string(jj))
	//}

	inst.m_lock.Lock()
	inst.m_infos[ptp] = infosst
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
