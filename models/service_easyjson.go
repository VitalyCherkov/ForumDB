// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonCd93bc43DecodeForumDBModels(in *jlexer.Lexer, out *ServiceStatus) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "forum":
			out.Forum = int(in.Int())
		case "post":
			out.Post = int(in.Int())
		case "thread":
			out.Thread = int(in.Int())
		case "user":
			out.User = int(in.Int())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonCd93bc43EncodeForumDBModels(out *jwriter.Writer, in ServiceStatus) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"forum\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Forum))
	}
	{
		const prefix string = ",\"post\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Post))
	}
	{
		const prefix string = ",\"thread\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Thread))
	}
	{
		const prefix string = ",\"user\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.User))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ServiceStatus) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCd93bc43EncodeForumDBModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ServiceStatus) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCd93bc43EncodeForumDBModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ServiceStatus) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCd93bc43DecodeForumDBModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ServiceStatus) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCd93bc43DecodeForumDBModels(l, v)
}
