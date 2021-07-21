### 這是練習[被選召的 Gopher 們，從零開始探索 Golang, Istio, K8s 數碼微服務世界 系列](https://ithelp.ithome.com.tw/users/20122925/ironman/3537) 的 Repository

# 遇到的問題

1. 在 build protobuf 時範例是採用 namely/gen-grpc-gateway 這個 image，但我試都會失敗，就算指定版本也是。後來改用 namely/protoc-all，因為沒有要使用 gateway 功能所以應該是不影響