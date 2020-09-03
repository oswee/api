/**
 * @fileoverview gRPC-Web generated client stub for oswee.module.v1
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as module_pb from './module_pb';


export class ModulesServiceClient {
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
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoListModules = new grpcWeb.AbstractClientBase.MethodInfo(
    module_pb.ListModulesResponse,
    (request: module_pb.ListModulesRequest) => {
      return request.serializeBinary();
    },
    module_pb.ListModulesResponse.deserializeBinary
  );

  listModules(
    request: module_pb.ListModulesRequest,
    metadata: grpcWeb.Metadata | null): Promise<module_pb.ListModulesResponse>;

  listModules(
    request: module_pb.ListModulesRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: module_pb.ListModulesResponse) => void): grpcWeb.ClientReadableStream<module_pb.ListModulesResponse>;

  listModules(
    request: module_pb.ListModulesRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: module_pb.ListModulesResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/oswee.module.v1.ModulesService/ListModules',
        request,
        metadata || {},
        this.methodInfoListModules,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/oswee.module.v1.ModulesService/ListModules',
    request,
    metadata || {},
    this.methodInfoListModules);
  }

}

