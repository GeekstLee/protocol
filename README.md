# protocol
用于存放所有的公共协议文件

### 生成桩文件
protoc --go_out=plugins=grpc:. filePath

如：protoc --go_out=plugins=grpc:. index.proto


引用文件：protoc -I=. -I=../ --go_out=plugins=grpc:. algorithm.proto