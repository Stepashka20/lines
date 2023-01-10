package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func lineCounter(filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}
	for {
		c, err := file.Read(buf)
		count += bytes.Count(buf[:c], lineSep)
		switch {
		case err == io.EOF:
			return count, nil
		case err != nil:
			return count, err
		}
	}
}

var rootCmd = &cobra.Command{
	Use:   "lines",
	Short: "How many lines of code are in this project?",
	Long:  `This is a simple tool to count the number of lines of code in a project.`,
	Run: func(cmd *cobra.Command, args []string) {
		dir, err := os.Getwd()
		if err != nil {
			fmt.Println("Error getting current directory: ", err)
			return
		}
		fmt.Println("Current directory: ", dir)
		display, _ := cmd.Flags().GetBool("info")
		ignoredFiles, _ := cmd.Flags().GetStringSlice("exclude")
		if display {
			fmt.Println("Exclude: ", ignoredFiles)
		}
		totalLines := 0
		filepath.Walk(".",
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if !info.IsDir() {
					filename := filepath.Base(path)
					for _, e := range ignoredFiles {
						if e == filename {

							if display {
								fmt.Println("Ignoring file: ", filename)
							}
							return nil
						}
					}
					count, err := lineCounter(path)
					if err != nil {
						fmt.Println("Error counting lines: ", err)
						return err
					}
					if display {
						fmt.Println("File: ", path, " has ", count, " lines.")
					}
					totalLines += count
				}

				return nil
			})
		fmt.Println("Total lines: ", totalLines)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().BoolP("info", "i", false, "Display information about each file while scanning.")
	rootCmd.Flags().StringSliceP("exclude", "e", []string{}, "Remove files from counting. Files separated by comma.")
	// rootCmd.Flags().BoolP("exclude", "e", false, "Remove files from counting")
}
