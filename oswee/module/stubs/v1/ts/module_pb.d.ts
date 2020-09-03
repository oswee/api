import * as jspb from 'google-protobuf'

import * as protoc$gen$swagger_options_annotations_pb from './protoc-gen-swagger/options/annotations_pb';
import * as google_api_annotations_pb from './google/api/annotations_pb';


export class Module extends jspb.Message {
  getId(): string;
  setId(value: string): Module;

  getTitle(): string;
  setTitle(value: string): Module;

  getPermalink(): string;
  setPermalink(value: string): Module;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Module.AsObject;
  static toObject(includeInstance: boolean, msg: Module): Module.AsObject;
  static serializeBinaryToWriter(message: Module, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Module;
  static deserializeBinaryFromReader(message: Module, reader: jspb.BinaryReader): Module;
}

export namespace Module {
  export type AsObject = {
    id: string,
    title: string,
    permalink: string,
  }
}

export class ListModulesRequest extends jspb.Message {
  getParent(): string;
  setParent(value: string): ListModulesRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListModulesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListModulesRequest): ListModulesRequest.AsObject;
  static serializeBinaryToWriter(message: ListModulesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListModulesRequest;
  static deserializeBinaryFromReader(message: ListModulesRequest, reader: jspb.BinaryReader): ListModulesRequest;
}

export namespace ListModulesRequest {
  export type AsObject = {
    parent: string,
  }
}

export class ListModulesResponse extends jspb.Message {
  getEntitiesMap(): jspb.Map<string, Module>;
  clearEntitiesMap(): ListModulesResponse;

  getIdsList(): Array<string>;
  setIdsList(value: Array<string>): ListModulesResponse;
  clearIdsList(): ListModulesResponse;
  addIds(value: string, index?: number): ListModulesResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListModulesResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListModulesResponse): ListModulesResponse.AsObject;
  static serializeBinaryToWriter(message: ListModulesResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListModulesResponse;
  static deserializeBinaryFromReader(message: ListModulesResponse, reader: jspb.BinaryReader): ListModulesResponse;
}

export namespace ListModulesResponse {
  export type AsObject = {
    entitiesMap: Array<[string, Module.AsObject]>,
    idsList: Array<string>,
  }
}

