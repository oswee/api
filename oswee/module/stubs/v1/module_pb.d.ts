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

