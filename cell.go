package cell

var defaultCell = Cell{make(map[string]func() interface{})}
var Regist = defaultCell.Regist
var Spilt = defaultCell.Split

type Cell struct {
	kvs map[string]func() interface{}
}

func (c *Cell) Regist(key string, f func() interface{}) {
	c.kvs[key] = f
}

type SplitOptions struct {
	Ks []string
	Vs []interface{}
}

type FillF interface {
	Fill(k string, v interface{})
}

func (c *Cell) Split(key string, o *SplitOptions) interface{} {
	if f, ok := c.kvs[key]; ok {
		dst := f()
		if o != nil {
			if f, ok := dst.(FillF); ok {
				for i, k := range o.Ks {
					f.Fill(k, o.Vs[i])
				}
			}
		}
		return dst
	}
	return nil
}
