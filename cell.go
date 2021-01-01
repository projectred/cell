package cell

var defaultCell = Cell{make(map[string]func() interface{}), nil}
var Regist = defaultCell.Regist
var Spilt = defaultCell.Split
var RegistDefault = defaultCell.RegistDefault

type Cell struct {
	kvs      map[string]func() interface{}
	defaultF func() interface{}
}

func New() *Cell { return &Cell{make(map[string]func() interface{})} }

func (c *Cell) Regist(keys []string, f func() interface{}) {
	for _, key := range keys {
		c.kvs[key] = f
	}
}

func (c *Cell) RegistDefault(f func() interface{}) { c.defaultF = f }

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
	if c.defaultF != nil {
		return c.defaultF()
	}
	return nil
}
