// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!

	if len(rhyme) == 0 {
		return []string{}
	}

	var output []string

	for i, val := range rhyme {
		if i < len(rhyme)-1 {
			output = append(output, "For want of a "+val+" the "+string(rhyme[i+1])+" was lost.")
		} else {
			output = append(output, "And all for the want of a "+string(rhyme[0])+".")
		}
	}

	return output
}
