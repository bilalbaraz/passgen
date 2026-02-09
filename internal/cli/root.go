package cli

import (
	"errors"
	"fmt"
	"strings"

	"github.com/skip2/go-qrcode"
	"github.com/spf13/cobra"

	"github.com/bilalbaraz/passgen/internal/passgen"
)

var cfg passgen.Config

var rootCmd = &cobra.Command{
	Use:   "passgen",
	Short: "Generate secure passwords",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Print(passgen.Banner)

		applyDefaultCharset(cmd, &cfg)
		if err := cfg.Validate(); err != nil {
			return err
		}

		excludeSet := cfg.ExcludeSet()
		charset := passgen.BuildCharset(cfg.Lower, cfg.Upper, cfg.Digits, cfg.Symbols, excludeSet)
		if len(charset) == 0 {
			return errors.New("character set is empty after exclusions")
		}

		passwords := make([]string, 0, cfg.Count)
		for i := 0; i < cfg.Count; i++ {
			pwd, err := passgen.GeneratePassword(cfg.Length, charset)
			if err != nil {
				return err
			}
			passwords = append(passwords, pwd)
			fmt.Println(pwd)
		}

		if cfg.QR {
			qr, err := qrcode.New(passwords[0], qrcode.Medium)
			if err != nil {
				return err
			}
			fmt.Println(qr.ToSmallString(false))
		}

		if cfg.Copy {
			if err := passgen.CopyToClipboard(strings.Join(passwords, "\n")); err != nil {
				return err
			}
		}

		return nil
	},
}

func applyDefaultCharset(cmd *cobra.Command, cfg *passgen.Config) {
	if cmd.Flags().Changed("lower") || cmd.Flags().Changed("upper") || cmd.Flags().Changed("digits") || cmd.Flags().Changed("symbols") {
		return
	}
	cfg.Lower = true
	cfg.Upper = true
	cfg.Digits = true
	cfg.Symbols = true
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.SilenceUsage = true

	rootCmd.Flags().IntVarP(&cfg.Length, "len", "l", 16, "password length")
	rootCmd.Flags().BoolVar(&cfg.Lower, "lower", false, "include lowercase letters")
	rootCmd.Flags().BoolVar(&cfg.Upper, "upper", false, "include uppercase letters")
	rootCmd.Flags().BoolVar(&cfg.Digits, "digits", false, "include digits")
	rootCmd.Flags().BoolVar(&cfg.Symbols, "symbols", false, "include symbols")
	rootCmd.Flags().IntVarP(&cfg.Count, "count", "c", 1, "number of passwords to generate")
	rootCmd.Flags().StringVarP(&cfg.Exclude, "exclude", "x", "", "exclude specific characters")
	rootCmd.Flags().BoolVar(&cfg.NoAmbiguous, "no-ambiguous", false, "exclude ambiguous characters like 0 O 1 l I")
	rootCmd.Flags().BoolVarP(&cfg.Copy, "copy", "p", false, "copy generated passwords to clipboard")
	rootCmd.Flags().BoolVarP(&cfg.QR, "qr", "q", false, "render first password as QR code in terminal")
}
