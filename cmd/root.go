/**
 * SPDX-License-Identifier: Apache-2.0
 * Copyright 2020 vorteil.io Pty Ltd
 */

package cmd

import (
	"fmt"
	"os"

	"code.vorteil.io/vorteil/os/tcpdump/pkg/tcpdump"
	"github.com/spf13/cobra"
)

var cfgFile string
var targetDeviceName string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tcpdump",
	Short: "A tcpdump Clone",
	Long:  `A tcpdump Clone`,

	Run: func(cmd *cobra.Command, args []string) {
		if targetDeviceName == "" {
			os.Stderr.WriteString(fmt.Sprintf("Flag error [--device=\"%s\"], device flag is missing or set to \"\"\n", targetDeviceName))
			os.Exit(1)
		}

		newCaptureManager := tcpdump.NewPacketCapture(tcpdump.DefaultSnapshotLen, tcpdump.DefaultPromiscuousMode, tcpdump.DefaultTimeout)
		if err := newCaptureManager.SetDevice(targetDeviceName); err != nil {
			os.Stderr.WriteString(fmt.Sprintf("Could not set target interface to device named \"%s\", error=%v\n", targetDeviceName, err))
			os.Exit(1)
		}

		if err := newCaptureManager.SetFilter(tcpdump.DefaultBPFFilter); err != nil {
			os.Stderr.WriteString(fmt.Sprintf("Could not set bpf filter on device named \"%s\", error=%v\n", targetDeviceName, err))
			os.Exit(2)
		}

		err := newCaptureManager.StartCapturing()
		if err != nil {
			os.Stderr.WriteString(fmt.Sprintf("packet capturing on device named \"%s\" failed, error=%v\n", targetDeviceName, err))
			os.Exit(3)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tcpdump.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringVarP(&targetDeviceName, "device", "d", "", "Target Device Name to Capture Packets")
}
