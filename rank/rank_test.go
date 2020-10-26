package rank

import (
	"testing"
)

var redisSetting = RedisSettings {
			Addr: ":6379",
		}

func TestRankMember(t *testing.T) {
	accountBalanceRank := NewRedRank("acccount_balance_rank", redisSetting)
	accountBalanceRank.RankMember("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",7413160.6451144)
	accountBalanceRank.RankMember("0x742d35cc6634c0532925a3b844bc454e4438f44e",3735282.85844998)
	accountBalanceRank.RankMember("0xbe0eb53f46cd790cd13851d5eff43d12404d33e8", 2235761.37283618)
	accountBalanceRank.RankMember("0xdc76cd25977e0a5ae17155770273ad58648900d3",1925760.47046488)

	if m:=accountBalanceRank.GetMember("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"); m.Rank != 1{
		t.Errorf("rankMember error, account rank %d, expected %d", m.Rank, 1)
	}
	if m:=accountBalanceRank.GetMember("0x742d35cc6634c0532925a3b844bc454e4438f44e"); m.Rank != 2{
		t.Errorf("rankMember error, account rank %d, expected %d", m.Rank, 2)
	}
}