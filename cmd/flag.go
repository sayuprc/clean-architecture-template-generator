package cmd

import "github.com/spf13/cobra"

func getParamBool(cmd *cobra.Command, paramName string) (bool, error) {
	if param, err := cmd.PersistentFlags().GetBool(paramName); err != nil {
		return param, err
	} else {
		return param, nil
	}
}

func getParamString(cmd *cobra.Command, paramName string) (string, error) {
	if param, err := cmd.PersistentFlags().GetString(paramName); err != nil {
		return param, err
	} else {
		return param, nil
	}
}

func getParamStringSlice(cmd *cobra.Command, paramName string) ([]string, error) {
	if param, err := cmd.PersistentFlags().GetStringSlice(paramName); err != nil {
		return nil, err
	} else {
		return param, nil
	}
}
