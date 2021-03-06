// Copyright 2015 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package admin

import (
	api "github.com/gogs/go-gogs-client"
	org2 "gogs.io/gogs/internal/route/api/v1/org"
	user2 "gogs.io/gogs/internal/route/api/v1/user"

	"gogs.io/gogs/internal/context"
)

func CreateOrg(c *context.APIContext, form api.CreateOrgOption) {
	org2.CreateOrgForUser(c, form, user2.GetUserByParams(c))
}
