package main

import (
	"fmt"
)

func main() {
	gameId := []string{}

	sql := "SELECT user_id, real_name, player_name, player_id, server FROM gs WHERE gs=? AND is_auth=1"

    if len(gameId) != 0 {
        sql += " AND main_game_id IN ("
        for _, v := range gameId {
            sql += fmt.Sprintf("%s,", v)
        }
    sql = sql[:len(sql)-1]
	sql += ")"
    }
    sql += " LIMIT ?,?"

    fmt.Print(sql)
    fmt.Println(len("62497163,62609368,62727117,62792151,62946023,63049570,63232347,62137341,62245453,62372787,62374947,62374976,62375810,62377333,62495508,62496876,62497124,61730617,61732263,61732322,61781396,61819315,61920574,61992436,62037577,62134063,62136435,61669498,61708296,61708316"))
}
