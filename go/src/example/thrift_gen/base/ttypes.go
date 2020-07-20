// Autogenerated by Thrift Compiler (0.9.2)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package base

import (
	"bytes"
	"fmt"

	"code.byted.org/gopkg/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var GoUnusedProtection__ int

type TrafficEnv struct {
	Open bool   `thrift:"Open,1" json:"Open"`
	Env  string `thrift:"Env,2" json:"Env"`
}

func NewTrafficEnv() *TrafficEnv {
	return &TrafficEnv{}
}

func (p *TrafficEnv) GetOpen() bool {
	return p.Open
}

func (p *TrafficEnv) GetEnv() string {
	return p.Env
}
func (p *TrafficEnv) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *TrafficEnv) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		p.Open = v
	}
	return nil
}

func (p *TrafficEnv) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 2: %s", err)
	} else {
		p.Env = v
	}
	return nil
}

func (p *TrafficEnv) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("TrafficEnv"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if p != nil {
		if err := p.writeField1(oprot); err != nil {
			return err
		}
		if err := p.writeField2(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *TrafficEnv) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("Open", thrift.BOOL, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:Open: %s", p, err)
	}
	if err := oprot.WriteBool(bool(p.Open)); err != nil {
		return fmt.Errorf("%T.Open (1) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:Open: %s", p, err)
	}
	return err
}

func (p *TrafficEnv) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("Env", thrift.STRING, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:Env: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Env)); err != nil {
		return fmt.Errorf("%T.Env (2) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:Env: %s", p, err)
	}
	return err
}

func (p *TrafficEnv) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TrafficEnv(%+v)", *p)
}

type Base struct {
	LogID      string            `thrift:"LogID,1" json:"LogID"`
	Caller     string            `thrift:"Caller,2" json:"Caller"`
	Addr       string            `thrift:"Addr,3" json:"Addr"`
	Client     string            `thrift:"Client,4" json:"Client"`
	TrafficEnv *TrafficEnv       `thrift:"TrafficEnv,5" json:"TrafficEnv"`
	Extra      map[string]string `thrift:"Extra,6" json:"Extra"`
}

func NewBase() *Base {
	return &Base{}
}

func (p *Base) GetLogID() string {
	return p.LogID
}

func (p *Base) GetCaller() string {
	return p.Caller
}

func (p *Base) GetAddr() string {
	return p.Addr
}

func (p *Base) GetClient() string {
	return p.Client
}

var Base_TrafficEnv_DEFAULT *TrafficEnv

func (p *Base) GetTrafficEnv() *TrafficEnv {
	if !p.IsSetTrafficEnv() {
		return Base_TrafficEnv_DEFAULT
	}
	return p.TrafficEnv
}

var Base_Extra_DEFAULT map[string]string

func (p *Base) GetExtra() map[string]string {
	return p.Extra
}
func (p *Base) IsSetTrafficEnv() bool {
	return p.TrafficEnv != nil
}

func (p *Base) IsSetExtra() bool {
	return p.Extra != nil
}

func (p *Base) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.ReadField3(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.ReadField4(iprot); err != nil {
				return err
			}
		case 5:
			if err := p.ReadField5(iprot); err != nil {
				return err
			}
		case 6:
			if err := p.ReadField6(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *Base) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		p.LogID = v
	}
	return nil
}

func (p *Base) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 2: %s", err)
	} else {
		p.Caller = v
	}
	return nil
}

func (p *Base) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 3: %s", err)
	} else {
		p.Addr = v
	}
	return nil
}

func (p *Base) ReadField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 4: %s", err)
	} else {
		p.Client = v
	}
	return nil
}

func (p *Base) ReadField5(iprot thrift.TProtocol) error {
	p.TrafficEnv = &TrafficEnv{}
	if err := p.TrafficEnv.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.TrafficEnv, err)
	}
	return nil
}

func (p *Base) ReadField6(iprot thrift.TProtocol) error {
	_, _, size, err := iprot.ReadMapBegin()
	if err != nil {
		return fmt.Errorf("error reading map begin: %s", err)
	}
	tMap := make(map[string]string, size)
	p.Extra = tMap
	for i := 0; i < size; i++ {
		var _key0 string
		if v, err := iprot.ReadString(); err != nil {
			return fmt.Errorf("error reading field 0: %s", err)
		} else {
			_key0 = v
		}
		var _val1 string
		if v, err := iprot.ReadString(); err != nil {
			return fmt.Errorf("error reading field 0: %s", err)
		} else {
			_val1 = v
		}
		p.Extra[_key0] = _val1
	}
	if err := iprot.ReadMapEnd(); err != nil {
		return fmt.Errorf("error reading map end: %s", err)
	}
	return nil
}

func (p *Base) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Base"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if p != nil {
		if err := p.writeField1(oprot); err != nil {
			return err
		}
		if err := p.writeField2(oprot); err != nil {
			return err
		}
		if err := p.writeField3(oprot); err != nil {
			return err
		}
		if err := p.writeField4(oprot); err != nil {
			return err
		}
		if err := p.writeField5(oprot); err != nil {
			return err
		}
		if err := p.writeField6(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *Base) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("LogID", thrift.STRING, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:LogID: %s", p, err)
	}
	if err := oprot.WriteString(string(p.LogID)); err != nil {
		return fmt.Errorf("%T.LogID (1) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:LogID: %s", p, err)
	}
	return err
}

func (p *Base) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("Caller", thrift.STRING, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:Caller: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Caller)); err != nil {
		return fmt.Errorf("%T.Caller (2) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:Caller: %s", p, err)
	}
	return err
}

