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
layout: layouts/post.njk
body_class: %s
critter_type:
  title: %s
  slug: %s
critter_id: %s
title: %s
permalink: %s/
---
`,
		item.Attributes.Type.Slug,
		item.Attributes.Type.Title,
		item.Attributes.Type.Slug,
		item.ID,
		title,
		permalink,
	)

	return path, []byte(content)
}
