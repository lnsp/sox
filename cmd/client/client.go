package main

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/lnsp/sox/api"
	"github.com/lnsp/sox/meta"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var endpoint string

var insecure bool

var timeout time.Duration

var rootCmd = cobra.Command{
	Use:     "virtm-cli",
	Short:   "CLI for interacting with a VirtM instance",
	Version: meta.Version,
}

var listIdsOnly bool

var imagesCmd = cobra.Command{
	Use:          "images",
	Short:        "Manage machine images",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := connect()
		if err != nil {
			return err
		}
		// create context
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		// list images
		resp, err := client.ListImages(ctx, &api.ListImagesRequest{})
		if err != nil {
			return err
		}
		if listIdsOnly {
			for i := range resp.Images {
				fmt.Println(resp.Images[i].Id)
			}
			return nil
		}
		// print out ssh keys in table format
		tw := tabwriter.NewWriter(os.Stdout, 1, 4, 1, ' ', 0)
		defer tw.Flush()

		fmt.Fprintf(tw, "ID\tNAME\tOPERATING SYSTEM\n")
		for _, img := range resp.Images {
			fmt.Fprintf(tw, "%s\t%s\t%s\n", img.Id, img.Name, img.System.String())
		}
		return nil
	},
}

var sshKeysCmd = cobra.Command{
	Use:          "ssh-keys",
	Short:        "Manage SSH keys",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := connect()
		if err != nil {
			return err
		}
		// create context
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		// list ssh keys
		resp, err := client.ListSSHKeys(ctx, &api.ListSSHKeysRequest{})
		if err != nil {
			return err
		}
		if listIdsOnly {
			for i := range resp.Keys {
				fmt.Println(resp.Keys[i].Id)
			}
			return nil
		}
		// print out ssh keys in table format
		tw := tabwriter.NewWriter(os.Stdout, 1, 4, 1, ' ', 0)
		defer tw.Flush()

		fmt.Fprintf(tw, "ID\tNAME\tFINGERPRINT\n")
		for _, key := range resp.Keys {
			fmt.Fprintf(tw, "%s\t%s\t%.32s\n", key.Id, key.Name, key.Pubkey)
		}
		return nil
	},
}

var machinesCmd = cobra.Command{
	Use:   "machines",
	Short: "Manage machine instances",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := connect()
		if err != nil {
			return err
		}
		// create context
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		// list machines
		resp, err := client.ListMachines(ctx, &api.ListMachinesRequest{})
		if err != nil {
			return err
		}
		if listIdsOnly {
			for i := range resp.Machines {
				fmt.Println(resp.Machines[i].Id)
			}
			return nil
		}
		// print out machines in table format
		tw := tabwriter.NewWriter(os.Stdout, 1, 4, 1, ' ', 0)
		defer tw.Flush()

		fmt.Fprintf(tw, "ID\tNAME\tSTATUS\n")
		for _, machine := range resp.Machines {
			fmt.Fprintf(tw, "%s\t%s\t%s\n", machine.Id, machine.Name, machine.Status)
		}
		return nil
	},
}

var machinesInspectCmd = cobra.Command{
	Use:   "inspect [name]",
	Short: "Display detailed information on the machine instance",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := connect()
		if err != nil {
			return err
		}
		// create context
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		// get details
		resp, err := client.GetMachineDetails(ctx, &api.GetMachineDetailsRequest{
			Id: args[0],
		})
		if err != nil {
			return err
		}
		// print out machine details
		tw := tabwriter.NewWriter(os.Stdout, 1, 4, 1, ' ', 0)
		defer tw.Flush()

		fmt.Fprintf(tw, "%s\t%s\n", "ID", resp.Machine.Id)
		fmt.Fprintf(tw, "%s\t%s\n", "Name", resp.Machine.Name)
		fmt.Fprintf(tw, "%s\t%s\n", "Status", resp.Machine.Status)
		fmt.Fprintf(tw, "%s\t%+v\n", "Specs", resp.Machine.Specs)
		for _, net := range resp.Machine.Networks {
			fmt.Fprintf(tw, "%s\t%+v\n", "Network", net)
		}
		return nil
	},
}

var machinesCreateCpu int64
var machinesCreateMemory int64
var machinesCreateDisk int64
var machinesCreateImage string
var machinesCreateSSHKeys []string
var machinesCreateNetworks []string
var machinesCreateUser string

