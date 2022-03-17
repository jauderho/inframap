package aws

var (
	// nodeTypes all the Nodes that we support right now
	nodeTypes = map[string]struct{}{
		"aws_api_gateway_resource":          {},
		"aws_athena_database":               {},
		"aws_autoscaling_group":             {},
		"aws_dax_cluster":                   {},
		"aws_db_instance":                   {},
		"aws_directory_service_directory":   {},
		"aws_dms_replication_instance":      {},
		"aws_dx_gateway":                    {},
		"aws_dynamodb_table":                {},
		"aws_ebs_volume":                    {},
		"aws_ecs_cluster":                   {},
		"aws_ecs_service":                   {},
		"aws_efs_file_system":               {},
		"aws_eip":                           {},
		"aws_eks_cluster":                   {},
		"aws_elasticache_cluster":           {},
		"aws_elastic_beanstalk_application": {},
		"aws_elasticsearch_domain":          {},
		"aws_elb":                           {},
		"aws_emr_cluster":                   {},
		"aws_iam_user":                      {},
		"aws_instance":                      {},
		"aws_internet_gateway":              {},
		"aws_kinesis_stream":                {},
		"aws_lambda_function":               {},
		"aws_lightsail_instance":            {},
		"aws_mq_broker":                     {},
		"aws_media_store_container":         {},
		"aws_nat_gateway":                   {},
		"aws_rds_cluster":                   {},
		"aws_rds_cluster_instance":          {},
		"aws_redshift_cluster":              {},
		"aws_s3_bucket":                     {},
		"aws_storagegateway_gateway":        {},
		"aws_sqs_queue":                     {},
		"aws_vpn_gateway":                   {},
		"aws_batch_job_definition":          {},
		"aws_neptune_cluster":               {},
		"aws_alb":                           {},
		"aws_lb":                            {},
		"aws_cloudfront_distribution":       {},
		"aws_elasticache_replication_group": {},
		"aws_launch_template":               {},
	}

	// noSecurityGroup is a map of all the resourcese that do not
	// have/relay on SecurityGroups
	noSecurityGroup = map[string]struct{}{
		"aws_s3_bucket":               {},
		"aws_cloudfront_distribution": {},
	}

	// edgeTypes map of all the supported Edges
	edgeTypes = map[string]struct{}{
		"aws_security_group":      {},
		"aws_security_group_rule": {},
	}
)
