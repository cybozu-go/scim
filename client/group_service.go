package client

func (call *CreateGroupCall) MemberFrom(v interface{}) *CreateGroupCall {
	call.builder.MemberFrom(v)
	return call
}
