package reply

import (
	"fmt"
	"testing"
)

// TestTulingReply
func TestTulingReply(t *testing.T) {
	r := NewTulingReplier()
	keys := []string{"你好", "北京到武汉的高铁", "我想睡你", "鱼香肉丝怎么做", "明天北京到拉萨的飞机", "关于风险的新闻"}
	for _, v := range keys {
		s := r.Reply(v)
		fmt.Printf("say: %s\nreply: %s\n\n", v, s)
	}
}
