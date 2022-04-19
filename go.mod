module execlt1

go 1.16

require (
	github.com/Shopify/sarama v1.29.1
	github.com/cihub/seelog v0.0.0-20170130134532-f561c5e57575 // indirect
	github.com/coreos/etcd v3.3.25+incompatible
	github.com/denisenkom/go-mssqldb v0.10.0 // indirect
	github.com/evanphx/json-patch v4.9.0+incompatible
	github.com/extrame/ole2 v0.0.0-20160812065207-d69429661ad7 // indirect
	github.com/extrame/xls v0.0.1
	github.com/fsnotify/fsnotify v1.4.9
	github.com/garyburd/redigo v1.6.2
	github.com/gin-gonic/gin v1.7.2
	github.com/go-redis/redis v6.14.2+incompatible
	github.com/go-sql-driver/mysql v1.6.0
	github.com/go-workflow/go-workflow v0.0.0-20200122043112-85255031ec8d
	github.com/gogo/protobuf v1.3.2
	github.com/gohouse/converter v0.0.3
	github.com/gohouse/gorose v1.0.5
	github.com/jinzhu/gorm v1.9.16
	github.com/lack-io/vine v0.20.14
	github.com/lib/pq v1.10.1 // indirect
	github.com/lithammer/dedent v1.1.0
	github.com/lucas-clemente/quic-go v0.19.3 // indirect
	github.com/mattn/go-sqlite3 v2.0.3+incompatible // indirect
	github.com/micro/go-micro v1.18.0
	github.com/mumushuiding/util v0.0.0-20210203080010-04699a081184
	github.com/opentracing/opentracing-go v1.1.1-0.20190913142402-a7454ce5950e
	github.com/qianlnk/pgbar v0.0.0-20210208085217-8c19b9f2477e
	github.com/qianlnk/to v0.0.0-20191230085244-91e712717368 // indirect
	github.com/robfig/cron v1.2.0
	github.com/smallnest/rpcx v1.6.2
	github.com/spf13/cobra v1.1.1
	github.com/spf13/pflag v1.0.5
	github.com/tealeg/xlsx v1.0.5
	github.com/zngw/kafka v0.0.0-20191214161347-4191e5f8683f
	github.com/zngw/log v0.0.0-20200327115753-04ba41d5c8f8
	go.etcd.io/bbolt v1.3.6
	golang.org/x/net v0.0.0-20210614182718-04defd469f4e
	golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c // indirect
	google.golang.org/genproto v0.0.0-20210701191553-46259e63a0a9 // indirect
	google.golang.org/grpc v1.39.0
	google.golang.org/protobuf v1.27.1
	gorm.io/driver/sqlite v1.1.4
	gorm.io/gorm v1.21.13
	k8s.io/api v0.21.3
	k8s.io/apimachinery v0.21.3
	k8s.io/cli-runtime v0.21.3
	k8s.io/client-go v0.21.3
	k8s.io/component-base v0.21.3
	k8s.io/klog/v2 v2.10.0
	k8s.io/kubectl v0.21.3
	k8s.io/utils v0.0.0-20210722164352-7f3ee0f31471
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
