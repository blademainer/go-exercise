package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

const DATA_FIELD = "datas"

var jsonStr = `{
  "datas": {
    "abc,123": {
      "code": {
        "0": 1
      }
    },
    "123": {
      "code": {
        "0": 2
      }
    },
    "111": {
      "code": {
        "0": 105
      }
    },
    "222": {
      "code": {
        "0": 20
      }
    },
    "444": {
      "code": {
        "0": 78
      }
    },
    "56677": {
      "code": {
        "0": 47
      }
    },
    "1234": {
      "code": {
        "0": 256
      }
    },
    "11": {
      "code": {
        "0": 177
      }
    },
    "222": {
      "code": {
        "0": 35
      }
    },
    "444": {
      "code": {
        "0": 3
      }
    },
    "55": {
      "code": {
        "0": 1
      }
    },
    "66": {
      "code": {
        "0": 5
      }
    },
    "77": {
      "code": {
        "0": 1
      }
    },
    "88": {
      "code": {
        "0": 16
      }
    },
    "11": {
      "code": {
        "0": 3
      }
    },
    "33": {
      "code": {
        "0": 17
      }
    },
    "44": {
      "code": {
        "0": 35
      }
    },
    "55": {
      "code": {
        "0": 11
      }
    },
    "66": {
      "code": {
        "0": 1
      }
    },
    "66": {
      "code": {
        "0": 2
      }
    },
    "123": {
      "code": {
        "0": 5
      }
    },
    "124": {
      "code": {
        "0": 5
      }
    },
    "124": {
      "code": {
        "0": 1
      }
    },
    "124": {
      "code": {
        "0": 4,
        "400": 10
      }
    },
    "111": {
      "code": {
        "0": 6,
        "400": 1
      }
    },
    "15225125": {
      "code": {
        "0": 12
      }
    },
    "1243123": {
      "code": {
        "0": 1
      }
    },
    "124": {
      "code": {
        "0": 5
      }
    },
    "124": {
      "code": {
        "0": 15
      }
    },
    "123": {
      "code": {
        "0": 39,
        "100": 100
      }
    }
  }
}`

type Results struct {
	Data    map[string]map[string]map[string]map[string]int
	Entries []Entry
}

type Entry struct {
	key   string
	value map[string]map[string]int
}

func (r *Results) ConvertToSlices() {
	datas := r.Data[DATA_FIELD]

	r.Entries = make([]Entry, 0)
	for k, v := range datas {
		e := Entry{}
		e.key = k
		e.value = v
		r.Entries = append(r.Entries, e)
	}
}

func (r *Results) Pack() map[string]map[string]map[string]map[string]int {
	m2 := make(map[string]map[string]map[string]map[string]int)

	datas := make(map[string]map[string]map[string]int)
	m2[DATA_FIELD] = datas
	for _, e := range r.Entries {
		datas[e.key] = e.value
	}
	return m2
}

func (r *Results) Len() int {
	return len(r.Entries)
}

func (r *Results) Less(i, j int) bool {
	iEntry := r.Entries[i]
	jEntry := r.Entries[j]
	irate := rate(iEntry.value)
	jrate := rate(jEntry.value)
	fmt.Println("irate: ", irate)
	fmt.Println("jrate: ", jrate)
	return irate < jrate
}

func (r *Results) Swap(i, j int) {
	r.Entries[i], r.Entries[j] = r.Entries[j], r.Entries[i]
}

func main() {
	m := &Results{}

	e := json.Unmarshal([]byte(jsonStr), &m.Data)

	if e != nil {
		panic(e)
	}
	m.ConvertToSlices()

	fmt.Println(m)

	fmt.Println("before: ", m.Entries)
	sort.Sort(m)
	fmt.Println("sorted: ", m.Entries)
	//pack := m.Pack() // nouse
	//fmt.Println("final: ", pack)
	//bytes, err := json.Marshal(pack)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("final json: ", string(bytes))
}

func rate(entries map[string]map[string]int) float64 {
	irate := 0.0
	sum := 0.0
	for _, ie := range entries {
		for k, v := range ie {
			if v <= 0 {
				continue
			}
			vf := float64(v)
			sum += vf
			if k == "0" {
				irate = (irate * sum) / sum
			} else {
				irate = ((irate * sum) + vf) / sum
			}
		}

	}
	return irate
}
