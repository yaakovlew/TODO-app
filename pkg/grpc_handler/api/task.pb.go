// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.0
// source: task.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_task_proto protoreflect.FileDescriptor

var file_task_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74, 0x61,
	0x73, 0x6b, 0x1a, 0x0c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x32, 0xac, 0x03, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x39, 0x0a, 0x0a, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x16, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a,
	0x13, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x54, 0x61,
	0x73, 0x6b, 0x49, 0x44, 0x12, 0x31, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x12,
	0x13, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x54, 0x61,
	0x73, 0x6b, 0x49, 0x44, 0x1a, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x32, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x1a, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x36, 0x0a, 0x0a, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x13, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x1a, 0x13,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x54, 0x61, 0x73,
	0x6b, 0x49, 0x44, 0x12, 0x41, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x42, 0x79,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x42, 0x79, 0x44, 0x61, 0x74, 0x65, 0x1a, 0x17, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x54, 0x61, 0x73, 0x6b,
	0x73, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73,
	0x6b, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x19, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x54, 0x61, 0x73, 0x6b, 0x73, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x42, 0x79, 0x50, 0x61, 0x67, 0x65, 0x12, 0x18, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x42,
	0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_task_proto_goTypes = []interface{}{
	(*TaskInput)(nil),    // 0: grpc_models.TaskInput
	(*TaskID)(nil),       // 1: grpc_models.TaskID
	(*Task)(nil),         // 2: grpc_models.Task
	(*TaskByDate)(nil),   // 3: grpc_models.TaskByDate
	(*EmptyRequest)(nil), // 4: grpc_models.EmptyRequest
	(*PageRequest)(nil),  // 5: grpc_models.PageRequest
	(*TasksSlice)(nil),   // 6: grpc_models.TasksSlice
}
var file_task_proto_depIdxs = []int32{
	0, // 0: task.Task.CreateTask:input_type -> grpc_models.TaskInput
	1, // 1: task.Task.GetTask:input_type -> grpc_models.TaskID
	2, // 2: task.Task.UpdateTask:input_type -> grpc_models.Task
	1, // 3: task.Task.DeleteTask:input_type -> grpc_models.TaskID
	3, // 4: task.Task.GetTaskByDate:input_type -> grpc_models.TaskByDate
	4, // 5: task.Task.GetTasksList:input_type -> grpc_models.EmptyRequest
	5, // 6: task.Task.GetTasksByPage:input_type -> grpc_models.PageRequest
	1, // 7: task.Task.CreateTask:output_type -> grpc_models.TaskID
	2, // 8: task.Task.GetTask:output_type -> grpc_models.Task
	2, // 9: task.Task.UpdateTask:output_type -> grpc_models.Task
	1, // 10: task.Task.DeleteTask:output_type -> grpc_models.TaskID
	6, // 11: task.Task.GetTaskByDate:output_type -> grpc_models.TasksSlice
	6, // 12: task.Task.GetTasksList:output_type -> grpc_models.TasksSlice
	6, // 13: task.Task.GetTasksByPage:output_type -> grpc_models.TasksSlice
	7, // [7:14] is the sub-list for method output_type
	0, // [0:7] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_task_proto_init() }
func file_task_proto_init() {
	if File_task_proto != nil {
		return
	}
	file_models_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_task_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_task_proto_goTypes,
		DependencyIndexes: file_task_proto_depIdxs,
	}.Build()
	File_task_proto = out.File
	file_task_proto_rawDesc = nil
	file_task_proto_goTypes = nil
	file_task_proto_depIdxs = nil
}
