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

func easyjson5a72dc82DecodeForumDBModels(in *jlexer.Lexer, out *PostDetailList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(PostDetailList, 0, 1)
			} else {
				*out = PostDetailList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 PostDetail
			(v1).UnmarshalEasyJSON(in)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson5a72dc82EncodeForumDBModels(out *jwriter.Writer, in PostDetailList) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v PostDetailList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5a72dc82EncodeForumDBModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostDetailList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5a72dc82EncodeForumDBModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostDetailList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5a72dc82DecodeForumDBModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostDetailList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5a72dc82DecodeForumDBModels(l, v)
}
func easyjson5a72dc82DecodeForumDBModels1(in *jlexer.Lexer, out *PostDetail) {
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
		case "author":
			out.Author = string(in.String())
		case "message":
			out.Message = string(in.String())
		case "parent":
			out.Parent = uint64(in.Uint64())
		case "id":
			out.Id = uint64(in.Uint64())
		case "created":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Created).UnmarshalJSON(data))
			}
		case "forum":
			out.Forum = string(in.String())
		case "thread":
			out.Thread = uint64(in.Uint64())
		case "isEdited":
			out.IsEdited = bool(in.Bool())
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
func easyjson5a72dc82EncodeForumDBModels1(out *jwriter.Writer, in PostDetail) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"author\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Author))
	}
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
	{
		const prefix string = ",\"parent\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint64(uint64(in.Parent))
	}
	if in.Id != 0 {
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint64(uint64(in.Id))
	}
	if true {
		const prefix string = ",\"created\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.Created).MarshalJSON())
	}
	if in.Forum != "" {
		const prefix string = ",\"forum\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Forum))
	}
	if in.Thread != 0 {
		const prefix string = ",\"thread\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint64(uint64(in.Thread))
	}
	if in.IsEdited {
		const prefix string = ",\"isEdited\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.IsEdited))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PostDetail) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5a72dc82EncodeForumDBModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostDetail) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5a72dc82EncodeForumDBModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostDetail) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5a72dc82DecodeForumDBModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostDetail) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5a72dc82DecodeForumDBModels1(l, v)
}
func easyjson5a72dc82DecodeForumDBModels2(in *jlexer.Lexer, out *PostCombined) {
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
		case "post":
			if in.IsNull() {
				in.Skip()
				out.Post = nil
			} else {
				if out.Post == nil {
					out.Post = new(PostDetail)
				}
				(*out.Post).UnmarshalEasyJSON(in)
			}
		case "author":
			if in.IsNull() {
				in.Skip()
				out.Author = nil
			} else {
				if out.Author == nil {
					out.Author = new(UserDetail)
				}
				(*out.Author).UnmarshalEasyJSON(in)
			}
		case "thread":
			if in.IsNull() {
				in.Skip()
				out.Thread = nil
			} else {
				if out.Thread == nil {
					out.Thread = new(ThreadDetail)
				}
				(*out.Thread).UnmarshalEasyJSON(in)
			}
		case "forum":
			if in.IsNull() {
				in.Skip()
				out.Forum = nil
			} else {
				if out.Forum == nil {
					out.Forum = new(ForumDetail)
				}
				(*out.Forum).UnmarshalEasyJSON(in)
			}
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
func easyjson5a72dc82EncodeForumDBModels2(out *jwriter.Writer, in PostCombined) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"post\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Post == nil {
			out.RawString("null")
		} else {
			(*in.Post).MarshalEasyJSON(out)
		}
	}
	if in.Author != nil {
		const prefix string = ",\"author\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Author).MarshalEasyJSON(out)
	}
	if in.Thread != nil {
		const prefix string = ",\"thread\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Thread).MarshalEasyJSON(out)
	}
	if in.Forum != nil {
		const prefix string = ",\"forum\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Forum).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PostCombined) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5a72dc82EncodeForumDBModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostCombined) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5a72dc82EncodeForumDBModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostCombined) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5a72dc82DecodeForumDBModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostCombined) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5a72dc82DecodeForumDBModels2(l, v)
}
