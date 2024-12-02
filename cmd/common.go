package cmd

import (
	"log"
	"os"

	"github.com/Masterminds/semver"
	"github.com/spf13/cobra"
)

type comparisonType int

const (
	equal comparisonType = iota
	notequal
	greaterthan
	lessthan
	constraint
)

const shouldbeunreachable = false

func versionCompare(t comparisonType, v1, v2 string) bool {
	v1V, err := semver.NewVersion(v1)
	if err != nil {
		log.Fatalf("invalid semver: %v", v1)
	}

	v2V, err := semver.NewVersion(v2)
	if err != nil {
		log.Fatalf("invalid semver: %v", v2)
	}

	switch t {
	case equal:
		return v1V.Equal(v2V)
	case notequal:
		return !v1V.Equal(v2V)
	case greaterthan:
		return v1V.GreaterThan(v2V)
	case lessthan:
		return v1V.LessThan(v2V)
	}

	return shouldbeunreachable
}

func constraintCheck(v, c string) bool {
	cC, err := semver.NewConstraint(c)
	if err != nil {
		log.Fatalf("invalid constraint: %v", c)
	}

	vV, err := semver.NewVersion(v)
	if err != nil {
		log.Fatalf("invalid version: %v", v)
	}

	return cC.Check(vV)
}

func newCommand(operator comparisonType) *cobra.Command {
	var name string
	var example string
	var aliases []string

	switch operator {
	case equal:
		name = "eq"
		aliases = []string{"equal"}
		example = "eq v1 v2"
	case notequal:
		name = "neq"
		aliases = []string{"not-equal", "notequal"}
		example = "neq v1 v2"
	case greaterthan:
		name = "gt"
		aliases = []string{"greaterthan", "greater-than"}
		example = "gt v1 v2"
	case lessthan:
		name = "lt"
		aliases = []string{"lessthan", "less-than"}
		example = "lt v1 v2"
	case constraint:
		name = "check"
		example = "check version constraint"
	}

	cmd := &cobra.Command{
		Use: name,
		Aliases: aliases,
		Example: example,
		Short:   example,
		Args: cobra.ExactArgs(2),
		Run: func(_cmd *cobra.Command, args []string) {
			if operator == constraint {
				if constraintCheck(args[0], args[1]) {
					os.Exit(0)
				}

				os.Exit(1)
			}

			if versionCompare(operator, args[0], args[1]) {
				os.Exit(0)
			} else {
				os.Exit(1)
			}
		},
	}

	return cmd
}
