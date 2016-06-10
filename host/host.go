package host

import (
	"encoding/xml"
  "fmt"
	"github.com/megamsys/opennebula-go/api"
)

const (
  SHOW = "show"
)
type HQuery struct {
	HostID  string
	T      *api.Rpc
}

type Host struct {
  Id  int         `xml:ID`
  HostName string     `xml:NAME`
	Temp    *Template `xml:"TEMPLATE"`
}

type Template struct {
  VMs []*VM        `xml:"VM"`
  Wildvms string `xml:"WILDS"`
}

type VM struct {
	Id   int          `xml:"ID"`
	Deploy_id  int    `xml:"DEPLOY_ID"`
	VMName string     `xml:"VM_NAME"`
}

// Given a name, this function will return the VM
func (v *HQuery) GetVMs(a int) ([]*VM, error) {
  args := []interface{}{v.T.Key,a}
  fmt.Println("Args for action :\n",args)
	HostVMs, err := v.T.Call(api.ONE_HOST_INFO, args)
	if err != nil {
    fmt.Println("*****************host***err***************",err)
		return nil, err
	}

	xmlVM := Host{}
	assert, _ := HostVMs[1].(string)
	if err = xml.Unmarshal([]byte(assert), &xmlVM); err != nil {
		return nil, err
	}

  fmt.Println("*****************host*****xmlvm***************")
  fmt.Printf("%#v",xmlVM.Temp)
	var matchedVM = make([]*VM, 2)

	return matchedVM, nil

}
