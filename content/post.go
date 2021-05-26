package content

import (
	"fmt"
	"markdown/targetstructures"
)

func Post(item targetstructures.Item, all targetstructures.Output, dir string) (string, []byte) {
	path := dir + item.ID + ".md"
	permalink := item.Attributes.URIS.URL
	title := item.Attributes.Titles.Safe
	content := fmt.Sprintf(`---
item:
  type:
    title: %s
    slug: %s
  id: %s
title: %s
permalink: %s
layout: page  
---
`,
		item.Attributes.Type.Title,
		item.Attributes.Type.Slug,
		item.ID,
		title,
		permalink,
	)

	return path, []byte(content)
}
