import * as jspb from 'google-protobuf'



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
  getModulesList(): Array<Module>;
  setModulesList(value: Array<Module>): ListModulesResponse;
  clearModulesList(): ListModulesResponse;
  addModules(value?: Module, index?: number): Module;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListModulesResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListModulesResponse): ListModulesResponse.AsObject;
  static serializeBinaryToWriter(message: ListModulesResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListModulesResponse;
  static deserializeBinaryFromReader(message: ListModulesResponse, reader: jspb.BinaryReader): ListModulesResponse;
}

export namespace ListModulesResponse {
  export type AsObject = {
    modulesList: Array<Module.AsObject>,
  }
}

export class Envelope extends jspb.Message {
  getA(): ListModulesRequest | undefined;
  setA(value?: ListModulesRequest): Envelope;
  hasA(): boolean;
  clearA(): Envelope;

  getB(): ListModulesResponse | undefined;
  setB(value?: ListModulesResponse): Envelope;
  hasB(): boolean;
  clearB(): Envelope;

  getKindCase(): Envelope.KindCase;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Envelope.AsObject;
  static toObject(includeInstance: boolean, msg: Envelope): Envelope.AsObject;
  static serializeBinaryToWriter(message: Envelope, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Envelope;
  static deserializeBinaryFromReader(message: Envelope, reader: jspb.BinaryReader): Envelope;
}

export namespace Envelope {
  export type AsObject = {
    a?: ListModulesRequest.AsObject,
    b?: ListModulesResponse.AsObject,
  }

  export enum KindCase { 
    KIND_NOT_SET = 0,
    A = 1,
    B = 2,
  }
}

