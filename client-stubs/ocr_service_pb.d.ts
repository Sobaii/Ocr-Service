import * as jspb from 'google-protobuf'



export class ExtractFileRequest extends jspb.Message {
  getEmailAddress(): string;
  setEmailAddress(value: string): ExtractFileRequest;

  getFolderName(): string;
  setFolderName(value: string): ExtractFileRequest;

  getBinary(): Uint8Array | string;
  getBinary_asU8(): Uint8Array;
  getBinary_asB64(): string;
  setBinary(value: Uint8Array | string): ExtractFileRequest;

  getMimeType(): MimeType;
  setMimeType(value: MimeType): ExtractFileRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ExtractFileRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ExtractFileRequest): ExtractFileRequest.AsObject;
  static serializeBinaryToWriter(message: ExtractFileRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ExtractFileRequest;
  static deserializeBinaryFromReader(message: ExtractFileRequest, reader: jspb.BinaryReader): ExtractFileRequest;
}

export namespace ExtractFileRequest {
  export type AsObject = {
    emailAddress: string,
    folderName: string,
    binary: Uint8Array | string,
    mimeType: MimeType,
  }
}

export class SearchFileRequest extends jspb.Message {
  getEmailAddress(): string;
  setEmailAddress(value: string): SearchFileRequest;

  getFolderName(): string;
  setFolderName(value: string): SearchFileRequest;

  getIndex(): string;
  setIndex(value: string): SearchFileRequest;

  getQuery(): string;
  setQuery(value: string): SearchFileRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SearchFileRequest.AsObject;
  static toObject(includeInstance: boolean, msg: SearchFileRequest): SearchFileRequest.AsObject;
  static serializeBinaryToWriter(message: SearchFileRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SearchFileRequest;
  static deserializeBinaryFromReader(message: SearchFileRequest, reader: jspb.BinaryReader): SearchFileRequest;
}

export namespace SearchFileRequest {
  export type AsObject = {
    emailAddress: string,
    folderName: string,
    index: string,
    query: string,
  }
}

export class FolderCreationRequest extends jspb.Message {
  getEmailAddress(): string;
  setEmailAddress(value: string): FolderCreationRequest;

  getFullName(): string;
  setFullName(value: string): FolderCreationRequest;

  getFolderName(): string;
  setFolderName(value: string): FolderCreationRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): FolderCreationRequest.AsObject;
  static toObject(includeInstance: boolean, msg: FolderCreationRequest): FolderCreationRequest.AsObject;
  static serializeBinaryToWriter(message: FolderCreationRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): FolderCreationRequest;
  static deserializeBinaryFromReader(message: FolderCreationRequest, reader: jspb.BinaryReader): FolderCreationRequest;
}

export namespace FolderCreationRequest {
  export type AsObject = {
    emailAddress: string,
    fullName: string,
    folderName: string,
  }
}

export class FolderSearchRequest extends jspb.Message {
  getEmailAddress(): string;
  setEmailAddress(value: string): FolderSearchRequest;

  getQuery(): string;
  setQuery(value: string): FolderSearchRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): FolderSearchRequest.AsObject;
  static toObject(includeInstance: boolean, msg: FolderSearchRequest): FolderSearchRequest.AsObject;
  static serializeBinaryToWriter(message: FolderSearchRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): FolderSearchRequest;
  static deserializeBinaryFromReader(message: FolderSearchRequest, reader: jspb.BinaryReader): FolderSearchRequest;
}

