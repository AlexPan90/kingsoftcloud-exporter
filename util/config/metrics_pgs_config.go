package config

var AllPGSMetricConfigs = []KscMetricConfig{
	{
		Namespace:        "PGS",
		MetricName:       "postgresql.riops",
		MetricDesc:       "磁盘读IOPS",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "次/秒",
	}, {
		Namespace:        "PGS",
		MetricName:       "postgresql.wiops",
		MetricDesc:       "磁盘写IOPS",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "次/秒",
	}, {
		Namespace:        "PGS",
		MetricName:       "postgresql.rbps",
		MetricDesc:       "读吞吐",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "字节/秒",
	}, {
		Namespace:        "PGS",
		MetricName:       "postgresql.wbps",
		MetricDesc:       "写吞吐",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "字节/秒",
	}, {
		Namespace:        "PGS",
		MetricName:       "postgresql.resident_memory_size",
		MetricDesc:       "内存使用量",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "MB",
	}, {
		Namespace:        "PGS",
		MetricName:       "postgresql.space_used_percent",
		MetricDesc:       "磁盘使用率",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "%",
	}, {
		Namespace:        "PGS",
		MetricName:       "postgresql.usage_in_percent",
		MetricDesc:       "内存使用率",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "%",
	}, {
		Namespace:        "PGS",
		MetricName:       "postgresql.cpu_used_percent",
		MetricDesc:       "CPU使用率",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "%",
	}, {
		Namespace:        "PGS",
		MetricName:       "postgresql.clientconns",
		MetricDesc:       "当前连接数",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "个",
	}, {
		Namespace:        "PGS",
		MetricName:       "postgresql.clientqps",
		MetricDesc:       "QPS",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "次/秒",
	}, {
		Namespace:        "PGS",
		MetricName:       "postgresql.tps",
		MetricDesc:       "TPS",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "次/秒",
	}, {
		Namespace:        "PGS",
		MetricName:       "postgresql.pg_xlog",
		MetricDesc:       "事务日志使用量",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "MB",
	}, {
		Namespace:        "PGS",
		MetricName:       "postgresql.max_used_xids",
		MetricDesc:       "事务最大已使用ID数",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "个",
	}, {
		Namespace:        "PGS",
		MetricName:       "postgresql.xlog_size_speed",
		MetricDesc:       "事务日志生成速率",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "MB/s",
	}, {
		Namespace:        "PGS",
		MetricName:       "postgresql.cache_hit",
		MetricDesc:       "缓冲区缓存命中率",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "%",
	}, {
		Namespace:        "PGS",
		MetricName:       "postgresql.top_ten_execute_delay",
		MetricDesc:       "最长TOP10执行时延",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "ms",
	}, {
		Namespace:        "PGS",
		MetricName:       "postgresql.unused_xids",
		MetricDesc:       "剩余XID数量",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "个",
	}, {
		Namespace:        "PGS",
		MetricName:       "postgresql.total_requests",
		MetricDesc:       "总请求数",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "次/秒",
	}, {
		Namespace:        "PGS",
		MetricName:       "postgresql.read_requests",
		MetricDesc:       "读请求数",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "次/秒",
	}, {
		Namespace:        "PGS",
		MetricName:       "postgresql.write_requests",
		MetricDesc:       "写请求数",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "次/秒",
	}, {
		Namespace:        "PGS",
		MetricName:       "postgresql.delay",
		MetricDesc:       "主从复制时延",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "秒",
	}, {
		Namespace:        "PGS",
		MetricName:       "postgresql.pg_xlog_location_diff",
		MetricDesc:       "主备XLOG同步差异",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "字节",
	}, {
		Namespace:        "PGS",
		MetricName:       "postgresql.replica_lagest_lag",
		MetricDesc:       "最滞后副本滞后量",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "MB",
	},
}