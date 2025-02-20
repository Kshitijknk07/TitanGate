
package middleware

import (
	"context"
	"net/http"

	"github.com/golang/groupcache"
)

const cacheSize = 64 << 20 

var TitanCache = groupcache.NewGroup("TitanGateCache", cacheSize, groupcache.GetterFunc(
	func(ctx groupcache.Context, key string, dest groupcache.Sink) error {
		
		return nil
	},
))


func CacheMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.RequestURI 
		if data, err := GetFromCache(key); err == nil {
			w.Write(data) 
			return
		}

		
		rec := &responseRecorder{ResponseWriter: w}
		next.ServeHTTP(rec, r)

		SetCache(key, rec.body) 
	})
}


func GetFromCache(key string) ([]byte, error) {
	var data []byte
	err := TitanCache.Get(context.TODO(), key, groupcache.AllocatingByteSliceSink(&data))
	if err != nil {
		return nil, err
	}
	return data, nil
}


func SetCache(key string, value []byte) {
	
	TitanCache.Get(context.TODO(), key, groupcache.AllocatingByteSliceSink(&value))
}

type responseRecorder struct {
	http.ResponseWriter
	body []byte
}

func (r *responseRecorder) Write(b []byte) (int, error) {
	r.body = append(r.body, b...)
	return r.ResponseWriter.Write(b)
}