export namespace FolderSearchRequest {
  export type AsObject = {
    emailAddress: string,
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

export class FileExtract extends jspb.Message {
  getFilePage(): ExpenseField | undefined;
  setFilePage(value?: ExpenseField): FileExtract;
  hasFilePage(): boolean;
  clearFilePage(): FileExtract;

  getFileName(): ExpenseField | undefined;
  setFileName(value?: ExpenseField): FileExtract;
  hasFileName(): boolean;
  clearFileName(): FileExtract;

  getInvoiceReceiptDate(): ExpenseField | undefined;
  setInvoiceReceiptDate(value?: ExpenseField): FileExtract;
  hasInvoiceReceiptDate(): boolean;
  clearInvoiceReceiptDate(): FileExtract;

  getVendorName(): ExpenseField | undefined;
  setVendorName(value?: ExpenseField): FileExtract;
  hasVendorName(): boolean;
  clearVendorName(): FileExtract;

  getVendorAddress(): ExpenseField | undefined;
  setVendorAddress(value?: ExpenseField): FileExtract;
  hasVendorAddress(): boolean;
  clearVendorAddress(): FileExtract;

  getTotal(): ExpenseField | undefined;
  setTotal(value?: ExpenseField): FileExtract;
  hasTotal(): boolean;
  clearTotal(): FileExtract;

  getSubtotal(): ExpenseField | undefined;
  setSubtotal(value?: ExpenseField): FileExtract;
  hasSubtotal(): boolean;
  clearSubtotal(): FileExtract;

  getTax(): ExpenseField | undefined;
  setTax(value?: ExpenseField): FileExtract;
  hasTax(): boolean;
  clearTax(): FileExtract;

  getVendorPhone(): ExpenseField | undefined;
  setVendorPhone(value?: ExpenseField): FileExtract;
  hasVendorPhone(): boolean;
  clearVendorPhone(): FileExtract;

  getStreet(): ExpenseField | undefined;
  setStreet(value?: ExpenseField): FileExtract;
  hasStreet(): boolean;
  clearStreet(): FileExtract;

  getGratuity(): ExpenseField | undefined;
  setGratuity(value?: ExpenseField): FileExtract;
  hasGratuity(): boolean;
  clearGratuity(): FileExtract;

  getCity(): ExpenseField | undefined;
  setCity(value?: ExpenseField): FileExtract;
  hasCity(): boolean;
  clearCity(): FileExtract;

  getState(): ExpenseField | undefined;
  setState(value?: ExpenseField): FileExtract;
  hasState(): boolean;
  clearState(): FileExtract;

  getCountry(): ExpenseField | undefined;
  setCountry(value?: ExpenseField): FileExtract;
  hasCountry(): boolean;
  clearCountry(): FileExtract;

  getZipCode(): ExpenseField | undefined;
  setZipCode(value?: ExpenseField): FileExtract;
  hasZipCode(): boolean;
  clearZipCode(): FileExtract;

  getCategory(): ExpenseField | undefined;
  setCategory(value?: ExpenseField): FileExtract;
  hasCategory(): boolean;
  clearCategory(): FileExtract;

  getObjectUrl(): string;
  setObjectUrl(value: string): FileExtract;

  getPreviewUrl(): string;
  setPreviewUrl(value: string): FileExtract;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): FileExtract.AsObject;
  static toObject(includeInstance: boolean, msg: FileExtract): FileExtract.AsObject;
  static serializeBinaryToWriter(message: FileExtract, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): FileExtract;
  static deserializeBinaryFromReader(message: FileExtract, reader: jspb.BinaryReader): FileExtract;
}

export namespace FileExtract {
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
    previewUrl: string,
  }
}

export class ExpenseItem extends jspb.Message {
  getFolderName(): string;
  setFolderName(value: string): ExpenseItem;

  getData(): FileExtract | undefined;
  setData(value?: FileExtract): ExpenseItem;
  hasData(): boolean;
  clearData(): ExpenseItem;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ExpenseItem.AsObject;
  static toObject(includeInstance: boolean, msg: ExpenseItem): ExpenseItem.AsObject;
  static serializeBinaryToWriter(message: ExpenseItem, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ExpenseItem;
  static deserializeBinaryFromReader(message: ExpenseItem, reader: jspb.BinaryReader): ExpenseItem;
}

export namespace ExpenseItem {
  export type AsObject = {
    folderName: string,
    data?: FileExtract.AsObject,
  }
}

export class Expenses extends jspb.Message {
  getInfoList(): Array<ExpenseItem>;
  setInfoList(value: Array<ExpenseItem>): Expenses;
  clearInfoList(): Expenses;
  addInfo(value?: ExpenseItem, index?: number): ExpenseItem;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Expenses.AsObject;
  static toObject(includeInstance: boolean, msg: Expenses): Expenses.AsObject;
  static serializeBinaryToWriter(message: Expenses, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Expenses;
  static deserializeBinaryFromReader(message: Expenses, reader: jspb.BinaryReader): Expenses;
}

export namespace Expenses {
  export type AsObject = {
    infoList: Array<ExpenseItem.AsObject>,
  }
}

export class FolderSearchResponse extends jspb.Message {
  getFolderFound(): boolean;
  setFolderFound(value: boolean): FolderSearchResponse;

  getActionDescription(): string;
  setActionDescription(value: string): FolderSearchResponse;

