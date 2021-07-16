package models

// CompareListRequest description: A single image object and a list of templates that it will be compared to.
type CompareListRequest struct {
	SingleTemplate Image
	TemplateList []Image
}
