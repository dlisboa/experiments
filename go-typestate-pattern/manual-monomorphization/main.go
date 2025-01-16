package main

func main() {
	draft := &DraftArticle{content: ""}
	draft.Edit("Humans (homo sapiens) are the most common primate...")

	// uncomment last line to test invariant:
	//   Only a Published Article is linkable (can call Link())
	// compiler:
	//   draft.Link undefined (type *DraftArticle has no field or method Link)
	// draft.Link()
	//
	// uncomment last line to test invariant:
	//   One cannot call Unpublish() on a Draft Article
	// compiler:
	//   draft.Unpublish undefined (type *DraftArticle has no field or method Unpublish)
	// draft.Unpublish()

	published := draft.Publish()
	published.Link()

	// uncomment last line to test invariant:
	//   Only a Draft Article is editable (can call Edit())
	// compiler:
	//   published.Edit undefined (type *PublishedArticle has no field or method Edit
	// published.Edit()

	// uncomment last line to test invariant:
	//   One cannot call Publish() on a Published Article
	// compiler:
	//   published.Publish undefined (type *PublishedArticle has no field or method Publish)
	// published.Publish()

	// test:
	// 	 Use After Publish:
	// compiler:
	// 	 allows use of draft after publish (does not "consume")
	draft.Edit("we shouldn't be able to edit after it's been published")
}
