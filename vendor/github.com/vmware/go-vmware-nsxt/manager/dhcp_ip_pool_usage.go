/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package manager

type DhcpIpPoolUsage struct {

	// allocated number. COULD BE INACCURATE, REFERENCE ONLY.
	AllocatedNumber int64 `json:"allocated_number"`

	// allocated percentage. COULD BE INACCURATE, REFERENCE ONLY.
	AllocatedPercentage int64 `json:"allocated_percentage"`

	// uuid of dhcp ip pool
	DhcpIpPoolId string `json:"dhcp_ip_pool_id"`

	// pool size
	PoolSize int64 `json:"pool_size"`
}