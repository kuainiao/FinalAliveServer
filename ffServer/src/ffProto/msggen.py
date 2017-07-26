# -*- coding: utf-8 -*-
# 使用 protoc-3.2.0.exe 生成协议

import os
import sys

configfile = open("gen_message.gen", "r")
message_cfg = configfile.read()
configfile.close()

configfile = open("gen_struct.gen", "r")
struct_cfg = configfile.read()
configfile.close()

msg_types = []

def parse(text, ismsg):
	output = ''
	result = ''
	while True:
		zsstart = text.find('/*')
		if zsstart != -1:
			result += text[:zsstart]
		else:
			result += text
			break
		zsend = text.find('*/', zsstart)
		if zsend != -1:
			text = text[zsend+2:]
		else:
			break
	text = result
	text = text.strip()
	msgs = text.split('@')
	for msg in msgs:
		msg = msg.strip()
		if msg == '':
			continue
		output += parse_msg(msg, ismsg)

	if not ismsg:
		return output

	enumstr = ''
	enumstr += '\t' + 'SpecialEventInvalid' + ' = 0;\n'
	type_i = 1
	for t in msg_types:
		enumstr += '\t' + t + ' = %d;\n' % type_i
		type_i += 1
	enumstr += '\t' + 'SpecialEventConnect' + ' = -1;\n'
	enumstr += '\t' + 'SpecialEventDisConnect' + ' = -2;\n'
	enumstr = '''enum MessageType {
%s}''' % enumstr
	output = enumstr + '\n\n' + output
	return output

def parse_msg(msg, ismsg):
	msg_top = {}
	top_sec = []
	msg_member = ''
	lines = msg.split('\n')
	msg_name = lines[0].strip()
	print(msg_name)
	mem_i = 1
	for line in lines[1:]:
		line = line.strip()
		if line == '':
			continue

		tags = line.split()
		if len(tags) == 1:
			pass
		elif len(tags) == 2:
			msg_member +='\t' + line + ' = %d;\n' % mem_i
		else:
			msg_member += '\t' + line + ' = %d;\n' % mem_i
		mem_i += 1



	if ismsg:
		if msg_name.find('Msg') != 0:
			go_error('Proto message must start with Msg')
		msg_types.append(msg_name[len('Msg'):])
		output = '''message %s{\n%s}\n\n''' % (msg_name, msg_member)
	else:
		output = '''message %s{\n%s}\n\n''' % (msg_name, msg_member)

	return output

message_proto_header = '''// Generated by the msggen.py message compiler.
'''

def save_go():
	t = '''
	package ffProto

	import "github.com/golang/protobuf/proto"

	'''


	t1 = 'var listProtoID = []MessageType{\n'
	t2 = 'var mapMessageCreator = map[MessageType]func() interface{}{\n'
	for mt in msg_types:
		t1 += 'MessageType_' + mt + ',\n'
		t2 += '''MessageType_%s: func() interface{} {
			return &Msg%s{}
		},
		''' % (mt, mt)
	t1 += '}\n\n'
	t2 += '}\n\n'

	text = '''
	package ffProto

	'''
	text += t1
	text += t2

	f = open('message.pb.go', 'w+')
	f.write(text)
	f.close()

	os.system("go fmt message.pb.go")

def saveoutput(text):
	global message_proto_header
	proto_header = 'syntax = "proto3";\n\n'
	text = proto_header + message_proto_header + text

	f = open('ffProto.proto', 'w+')
	f.write(text)
	f.close()

def go_error(error_msg):
	print error_msg
	os.system('color 4a')
	exit(1)

if __name__ == '__main__':
	saveoutput(parse(struct_cfg, False) + parse(message_cfg, True))
	print('\n')

	save_go()
	print('\n')
