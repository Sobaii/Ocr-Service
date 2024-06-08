/**
 * @fileoverview gRPC-Web generated client stub for ocr_service
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.5.0
// 	protoc              v3.12.4
// source: ocr_service.proto


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as ocr_service_pb from './ocr_service_pb'; // proto import: "ocr_service.proto"


export class OcrServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname.replace(/\/+$/, '');
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodDescriptorTestConnection = new grpcWeb.MethodDescriptor(
    '/ocr_service.OcrService/TestConnection',
    grpcWeb.MethodType.UNARY,
    ocr_service_pb.TestRequest,
    ocr_service_pb.TestResponse,
    (request: ocr_service_pb.TestRequest) => {
      return request.serializeBinary();
    },
    ocr_service_pb.TestResponse.deserializeBinary
  );

  testConnection(
    request: ocr_service_pb.TestRequest,
    metadata?: grpcWeb.Metadata | null): Promise<ocr_service_pb.TestResponse>;

  testConnection(
    request: ocr_service_pb.TestRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: ocr_service_pb.TestResponse) => void): grpcWeb.ClientReadableStream<ocr_service_pb.TestResponse>;

  testConnection(
    request: ocr_service_pb.TestRequest,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: ocr_service_pb.TestResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/ocr_service.OcrService/TestConnection',
        request,
        metadata || {},
        this.methodDescriptorTestConnection,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/ocr_service.OcrService/TestConnection',
    request,
    metadata || {},
    this.methodDescriptorTestConnection);
  }

}

