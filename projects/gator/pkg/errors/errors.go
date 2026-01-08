package errors

import "fmt"

var (
	ErrNotLoggedIn = fmt.Errorf("not logged in. run %s", UsageLogin)

	ErrInvalidLoginArgsLength    = fmt.Errorf("usage: %s", UsageLogin)
	ErrInvalidRegisterArgsLength = fmt.Errorf("usage: %s", UsageRegister)
	ErrInvalidAddFeedArgsLength  = fmt.Errorf("usage: %s", UsageAddFeed)
	ErrInvalidFollowArgsLength   = fmt.Errorf("usage: %s", UsageFollow)
	ErrInvalidUnfollowArgsLength = fmt.Errorf("usage: %s", UsageUnfollow)
	ErrInvalidAggArgsLength      = fmt.Errorf("usage: %s", UsageAgg)
)
