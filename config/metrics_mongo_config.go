package config

var AllMongoDBMetricConfigs = []KscMetricConfig{
	{
		Namespace:        "MONGO",
		MetricName:       "mongo.cpu.used",
		MetricDesc:       "CPU使用率",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "%",
	},
	{
		Namespace:        "MONGO",
		MetricName:       "mongo.cursor.open_total",
		MetricDesc:       "cursor.open_total",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "",
	},
	{
		Namespace:        "MONGO",
		MetricName:       "mongo.cursor.timeout",
		MetricDesc:       "cursor.timeout",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "",
	},
	{
		Namespace:        "MONGO",
		MetricName:       "mongo.disk.percent",
		MetricDesc:       "disk.percent",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "",
	},
	{
		Namespace:        "MONGO",
		MetricName:       "mongo.disk.used",
		MetricDesc:       "磁盘使用空间",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "GB",
	},
	{
		Namespace:        "MONGO",
		MetricName:       "mongo.iops.amount",
		MetricDesc:       "iops.amount",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "次/s",
	},
	{
		Namespace:        "MONGO",
		MetricName:       "mongo.iops.percent",
		MetricDesc:       "iops.percent",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "%",
	},
	{
		Namespace:        "MONGO",
		MetricName:       "mongo.lock.readers",
		MetricDesc:       "lock.readers",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "个",
	},
	{
		Namespace:        "MONGO",
		MetricName:       "mongo.lock.total",
		MetricDesc:       "lock.total",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "个",
	},
	{
		Namespace:        "MONGO",
		MetricName:       "mongo.lock.writers",
		MetricDesc:       "lock.writers",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "个",
	},
	{
		Namespace:        "MONGO",
		MetricName:       "mongo.mem.memused.amount",
		MetricDesc:       "内存使用量",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "GB",
	},
	{
		Namespace:        "MONGO",
		MetricName:       "mongo.mem.memused.percent",
		MetricDesc:       "内存使用率",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "%",
	},
	{
		Namespace:        "MONGO",
		MetricName:       "mongo.mongo_alive",
		MetricDesc:       "mongodb实例存活监控",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "",
	},
	{
		Namespace:        "MONGO",
		MetricName:       "mongo.network.bytesin",
		MetricDesc:       "入流量",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "",
	},
	{
		Namespace:        "MONGO",
		MetricName:       "mongo.network.bytesout",
		MetricDesc:       "出流量",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "",
	},
	{
		Namespace:        "MONGO",
		MetricName:       "mongo.network.connections",
		MetricDesc:       "连接数",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "",
	},
	{
		Namespace:        "MONGO",
		MetricName:       "mongo.network.connections_percent",
		MetricDesc:       "connections_percent",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "%",
	},
	{
		Namespace:        "MONGO",
		MetricName:       "mongo.network.numrequesets",
		MetricDesc:       "numrequesets",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "B/s",
	},
	{
		Namespace:        "MONGO",
		MetricName:       "mongo.qps.command",
		MetricDesc:       "总qps",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "",
	},
	{
		Namespace:        "MONGO",
		MetricName:       "mongo.qps.count",
		MetricDesc:       "Count查询次数",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "",
	},
	{
		Namespace:        "MONGO",
		MetricName:       "mongo.qps.delete",
		MetricDesc:       "删除次数",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "",
	},
	{
		Namespace:        "MONGO",
		MetricName:       "mongo.qps.getmore",
		MetricDesc:       "聚合查询次数",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "",
	},
	{
		Namespace:        "MONGO",
		MetricName:       "mongo.qps.insert",
		MetricDesc:       "插入次数",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "",
	},
	{
		Namespace:        "MONGO",
		MetricName:       "mongo.qps.query",
		MetricDesc:       "读取次数",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "",
	},
	{
		Namespace:        "MONGO",
		MetricName:       "mongo.qps.update",
		MetricDesc:       "更新次数",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "",
	},
	{
		Namespace:        "MONGO",
		MetricName:       "mongo.replica_health",
		MetricDesc:       "mongodb副本健康状态",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "",
	},
	{
		Namespace:        "MONGO",
		MetricName:       "mongo.replica_latency",
		MetricDesc:       "mongodb复制延迟",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "",
	},
	{
		Namespace:        "MONGO",
		MetricName:       "mongo.rs.slavedelay",
		MetricDesc:       "主从延迟",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "秒",
	},
	{
		Namespace:        "MONGO",
		MetricName:       "mongo.wtcache.readinto_b",
		MetricDesc:       "bytes_read_into_cache",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "B/s",
	},
	{
		Namespace:        "MONGO",
		MetricName:       "mongo.wtcache.writtenfrom_b",
		MetricDesc:       "bytes_written_from_cache",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "B/s",
	},
}
