package model

import (
	"log"
	"testing"
)
import "github.com/go-redis/redis/v7"

func TestNewScript(t *testing.T) {
	mapC := make(map[string]*Character)
	mapS := make(map[string]*Scene)
	character := NewCharacter("1号", "1号的剧本", "1号的目标")
	scene := NewScene("场景1", "这是场景1的描述")
	mapC[character.Name] = character
	mapS[scene.Name] = scene
	script := NewScript("剧本1", "这是剧本1描述", mapC, mapS)
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	err := client.Set(script.Name, script.Encode(), 0).Err()
	if err != nil {
		log.Fatal(err.Error())
	}
	val, err := client.Get(script.Name).Bytes()
	if err != nil {
		log.Fatal(err.Error())
	}
	retrieved := &Script{}
	retrieved.Decode(val)
	t.Log(retrieved.String())
}
