package metadata

import "context"

type InstanceBackupsData struct {
	Enabled bool    `json:"enabled"`
	Status  *string `json:"status"`
}

type InstanceSpecsData struct {
	VCPUs    int `json:"vcpus"`
	Memory   int `json:"memory"`
	GPUs     int `json:"gpus"`
	Transfer int `json:"transfer"`
	Disk     int `json:"disk"`
}

type InstanceData struct {
	ID       int                 `json:"id"`
	Label    string              `json:"label"`
	Region   string              `json:"region"`
	Type     string              `json:"type"`
	HostUUID string              `json:"host_uuid"`
	Tags     []string            `json:"tags"`
	Specs    InstanceSpecsData   `json:"specs"`
	Backups  InstanceBackupsData `json:"backups"`
}

// GetInstance gets various information about the current instance.
func (c *Client) GetInstance(ctx context.Context) (*InstanceData, error) {
	req := c.R(ctx).SetResult(&InstanceData{})

	resp, err := coupleAPIErrors(req.Get("instance"))
	if err != nil {
		return nil, err
	}

	return resp.Result().(*InstanceData), nil
}
