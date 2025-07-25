package godata

const (
	ALLPAGES = "allpages"
	NONE     = "none"
)

func ParseInlineCountString(inlinecount string) (*GoDataInlineCountQuery, error) {
	result := GoDataInlineCountQuery(inlinecount)
	switch inlinecount {
	case ALLPAGES:
		return &result, nil
	case NONE:
		return &result, nil
	default:
		return nil, BadRequestError("Invalid inlinecount query.")
	}
}