  getFoldersList(): Array<string>;
  setFoldersList(value: Array<string>): FolderSearchResponse;
  clearFoldersList(): FolderSearchResponse;
  addFolders(value: string, index?: number): FolderSearchResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): FolderSearchResponse.AsObject;
  static toObject(includeInstance: boolean, msg: FolderSearchResponse): FolderSearchResponse.AsObject;
  static serializeBinaryToWriter(message: FolderSearchResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): FolderSearchResponse;
  static deserializeBinaryFromReader(message: FolderSearchResponse, reader: jspb.BinaryReader): FolderSearchResponse;
}

export namespace FolderSearchResponse {
  export type AsObject = {
    folderFound: boolean,
    actionDescription: string,
    foldersList: Array<string>,
  }
}

export class FolderCreationResponse extends jspb.Message {
  getFolderCreated(): boolean;
  setFolderCreated(value: boolean): FolderCreationResponse;

  getActionDescription(): string;
  setActionDescription(value: string): FolderCreationResponse;

  getEmailAddress(): string;
  setEmailAddress(value: string): FolderCreationResponse;

  getFullName(): string;
  setFullName(value: string): FolderCreationResponse;

  getFolderName(): string;
  setFolderName(value: string): FolderCreationResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): FolderCreationResponse.AsObject;
  static toObject(includeInstance: boolean, msg: FolderCreationResponse): FolderCreationResponse.AsObject;
  static serializeBinaryToWriter(message: FolderCreationResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): FolderCreationResponse;
  static deserializeBinaryFromReader(message: FolderCreationResponse, reader: jspb.BinaryReader): FolderCreationResponse;
}

export namespace FolderCreationResponse {
  export type AsObject = {
    folderCreated: boolean,
    actionDescription: string,
    emailAddress: string,
    fullName: string,
    folderName: string,
  }
}

export class SearchFileResponse extends jspb.Message {
  getFileFound(): boolean;
  setFileFound(value: boolean): SearchFileResponse;

  getActionDescription(): string;
  setActionDescription(value: string): SearchFileResponse;

  getEmailAddress(): string;
  setEmailAddress(value: string): SearchFileResponse;

  getFolderName(): string;
  setFolderName(value: string): SearchFileResponse;

  getExpenses(): Expenses | undefined;
  setExpenses(value?: Expenses): SearchFileResponse;
  hasExpenses(): boolean;
  clearExpenses(): SearchFileResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SearchFileResponse.AsObject;
  static toObject(includeInstance: boolean, msg: SearchFileResponse): SearchFileResponse.AsObject;
  static serializeBinaryToWriter(message: SearchFileResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SearchFileResponse;
  static deserializeBinaryFromReader(message: SearchFileResponse, reader: jspb.BinaryReader): SearchFileResponse;
}

export namespace SearchFileResponse {
  export type AsObject = {
    fileFound: boolean,
    actionDescription: string,
    emailAddress: string,
    folderName: string,
    expenses?: Expenses.AsObject,
  }
}

export class ExtractFileResponse extends jspb.Message {
  getFileExtracted(): boolean;
  setFileExtracted(value: boolean): ExtractFileResponse;

  getActionDescription(): string;
  setActionDescription(value: string): ExtractFileResponse;

  getEmailAddress(): string;
  setEmailAddress(value: string): ExtractFileResponse;

  getFolderName(): string;
  setFolderName(value: string): ExtractFileResponse;

  getFile(): ExpenseItem | undefined;
  setFile(value?: ExpenseItem): ExtractFileResponse;
  hasFile(): boolean;
  clearFile(): ExtractFileResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ExtractFileResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ExtractFileResponse): ExtractFileResponse.AsObject;
  static serializeBinaryToWriter(message: ExtractFileResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ExtractFileResponse;
  static deserializeBinaryFromReader(message: ExtractFileResponse, reader: jspb.BinaryReader): ExtractFileResponse;
}

export namespace ExtractFileResponse {
  export type AsObject = {
    fileExtracted: boolean,
    actionDescription: string,
    emailAddress: string,
    folderName: string,
    file?: ExpenseItem.AsObject,
  }
}

export enum MimeType { 
  MIME_TYPE_UNSPECIFIED = 0,
  IMAGE_JPEG = 1,
  IMAGE_PNG = 2,
  APPLICATION_PDF = 3,
}
