package finder

type ContentGetter interface {
	 GetContent() (string, error)
}