module github.com/futurez/freedom/examples/client

go 1.17

require github.com/futurez/freedom v0.0.0-00010101000000-000000000000

require (
	github.com/bwmarrin/snowflake v0.3.0 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.21.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
)

replace github.com/futurez/freedom => ../../../freedom
