## map 에서 아이템 검색 방법 중에서 재귀 구문 사용하는 방법 (주석처리해서 없애버려서 일단 메모 해둠.)

```
// 재귀구문으로 처리	

j.subscribers.Range(func(k, v interface{}) bool {
			id, ok := k.(int64)
			if !ok {
				log.Printf("Failed to cast subscriber key: %T", k)
				return false
			}
			sub, ok := v.(sub)
			if !ok {
				log.Printf("Failed to cast subscriber value: %T", v)
				return false
			}

			if err := sub.stream.Send(&pb.JobsResponse{JobResId: id, OutputMessage: s}); err != nil {
				log.Printf("Failed to send data to client: %v", err)

				select {
				case sub.finished <- true:
					log.Printf("Unsubscribed client: %d", id)
				default:
					// Default case is to avoid blocking in case client has already unsubscribed
				}
				// In case of error the client would re-subscribe so close the subscriber stream
				unsubscribe = append(unsubscribe, id)
			}
			return true
		})

// for 구문으로 처리

     for k,k := range j.subscribers {
      // 어쩌구 저쩌구
     }
     
```

