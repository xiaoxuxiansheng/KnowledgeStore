package north

type QueryOptions struct {
	Page      int
	Limit     int
	StoreName string
}

type QueryOption func(*QueryOptions)

func WithPageLimit(page, limit int) QueryOption {
	return func(qo *QueryOptions) {
		qo.Page = page
		qo.Limit = limit
	}
}

func WithStoreName(storeName string) QueryOption {
	return func(qo *QueryOptions) {
		qo.StoreName = storeName
	}
}