var machinesCreateCmd = cobra.Command{
	Use:   "create [name]",
	Short: "Create new virtual machine instance",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := connect()
		if err != nil {
			return err
		}
		// create context
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		// submit request
		resp, err := client.CreateMachine(ctx, &api.CreateMachineRequest{
			Name: args[0],
			Specs: &api.Machine_Specs{
				Cpus:   machinesCreateCpu,
				Memory: machinesCreateMemory,
				Disk:   machinesCreateDisk,
			},
			ImageId:    machinesCreateImage,
			SshKeyIds:  machinesCreateSSHKeys,
			NetworkIds: machinesCreateNetworks,
			User:       machinesCreateUser,
		})
		if err != nil {
			return err
		}
		fmt.Fprintln(os.Stdout, resp.Id)
		return nil
	},
}

var machinesDeleteCmd = cobra.Command{
	Use:   "delete [id | name]",
	Short: "Delete an existing virtual machine instance",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := connect()
		if err != nil {
			return err
		}
		// create context
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		// submit request
		if _, err := client.DeleteMachine(ctx, &api.DeleteMachineRequest{
			Id: args[0],
		}); err != nil {
			return err
		}
		return nil
	},
}

var machinesStartCmd = cobra.Command{
	Use:   "start [id]",
	Short: "Boot a powered-off machine",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := connect()
		if err != nil {
			return err
		}
		// create context
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		// submit request
		resp, err := client.TriggerMachine(ctx, &api.TriggerMachineRequest{
			Id:    args[0],
			Event: api.TriggerMachineRequest_POWERON,
		})
		if err != nil {
			return err
		}
		fmt.Println(resp.Status)
		return nil
	},
}

var machinesStopCmd = cobra.Command{
	Use:   "stop [id]",
	Short: "Power off a running machine",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := connect()
		if err != nil {
			return err
		}
		// create context
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		// submit request
		resp, err := client.TriggerMachine(ctx, &api.TriggerMachineRequest{
			Id:    args[0],
			Event: api.TriggerMachineRequest_POWEROFF,
		})
		if err != nil {
			return err
		}
		fmt.Println(resp.Status)
		return nil
	},
}

var machinesRebootCmd = cobra.Command{
	Use:   "reboot [id]",
	Short: "Reboot a running machine",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := connect()
		if err != nil {
			return err
		}
		// create context
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		// submit request
		resp, err := client.TriggerMachine(ctx, &api.TriggerMachineRequest{
			Id:    args[0],
			Event: api.TriggerMachineRequest_REBOOT,
		})
		if err != nil {
			return err
		}
		fmt.Println(resp.Status)
		return nil
	},
}

var networksCmd = cobra.Command{
	Use:          "networks",
	Short:        "Manage virtual networks",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := connect()
		if err != nil {
			return err
		}
		// create context
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		// list networks
		resp, err := client.ListNetworks(ctx, &api.ListNetworksRequest{})
		if err != nil {
			return err
		}
		if listIdsOnly {
			for i := range resp.Networks {
				fmt.Println(resp.Networks[i].Id)
			}
			return nil
		}
		// print out machine details
		tw := tabwriter.NewWriter(os.Stdout, 1, 4, 1, ' ', 0)
		defer tw.Flush()

		fmt.Fprintf(tw, "ID\tNAME\tSUBNET\tGATEWAY\n")
		for _, net := range resp.Networks {
			fmt.Fprintf(
				tw,
				"%s\t%s\t%s\t%s\n",
				net.Id, net.Name, net.IpV4.Subnet, net.IpV4.Gateway,
			)
		}
		return nil
	},
}

var networksCreateName string
var networksCreateBridgeId uint32
var networksCreateIPV4Subnet string
var networksCreateIPV4Gateway string
var networksCreateIPV6Subnet string
var networksCreateIPV6Gateway string

var networksCreateCmd = cobra.Command{
	Use:          "create",
	Short:        "Create a virtual network",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := connect()
		if err != nil {
			return err
		}
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		resp, err := client.CreateNetwork(ctx, &api.CreateNetworkRequest{
			Name:     networksCreateName,
			BridgeId: networksCreateBridgeId,
			IpV4: &api.IpNetwork{
				Subnet:  networksCreateIPV4Subnet,
				Gateway: networksCreateIPV4Gateway,
			},
			IpV6: &api.IpNetwork{
				Subnet:  networksCreateIPV6Subnet,
				Gateway: networksCreateIPV6Gateway,
			},
		})
		if err != nil {
			return err
		}
		fmt.Println(resp.Id)
		return nil
	},
}

