import * as jspb from 'google-protobuf'



export class TestRequest extends jspb.Message {
  getMessage(): string;
  setMessage(value: string): TestRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TestRequest.AsObject;
  static toObject(includeInstance: boolean, msg: TestRequest): TestRequest.AsObject;
  static serializeBinaryToWriter(message: TestRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TestRequest;
  static deserializeBinaryFromReader(message: TestRequest, reader: jspb.BinaryReader): TestRequest;
}

export namespace TestRequest {
  export type AsObject = {
    message: string,
  }
}

export class TestResponse extends jspb.Message {
  getResponse(): string;
  setResponse(value: string): TestResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TestResponse.AsObject;
  static toObject(includeInstance: boolean, msg: TestResponse): TestResponse.AsObject;
  static serializeBinaryToWriter(message: TestResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TestResponse;
  static deserializeBinaryFromReader(message: TestResponse, reader: jspb.BinaryReader): TestResponse;
}

export namespace TestResponse {
  export type AsObject = {
    response: string,
  }
}

export class ExtractRequest extends jspb.Message {
  getBinary(): Uint8Array | string;
  getBinary_asU8(): Uint8Array;
  getBinary_asB64(): string;
  setBinary(value: Uint8Array | string): ExtractRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ExtractRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ExtractRequest): ExtractRequest.AsObject;
  static serializeBinaryToWriter(message: ExtractRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ExtractRequest;
  static deserializeBinaryFromReader(message: ExtractRequest, reader: jspb.BinaryReader): ExtractRequest;
}

export namespace ExtractRequest {
  export type AsObject = {
    binary: Uint8Array | string,
  }
}

export class SearchRequest extends jspb.Message {
  getIndex(): string;
  setIndex(value: string): SearchRequest;

  getQuery(): string;
  setQuery(value: string): SearchRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SearchRequest.AsObject;
  static toObject(includeInstance: boolean, msg: SearchRequest): SearchRequest.AsObject;
  static serializeBinaryToWriter(message: SearchRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SearchRequest;
  static deserializeBinaryFromReader(message: SearchRequest, reader: jspb.BinaryReader): SearchRequest;
}

export namespace SearchRequest {
  export type AsObject = {
    index: string,
    query: string,
  }
}

export class ExpenseField extends jspb.Message {
  getText(): string;
  setText(value: string): ExpenseField;

  getConfidence(): number;
  setConfidence(value: number): ExpenseField;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ExpenseField.AsObject;
  static toObject(includeInstance: boolean, msg: ExpenseField): ExpenseField.AsObject;
  static serializeBinaryToWriter(message: ExpenseField, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ExpenseField;
  static deserializeBinaryFromReader(message: ExpenseField, reader: jspb.BinaryReader): ExpenseField;
}

export namespace ExpenseField {
  export type AsObject = {
    text: string,
    confidence: number,
  }
}

export class ExtractResponse extends jspb.Message {
  getFilePage(): ExpenseField | undefined;
  setFilePage(value?: ExpenseField): ExtractResponse;
  hasFilePage(): boolean;
  clearFilePage(): ExtractResponse;

  getFileName(): ExpenseField | undefined;
  setFileName(value?: ExpenseField): ExtractResponse;
  hasFileName(): boolean;
  clearFileName(): ExtractResponse;

  getInvoiceReceiptDate(): ExpenseField | undefined;
  setInvoiceReceiptDate(value?: ExpenseField): ExtractResponse;
  hasInvoiceReceiptDate(): boolean;
  clearInvoiceReceiptDate(): ExtractResponse;

  getVendorName(): ExpenseField | undefined;
  setVendorName(value?: ExpenseField): ExtractResponse;
  hasVendorName(): boolean;
  clearVendorName(): ExtractResponse;

  getVendorAddress(): ExpenseField | undefined;
  setVendorAddress(value?: ExpenseField): ExtractResponse;
  hasVendorAddress(): boolean;
  clearVendorAddress(): ExtractResponse;

