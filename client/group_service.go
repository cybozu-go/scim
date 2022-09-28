package client

func (call *CreateGroupCall) MembersFrom(v ...interface{}) *CreateGroupCall {
	call.builder.MembersFrom(v...)
	return call
}
