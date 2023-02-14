package body_types

type ImportExportRequest struct {
	TarFile   string `json:"tar"`
	Target    string `json:"target"`
	IgnoreSSL bool   `json:"ignore_ssl"`
}

type PackageInfo struct {
	ImageName string `json:"image_name_tag"`
	TarPath   string `json:"tar_path"`
}
