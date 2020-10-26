package rank

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/prometheus/common/log"
	"strings"
)

type Member struct {
	Name  string
	Score float64
	Rank  int
}
type RedisSettings struct {
	Addr string
}
type RedRank struct {
	Name     string
	PageSize int

	redisSettings RedisSettings
}

func NewRedRank(name string, redisSettings RedisSettings) *RedRank {
	return &RedRank{
		Name:          name,
		PageSize:      50,
		redisSettings: redisSettings,
	}
}

func getRedisDB(settings RedisSettings) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: settings.Addr,
	})
}


func (rl *RedRank) RankMember(name string, score float64) (Member) {
	rdb := getRedisDB(rl.redisSettings)
	defer rdb.Close()
	_,err := rdb.Do("zadd",rl.Name, score, name).Result()
	if err!=nil {
		log.Errorf("error on add member %s to rank!\n", name)
	}

	rank, err := rdb.ZRank(rl.Name, name).Result()
	if err!=nil {
		log.Errorf("error on get member %s rank\n", name)
	}

	member := Member{
		Name:  name,
		Score: score,
		Rank:  int(rank)+1,
	}
	return member
}

func (rl *RedRank) GetRankingList() []Member {
	rdb := getRedisDB(rl.redisSettings)
	defer rdb.Close()

	zs, _ := rdb.ZRevRangeWithScores(rl.Name, 0, -1).Result()
	members := make([]Member, 0, len(zs))
	for i:=0; i< len(zs); i++{
		member := Member{}
		member.Name = zs[i].Member.(string)
		member.Score = zs[i].Score

		rank ,_ := rdb.ZRank(rl.Name, member.Name).Result()
		member.Rank = int(rank) + 1

		members = append(members, member)
	}

	return members
}

func (rl *RedRank) TotalMembers() int {
	rdb := getRedisDB(rl.redisSettings)
	defer rdb.Close()

	total,err := rdb.ZCard(rl.Name).Result()
	if err!=nil {
		log.Errorf("error on get ranking list %d, %v\n", total, err)
	}
	return  int(total)
}

func (rl *RedRank) GetMember(name string) Member {
	rdb := getRedisDB(rl.redisSettings)
	defer rdb.Close()

	score, err := rdb.ZScore(rl.Name, name).Result()
	if err!=nil {
		log.Errorf("error on get member %s score", name)
	}
	rank, err := rdb.ZRevRank(rl.Name, name).Result()
	if err!= nil {
		log.Errorf("error on get member %s rank", name)
	}

	return Member{
		Name:  name,
		Score: score,
		Rank:  int(rank) +1,
	}
}

func (rl *RedRank) GetMemberByRank(rank int) Member {
	rdb := getRedisDB(rl.redisSettings)
	defer rdb.Close()

	zs,err := rdb.ZRevRangeWithScores(rl.Name, int64(rank-1), int64(rank-1)).Result()
	if err != nil {
		log.Errorf("error on get memeberByRank %v", err)
	}
	if len(zs) == 0{
		return Member{}
	}

	return Member{
		Name:  zs[0].Member.(string),
		Score: zs[0].Score,
		Rank:  rank,
	}
}

func (rl *RedRank) String() string {
	members := rl.GetRankingList()
	var builder strings.Builder
	for _,m := range members {
		builder.WriteString(fmt.Sprintf("%d    %s   %f\n", m.Rank, m.Name, m.Score))
	}
	fmt.Println(len(members))
	return builder.String()
}

