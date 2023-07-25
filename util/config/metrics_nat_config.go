package config

var AllNATMetricConfigs = []KscMetricConfig{
	{
		Namespace:        "NAT",
		MetricName:       "vpc.nat.bps.in",
		MetricDesc:       "NAT入网流量",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "bps",
	},
	{
		Namespace:        "NAT",
		MetricName:       "vpc.nat.bps.out",
		MetricDesc:       "NAT出网流量",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "bps",
	},
	{
		Namespace:        "NAT",
		MetricName:       "vpc.nat.pps.in",
		MetricDesc:       "NAT每秒入包数",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "count",
	},
	{
		Namespace:        "NAT",
		MetricName:       "vpc.nat.pps.out",
		MetricDesc:       "NAT每秒出包数",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "count",
	},
	{
		Namespace:        "NAT",
		MetricName:       "vpc.nat.public.bps.in",
		MetricDesc:       "NAT(公网)入网流量",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "bps",
	}, {
		Namespace:        "NAT",
		MetricName:       "vpc.nat.public.bps.out",
		MetricDesc:       "NAT(公网)出网流量",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "bps",
	},
	{
		Namespace:        "NAT",
		MetricName:       "vpc.nat.public.pps.in",
		MetricDesc:       "NAT(公网)每秒入包数",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "count",
	},
	{
		Namespace:        "NAT",
		MetricName:       "vpc.nat.public.pps.out",
		MetricDesc:       "NAT(公网)每秒出包数",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "count",
	},
	{
		Namespace:        "NAT",
		MetricName:       "vpc.nat.public.utilization.in",
		MetricDesc:       "NAT（公网）入网带宽使用率",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "%",
	},
	{
		Namespace:        "NAT",
		MetricName:       "vpc.nat.public.utilization.out",
		MetricDesc:       "NAT（公网）出网带宽使用率",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "%",
	},
	{
		Namespace:        "NAT",
		MetricName:       "vpc.nat.ipconflict",
		MetricDesc:       "NAT（IP）端口占满",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "",
	},
	{
		Namespace:        "NAT",
		MetricName:       "vpc.nat.concur_connect_num",
		MetricDesc:       "NAT并发连接数",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "count",
	},
	{
		Namespace:        "NAT",
		MetricName:       "vpc.nat.concur_dropped_connect_rate",
		MetricDesc:       "NAT并发丢弃连接速率",
		MetricType:       2,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "count",
	},
}