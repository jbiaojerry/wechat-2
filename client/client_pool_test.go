// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"bytes"
	"testing"
)

// NOTE: 这个测试的结果如果很小则说明 pool 也是正常工作
func BenchmarkGetBufferFromPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() {
			buf := _test_client.getBufferFromPool()
			defer _test_client.putBufferToPool(buf)
		}()
	}
}

func BenchmarkGetBufferFromNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() {
			_ = bytes.NewBuffer(make([]byte, 2<<20)) // 默认 2MB
		}()
	}
}
