package go_ical

import "strings"


var textEscaper = strings.NewReplacer(
	`\`, `\\`,
	"\n", `\n`,
	`;`, `\;`,
	`,`, `\,`,
)

var textUnescaper = strings.NewReplacer(
	`\\`, `\`,
	`\n`, "\n",
	`\N`, "\n",
	`\;`, `;`,
	`\,`, `,`,
)

func ToText(s string) string {
	// process special characters
	return textEscaper.Replace(s)
}

func FromText(s string) string {
	return textUnescaper.Replace(s)
}
