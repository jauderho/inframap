package generate_test

import (
	"testing"

	"github.com/cycloidio/inframap/generate"
	"github.com/cycloidio/inframap/graph"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
)

func TestFromHCL_AWS(t *testing.T) {
	t.Run("SuccessSG", func(t *testing.T) {
		fs := afero.NewOsFs()

		g, err := generate.FromHCL(fs, "./testdata/aws_hcl_sg.tf", generate.Options{Clean: true, Connections: true, ExternalNodes: true})
		require.NoError(t, err)
		require.NotNil(t, g)

		eg := &graph.Graph{
			Nodes: []*graph.Node{
				{
					Canonical: "im_out.tcp/443->443",
				},
				{
					Canonical: "im_out.tcp/80->80",
				},
				{
					Canonical: "aws_lb.front",
				},
				{
					Canonical: "aws_launch_template.front",
				},
				{
					Canonical: "aws_db_instance.application",
				},
				{
					Canonical: "aws_elasticache_cluster.redis",
				},
			},
			Edges: []*graph.Edge{
				{
					Source:     "im_out.tcp/443->443",
					Target:     "aws_lb.front",
					Canonicals: []string(nil),
				},
				{
					Source:     "im_out.tcp/80->80",
					Target:     "aws_lb.front",
					Canonicals: []string(nil),
				},
				{
					Source:     "aws_lb.front",
					Target:     "aws_launch_template.front",
					Canonicals: []string{"aws_security_group.lb-front", "aws_security_group.front"},
				},
				{
					Source:     "aws_launch_template.front",
					Target:     "aws_db_instance.application",
					Canonicals: []string{"aws_security_group.front", "aws_security_group.rds"},
				},
				{
					Source:     "aws_launch_template.front",
					Target:     "aws_elasticache_cluster.redis",
					Canonicals: []string{"aws_security_group.redis", "aws_security_group.front"},
				},
			},
		}

		assertEqualGraph(t, eg, g, nil)
	})
}

func TestFromHCL_FlexibleEngine(t *testing.T) {
	t.Run("SuccessSG", func(t *testing.T) {
		fs := afero.NewOsFs()

		g, err := generate.FromHCL(fs, "./testdata/flexibleengine_hcl.tf", generate.Options{Clean: true, Connections: true, ExternalNodes: true})
		require.NoError(t, err)
		require.NotNil(t, g)

		eg := &graph.Graph{
			Nodes: []*graph.Node{
				{
					Canonical: "flexibleengine_compute_instance_v2.instance_one",
				},
				{
					Canonical: "flexibleengine_compute_instance_v2.instance_two",
				},
			},
			Edges: []*graph.Edge{
				{
					Target: "flexibleengine_compute_instance_v2.instance_one",
					Source: "flexibleengine_compute_instance_v2.instance_two",
					Canonicals: []string{
						"flexibleengine_networking_port_v2.port_instance_two",
						"flexibleengine_networking_port_v2.port_instance_one",
						"flexibleengine_networking_secgroup_rule_v2.ssh_two_to_one",
						"flexibleengine_networking_secgroup_v2.secgroup_instance_two",
						"flexibleengine_networking_secgroup_v2.secgroup_instance_one",
					},
				},
			},
		}

		assertEqualGraph(t, eg, g, nil)
	})
}

func TestFromHCL_Google(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		fs := afero.NewOsFs()

		g, err := generate.FromHCL(fs, "./testdata/google_hcl.tf", generate.Options{Clean: true, Connections: true, ExternalNodes: true})
		require.NoError(t, err)
		require.NotNil(t, g)

		eg := &graph.Graph{
			Nodes: []*graph.Node{
				{
					Canonical: "google_compute_instance.inframap-tmp-two",
				},
				{
					Canonical: "google_compute_instance.inframap-tmp",
				},
			},
			Edges: []*graph.Edge{
				{
					Target: "google_compute_instance.inframap-tmp",
					Source: "google_compute_instance.inframap-tmp-two",
					Canonicals: []string{
						"google_compute_firewall.allow-ssh",
					},
				},
			},
		}

		assertEqualGraph(t, eg, g, nil)
	})
}

func TestFromHCL_Module(t *testing.T) {
	t.Run("SuccessSG", func(t *testing.T) {
		fs := afero.NewOsFs()

		g, err := generate.FromHCL(fs, "./testdata/tf-module/", generate.Options{Clean: true, Connections: true, ExternalNodes: true})
		require.NoError(t, err)
		require.NotNil(t, g)

		eg := &graph.Graph{
			Nodes: []*graph.Node{
				{
					Canonical: "im_out.tcp/443->443",
				},
				{
					Canonical: "im_out.tcp/80->80",
				},
				{
					Canonical: "aws_lb.front",
				},
				{
					Canonical: "aws_launch_template.front",
				},
				{
					Canonical: "aws_db_instance.application",
				},
				{
					Canonical: "aws_elasticache_cluster.redis",
				},
			},
			Edges: []*graph.Edge{
				{
					Source:     "im_out.tcp/443->443",
					Target:     "aws_lb.front",
					Canonicals: []string(nil),
				},
				{
					Source:     "im_out.tcp/80->80",
					Target:     "aws_lb.front",
					Canonicals: []string(nil),
				},
				{
					Source:     "aws_lb.front",
					Target:     "aws_launch_template.front",
					Canonicals: []string{"aws_security_group.lb-front", "aws_security_group.front"},
				},
				{
					Source:     "aws_launch_template.front",
					Target:     "aws_db_instance.application",
					Canonicals: []string{"aws_security_group.front", "aws_security_group.rds"},
				},
				{
					Source:     "aws_launch_template.front",
					Target:     "aws_elasticache_cluster.redis",
					Canonicals: []string{"aws_security_group.redis", "aws_security_group.front"},
				},
			},
		}

		assertEqualGraph(t, eg, g, nil)
	})
}
