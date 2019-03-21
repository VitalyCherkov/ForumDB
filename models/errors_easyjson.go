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

func easyjsonD31a5a85DecodeForumDBModels(in *jlexer.Lexer, out *ErrorNotFound) {
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
		case "message":
			out.Message = string(in.String())
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
func easyjsonD31a5a85EncodeForumDBModels(out *jwriter.Writer, in ErrorNotFound) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"message\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Message))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ErrorNotFound) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD31a5a85EncodeForumDBModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ErrorNotFound) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD31a5a85EncodeForumDBModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ErrorNotFound) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD31a5a85DecodeForumDBModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ErrorNotFound) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD31a5a85DecodeForumDBModels(l, v)
}
func easyjsonD31a5a85DecodeForumDBModels1(in *jlexer.Lexer, out *ErrorConflict) {
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
		case "message":
			out.Message = string(in.String())
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
func easyjsonD31a5a85EncodeForumDBModels1(out *jwriter.Writer, in ErrorConflict) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"message\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Message))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ErrorConflict) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD31a5a85EncodeForumDBModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ErrorConflict) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD31a5a85EncodeForumDBModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ErrorConflict) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD31a5a85DecodeForumDBModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ErrorConflict) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD31a5a85DecodeForumDBModels1(l, v)
}
func easyjsonD31a5a85DecodeForumDBModels2(in *jlexer.Lexer, out *DatabaseError) {
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
		case "Message":
			out.Message = string(in.String())
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
func easyjsonD31a5a85EncodeForumDBModels2(out *jwriter.Writer, in DatabaseError) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Message\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Message))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DatabaseError) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD31a5a85EncodeForumDBModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DatabaseError) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD31a5a85EncodeForumDBModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DatabaseError) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD31a5a85DecodeForumDBModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DatabaseError) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD31a5a85DecodeForumDBModels2(l, v)
}
