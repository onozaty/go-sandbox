package main

import (
	"fmt"
	"strings"

	"github.com/antchfx/xmlquery"
)

func main() {

	s := `<?xml version="1.0" encoding="UTF-8" ?>
	<rss version="2.0">
	<channel>
	  <title>W3Schools Home Page</title>
	  <link>https://www.w3schools.com</link>
	  <description>Free web building tutorials</description>
	  <item id="1">
		<title>RSS Tutorial</title>
		<link name="xxx">https://www.w3schools.com/xml/xml_rss.asp</link>
		<description>New RSS tutorial on W3Schools</description>
	  </item>
	  <item id="2">
		<title>XML Tutorial</title>
		<link>https://www.w3schools.com/xml</link>
		<description>New XML tutorial on W3Schools</description>
	  </item>
	</channel>
	</rss>`

	doc, err := xmlquery.Parse(strings.NewReader(s))
	if err != nil {
		panic(err)
	}

	for i, item := range xmlquery.Find(doc, "//item") {

		id := xmlquery.FindOne(item, "@id")
		fmt.Printf("#%d id:%s\n", i, id.InnerText())

		title := xmlquery.FindOne(item, "/title")
		fmt.Printf("#%d title:%s\n", i, title.InnerText())

		name := xmlquery.FindOne(item, "/link/@name")
		if name != nil {
			fmt.Printf("#%d name:%s\n", i, name.InnerText())
		}
	}
}