  getTotal(): ExpenseField | undefined;
  setTotal(value?: ExpenseField): ExtractResponse;
  hasTotal(): boolean;
  clearTotal(): ExtractResponse;

  getSubtotal(): ExpenseField | undefined;
  setSubtotal(value?: ExpenseField): ExtractResponse;
  hasSubtotal(): boolean;
  clearSubtotal(): ExtractResponse;

  getTax(): ExpenseField | undefined;
  setTax(value?: ExpenseField): ExtractResponse;
  hasTax(): boolean;
  clearTax(): ExtractResponse;

  getVendorPhone(): ExpenseField | undefined;
  setVendorPhone(value?: ExpenseField): ExtractResponse;
  hasVendorPhone(): boolean;
  clearVendorPhone(): ExtractResponse;

  getStreet(): ExpenseField | undefined;
  setStreet(value?: ExpenseField): ExtractResponse;
  hasStreet(): boolean;
  clearStreet(): ExtractResponse;

  getGratuity(): ExpenseField | undefined;
  setGratuity(value?: ExpenseField): ExtractResponse;
  hasGratuity(): boolean;
  clearGratuity(): ExtractResponse;

  getCity(): ExpenseField | undefined;
  setCity(value?: ExpenseField): ExtractResponse;
  hasCity(): boolean;
  clearCity(): ExtractResponse;

  getState(): ExpenseField | undefined;
  setState(value?: ExpenseField): ExtractResponse;
  hasState(): boolean;
  clearState(): ExtractResponse;

  getCountry(): ExpenseField | undefined;
  setCountry(value?: ExpenseField): ExtractResponse;
  hasCountry(): boolean;
  clearCountry(): ExtractResponse;

  getZipCode(): ExpenseField | undefined;
  setZipCode(value?: ExpenseField): ExtractResponse;
  hasZipCode(): boolean;
  clearZipCode(): ExtractResponse;

  getCategory(): ExpenseField | undefined;
  setCategory(value?: ExpenseField): ExtractResponse;
  hasCategory(): boolean;
  clearCategory(): ExtractResponse;

  getObjectUrl(): string;
  setObjectUrl(value: string): ExtractResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ExtractResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ExtractResponse): ExtractResponse.AsObject;
  static serializeBinaryToWriter(message: ExtractResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ExtractResponse;
  static deserializeBinaryFromReader(message: ExtractResponse, reader: jspb.BinaryReader): ExtractResponse;
}

export namespace ExtractResponse {
  export type AsObject = {
    filePage?: ExpenseField.AsObject,
    fileName?: ExpenseField.AsObject,
    invoiceReceiptDate?: ExpenseField.AsObject,
    vendorName?: ExpenseField.AsObject,
    vendorAddress?: ExpenseField.AsObject,
    total?: ExpenseField.AsObject,
    subtotal?: ExpenseField.AsObject,
    tax?: ExpenseField.AsObject,
    vendorPhone?: ExpenseField.AsObject,
    street?: ExpenseField.AsObject,
    gratuity?: ExpenseField.AsObject,
    city?: ExpenseField.AsObject,
    state?: ExpenseField.AsObject,
    country?: ExpenseField.AsObject,
    zipCode?: ExpenseField.AsObject,
    category?: ExpenseField.AsObject,
    objectUrl: string,
  }
}

export class Expenses extends jspb.Message {
  getInfoList(): Array<ExtractResponse>;
  setInfoList(value: Array<ExtractResponse>): Expenses;
  clearInfoList(): Expenses;
  addInfo(value?: ExtractResponse, index?: number): ExtractResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Expenses.AsObject;
  static toObject(includeInstance: boolean, msg: Expenses): Expenses.AsObject;
  static serializeBinaryToWriter(message: Expenses, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Expenses;
  static deserializeBinaryFromReader(message: Expenses, reader: jspb.BinaryReader): Expenses;
}

export namespace Expenses {
  export type AsObject = {
    infoList: Array<ExtractResponse.AsObject>,
  }
}

