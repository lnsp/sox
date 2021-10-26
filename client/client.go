package main

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/spf13/cobra"
	"github.com/valar/virtm/api"
	"google.golang.org/grpc"
)

var endpoint string

var insecure bool

var timeout time.Duration

var rootCmd = cobra.Command{
	Use:   "virtm-cli",
	Short: "CLI for interacting with a VirtM instance",
}

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
		// list ssh keys
		resp, err := client.ListImages(ctx, &api.ListImagesRequest{})
		if err != nil {
			return err
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
		// print out ssh keys in table format
		tw := tabwriter.NewWriter(os.Stdout, 1, 4, 1, ' ', 0)
		defer tw.Flush()

		fmt.Fprintf(tw, "ID\tNAME\tFOOTPRINT\n")
		for _, key := range resp.Keys {
			fmt.Fprintf(tw, "%s\t%s\t%.32s\n", key.Id, key.Name, key.Pubkey)
		}
		return nil
	},
}

var machinesCmd = cobra.Command{
	Use:   "machines",
	Short: "Manage machine instances",
}

var machinesCreateCpu int64
var machinesCreateMemory int64
var machinesCreateDisk int64
var machinesCreateImage string
var machinesCreateSSHKey string

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
			ImageId:  machinesCreateImage,
			SshKeyId: machinesCreateSSHKey,
		})
		if err != nil {
			return err
		}
		fmt.Fprintln(os.Stdout, resp.Id)
		return nil
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&endpoint, "endpoint", "p", "localhost:9876", "VirtM endpoint address")
	rootCmd.PersistentFlags().BoolVar(&insecure, "insecure", true, "Connect to insecure endpoint")
	rootCmd.PersistentFlags().DurationVar(&timeout, "timeout", time.Minute, "Client connection timeout")
	rootCmd.AddCommand(&imagesCmd)
	rootCmd.AddCommand(&sshKeysCmd)
	rootCmd.AddCommand(&machinesCmd)
	machinesCmd.AddCommand(&machinesCreateCmd)
	machinesCreateCmd.Flags().StringVarP(&machinesCreateImage, "image", "i", "", "Operating system image")
	machinesCreateCmd.Flags().StringVarP(&machinesCreateSSHKey, "ssh-key", "k", "", "SSH key for login")
	machinesCreateCmd.Flags().Int64Var(&machinesCreateCpu, "cpu", 2, "Number of vCPUs")
	machinesCreateCmd.Flags().Int64Var(&machinesCreateDisk, "disk", 10000, "Disk size in MB")
	machinesCreateCmd.Flags().Int64Var(&machinesCreateMemory, "memory", 2, "Memory size in MB")
	machinesCreateCmd.MarkFlagRequired("image")
	machinesCreateCmd.MarkFlagRequired("ssh-key")
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
