# gologstash
logstash version with golang coding

运行：生成应用程序后如hekad -f config配置文件

配置文件用法示例：
#协程个数...
maxprocs = 4

[LogstreamerInput]
decoder = "excel_transform_decoder"
log_directory = "/data/servers/BetTest/GolangWork/heka/log"
file_match = 'CaiNiXiHuan\.txt'

[excel_transform_decoder]
type = "PayloadRegexDecoder"
match_regex = '^(?P<Date>\S+)\s+(?P<DetailVisitUsers>\d+)\s+(?P<DetailVisitCount>\d+)\s+(?P<VisitUsers>\d+)\s+(?P<VisitCount>\d+)\s+(?P<TryDownLoadUsers>\d+)\s+(?P<TryDownLoadCount>\d+)\s+(?P<DownLoadUsers>\d+)\s+(?P<DownLoadCount>\d+)\s+(?P<TryRate>\S+)'
timestamp_layout = "02/Jan/2006:15:04:05 -0700"

[excel_transform_decoder.message_fields]
Type = "CaiNiXiHuanfile"
Logger = "/data/servers/BetTest/GolangWork/heka/log/CaiNiXiHuan.log"
TryRate = "%TryRate%"
Date|date-time = "%Date%"
DetailVisitUsers|int|INTEGER = "%DetailVisitUsers%"
DetailVisitCount|int|2 = "%DetailVisitCount%"
VisitUsers|int|uint = "%VisitUsers%"
VisitCount|int|uint = "%VisitCount%"
TryDownLoadUsers|int|count = "%TryDownLoadUsers%"
TryDownLoadCount|int|count = "%TryDownLoadCount%"
DownLoadUsers|int|count = "%DownLoadUsers%"
DownLoadCount|int|count = "%DownLoadCount%"


[ESJsonEncoder]
###append_newlines = false
#es_index_from_timestamp = true
index = "%{Type}-%{%Y.%m.%d}"
type_name = "%{Type}"

#[LogOutput]
#message_matcher = "TRUE"
#encoder = "ESJsonEncoder"

[ElasticSearchOutput]
message_matcher = "TRUE"
server = "http://183.131.145.158:9200"
encoder = "ESJsonEncoder"