var activityCmd = cobra.Command{
	Use:          "activity",
	Short:        "List recent activities",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := connect()
		if err != nil {
			return err
		}
		// create context
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		// list activities
		resp, err := client.ListActivities(ctx, &api.ListActivitiesRequest{})
		if err != nil {
			return err
		}
		// print out activities
		tw := tabwriter.NewWriter(os.Stdout, 1, 4, 1, ' ', 0)
		defer tw.Flush()

		fmt.Fprintf(tw, "TIME\tACTIVITY\tSUBJECT\n")
		for _, act := range resp.Activities {
			fmt.Fprintf(
				tw,
				"%s\t%s\t%s\n",
				humanize.Time(act.Timestamp.AsTime()),
				act.Type, act.Subject,
			)
		}

		return nil
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&endpoint, "endpoint", "p", "localhost:9876", "VirtM endpoint address")
	rootCmd.PersistentFlags().BoolVar(&insecure, "insecure", true, "Connect to insecure endpoint")
	rootCmd.PersistentFlags().DurationVar(&timeout, "timeout", time.Minute, "Client connection timeout")
	rootCmd.AddCommand(&imagesCmd)
	imagesCmd.Flags().BoolVarP(&listIdsOnly, "ids-only", "1", false, "Only display IDs")
	rootCmd.AddCommand(&sshKeysCmd)
	sshKeysCmd.Flags().BoolVarP(&listIdsOnly, "ids-only", "1", false, "Only display IDs")
	rootCmd.AddCommand(&machinesCmd)
	machinesCmd.Flags().BoolVarP(&listIdsOnly, "ids-only", "1", false, "Only display IDs")
	rootCmd.AddCommand(&networksCmd)
	networksCmd.Flags().BoolVarP(&listIdsOnly, "ids-only", "1", false, "Only display IDs")
	networksCmd.AddCommand(&networksCreateCmd)
	networksCreateCmd.Flags().StringVarP(&networksCreateName, "name", "n", "", "Unique name of the network")
	networksCreateCmd.Flags().Uint32VarP(&networksCreateBridgeId, "bridge-id", "b", 100, "VXLAN ID to attach to")
	networksCreateCmd.Flags().StringVar(&networksCreateIPV4Subnet, "ipv4-subnet", "10.0.0.0/24", "IPv4 subnet")
	networksCreateCmd.Flags().StringVar(&networksCreateIPV4Gateway, "ipv4-gateway", "", "IPv4 gateway")
	networksCreateCmd.Flags().StringVar(&networksCreateIPV6Subnet, "ipv6-subnet", "", "IPv6 subnet")
	networksCreateCmd.Flags().StringVar(&networksCreateIPV6Gateway, "ipv6-gateway", "", "IPv6 gateway")
	rootCmd.AddCommand(&activityCmd)
	machinesCmd.AddCommand(&machinesCreateCmd)
	machinesCmd.AddCommand(&machinesInspectCmd)
	machinesCmd.AddCommand(&machinesDeleteCmd)
	machinesCmd.AddCommand(&machinesStartCmd)
	machinesCmd.AddCommand(&machinesStopCmd)
	machinesCmd.AddCommand(&machinesRebootCmd)
	machinesCreateCmd.Flags().StringVarP(&machinesCreateImage, "image", "i", "", "Operating system image")
	machinesCreateCmd.Flags().StringArrayVarP(&machinesCreateSSHKeys, "ssh-keys", "k", nil, "SSH keys for login")
	machinesCreateCmd.Flags().StringArrayVarP(&machinesCreateNetworks, "networks", "n", nil, "Network to connect to")
	machinesCreateCmd.Flags().Int64Var(&machinesCreateCpu, "cpu", 2, "Number of vCPUs")
	machinesCreateCmd.Flags().Int64Var(&machinesCreateDisk, "disk", 10000, "Disk size in GB")
	machinesCreateCmd.Flags().Int64Var(&machinesCreateMemory, "memory", 2000, "Memory size in MB")
	machinesCreateCmd.Flags().StringVarP(&machinesCreateUser, "user", "u", "ken", "User account to be created")
	machinesCreateCmd.MarkFlagRequired("image")
	machinesCreateCmd.MarkFlagRequired("ssh-key")
	machinesCreateCmd.MarkFlagRequired("networks")
}

func connect() (api.VirtMClient, error) {
	var grpcOpts []grpc.DialOption
	if insecure {
		grpcOpts = append(grpcOpts, grpc.WithInsecure())
	}
	grpcClient, err := grpc.Dial(endpoint, grpcOpts...)
	if err != nil {
		return nil, fmt.Errorf("dial endpoint: %w", err)
	}
	return api.NewVirtMClient(grpcClient), nil
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
