package cell

var defaultCell = Cell{make(map[string][]interface{})}
var Regist = defaultCell.Regist
var Spilt = defaultCell.Split

type Cell struct {
	kvs map[string][]interface{}
}

func (c *Cell) Regist(key string, main interface{}) {
	c.kvs[key] = []interface{}{main}
}

type SplitOptions struct {
	Ks []string
	Vs []interface{}
}

type FillF interface {
	Fill(k string, v interface{})
}

func (c *Cell) Split(key string, o *SplitOptions) interface{} {
	if source, ok := c.kvs[key]; ok {
		var dst [1]interface{}
		copy(dst[:], source)
		if o != nil {
			if f, ok := dst[0].(FillF); ok {
				for i, k := range o.Ks {
					f.Fill(k, o.Vs[i])
				}
			}
		}
		return dst[0]
	}
	return nil
}
