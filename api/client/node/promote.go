package node

import (
	"fmt"

	"github.com/docker/docker/api/client"
	"github.com/docker/docker/cli"
	"github.com/docker/engine-api/types/swarm"
	"github.com/spf13/cobra"
)

func newPromoteCommand(dockerCli *client.DockerCli) *cobra.Command {
	return &cobra.Command{
		Use:   "promote NODE [NODE...]",
		Short: "在Swarm集群中升级一个或多个到管理者角色",
		Args:  cli.RequiresMinArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runPromote(dockerCli, args)
		},
	}
}

func runPromote(dockerCli *client.DockerCli, nodes []string) error {
	promote := func(node *swarm.Node) error {
		node.Spec.Role = swarm.NodeRoleManager
		return nil
	}
	success := func(nodeID string) {
		fmt.Fprintf(dockerCli.Out(), "成功将Swarm集群中的节点 %s 升级到管理者角色。\n", nodeID)
	}
	return updateNodes(dockerCli, nodes, promote, success)
}