func (p *Base) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("Addr", thrift.STRING, 3); err != nil {
		return fmt.Errorf("%T write field begin error 3:Addr: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Addr)); err != nil {
		return fmt.Errorf("%T.Addr (3) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 3:Addr: %s", p, err)
	}
	return err
}

func (p *Base) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("Client", thrift.STRING, 4); err != nil {
		return fmt.Errorf("%T write field begin error 4:Client: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Client)); err != nil {
		return fmt.Errorf("%T.Client (4) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 4:Client: %s", p, err)
	}
	return err
}

func (p *Base) writeField5(oprot thrift.TProtocol) (err error) {
	if p.IsSetTrafficEnv() {
		if err := oprot.WriteFieldBegin("TrafficEnv", thrift.STRUCT, 5); err != nil {
			return fmt.Errorf("%T write field begin error 5:TrafficEnv: %s", p, err)
		}
		if err := p.TrafficEnv.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", p.TrafficEnv, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 5:TrafficEnv: %s", p, err)
		}
	}
	return err
}

func (p *Base) writeField6(oprot thrift.TProtocol) (err error) {
	if p.IsSetExtra() {
		if err := oprot.WriteFieldBegin("Extra", thrift.MAP, 6); err != nil {
			return fmt.Errorf("%T write field begin error 6:Extra: %s", p, err)
		}
		if err := oprot.WriteMapBegin(thrift.STRING, thrift.STRING, len(p.Extra)); err != nil {
			return fmt.Errorf("error writing map begin: %s", err)
		}
		for k, v := range p.Extra {
			if err := oprot.WriteString(string(k)); err != nil {
				return fmt.Errorf("%T. (0) field write error: %s", p, err)
			}
			if err := oprot.WriteString(string(v)); err != nil {
				return fmt.Errorf("%T. (0) field write error: %s", p, err)
			}
		}
		if err := oprot.WriteMapEnd(); err != nil {
			return fmt.Errorf("error writing map end: %s", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 6:Extra: %s", p, err)
		}
	}
	return err
}

func (p *Base) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Base(%+v)", *p)
}

type BaseResp struct {
	StatusMessage string            `thrift:"StatusMessage,1" json:"StatusMessage"`
	StatusCode    int32             `thrift:"StatusCode,2" json:"StatusCode"`
	Extra         map[string]string `thrift:"Extra,3" json:"Extra"`
}

func NewBaseResp() *BaseResp {
	return &BaseResp{}
}

func (p *BaseResp) GetStatusMessage() string {
	return p.StatusMessage
}

func (p *BaseResp) GetStatusCode() int32 {
	return p.StatusCode
}

var BaseResp_Extra_DEFAULT map[string]string

func (p *BaseResp) GetExtra() map[string]string {
	return p.Extra
}
func (p *BaseResp) IsSetExtra() bool {
	return p.Extra != nil
}

func (p *BaseResp) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.ReadField3(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *BaseResp) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		p.StatusMessage = v
	}
	return nil
}

func (p *BaseResp) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 2: %s", err)
	} else {
		p.StatusCode = v
	}
	return nil
}

func (p *BaseResp) ReadField3(iprot thrift.TProtocol) error {
	_, _, size, err := iprot.ReadMapBegin()
	if err != nil {
		return fmt.Errorf("error reading map begin: %s", err)
	}
	tMap := make(map[string]string, size)
	p.Extra = tMap
	for i := 0; i < size; i++ {
		var _key2 string
		if v, err := iprot.ReadString(); err != nil {
			return fmt.Errorf("error reading field 0: %s", err)
		} else {
			_key2 = v
		}
		var _val3 string
		if v, err := iprot.ReadString(); err != nil {
			return fmt.Errorf("error reading field 0: %s", err)
		} else {
			_val3 = v
		}
		p.Extra[_key2] = _val3
	}
	if err := iprot.ReadMapEnd(); err != nil {
		return fmt.Errorf("error reading map end: %s", err)
	}
	return nil
}

func (p *BaseResp) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("BaseResp"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if p != nil {
		if err := p.writeField1(oprot); err != nil {
			return err
		}
		if err := p.writeField2(oprot); err != nil {
			return err
		}
		if err := p.writeField3(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *BaseResp) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("StatusMessage", thrift.STRING, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:StatusMessage: %s", p, err)
	}
	if err := oprot.WriteString(string(p.StatusMessage)); err != nil {
		return fmt.Errorf("%T.StatusMessage (1) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:StatusMessage: %s", p, err)
	}
	return err
}

func (p *BaseResp) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("StatusCode", thrift.I32, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:StatusCode: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.StatusCode)); err != nil {
		return fmt.Errorf("%T.StatusCode (2) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:StatusCode: %s", p, err)
	}
	return err
}

func (p *BaseResp) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetExtra() {
		if err := oprot.WriteFieldBegin("Extra", thrift.MAP, 3); err != nil {
			return fmt.Errorf("%T write field begin error 3:Extra: %s", p, err)
		}
		if err := oprot.WriteMapBegin(thrift.STRING, thrift.STRING, len(p.Extra)); err != nil {
			return fmt.Errorf("error writing map begin: %s", err)
		}
		for k, v := range p.Extra {
			if err := oprot.WriteString(string(k)); err != nil {
				return fmt.Errorf("%T. (0) field write error: %s", p, err)
			}
			if err := oprot.WriteString(string(v)); err != nil {
				return fmt.Errorf("%T. (0) field write error: %s", p, err)
			}
		}
		if err := oprot.WriteMapEnd(); err != nil {
			return fmt.Errorf("error writing map end: %s", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 3:Extra: %s", p, err)
		}
	}
	return err
}

func (p *BaseResp) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("BaseResp(%+v)", *p)
}
