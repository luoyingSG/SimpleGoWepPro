package middleware

import (
	"context"
	"net/http"
	"time"
)

type TimeoutMiddleware struct {
	Next http.Handler
}

func (tm TimeoutMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if tm.Next == nil {
		tm.Next = http.DefaultServeMux
	}

	ctx := r.Context()                               // 找到原有的 Context
	ctx, _ = context.WithTimeout(ctx, 3*time.Second) // 添加一个 timeout
	r.WithContext(ctx)                               // 基于 Context 进行修改
	ch := make(chan struct{})
	go func() {
		tm.Next.ServeHTTP(w, r)
		ch <- struct{}{} // 如果顺利处理完成请求，那么在管道中添加一个 struct{}{}
	}()
	select {
	case <-ch:
		return
	case <-ctx.Done():
		w.WriteHeader(http.StatusRequestTimeout)
	}
	ctx.Done()
}
