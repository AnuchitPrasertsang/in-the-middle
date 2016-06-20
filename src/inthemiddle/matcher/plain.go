package matcher

import (
	httper "inthemiddle/httper"
)

type PlainTextRequestMatcher struct {
}

func (m PlainTextRequestMatcher) Match(req *httper.Request, matcher *MatchOption) bool {
    result := (req.URL() == matcher.Pattern)
	return result
}
